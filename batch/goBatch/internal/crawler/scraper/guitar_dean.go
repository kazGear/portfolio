package scraper

import (
	"context"
	"fmt"
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

type CrawlerDean struct {
    name     string
    gScraper Crawler[*model.Guitar]
}

type CallBacksDean struct {
    funcs CallBacks
}

func NewScraperDean() Scraper[*model.Guitar] {
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
    return &CrawlerDean{
        "DEAN",
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksDean() *CallBacksDean {
    return &CallBacksDean{
        CallBacks{},
    }
}

func (g *CrawlerDean) CollectLinks(parentCtx context.Context) ([]string, error) {
    // タブごとに独立した context を作る
    tabCtx, tabCancel := chromedp.NewContext(parentCtx)
    defer tabCancel()

    // タブにだけ timeout を付ける
    ctx, cancel := context.WithTimeout(tabCtx, 300 * time.Second) // autoScroll があるので長めに確保
    defer cancel()

    html := ""

    err := chromedp.Run(ctx,
        chromedp.Navigate("https://deanguitars.com/collections/all-electric-guitars"), // ギター一覧ページ
        autoScroll(), // ページ下部の詳細リンクまで表示させる
        chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
    )

    if err != nil {
        return []string{}, fmt.Errorf("[Chromedp error from CollectLinks DEAN guitars]: %w", err)
    }

    doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

    if err != nil {
        return []string{}, fmt.Errorf("[NewDocumentFromReader error from CollectLinks DEAN guitars]: %w", err)
    }

    mutex := &sync.Mutex{}
    urls  := make([]string, 0, 120)

    doc.Find(`.product-grid-container a[href*="/products/"]`).Each(func(_ int, selector *goquery.Selection) {
        href, _ := selector.Attr(`href`)
        urls     = utils.LockedAppend(mutex, urls, "https://deanguitars.com" + href)
    })

    g.gScraper.urls = urls

    return g.gScraper.urls, nil
}

func (g *CrawlerDean) Scrape(provider  PageProvider,
                             parser    ModelParser[*model.Guitar],
                             parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *CallBacksDean) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
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

var _exchangeRateDean     = utils.GetExchangeUSDtoJPY()
var _regPriceFractionDean = regexp.MustCompile(`\.\d+`)

func (c *CallBacksDean) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{} // 捨てる属性は基本的に空文字を割り当てる

        spec[C.Maker] = strconv.Itoa(C.DEAN)
        spec[C.Name]  = doc.Find(`.b-main-title`).Text()
        spec[C.Color] = doc.Find(`.rte.text-left:contains("Top Color")`).Next().Text()

        spec[C.BodyFinish]       = doc.Find(`.rte.text-left:contains("Body Finish")`).Next().Text()
        spec[C.BodyMaterialBack] = doc.Find(`.rte.text-left:contains("Body Material")`).Next().Text()
        spec[C.BodyMaterialTop]  = doc.Find(`.rte.text-left:contains("Top Material")`).Next().Text()

        spec[C.Bridge]   = doc.Find(`.rte.text-left:contains("Bridge")`).Next().Text()
        spec[C.Controls] = doc.Find(`.rte.text-left:contains("Controls")`).Next().Text()
        spec[C.Comment]  = doc.Find(`.product__description`).Text()

        spec[C.Fingerboard]  = doc.Find(`.rte.text-left:contains("Fingerboard Material")`).Next().Text()
        spec[C.FretCount]    = doc.Find(`.rte.text-left:contains("Number of Frets")`).Next().Text()
        spec[C.Inlays]       = doc.Find(`.rte.text-left:contains("Position Inlays")`).Next().Text() + " / " +
                               doc.Find(`.rte.text-left:contains("Side Dots")`).Next().Text()
        spec[C.Joint]        = doc.Find(`.rte.text-left:contains("Neck Construction")`).Next().Text()
        spec[C.NeckMaterial] = doc.Find(`.rte.text-left:contains("Neck Material")`).Next().Text()

        // 基本的にはNeck, Center, Bridgeを取得してフレーム側で組み立てる。分割して取得できなければ Pickups へ
        spec[C.Pickups]      = ""
        spec[C.NeckPickup]   = doc.Find(`.rte.text-left:contains("Neck Pickup")`).Next().Text()
        spec[C.CenterPickup] = doc.Find(`.rte.text-left:contains("Middle Pickup")`).Next().Text()
        spec[C.BridgePickup] = doc.Find(`.rte.text-left:contains("Bridge Pickup")`).Next().Text()

        foreignPrice         := doc.Find(`span.price-item--regular`).Text()
        foreignPrice          = _regPriceFractionDean.ReplaceAllString(foreignPrice, "") // 小数点以下を切り捨てる
        spec[C.Price]         = utils.CalcExchangedPrice(foreignPrice, _exchangeRateDean)
        spec[C.ScaleLengthMM] = doc.Find(`.rte.text-left:contains("Scale Length")`).Next().Text()
        spec[C.Series]        = doc.Find(`.rte.text-left:contains("Series")`).Next().Text()

        src, _        := doc.Find(`.is-active img`).Attr(`src`)
        spec[C.Src]    = src
        spec[C.Weight] = strconv.Itoa(C.InvalidNumber)

        specs = utils.LockedAppend(mutex, specs, spec)

        return specs
    }
}

func (c *CallBacksDean) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url)
    }
}

func (c *CallBacksDean) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, `body`)
    }
}