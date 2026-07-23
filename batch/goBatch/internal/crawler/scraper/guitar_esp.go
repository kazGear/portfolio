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

type guitarScraperEsp struct {
    gScraper guitarScraper
}

type callBacksEsp struct {
    funcs callBacks
}


func NewScraperEsp(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(4),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &guitarScraperEsp{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksEsp(logger *log.Logger) *callBacksEsp {
    return &callBacksEsp{
        callBacks{
            logger: logger,
        },
    }
}


var regNeedPatterEsp = regexp.MustCompile(`https://espguitars.co.jp/product/\d+/`)

func (g *guitarScraperEsp) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(c ,crawlStats, g.gScraper.logger)

    // URL収集、クロール
    visited := make(map[string]struct{}, 500)
    mutex   := &sync.Mutex{}

    c.OnHTML("#item .figcap a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML("#inner_content .figcap a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML("section.color_variation a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.Visit("https://espguitars.co.jp/products/esp")
    c.Wait()

    loggingCrawlStats(crawlStats, g.gScraper.logger)

    g.gScraper.urls = utils.MapToSliceUrl(visited)
    g.gScraper.urls = utils.GetNeedLinks(g.gScraper.urls, regNeedPatterEsp, 400)
    return g.gScraper.urls, nil
}

func (g *guitarScraperEsp) Scrape(provider  PageProvider,
                                  parser    GuitarParser,
                                  parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksEsp) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://espguitars\.co\.jp/product/\d{4,}/?$`, url) {
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
               chromedp.WaitVisible("#main", chromedp.ByQuery), // 求める要素が出るまで待つ
               chromedp.Sleep(300 * time.Millisecond), // JSが動く猶予を与える
               tryWaitReady("h1.header_title"), // 必要な要素が生成されるのを待つ
               tryWaitReady(".tbl_spec"),
               tryWaitReady("p.detail_price"),
               chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )
        if err != nil {
            return "", fmt.Errorf("[chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

var regSeries = regexp.MustCompile(`^[\w-]+`)

func (c *callBacksEsp) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec[C.Maker]   = strconv.Itoa(C.Esp)
        spec[C.Name]    = doc.Find("h1.header_title").Text()
        spec[C.Color]   = doc.Find(".header_content h3.clr_name").Text()
        spec[C.Comment] = doc.Find("#specialfeatures .container_small p").Text()
        spec[C.Price]   = doc.Find("p.detail_price").Text()
        src, _         := doc.Find("#main .header_content img.transform-5").Attr("src")
        spec[C.Src]     = src
        spec[C.Series]  = regSeries.FindString(spec[C.Name])

        doc.Find("#specifications table.tbl_spec tr").Each(func(idx int, selector *goquery.Selection) {
            th      := selector.Find("th").Text()
            td      := selector.Find("td").Text()
            th, _    = utils.ConvertLabel(th, specFieldMap)
            spec[th] = td
        })

        materials    := strings.Split(spec[C.BodyMaterialBack], ",")

        if len(materials) == 1 {
            spec[C.BodyMaterialBack] = materials[0]
        } else if len(materials) == 2 {
            spec[C.BodyMaterialTop]  = materials[0]
            spec[C.BodyMaterialBack] = materials[1]
        } else {
            spec[C.BodyMaterialBack] = materials[0]
        }

        pickups := strings.Split(spec[C.Pickups], ",")

        if len(pickups) <= 1 {
            spec[C.BridgePickup] = pickups[0]
        } else if len(pickups) == 2 {
            spec[C.NeckPickup]   = pickups[0]
            spec[C.BridgePickup] = pickups[1]
        } else {
            spec[C.NeckPickup]   = pickups[0]
            spec[C.CenterPickup] = pickups[1]
            spec[C.BridgePickup] = pickups[2]
        }

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksEsp) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksEsp) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "tbl_spec")
    }
}