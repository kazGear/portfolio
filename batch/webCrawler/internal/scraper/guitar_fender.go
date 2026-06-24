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
    // タブごとに独立した context を作る
    tabCtx, tabCancel := chromedp.NewContext(parentCtx)
    defer tabCancel()
    // タブにだけ timeout を付ける
    ctx, cancel := context.WithTimeout(tabCtx, 300 * time.Second)
    defer cancel()

    startLinks := []string{
        `/collections/electric-guitars`,
        `/collections/electric-basses`,
        `/collections/artist`,
        `/collections/acoustic-guitars`,
    }

    // var prevHtml = ""
    var html     = ""

    // for _, link := range startLinks {
log.Println("navigate start")
        err := chromedp.Run(ctx,
            chromedp.Navigate(fmt.Sprintf("https://jp.fender.com%v", startLinks[0])),
        )
log.Println("navigate end", err)
log.Println("wait start")
        err = chromedp.Run(ctx,
            autoScroll(),
            chromedp.Sleep(10 * time.Second),
            // chromedp.WaitReady(`.SearchspringPLP__ProductsContainer-sc-15dyhu-3`),
            chromedp.WaitReady(`#MainContent #react-searchspring-plp`),
            // chromedp.WaitVisible(`#MainContent`),
            // chromedp.WaitVisible(`body`),
            // tryClick(`.SearchspringPLP__ListFooter-sc-15dyhu-4 button`),
            // chromedp.OuterHTML(`#MainContent #react-searchspring-plp`, &html, chromedp.ByQuery),
            chromedp.OuterHTML(`html`, &html, chromedp.ByQuery),
        )
log.Println("wait end", err)
g.gScraper.logger.Printf("[chromedp error]: %v\n", err)
        if err != nil {
            return []string{}, fmt.Errorf("[Chromedp error]: %w\n", err)
        }
g.gScraper.logger.Printf("[html]: len %v\n%v\n", len(html), html)
    // }

// utils.LoggingCollectedLinks(g.gScraper.urls, g.gScraper.logger)
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
        spec[C.Name]    = doc.Find(`h1.product-title`).Text()
        spec[C.Comment] = doc.Find(`.product-content-description__text p`).Text()
        spec[C.Joint]   = ""
        spec[C.Price]   = doc.Find(`.price__regular .price-item`).Text()
        src, _         := doc.Find(`#fender-react div div img`).Attr(`src`)
        spec[C.Src]     = src
        spec[C.Weight]  = strconv.Itoa(C.InvalidNumber)

        doc.Find(`.product-specs-wrap .product-specs-lists .product-specs-props`).Each(func(idx int, selector *goquery.Selection) {
            label      := selector.Find(`h4`).Text()
            elem       := selector.Find(`p`).Text()
            field, _   := utils.ConvertLabel(label, specFieldMap)
            spec[field] = elem
        })

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
        return strings.Contains(html, "product-specs-wrap")
    }
}