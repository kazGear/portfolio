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

type CrawlerKiesel struct {
    name     string
    gScraper Crawler[*model.Guitar]
}

type CallBacksKiesel struct {
    funcs CallBacks
}

func NewScraperKiesel() Scraper[*model.Guitar] {
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
    return &CrawlerKiesel{
        "KIESEL",
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksKiesel() *CallBacksKiesel {
    return &CallBacksKiesel{
        CallBacks{},
    }
}

func (g *CrawlerKiesel) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(c ,crawlStats)

    // URL収集、クロール
    visited := make(map[string]struct{}, 100)

    mutex := &sync.Mutex{}

    // 詳細ページ
    c.OnHTML(`a[href^="/series/guitar/"], a[href^="/series/bass/"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    c.Visit("https://www.kieselguitars.com/models/guitar")
    c.Visit("https://www.kieselguitars.com/models/bass")
    c.Wait()

    loggingCrawlStats(g.name, crawlStats)

    g.gScraper.urls = utils.MapToSliceUrl(visited)

    return g.gScraper.urls, nil
}

func (g *CrawlerKiesel) Scrape(provider  PageProvider,
                               parser    ModelParser[*model.Guitar],
                               parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *CallBacksKiesel) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        // 動的ページを取得しない場合、引数のパターンは記載しないで良い
        if !isDetailPage(`https://www.kieselguitars.com/series/(guitar|bass)/\w+`, url) {
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

var _exchangeRateKiesel = utils.GetExchangeUSDtoJPY()

func (c *CallBacksKiesel) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{} // 捨てる属性は基本的に空文字を割り当てる

        spec[C.Maker] = strconv.Itoa(C.KIESEL)
        spec[C.Name]  = doc.Find(`title`).Text()
        spec[C.Color] = "" // 記載なし

        // First(): 通常のセレクターだと大量の後続謎要素を引っ張ってきてしまうので、先頭だけ取得しノイズ除去

        spec[C.BodyFinish]       = ""
        spec[C.BodyMaterialBack] = doc.Find(`div[class*="SpecCategoryColumn"]:contains("Body")`).Next().First().Text()
        spec[C.BodyMaterialTop]  = "" // 記載なし

        spec[C.Bridge]   = doc.Find(`div[class*="SpecCategoryColumn"]:contains("Bridge")`).Next().First().Text()
        spec[C.Controls] = doc.Find(`div[class*="SpecCategoryColumn"]:contains("Electronics")`).Next().First().Text()
        spec[C.Comment]  = doc.Find(`div[class*="ProductDescription"]`).Text()

        spec[C.Fingerboard] = doc.Find(`div[class*="SpecCategoryColumn"]:contains("Fingerboard")`).Next().First().Text()
        spec[C.FretCount]    = doc.Find(`div[class*="SpecCategoryColumn"]:contains("Frets")`).Next().First().Text()
        spec[C.Inlays]       = doc.Find(`div[class*="SpecCategoryColumn"]:contains("Inlays")`).Next().First().Text()
        spec[C.Joint]        = doc.Find(`div[class*="SpecCategoryColumn"]:contains("Construction")`).Next().First().Text()
        spec[C.NeckMaterial] = doc.Find(`div[class*="SpecCategoryColumn"]:contains("Neck")`).Next().First().Text()

        // 基本的にはNeck, Center, Bridgeを取得してフレーム側で組み立てる。分割して取得できなければ Pickups へ
        spec[C.Pickups]      = doc.Find(`div[class*="SpecCategoryColumn"]:contains("Pickups")`).Next().First().Text()
        spec[C.NeckPickup]   = ""
        spec[C.CenterPickup] = ""
        spec[C.BridgePickup] = ""

        foreignPrice         := strings.TrimSpace(doc.Find(`#price`).Text())
        spec[C.Price]         = utils.CalcExchangedPrice(foreignPrice, _exchangeRateKiesel)
        spec[C.ScaleLengthMM] = doc.Find(`div[class*="SpecCategoryColumn"]:contains("Scale Length")`).
                                    Next().Children().First().Text()
        spec[C.Series]        = ""

        src, _        := doc.Find(`button[aria-label="Active product image"] img`).Attr(`src`)
        spec[C.Src]    = src

        weightLbsStr  := doc.Find(`div[class*="SpecCategoryColumn"]:contains("Weight")`).
                             Next().Children().First().Text()
        weightLbsStr   = strings.ReplaceAll(weightLbsStr, " lbs", "")
        weightLbs, _  := strconv.ParseFloat(weightLbsStr, 64)
        spec[C.Weight] = utils.PoundsToKilograms(weightLbs)

        specs = utils.LockedAppend(mutex, specs, spec)

        return specs
    }
}

func (c *CallBacksKiesel) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url)
    }
}

func (c *CallBacksKiesel) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, `id="price"`)
    }
}