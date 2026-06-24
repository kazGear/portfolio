package scraper

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	C "github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type guitarScraperFender struct {
    gScraper guitarScraper
}

type callBacksFender struct {
    funcs callBacks
}

func NewScraperFender(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &guitarScraperFender{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksFender(logger *log.Logger) *callBacksFender {
    return &callBacksFender{
        callBacks{
            logger: logger,
        },
    }
}

func (g *guitarScraperFender) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(c ,crawlStats, g.gScraper.logger)

    // URL収集、クロール
    visited := make(map[string]struct{}, 130)
    mutex   := &sync.Mutex{}

    c.OnHTML("ol.product-items li a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })

    c.Visit("https://shop.music-man.com/instruments.html")
    c.Wait()

    loggingCrawlStats(crawlStats, g.gScraper.logger)

    g.gScraper.urls = utils.MapToSliceUrl(visited)
    return g.gScraper.urls, nil
}

func (g *guitarScraperFender) Scrape(provider  PageProvider,
                                       parser    GuitarParser,
                                       parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksFender) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://shop.music-man.com/.+`, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 15*time.Second)
        defer cancel()

        var html string

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.WaitVisible("body", chromedp.ByQuery), // 求める要素が出るまで待つ
            chromedp.Sleep(400 * time.Millisecond), // JSが動く猶予を与える
            chromedp.WaitReady(`div.fotorama__stage__shaft img[src]`),
            chromedp.WaitReady(`.additional-attributes-wrapper`),
            chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )
        if err != nil {
            return "", fmt.Errorf("[chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

func (c *callBacksFender) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec[C.Maker]   = strconv.Itoa(C.Fender)
        spec[C.Name]    = doc.Find(`h1.page-title span`).Text()
        spec[C.Comment] = doc.Find(`div.product.attribute.overview div p`).Text()
        spec[C.Joint]   = doc.Find(`h3:contains("Neck Construction")`).Text()
        spec[C.Price]   = doc.Find(`div[data-role="priceBox"]`).Text()
        spec[C.Series]  = ""
        src, _         := doc.Find(`div.fotorama__stage__shaft img`).Attr(`src`)
        spec[C.Src]     = src
        spec[C.Weight]  = strconv.Itoa(C.InvalidNumber)

        doc.Find(`table#product-attribute-specs-table tbody tr`).Each(func(idx int, selector *goquery.Selection) {
            label      := selector.Find(`th`).Text()
            elem       := selector.Find(`td`).Text()
            field, _   := utils.ConvertLabel(label, specFieldMap)
            spec[field] = elem
        })

        // 木材 整形
        if strings.Contains(spec[C.BodyMaterialBack], "and") {
            topAndBack := strings.Split(spec[C.BodyMaterialBack], "and")
            spec[C.BodyMaterialTop]  = topAndBack[0]
            spec[C.BodyMaterialBack] = topAndBack[1]
        } else if strings.Contains(spec[C.BodyMaterialBack], "with") {
            backAndTop := strings.Split(spec[C.BodyMaterialBack], "with")
            spec[C.BodyMaterialBack] = backAndTop[0]
            spec[C.BodyMaterialTop]  = backAndTop[1]
        }

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksFender) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksFender) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "fotorama__stage__shaft")
    }
}