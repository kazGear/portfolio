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

type CrawlerMayones struct {
    name     string
    gScraper Crawler[*model.Guitar]
}

type CallBacksMayones struct {
    funcs CallBacks
}

func NewScraperMayones() Scraper[*model.Guitar] {
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
    return &CrawlerMayones{
        "MAYONES",
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksMayones() *CallBacksMayones {
    return &CallBacksMayones{
        CallBacks{},
    }
}

func (g *CrawlerMayones) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(c ,crawlStats)

    // URL収集、クロール
    visited := make(map[string]struct{}, 130)
    mutex   := &sync.Mutex{}

    c.OnHTML(`a[href^="https://mayones.com/page/"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })

    c.Visit("https://mayones.com/guitars/mbc_guitars/") // Master Builder Collection
    c.Visit("https://mayones.com/custom-shop/custom-shop-guitar-gallery/")
    c.Visit("https://mayones.com/basses/mbc_basses/") // Master Builder Collection
    c.Visit("https://mayones.com/custom-shop/custom-shop-bass-gallery/")
    c.Wait()

    loggingCrawlStats(g.name, crawlStats)

    g.gScraper.urls = utils.MapToSliceUrl(visited)
    return g.gScraper.urls, nil
}

func (g *CrawlerMayones) Scrape(provider  PageProvider,
                                parser    ModelParser[*model.Guitar],
                                parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *CallBacksMayones) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
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

func (c *CallBacksMayones) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{} // 捨てる属性は基本的に空文字を割り当てる

        spec[C.Maker] = strconv.Itoa(C.MAYONES)
        spec[C.Name]  = doc.Find(`.model_name h1`).Text()
        spec[C.Color] = " " // pk制約回避。カラー取れないけど登録したい

        spec[C.BodyFinish]       = doc.Find(`div:contains("Available finishes")`).Next().Find(`h3`).Text()
        spec[C.BodyMaterialBack] = doc.Find(`.acf_label:contains("Body:")`).Next().Find(`h3`).Text()
        spec[C.BodyMaterialTop]  = doc.Find(`.acf_label:contains("Top:")`).Next().Find(`h3`).Text()

        spec[C.Bridge]   = doc.Find(`.acf_label:contains("Bridge")`).Next().Find(`h3`).Text()
        spec[C.Controls] = doc.Find(`.acf_label:contains("Control:")`).Next().Find(`h3`).Text()
        spec[C.Comment]  = ""

        spec[C.Fingerboard]  = doc.Find(`.acf_label:contains("Fingerboard:")`).Next().Find(`h3`).Text()
        spec[C.FretCount]    = doc.Find(`.acf_label:contains("Frets:")`).Next().Find(`h3`).Text()
        spec[C.Inlays]       = doc.Find(`.acf_label:contains("Markers & Inlays:")`).Next().Find(`h3`).Text()
        spec[C.Joint]        = doc.Find(`.acf_label:contains("Construction:")`).Next().Find(`h3`).Text()
        spec[C.NeckMaterial] = doc.Find(`.acf_label:contains("Neck:")`).Next().Find(`h3`).Text()

        // 基本的にはNeck, Center, Bridgeを取得してフレーム側で組み立てる。分割して取得できなければ Pickups へ
        spec[C.Pickups]      = doc.Find(`.acf_label:contains("Pickups / Electronics:")`).Next().Find(`h3`).Text()
        spec[C.NeckPickup]   = ""
        spec[C.CenterPickup] = ""
        spec[C.BridgePickup] = ""

        spec[C.Price]         = ""
        spec[C.ScaleLengthMM] = doc.Find(`.acf_label:contains("Scale:")`).Next().Find(`h3`).Text()
        spec[C.Series]        = strings.Split(spec[C.Name], " ")[0]

        src, _        := doc.Find(`.wpb_single_image img[src*="/wp-content/uploads/"]`).Attr(`src`)
        spec[C.Src]    = src
        spec[C.Weight] = strconv.Itoa(C.InvalidNumber)

        specs = utils.LockedAppend(mutex, specs, spec)

        return specs
    }
}

func (c *CallBacksMayones) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url)
    }
}

func (c *CallBacksMayones) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, "body")
    }
}