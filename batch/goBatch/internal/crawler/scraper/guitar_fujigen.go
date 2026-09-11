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

type CrawlerFujigen struct {
    name     string
    gScraper Crawler[*model.Guitar]
}

type CallBacksFujigen struct {
    funcs CallBacks
}

func NewScraperFujigen() Scraper[*model.Guitar] {
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
    return &CrawlerFujigen{
        "FUJIGEN",
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksFujigen() *CallBacksFujigen {
    return &CallBacksFujigen{
        CallBacks{},
    }
}

func (g *CrawlerFujigen) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(c ,crawlStats)

    // URL収集、クロール
    visited := make(map[string]struct{}, 300)

    mutex := &sync.Mutex{}

    // 詳細ページ
    c.OnHTML(`.modellist_item_img a[href*="detail.php?product_id="]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    c.Visit("https://fujigen.shop/products/list.php?category_id=7") // guitar
    c.Visit("https://fujigen.shop/products/list.php?category_id=8") // base
    c.Wait()

    loggingCrawlStats(g.name, crawlStats)

    g.gScraper.urls = utils.MapToSliceUrl(visited)

    return g.gScraper.urls, nil
}

func (g *CrawlerFujigen) Scrape(provider  PageProvider,
                                parser    ModelParser[*model.Guitar],
                                parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *CallBacksFujigen) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
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

func (c *CallBacksFujigen) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{} // 捨てる属性は基本的に空文字を割り当てる

        spec[C.Maker] = strconv.Itoa(C.FUJIGEN)
        spec[C.Name]  = doc.Find(`.item_detail_info_name`).Text()
        // Hardware Color への誤爆。Color だと複数の情報を拾ってしまうため Accessories から迂回して取得
        spec[C.Color] = doc.Find(`.item_spec_list div:contains("Accessories")`).Next().Children().Next().Text()

        spec[C.BodyFinish]       = doc.Find(`.item_spec_list dt:contains("Body Finish")`).Next().Text()
        // Body Finish への誤爆。Body だと複数の情報を拾ってしまうため Construction から迂回して取得
        spec[C.BodyMaterialBack] =
            doc.Find(`.item_spec_list div:contains("Construction")`).Next().Children().Next().Text()
        spec[C.BodyMaterialTop]  = "" // 記載なし

        // Pickup (Bridge) への誤爆。Bridge だと複数の情報を拾ってしまうため Tuners から迂回して取得
        spec[C.Bridge]   = doc.Find(`.item_spec_list div:contains("Tuners")`).Next().Children().Next().Text()
        spec[C.Controls] = doc.Find(`.item_spec_list dt:contains("Controls")`).Next().Text()
        spec[C.Comment]  = strings.Split(doc.Find(`#item_comment`).Text(), "=====")[0]

        spec[C.Fingerboard]  = doc.Find(`.item_spec_list dt:contains("Fingerboard")`).Next().Text()
        spec[C.FretCount]    = doc.Find(`.item_spec_list dt:contains("Frets")`).Next().Text()
        spec[C.Inlays]       = ""
        spec[C.Joint]        = doc.Find(`.item_spec_list dt:contains("Construction")`).Next().Text()
        spec[C.NeckMaterial] = doc.Find(`.item_spec_list dt:contains("Neck")`).Next().Text()

        // 基本的にはNeck, Center, Bridgeを取得してフレーム側で組み立てる。分割して取得できなければ Pickups へ
        spec[C.Pickups]      = ""
        spec[C.NeckPickup]   = doc.Find(`.item_spec_list dt:contains("Pickup (Neck)")`).Next().Text()
        spec[C.CenterPickup] = doc.Find(`.item_spec_list dt:contains("Pickup (Middle)")`).Next().Text()
        spec[C.BridgePickup] = doc.Find(`.item_spec_list dt:contains("Pickup (Bridge)")`).Next().Text()

        spec[C.Price]         = doc.Find(`.item_detail_info_price`).Text()
        spec[C.ScaleLengthMM] = doc.Find(`.item_spec_list dt:contains("Scale")`).Next().Text()
        spec[C.Series]        = doc.Find(`.item_detail_info_series`).Children().Text()

        src, _        := doc.Find(`.item_img_wrap img[src^="/upload/save_image/"]`).Attr(`src`)
        spec[C.Src]    = src
        spec[C.Weight] = strconv.Itoa(C.InvalidNumber)

        specs = utils.LockedAppend(mutex, specs, spec)

        return specs
    }
}

func (c *CallBacksFujigen) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url)
    }
}

func (c *CallBacksFujigen) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, "body")
    }
}