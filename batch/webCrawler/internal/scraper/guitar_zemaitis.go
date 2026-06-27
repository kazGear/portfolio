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

type guitarScraperZemaitis struct {
    gScraper guitarScraper
}

type callBacksZemaitis struct {
    funcs callBacks
}

func NewScraperZemaitis(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &guitarScraperZemaitis{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksZemaitis(logger *log.Logger) *callBacksZemaitis {
    return &callBacksZemaitis{
        callBacks{
            logger: logger,
        },
    }
}

var needPatterZemaitis =
    `https://www.zemaitis-guitars.jp/(metal|disc|pearl|superior|z|jumbo|small|large|jumbo|grand|orchestra|new).*`
var regNeedPatterZemaitis = regexp.MustCompile(needPatterZemaitis)

func (g *guitarScraperZemaitis) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(c ,crawlStats, g.gScraper.logger)

    // URL収集、クロール
    visited := make(map[string]struct{}, 130)
    mutex   := &sync.Mutex{}

    c.OnHTML(".series ul li:nth-child(2) a, .series ul li:nth-child(3) a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML("#mainList ul li a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })

    c.Visit("https://www.zemaitis-guitars.jp/")
    c.Wait()

    loggingCrawlStats(crawlStats, g.gScraper.logger)

    g.gScraper.urls = utils.MapToSliceUrl(visited)
    g.gScraper.urls = utils.GetNeedLinks(g.gScraper.urls, regNeedPatterZemaitis, 110)
    return g.gScraper.urls, nil
}

func (g *guitarScraperZemaitis) Scrape(provider  PageProvider,
                                       parser    GuitarParser,
                                       parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksZemaitis) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(needPatterZemaitis, url) {
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
            return "", fmt.Errorf("[chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

func (c *callBacksZemaitis) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec[C.Maker]      = strconv.Itoa(C.ZEMAITIS)
        spec[C.Name]       = doc.Find(`p.category`).Next().Text()

        if strings.Contains(spec[C.Name], "モデル") {
            return []map[string]string{}
        }

        spec[C.BodyFinish] = doc.Find(`h3:contains("Body Finish Type")`).Text()
        spec[C.Bridge]     = ""
        spec[C.Comment]    = doc.Find(`.inner_left .text`).Text()
        spec[C.Price]      = doc.Find(`.inner_left .price`).Text()
        spec[C.Series]     = doc.Find(`.inner_left .category`).Text()
        src, _            := doc.Find(`.detail_left a:nth-child(1)`).Attr(`href`)
        spec[C.Src]        = src
        spec[C.Weight]     = strconv.Itoa(C.InvalidNumber)

        detail  := doc.Find(`.detail_right dl`).Text()
        details := strings.Split(detail, "\n")

        // details >>> [split]: [       Body Top, Sandblasted Ash, Body Back, Mahogany, Neck, Mahogany ...
        for idx, elem := range details {
            elem          = strings.TrimSpace(elem)
            field, exist := utils.ConvertLabel(elem, specFieldMap)

            if exist {
                spec[field] = details[idx + 1] // ラベルの次の要素が内容
            }
        }

        if len(spec[C.Color]) <= 0 {
            spec[C.Color] = "undefined" // アコギは色の定義がない
        }
        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksZemaitis) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksZemaitis) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "detail_right")
    }
}