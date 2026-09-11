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

type CrawlerGretsch struct {
    name     string
    gScraper Crawler[*model.Guitar]
}

type CallBacksGretsch struct {
    funcs CallBacks
}

func NewScraperGretsch() Scraper[*model.Guitar] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
        Delay:       250 * time.Millisecond,
        RandomDelay: 750 * time.Millisecond,
	})
    return &CrawlerGretsch{
        "GRETSCH",
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksGretsch() *CallBacksGretsch {
    return &CallBacksGretsch{
        CallBacks{},
    }
}

func (g *CrawlerGretsch) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(c ,crawlStats)

    // URL収集、クロール
    visited := make(map[string]struct{}, 300)

    mutex := &sync.Mutex{}

    // ページネーション
    c.OnHTML(`.pagination a[href^="?page="]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    // 詳細ページ
    c.OnHTML(`#product-grid-anchor a[href*="/gear/"][data-product-id]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })
    c.Wait()

    c.Visit("https://www.gretschguitars.jp/gear/")
    c.Wait()

    loggingCrawlStats(g.name, crawlStats)

    g.gScraper.urls = utils.MapToSliceUrl(visited)

    return g.gScraper.urls, nil
}

func (g *CrawlerGretsch) Scrape(provider  PageProvider,
                                parser    ModelParser[*model.Guitar],
                                parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *CallBacksGretsch) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        // 動的ページを取得しない場合、引数のパターンは記載しないで良い
        if !isDetailPage(``, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 3 * time.Second)
        defer cancel()

        var html string

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.WaitVisible(".shrink", chromedp.ByQuery), // 求める要素が出るまで待つ
            chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )

        if err != nil {
            log.Printf("[Chromedp error]: %v", err)
        }
        return html, nil
    }
}

func (c *CallBacksGretsch) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{} // 捨てる属性は基本的に空文字を割り当てる

        spec[C.Maker] = strconv.Itoa(C.GRETSCH)
        spec[C.Name]  = doc.Find(`.sku-number`).Prev().Text()
        spec[C.Color] = doc.Find(`.spec-name:contains("Color")`).Next().Text()

        spec[C.BodyFinish]       = doc.Find(`.spec-name:contains("Body Finish")`).Next().Text()

        // Mahogany with Arched Maple Top << ボディ材とトップ材は左記の様にまとめられている
        backWithTop    := doc.Find(`.spec-name:contains("Body Material")`).Next().Text()
        backWithTopArr := strings.Split(backWithTop, "with")

        spec[C.BodyMaterialBack] = backWithTopArr[0]

        if len(backWithTopArr) > 1 {
            spec[C.BodyMaterialTop] = backWithTopArr[1]
        } else {
            spec[C.BodyMaterialTop]  = ""
        }

        // Bridge Pickup も取得してしまうため、他の要素から迂回して取得
        spec[C.Bridge]   = doc.Find(`tr:contains("Tuning Machines")`).Prev().Find(`.spec-value`).Text()
        spec[C.Controls] = doc.Find(`.spec-name:contains("Controls")`).Next().Text()

        // コメントは複数の pタグ に分割されている。数は不明確。おおまかいくつかに取得する
        spec[C.Comment]  = doc.Find(`.sku-description`).Next().Text() +
                           doc.Find(`.sku-description`).Next().Next().Text() +
                           doc.Find(`.sku-description`).Next().Next().Next().Text() +
                           doc.Find(`.sku-description`).Next().Next().Next().Next().Text()

        spec[C.Fingerboard]  = doc.Find(`.spec-name:contains("Fingerboard Material")`).Next().Text()
        spec[C.FretCount]    = doc.Find(`.spec-name:contains("Number of Frets")`).Next().Text()
        spec[C.Inlays]       = doc.Find(`.spec-name:contains("Position Inlays")`).Next().Text()
        spec[C.Joint]        = ""
        spec[C.NeckMaterial] = doc.Find(`.spec-name:contains("Neck Material")`).Next().Text()

        // 基本的にはNeck, Center, Bridgeを取得してフレーム側で組み立てる。分割して取得できなければ Pickups へ
        spec[C.Pickups]      = ""
        spec[C.NeckPickup]   = doc.Find(`.spec-name:contains("Neck Pickup")`).Next().Text()
        spec[C.CenterPickup] = ""
        spec[C.BridgePickup] = doc.Find(`.spec-name:contains("Bridge Pickup")`).Next().Text()

        spec[C.Price]         = doc.Find(`.sku-price`).Text()
        spec[C.ScaleLengthMM] = ""
        spec[C.Series]        = ""

        src, _        := doc.Find(`.sku-main-image img[src*="/media/CACHE/images/products/"]`).Attr(`src`)
        spec[C.Src]    = src
        spec[C.Weight] = strconv.Itoa(C.InvalidNumber)

        specs = utils.LockedAppend(mutex, specs, spec)

        return specs
    }
}

func (c *CallBacksGretsch) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url)
    }
}

func (c *CallBacksGretsch) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, "body")
    }
}