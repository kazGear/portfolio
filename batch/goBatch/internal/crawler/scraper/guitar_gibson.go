package scraper

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

type CrawlerGibson struct {
    gScraper Crawler[*model.Guitar]
}

type CallBacksGibson struct {
    funcs CallBacks
}


func NewScraperGibson(logger *log.Logger) Scraper[*model.Guitar] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &CrawlerGibson{
        Crawler[*model.Guitar]{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksGibson(logger *log.Logger) *CallBacksGibson {
    return &CallBacksGibson{
        CallBacks{
            logger: logger,
        },
    }
}

var regNeedPatterGibson = regexp.MustCompile(`https://gibson.jp/(electric|acoustic)/`)

func (g *CrawlerGibson) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(c ,crawlStats, g.gScraper.logger)

    // URL収集、クロール
    visited := make(map[string]struct{}, 600)
    mutex   := &sync.Mutex{}

    c.OnHTML(".body-types a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML(".category-wrapper .model-card a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.Visit("https://gibson.jp/")
    c.Wait()

    loggingCrawlStats(crawlStats, g.gScraper.logger)

    g.gScraper.urls = utils.MapToSliceUrl(visited)
    g.gScraper.urls = utils.GetNeedLinks(g.gScraper.urls, regNeedPatterGibson, 490)
    return g.gScraper.urls, nil
}

func (g *CrawlerGibson) Scrape(provider  PageProvider,
                               parser    ModelParser[*model.Guitar],
                               parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    utils.AutoDownLoader(guitars, "images/gibson")
    return guitars
}

func (c *CallBacksGibson) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`https://gibson.jp/(electric|acoustic)/[a-z0-9\-]+`, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 20 * time.Second)
        defer cancel()

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            tryWaitVisible("#body-wrap"), // 求める要素が出るまで待つ
            chromedp.Sleep(200 * time.Millisecond), // JSが動く猶予を与える
        )
        if err != nil {
            return "", fmt.Errorf("[Chromedp error]: %v [url]: %v\n", err, url)
        }
        // HTMLを取得、マージ
        var html string
        err = chromedp.Run(ctx,
            chromedp.OuterHTML("html", &html, chromedp.ByQuery),
        )
        if err != nil {
            return "", fmt.Errorf("[Chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

// シリーズ名の抽出用
var regSeriesGibson = regexp.MustCompile(
    `(Les Paul|SG|ES-\d+|Flying V|Explorer|Firebird|Hummingbird|J\-\d+)+\s[A-Za-z]+\b`,
)

func (c *CallBacksGibson) CollectAttributes() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := []map[string]string{}
        mutex := &sync.Mutex{}

        spec        := map[string]string{}

        doc.Find(`#cart-options h2.marketing-headline small`).Remove() // Nameからノイズを除去

        spec[C.Maker]        = strconv.Itoa(C.Gibson)
        spec[C.Name]         = doc.Find(`h2.marketing-headline`).Text()
        spec[C.Color]        = doc.Find(`div#displayed-finish`).Text()
        spec[C.Comment]      = doc.Find(`#cart-options .marketing-copy p`).Text()
        spec[C.NeckPickup]   = doc.Find(`.spec-item div:contains("Neck pickup")`).Next().Text()
        spec[C.CenterPickup] = doc.Find(`.spec-item div:contains("Middle Pickup")`).Next().Text()
        spec[C.BridgePickup] = doc.Find(`.spec-item div:contains("Bridge pickup")`).Next().Text()
        src, _              := doc.Find(`img#gallery-front`).Attr(`src`)
        spec[C.Src]          = src
        spec[C.Series]       = regSeriesGibson.FindString(spec[C.Name])
        spec[C.Weight]       = strconv.Itoa(C.InvalidNumber)

        doc.Find(`#product-overview .spec-section .spec-item`).Each(func(idx int, selector *goquery.Selection) {
            label      := selector.Find(`div:nth-child(1)`).Text()
            elem       := selector.Find(`div:nth-child(2)`).Text()
            field, _   := utils.ConvertLabel(label, specFieldMap)
            spec[field] = elem
        })
        spec[C.Price] = ""

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *CallBacksGibson) BuildModel(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *CallBacksGibson) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "product-overview")
    }
}
