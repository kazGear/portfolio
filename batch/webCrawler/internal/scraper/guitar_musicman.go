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

type guitarScraperMusicMan struct {
    gScraper guitarScraper
}

type callBacksMusicMan struct {
    funcs callBacks
}

func NewScraperMusicMan(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &guitarScraperMusicMan{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksMusicMan(logger *log.Logger) *callBacksMusicMan {
    return &callBacksMusicMan{
        callBacks{
            logger: logger,
        },
    }
}

var needPatterMusicMan =
    `https://www.zemaitis-guitars.jp/(metal|disc|pearl|superior|z|jumbo|small|large|jumbo|grand|orchestra|new).*`
var regNeedPatterMusicMan = regexp.MustCompile(needPatterMusicMan)

func (g *guitarScraperMusicMan) CollectLinks(parentCtx context.Context) ([]string, error) {
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
    g.gScraper.urls = utils.GetNeedLinks(g.gScraper.urls, regNeedPatterMusicMan, 110)
    return g.gScraper.urls, nil
}

func (g *guitarScraperMusicMan) Scrape(provider  PageProvider,
                                       parser    GuitarParser,
                                       parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksMusicMan) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(needPatterMusicMan, url) {
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

func (c *callBacksMusicMan) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec[C.Maker]            = strconv.Itoa(C.MusicMan)
        spec[C.Name]             = strings.TrimSpace(doc.Find(
                                    `div[data-sentry-component="ProductInfo"] div div h1`,
                                   ).Text())
        spec[C.Color]            = doc.Find(`h3:contains("Body finish color")`).Text()
        spec[C.BodyFinish]       = doc.Find(`h3:contains("Body Finish Type")`).Text()
        spec[C.BodyMaterialBack] = doc.Find(`h3:contains("Body Material")`).Text()
        spec[C.BodyMaterialTop]  = doc.Find(`h3:contains("Body Top Material")`).Text()
        spec[C.Bridge]           = doc.Find(`h3:contains("Bridge")`).Text()
        spec[C.Controls]         = doc.Find(`h3:contains("Control Set")`).Text()
        spec[C.Comment]          = ""
        spec[C.Fingerboard]      = doc.Find(`h3:contains("Fretboard Material")`).Text()
        spec[C.FretCount]        = doc.Find(`h3:contains("Number of Frets")`).Text()
        spec[C.Inlays]           = doc.Find(`h3:contains("Fretboard Inlays")`).Text()
        spec[C.Joint]            = doc.Find(`h3:contains("Neck Construction")`).Text()
        spec[C.NeckMaterial]     = doc.Find(`h3:contains("Neck Material")`).Text()
        neckPickup              := doc.Find(`h3:contains("Neck pickup")`).Text()
        bridgePickup            := doc.Find(`h3:contains("Bridge pickup")`).Text()
        spec[C.Pickups]          = fmt.Sprintf("%v / %v", neckPickup, bridgePickup)
        priceStr                := doc.Find(`span:contains("Excluding vat")`).Text()
        spec[C.Price]            = utils.CalcExchangedPrice(priceStr, exchangeRate)
        spec[C.ScaleLengthMM]    = doc.Find(`h3:contains("Instrument Length Global")`).Text()
        spec[C.Series]           = regSeriesStrandberg.FindString(spec[C.Name])
        proxyPath, _            := doc.Find(`img[width="1200"][height="1200"]`).Attr(`src`)
        spec[C.Src], _           = utils.ConvertRealUrl(proxyPath)
        spec[C.Weight]           = doc.Find(`h3:contains("Instrument Weight Global")`).Text()

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksMusicMan) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksMusicMan) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "detail_right")
    }
}