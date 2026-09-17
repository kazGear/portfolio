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

type CrawlerTomAnderson struct {
    name     string
    gScraper Crawler[*model.Guitar]
}

type CallBacksTomAnderson struct {
    funcs CallBacks
}

func NewScraperTomAnderson() Scraper[*model.Guitar] {
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
    return &CrawlerTomAnderson{
        "TOM ANDERSON",
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksTomAnderson() *CallBacksTomAnderson {
    return &CallBacksTomAnderson{
        CallBacks{},
    }
}

func (g *CrawlerTomAnderson) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(c ,crawlStats)

    // URL収集、クロール
    visited := make(map[string]struct{}, 100)

    mutex := &sync.Mutex{}

    // 詳細ページ
    c.OnHTML(`ul.aColor_white a[href$=".html"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
        c.Visit(link)
    })

    c.Visit("https://jes1988.com/tomandersonguitar/products/")
    c.Wait()

    loggingCrawlStats(g.name, crawlStats)

    g.gScraper.urls = utils.MapToSliceUrl(visited)

    return g.gScraper.urls, nil
}

func (g *CrawlerTomAnderson) Scrape(provider  PageProvider,
                                    parser    ModelParser[*model.Guitar],
                                    parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *CallBacksTomAnderson) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
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

var _regTopAndBackTomAnderson = regexp.MustCompile(`(Top on| on )`)

func (c *CallBacksTomAnderson) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{} // 捨てる属性は基本的に空文字を割り当てる

        spec[C.Maker] = strconv.Itoa(C.TOM_ANDERSON)
        spec[C.Name]  = doc.Find(`#titleBox h1`).Text()
        spec[C.Color] = getExactMatchedDoc(doc, "dt", "FINISH").Next().Text()

        spec[C.BodyFinish] = ""

        // BODY WOOD: Quilt Maple Top on Mahogany | Alder といった構造
        topAndBack := _regTopAndBackTomAnderson.Split(getExactMatchedDoc(doc, "dt", "BODY WOOD").Next().Text(), -1)

        if len(topAndBack) == 1 {
            spec[C.BodyMaterialBack] = topAndBack[0]
        } else if len(topAndBack) == 2 {
            spec[C.BodyMaterialBack] = topAndBack[1]
        }

        if len(topAndBack) == 1 {
            spec[C.BodyMaterialTop]  = ""
        } else if len(topAndBack) == 2 {
            spec[C.BodyMaterialTop]  = topAndBack[0]
        }

        spec[C.Bridge]   = getExactMatchedDoc(doc, "dt", "BRIDGE").Next().Text()
        spec[C.Controls] = doc.Find(`dt:contains("SWITCHING")`).Next().Text()
        spec[C.Comment]  = doc.Find(`#titleBox h1`).Next().Text() +
                           strings.TrimSpace(doc.Find(`#feature section`).Text())

        // NECK WOOD: Maple | Mahogany, Rosewood Fingerboard といった構造
        neckAndFingerboard := strings.Split(doc.Find(`dt:contains("NECK WOOD")`).Next().Text(), ",")

        if len(neckAndFingerboard) == 1 {
            spec[C.Fingerboard] = neckAndFingerboard[0]
        } else if len(neckAndFingerboard) == 2 {
            spec[C.Fingerboard] = neckAndFingerboard[1]
        }

        spec[C.FretCount]    = ""
        spec[C.Inlays]       = ""
        spec[C.Joint]        = ""
        spec[C.NeckMaterial] = neckAndFingerboard[0]

        // 基本的にはNeck, Center, Bridgeを取得してフレーム側で組み立てる。分割して取得できなければ Pickups へ
        spec[C.Pickups]      = ""
        spec[C.NeckPickup]   = doc.Find(`dt:contains("NECK PICKUP")`).Next().Text()
        spec[C.CenterPickup] = doc.Find(`dt:contains("MIDDLE PICKUP")`).Next().Text()
        spec[C.BridgePickup] = doc.Find(`dt:contains("BRIDGE PICKUP")`).Next().Text()

        spec[C.Price]         = ""
        spec[C.ScaleLengthMM] = doc.Find(`dt:contains("SCALE LENGTH")`).Next().Text()
        spec[C.Series]        = strings.Split(spec[C.Name], " ")[0]

        src, _        := doc.Find(`.mainImage img`).Attr(`src`)
        spec[C.Src]    = "https://jes1988.com/tomandersonguitar/products/" + src
        spec[C.Weight] = strconv.Itoa(C.InvalidNumber)

        specs = utils.LockedAppend(mutex, specs, spec)

        return specs
    }
}

func (c *CallBacksTomAnderson) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url)
    }
}

func (c *CallBacksTomAnderson) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, `body`)
    }
}