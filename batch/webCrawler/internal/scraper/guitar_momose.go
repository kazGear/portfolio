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

func (g *guitarScraperMomose) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(c ,crawlStats, g.gScraper.logger)

    // URL収集、クロール
    visited := make(map[string]struct{}, 500)
    mutex   := &sync.Mutex{}

    // product >>> guitars, base, accessory
    c.OnHTML("ul.c-gnav__items #menu-item-39605 .sub-menu a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    // special model >>> limited, premium
    c.OnHTML("ul.c-gnav__items #menu-item-142109 .sub-menu a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    // ページネーション
    c.OnHTML("div.pagination a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    // 商品カード custom guitars
    c.OnHTML("div.p-product-list .p-product-list__item a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    // 商品カード limited, premium  guitars
    c.OnHTML("article div.p-product-list a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })

    c.Visit("https://www.deviser.co.jp/momose")
    c.Wait()

    loggingCrawlStats(crawlStats, g.gScraper.logger)

    g.gScraper.urls = utils.MapToSliceUrl(visited)
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
        if !isDetailPage(`^https://www.deviser.co.jp/products/.+`, url) {
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
               chromedp.WaitVisible("main", chromedp.ByQuery), // 求める要素が出るまで待つ
               chromedp.Sleep(300 * time.Millisecond), // JSが動く猶予を与える
               chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )
        if err != nil {
            return "", fmt.Errorf("[chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

var regSplitMomose = regexp.MustCompile(`(/| |’)`)

func (c *callBacksMomose) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        doc.Find(`.p-color-list__jan`).Remove() // colorからノイズを除去

        spec[C.Maker]     = strconv.Itoa(C.Momose)
        spec[C.Name]      = doc.Find("h1.p-product__title").Text()
        spec[C.Color]     = doc.Find(`h2:contains("COLOR")`).Next().Children().Children().Next().Text()
        spec[C.Comment]   = doc.Find(".wp-block-group__inner-container .wp-block-group__inner-container p").Text()
        spec[C.FretCount] = strconv.Itoa(C.InvalidNumber)
        spec[C.Inlays]    = ""
        spec[C.Joint]     = ""
        spec[C.Price]     = doc.Find(`.p-product__price strong`).Text()
        spec[C.Series]    = regSplitMomose.Split(spec[C.Name], 2)[0]
        src, _           := doc.Find(`.p-product__content div div div div img`).Attr(`src`)
        spec[C.Src]       = src
        spec[C.Weight]    = strconv.Itoa(C.InvalidNumber)

        doc.Find("div.p-spec div div div:nth-child(2) dl").Each(func(idx int, selector *goquery.Selection) {
            label      := selector.Find("dt").Text()
            elem       := selector.Find("dd").Text()
            field, _   := utils.ConvertLabel(label, specFieldMap)
            spec[field] = elem
        })

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

func (c *callBacksMomose) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksMomose) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "Finish")
    }
}