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

type CrawlerBCRich struct {
    name     string
    gScraper Crawler[*model.Guitar]
}

type CallBacksBCRich struct {
    funcs CallBacks
}

func NewScraperBCRich() Scraper[*model.Guitar] {
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
    return &CrawlerBCRich{
        "B.C.Rich",
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksBCRich() *CallBacksBCRich {
    return &CallBacksBCRich{
        CallBacks{},
    }
}

func (g *CrawlerBCRich) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(c ,crawlStats)

    // URL収集、クロール
    visited := make(map[string]struct{}, 100)

    mutex := &sync.Mutex{}

    // ページネーション
    c.OnHTML(`a[aria-label^="Page"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    // 詳細ページ
    c.OnHTML(`.product-feed-container a[href^="https://bcrich.com/product/"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    c.Visit("https://bcrich.com/product-category/guitars/")
    c.Wait()

    loggingCrawlStats(g.name, crawlStats)

    g.gScraper.urls = utils.MapToSliceUrl(visited)

    return g.gScraper.urls, nil
}

func (g *CrawlerBCRich) Scrape(provider  PageProvider,
                               parser    ModelParser[*model.Guitar],
                               parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *CallBacksBCRich) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
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
            chromedp.WaitVisible("body", chromedp.ByQuery), // 求める要素が出るまで待つ
            chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )

        if err != nil {
            log.Printf("[Chromedp error]: %v", err)
        }
        return html, nil
    }
}

var _exchangeRateBCRich     = utils.GetExchangeUSDtoJPY()
var _regPriceFractionBCRich = regexp.MustCompile(`\.\d+`)

func (c *CallBacksBCRich) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{} // 捨てる属性は基本的に空文字を割り当てる

        spec[C.Maker] = strconv.Itoa(C.BCRich)
        spec[C.Name]  = doc.Find(`#specs h1`).Text()
        spec[C.Color] = doc.Find(`#color option[selected]`).Text()

        // ギター以外の情報も流れてくる。それらはカラーを取得できておらず、そのまま破棄する
        if len(spec[C.Color]) >= 1 {
            // pk制約回避。カラーを画像と対応させて取得できないためカラー情報を破棄
            spec[C.Color] = C.ColorUndefined
        }

        spec[C.BodyFinish]       = doc.Find(`.product-spec-title:contains("Body Finish")`).Next().Text()
        spec[C.BodyMaterialBack] = doc.Find(`.product-spec-title:contains("Body Wood")`).Next().Text()
        spec[C.BodyMaterialTop]  = doc.Find(`.product-spec-title:contains("Top Wood")`).Next().Text()

        spec[C.Bridge]   = doc.Find(`.product-spec-title:contains("Bridge Design")`).Next().Text()
        spec[C.Controls] = doc.Find(`.product-spec-title:contains("Control Layout")`).Next().Text()
        spec[C.Comment]  = doc.Find(`.has-default-attributes`).Children().Next().Next().Find(`.fl-rich-text p`).Text()

        spec[C.Fingerboard]  = doc.Find(`.product-spec-title:contains("Material")`).Next().Text()
        spec[C.FretCount]    = doc.Find(`.product-spec-title:contains("Number of Frets")`).Next().Text()
        spec[C.Inlays]       = doc.Find(`.product-spec-title:contains("Inlays")`).Next().Text()
        spec[C.Joint]        = doc.Find(`.product-spec-title:contains("Construction")`).Next().Text()
        spec[C.NeckMaterial] = doc.Find(`.product-spec-title:contains("Neck Wood")`).Next().Text()

        // 基本的にはNeck, Center, Bridgeを取得してフレーム側で組み立てる。分割して取得できなければ Pickups へ
        spec[C.Pickups]      = ""
        spec[C.NeckPickup]   = doc.Find(`.product-spec-title:contains("Neck Pickup")`).Next().Text()
        spec[C.CenterPickup] = doc.Find(`.product-spec-title:contains("Middle Pickup")`).Next().Text()
        spec[C.BridgePickup] = doc.Find(`.product-spec-title:contains("Bridge Pickup")`).Next().Text()

        foreignPrice, _      := doc.Find(`input[name="wc_braintree_paypal_amount"]`).Attr(`value`)
        foreignPrice          = _regPriceFractionBCRich.ReplaceAllString(foreignPrice, "") // 小数点以下を切り捨てる
        spec[C.Price]         = utils.CalcExchangedPrice(foreignPrice, _exchangeRateBCRich)
        spec[C.ScaleLengthMM] = doc.Find(`.product-spec-title:contains("scale")`).Next().Text()
        spec[C.Series]        = doc.Find(`.product-spec-title:contains("BODY SHAPE")`).Next().Text()

        src, _        := doc.Find(`div[data-thumb] img`).Attr(`src`)
        spec[C.Src]    = src
        spec[C.Weight] = strconv.Itoa(C.InvalidNumber)

        specs = utils.LockedAppend(mutex, specs, spec)

        return specs
    }
}

func (c *CallBacksBCRich) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url)
    }
}

func (c *CallBacksBCRich) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, `body`)
    }
}