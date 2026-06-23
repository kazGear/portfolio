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
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	C "github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type guitarScraperMomose struct {
    gScraper guitarScraper
}

type callBacksMomose struct {
    funcs callBacks
}


func NewScraperMomose(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(4),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &guitarScraperMomose{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksMomose(logger *log.Logger) *callBacksMomose {
    return &callBacksMomose{
        callBacks{
            logger: logger,
        },
    }
}


var regNeedPatterMomose = regexp.MustCompile(`https://Momoseguitars.co.jp/product/\d+/`)

func (g *guitarScraperMomose) CollectLinks(parentCtx context.Context) ([]string, error) {
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
    c.Visit("https://Momoseguitars.co.jp/products/Momose")
    c.Wait()

    loggingCrawlStats(crawlStats, g.gScraper.logger)


    g.gScraper.urls = utils.MapToSliceUrl(visited)
    g.gScraper.urls = utils.GetNeedLinks(g.gScraper.urls, regNeedPatterMomose, 400)
    return g.gScraper.urls, nil
}

func (g *guitarScraperMomose) Scrape(provider  PageProvider,
                                  parser    GuitarParser,
                                  parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksMomose) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://Momoseguitars\.co\.jp/product/\d{4,}/?$`, url) {
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

func (c *callBacksMomose) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec[C.Maker]   = strconv.Itoa(C.Momose)
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

        bodyMaterial := spec[C.BodyMaterialBack]
        materials    := strings.Split(bodyMaterial, ",")

        if len(materials) == 1 {
            spec[C.BodyMaterialBack] = materials[0]
        } else if len(materials) == 2 {
            spec[C.BodyMaterialTop]  = materials[0]
            spec[C.BodyMaterialBack] = materials[1]
        } else {
            spec[C.BodyMaterialBack] = materials[0]
        }

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksMomose) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksMomose) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "tbl_spec")
    }
}