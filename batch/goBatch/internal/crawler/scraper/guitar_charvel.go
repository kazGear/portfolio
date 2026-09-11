package scraper

import (
	"context"
	"log"
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

type CrawlerCharvel struct {
    name     string
    gScraper Crawler[*model.Guitar]
}

type CallBacksCharvel struct {
    funcs CallBacks
}

func NewScraperCharvel() Scraper[*model.Guitar] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(4),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
        Delay:       250 * time.Millisecond,
        RandomDelay: 750 * time.Millisecond,
	})
    return &CrawlerCharvel{
        "CHARVEL",
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksCharvel() *CallBacksCharvel {
    return &CallBacksCharvel{
        CallBacks{},
    }
}

func (g *CrawlerCharvel) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(c ,crawlStats)

    // URL収集、クロール
    visited := make(map[string]struct{}, 250)

    mutex := &sync.Mutex{}

    // ページネーション
    c.OnHTML(`a[href^="?sort=latest&page="]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    // 詳細ページ
    c.OnHTML(`.product-tile a[href^="/gear/"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    // 詳細ぺージ > カラーバリエーション
    c.OnHTML(`.tile-placeholder[href^="/gear/"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    c.Visit("https://www.charvel.jp/gear/?sort=latest")
    c.Wait()

    loggingCrawlStats(g.name, crawlStats)

    g.gScraper.urls = utils.MapToSliceUrl(visited)

    return g.gScraper.urls, nil
}

func (g *CrawlerCharvel) Scrape(provider  PageProvider,
                                parser    ModelParser[*model.Guitar],
                                parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *CallBacksCharvel) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
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

func (c *CallBacksCharvel) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{} // 捨てる属性は基本的に空文字を割り当てる

        spec[C.Maker] = strconv.Itoa(C.CHARVEL)
        spec[C.Name]  = doc.Find(`.product-title`).Text()
        spec[C.Color] = doc.Find(`.spec-name:contains("Color")`).Next().Text()

        spec[C.BodyFinish]       = doc.Find(`.spec-name:contains("Body Finish")`).Next().Text()
        spec[C.BodyMaterialBack] = doc.Find(`.spec-name:contains("Body Material")`).Next().Text()
        spec[C.BodyMaterialTop]  = "" // 記載なし

        spec[C.Bridge]   = doc.Find(`#product-spec-tuning-machines`).Prev().Find(`.spec-value`).Text()
        spec[C.Controls] = doc.Find(`.spec-name:contains("Controls")`).Next().Text()
        spec[C.Comment]  = doc.Find(`.overview-desc`).Text()

        spec[C.Fingerboard]  = doc.Find(`.spec-name:contains("Fingerboard Material")`).Next().Text()
        spec[C.FretCount]    = doc.Find(`.spec-name:contains("Number of Frets")`).Next().Text()
        spec[C.Inlays]       = doc.Find(`.spec-name:contains("Position Inlays")`).Next().Text()
        spec[C.Joint]        = doc.Find(`.spec-name:contains("Neck Plate")`).Next().Text()
        spec[C.NeckMaterial] = doc.Find(`.spec-name:contains("Neck Material")`).Next().Text()

        // 基本的にはNeck, Center, Bridgeを取得してフレーム側で組み立てる。分割して取得できなければ Pickups へ
        spec[C.Pickups]      = ""
        spec[C.NeckPickup]   = doc.Find(`.spec-name:contains("Neck Pickup")`).Next().Text()

        pickupLayout := doc.Find(`.spec-name:contains("Configuration")`).Next().Text()

        if strings.Contains(pickupLayout, "HSS") {
            // 3 pickup 構成でも記載なし。センターが存在することだけを示す
            spec[C.CenterPickup] = "??"
        } else {
            spec[C.CenterPickup] = ""
        }
        spec[C.BridgePickup] = doc.Find(`.spec-name:contains("Bridge Pickup")`).Next().Text()

        spec[C.Price]         = doc.Find(`.currency`).Text()
        spec[C.ScaleLengthMM] = ""
        spec[C.Series]        = doc.Find(`.spec-name:contains("Series")`).Next().Text()

        src, _        := doc.Find(`.pdp-image img[src*="/media/CACHE/images/products/"]`).Attr(`src`)
        spec[C.Src]    = src
        spec[C.Weight] = strconv.Itoa(C.InvalidNumber)

        specs = utils.LockedAppend(mutex, specs, spec)

        return specs
    }
}

func (c *CallBacksCharvel) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url)
    }
}

func (c *CallBacksCharvel) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, "body")
    }
}