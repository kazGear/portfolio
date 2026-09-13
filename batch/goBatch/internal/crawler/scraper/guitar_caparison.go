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
)

type CrawlerCaparison struct {
    name     string
    gScraper Crawler[*model.Guitar]
}

type CallBacksCaparison struct {
    funcs CallBacks
}

func NewScraperCaparison() Scraper[*model.Guitar] {
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
    return &CrawlerCaparison{
        "Caparison",
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksCaparison() *CallBacksCaparison {
    return &CallBacksCaparison{
        CallBacks{},
    }
}

func (g *CrawlerCaparison) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(c ,crawlStats)

    // URL収集、クロール
    visited := make(map[string]struct{}, 100)
    mutex   := &sync.Mutex{}

    c.OnHTML(`a[href^="/ja/products/"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    c.Visit("https://www.caparisonguitars.com/ja/pages/products")
    c.Wait()

    loggingCrawlStats(g.name, crawlStats)

    g.gScraper.urls = utils.MapToSliceUrl(visited)

    return g.gScraper.urls, nil
}

func (g *CrawlerCaparison) Scrape(provider  PageProvider,
                                  parser    ModelParser[*model.Guitar],
                                  parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *CallBacksCaparison) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        // 動的ページを取得しない場合、引数のパターンは記載しないで良い
        if !isDetailPage(``, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 4*time.Second)
        defer cancel()

        var html string

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.WaitVisible("body", chromedp.ByQuery), // 求める要素が出るまで待つ
            chromedp.Sleep(300 * time.Millisecond), // JSが動く猶予を与える
            chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )

        if err != nil {
            log.Printf("[Chromedp error]: %v", err)
        }
        return html, nil
    }
}

var _regColorsStartTagCaparison  = regexp.MustCompile(`\<span.*\"\>`)
var _regColorsEndTagCaparison    = regexp.MustCompile(`\</span\>`)
var _regImageEndPatternCaparison = regexp.MustCompile(`(Full|-?Front)?\.(jpg|png).+`)

func (c *CallBacksCaparison) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        specs := make([]map[string]string, 0, 5)
        mutex := &sync.Mutex{}

        organizedSpec := organizeSpecCaparison(doc)

        colors := getColorsCaparison(organizedSpec["Finish"])

        comment := doc.Find(`.isolate`).First().Text()

        // 価格算出
        exchangeRate := utils.GetExchangeUSDtoJPY()
        foreignPrice, _ := doc.Find(`meta[property="og:price:amount"]`).Attr(`content`)
        price := utils.CalcExchangedPrice(foreignPrice, exchangeRate)

        // 画像URLセットを作成
        urls := make(map[string]struct{})

        doc.Find(`.collage-card img`).Each(func(index int, selector *goquery.Selection) {
            url, _ := selector.Attr("src")
            urls[url] = struct{}{}
        })

        // 名称、カラー、画像を動的に設定する
        for _, color := range colors {
            spec := map[string]string{} // 捨てる属性は基本的に空文字を割り当てる

            url := getUrlCaparison(urls, color)

            spec[C.Maker] = strconv.Itoa(C.Caparison)
            spec[C.Name]  = getGuitarNameCaparison(url)
            spec[C.Color] = color

            spec[C.BodyFinish]       = ""
            spec[C.BodyMaterialBack] = organizedSpec["Body"]
            spec[C.BodyMaterialTop]  = organizedSpec["Body Top"]

            spec[C.Bridge]   = organizedSpec["Bridge"]
            spec[C.Controls] = organizedSpec["Controls"]
            spec[C.Comment]  = comment

            spec[C.Fingerboard]  = organizedSpec["Fretboard"]
            spec[C.FretCount]    = organizedSpec["Frets"]
            spec[C.Inlays]       = organizedSpec["Position Inlay"]
            spec[C.Joint]        = organizedSpec["Neck Joint"]
            spec[C.NeckMaterial] = organizedSpec["Neck Material"]

            // 基本的にはNeck, Center, Bridgeを取得してフレーム側で組み立てる。分割して取得できなければ Pickups へ
            spec[C.Pickups]      = ""
            spec[C.NeckPickup]   = organizedSpec["Neck PU"]
            spec[C.CenterPickup] = organizedSpec["Middle PU"]
            spec[C.BridgePickup] = organizedSpec["Bridge PU"]

            spec[C.Price]         = price
            spec[C.ScaleLengthMM] = organizedSpec["Scale Length"]
            spec[C.Series]        = organizedSpec["Body Shape"]

            spec[C.Src]    = url
            spec[C.Weight] = strconv.Itoa(C.InvalidNumber)

            specs = utils.LockedAppend(mutex, specs, spec)
        }
        return specs
    }
}

/* html構造が粗いので、予めデータを整理しておく
    <div class="image-with-text__text rte body">
        <p>
            Model: TAT Custom
            <br/>
            Body Shape: TAT
            <br/>
            ...
        </p>
    </div>
    ...(上記が複数ブロックある)
*/
func organizeSpecCaparison(doc *goquery.Document) map[string]string {
    organizedSpec := make(map[string]string)

    doc.Find(`.image-with-text__text p`).Each(func(index int, selector *goquery.Selection) {
        html,  _ := selector.Html()
        specRows := strings.Split(html, "<br/>")

        for _, specRow := range specRows {
            keyAndValue := strings.Split(specRow, ":")

            if len(keyAndValue) > 1 {
                organizedSpec[strings.TrimSpace(keyAndValue[0])] = strings.TrimSpace(keyAndValue[1])
            }
        }
    })
    return organizedSpec
}

func getGuitarNameCaparison(imageSrc string) string {
    guitarName := strings.ReplaceAll(imageSrc, `//www.caparisonguitars.com/cdn/shop/files/`, "")
    guitarName  = _regImageEndPatternCaparison.ReplaceAllString(guitarName, "")
    guitarName  = strings.ReplaceAll(guitarName, "__", "-")
    guitarName  = strings.ReplaceAll(guitarName, "_", "")

    return guitarName
}

// arg: <span>..., ...</span>の形式で取得されている
func getColorsCaparison(colorHtml string) []string {
    colorHtml  = _regColorsStartTagCaparison.ReplaceAllString(colorHtml, "")
    colorHtml  = _regColorsEndTagCaparison.ReplaceAllString(colorHtml, "")
    colors    := strings.Split(colorHtml, ",")

    return colors
}

var _regNormalizeColorCaparison = regexp.MustCompile(`[a-z. ]`)

// 段階的に加工しながらマッチする URL を探す
func getUrlCaparison(urls map[string]struct{}, color string) string {
    // urlがひとつだけならそれが対応するurl
    if len(urls) == 1 {
        for url := range urls {
            return url
        }
    }
    color            = strings.TrimSpace(color)
    normalizedColor := normalizeForCompareCaparison(color)

    for url := range urls {
        // カラー名を完全に含むか
        normalizedUrl := normalizeForCompareCaparison(url)

        if strings.Contains(normalizedUrl, normalizedColor) {
            return url
        }

        // 正規化したカラー名が含まれているか
        if strings.Contains(url, _regNormalizeColorCaparison.ReplaceAllString(color, "")) {
            return url
        }

        // 部分的なカラー名が含まれるか
        colorParts := strings.Split(color, " ")

        for _, colorPart := range colorParts {
            if strings.Contains(url, colorPart) {
                return url
            }
        }
    }
    return "" // url (名称の源泉) がないため、DBには登録されない
}

func normalizeForCompareCaparison(target string) string {
    normalized := strings.TrimSpace(target)
    normalized  = strings.ReplaceAll(normalized, " ", "")
    normalized  = strings.ReplaceAll(normalized, "-", "")
    normalized  = strings.ReplaceAll(normalized, "_", "")

    return normalized
}

func (c *CallBacksCaparison) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url)
    }
}

func (c *CallBacksCaparison) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, "body")
    }
}