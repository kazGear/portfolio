package scraper

import (
	"context"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
	"golang.org/x/text/width"
)

type CrawlerKiller struct {
    name     string
    gScraper Crawler[*model.Guitar]
}

type CallBacksKiller struct {
    funcs CallBacks
}

func NewScraperKiller() Scraper[*model.Guitar] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(2),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
        Delay:       250 * time.Millisecond,
        RandomDelay: 750 * time.Millisecond,
	})
    return &CrawlerKiller{
        "Killer",
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksKiller() *CallBacksKiller {
    return &CallBacksKiller{
        CallBacks{},
    }
}

func (g *CrawlerKiller) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(c ,crawlStats)

    // URL収集、クロール
    visited := make(map[string]struct{}, 100)

    mutex := &sync.Mutex{}

    // 詳細ページ（精度粗め）
    c.OnHTML(`a[href^="kg-"], a[href^="kb-"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    c.Visit("https://killer.jp/guitar/index.html")
    c.Visit("https://killer.jp/bass/index.html")
    c.Wait()

    loggingCrawlStats(g.name, crawlStats)

    g.gScraper.urls = utils.MapToSliceUrl(visited)
    return g.gScraper.urls, nil
}

func (g *CrawlerKiller) Scrape(provider  PageProvider,
                               parser    ModelParser[*model.Guitar],
                               parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *CallBacksKiller) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        // 動的ページを取得しない場合、引数のパターンは記載しないで良い
        if !isDetailPage(``, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()

        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 5 * time.Second)
        defer cancel()

        var html string

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.WaitVisible("#price", chromedp.ByQuery), // 求める要素が出るまで待つ
            chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )

        if err != nil {
            log.Printf("[Chromedp error]: %v", err)
        }
        return html, nil
    }
}

func (c *CallBacksKiller) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        images := getAllImagesKiller(doc)

        pickups := getPickupsKiller(doc)

        colorCount  := getColorCountKiller(doc)
        dirtyColors := getAllColorsKiller(doc, colorCount)
        colors      := cleanupColorsKiller(dirtyColors)

        // カラバリごとにギター情報を収集
        for color := range colors {
            src := getGuitarImageKiller(color, images)

            if src == "" {
                return []map[string]string{}
            }

            spec := map[string]string{} // 捨てる属性は基本的に空文字を割り当てる

            spec[C.Maker] = strconv.Itoa(C.Killer)
            spec[C.Name]  = doc.Find(`h1`).Text()
            spec[C.Color] = color

            spec[C.BodyFinish]       = getExactMatchedDoc(doc, "td", "Paint").Next().Text()
            spec[C.BodyMaterialBack] = getExactMatchedDoc(doc, "td", "Body").Next().Text()
            spec[C.BodyMaterialTop]  = ""

            spec[C.Bridge]   = ""
            spec[C.Controls] = doc.Find(`td:contains("Controls")`).Next().Text()
            spec[C.Comment]  = ""

            spec[C.Fingerboard]  = doc.Find(`td:contains("Fingerboard")`).Next().Text()
            // Fingerboard にはフレット情報を含む
            spec[C.FretCount]    = doc.Find(`td:contains("Fingerboard")`).Next().Text()
            spec[C.Inlays]       = doc.Find(`td:contains("Position mark")`).Next().Text()
            spec[C.Joint]        = ""
            // 誤爆するので迂回して取得
            spec[C.NeckMaterial] = doc.Find(`tr:contains("Fingerboard")`).Prev().Children().Next().Text()

            // 基本的にはNeck, Center, Bridgeを取得してフレーム側で組み立てる。分割して取得できなければ Pickups へ
            spec[C.Pickups]      = pickups
            spec[C.NeckPickup]   = ""
            spec[C.CenterPickup] = ""
            spec[C.BridgePickup] = ""

            spec[C.Price]         = doc.Find(`div.price`).Text()
            spec[C.ScaleLengthMM] = doc.Find(`td:contains("Scale length")`).Next().Text()
            spec[C.Series]        = strings.Split(spec[C.Name], " ")[0]

            // ギター名の先頭は KG 、ベース名の先頭は KB
            if strings.Contains(spec[C.Name], "KG") {
                spec[C.Src] = "https://killer.jp/guitar/" + src
            } else if strings.Contains(spec[C.Name], "KB") {
                spec[C.Src] = "https://killer.jp/bass/" + src
            }

            spec[C.Weight] = strconv.Itoa(C.InvalidNumber)

            specs = utils.LockedAppend(mutex, specs, spec)
        }
        return specs
    }
}

// rowspan があるラベル名が統一されていない。Body Color, Body color, color, ...
func getColorCountKiller(doc *goquery.Document) int {
    // rowspan = カラー数
    colorCountStr, _ := doc.Find(`td:contains("Body color")`).Attr("rowspan")
    colorCount, err  := strconv.Atoi(colorCountStr)

    if err == nil {
        return colorCount
    }

    colorCountStr, _ = doc.Find(`td:contains("Body Color")`).Attr("rowspan")
    colorCount, err  = strconv.Atoi(colorCountStr)

    if err == nil {
        return colorCount
    }

    selector        := getExactMatchedDoc(doc, "td", "Color")
    colorCountStr, _ = selector.Attr("rowspan")
    colorCount, err  = strconv.Atoi(colorCountStr)

    if err == nil {
        return colorCount
    }
    // 1 color の場合、rowspan 属性は存在せず取れない
    return 1
}

// ラベル名が統一されていない。Body Color, Body color, ...
func getAllColorsKiller(doc *goquery.Document, colorCount int) map[string]struct{} {
    colors := make(map[string]struct{})

    colorStartingPoint := doc.Find(`tr:contains("Body Color")`)

    // ラベル揺れがあるので手当たり次第要素を拾ってみる
    if len(colorStartingPoint.Text()) <= 0 {
        colorStartingPoint = doc.Find(`tr:contains("Body color")`)

        if len(colorStartingPoint.Text()) <= 0 {
            colorStartingPoint = getExactMatchedDoc(doc, "td", "Color")
        }
    }

    colorsWithGarbagesHtml := colorStartingPoint.AddSelection(colorStartingPoint.NextAll())

    // 最初の要素だけ取得しておく（"Body color" のテキストが入らないように）
    colors[strings.TrimSpace(colorStartingPoint.Children().Next().Text())] = struct{}{}

    // 最初のカラー名は取得済
    for i := 1; i < colorCount; i++ {
        color := strings.TrimSpace(colorsWithGarbagesHtml.Eq(i).Text())
        colors[color] = struct{}{}
    }
    return colors
}

func getAllImagesKiller(doc *goquery.Document) map[string]struct{} {
    images := make(map[string]struct{})

    imagesHtml := doc.Find(`h2:contains("製品画像")`).Next().Find(`img`)

    imagesHtml.Each(func(index int, selector *goquery.Selection) {
        // 画像が表面、裏面、... と並んでおり、表面の画像だけが必要
        if index % 2 == 0 {
            image, _ := selector.Attr("src")
            images[strings.TrimSpace(image)] = struct{}{}
        }
    })
    return images
}

func getPickupsKiller(doc *goquery.Document) string {
    pickupCountStr, _ := doc.Find(`tr:contains("Scale length")`).Next().Children().Attr("rowspan")
    pickupCount, err  := strconv.Atoi(pickupCountStr)

    if err != nil {
        return ""
    }

    pickups := ""

    // pickup を直接取得できないので迂回して取得
    pickupsStartingPoint    := doc.Find(`tr:contains("Scale length")`).Next()
    pickupsWithGarbagesHtml := pickupsStartingPoint.AddSelection(pickupsStartingPoint.NextAll())

    // 最初の要素だけ取得しておく（"Pickup" のテキストが入らないように）
    pickups = strings.TrimSpace(pickupsStartingPoint.Children().Next().Text())

    // 最初のカラー名は取得済
    for i := 1; i < pickupCount; i++ {
        pickup  := strings.TrimSpace(pickupsWithGarbagesHtml.Eq(i).Text())
        pickups += " / " + pickup
    }
    return pickups
}

func normalizeForCompareKiller(target string) string {
    normalized := width.Narrow.String(target)
    normalized  = strings.TrimSpace(target)
    normalized  = strings.ToLower(target)
    normalized  = strings.ReplaceAll(normalized, " ", "")
    normalized  = strings.ReplaceAll(normalized, "-", "")

    return normalized
}

var _regCleanupColorKiller1 = regexp.MustCompile(`\(.*`)
var _regCleanupColorKiller2 = regexp.MustCompile(`(JAN|jan):\d*`)
var _regCleanupColorKiller3 = regexp.MustCompile(`\*.*`)
var _regCleanupColorKiller4 = regexp.MustCompile(`（.*`)
var _regCleanupColorKiller5 = regexp.MustCompile(`\/.*`)

func cleanupColorsKiller(colors map[string]struct{}) map[string]struct{} {
    cleanedColors := map[string]struct{}{}

    for k := range colors {
        k = _regCleanupColorKiller1.ReplaceAllString(k, "")
        k = _regCleanupColorKiller2.ReplaceAllString(k, "")
        k = _regCleanupColorKiller3.ReplaceAllString(k, "")
        k = _regCleanupColorKiller4.ReplaceAllString(k, "")
        k = _regCleanupColorKiller5.ReplaceAllString(k, "")

        cleanedColors[k] = struct{}{}
    }
    return cleanedColors
}

func getGuitarImageKiller(color string, images map[string]struct{}) string {
    // カラーからギター画像を選択する
    normalizedColor := normalizeForCompareKiller(color)
    src := ""

    // image にはカラー名が含まれている（無いことも）
    for image := range images {
        normalizedImage := normalizeForCompareKiller(image)

        if strings.Contains(normalizedImage, normalizedColor) {
            src = image
            break
        }
    }
    return src
}

func (c *CallBacksKiller) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url)
    }
}

func (c *CallBacksKiller) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, `body`)
    }
}