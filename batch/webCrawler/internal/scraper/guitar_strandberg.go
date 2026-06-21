package scraper

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	C "github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type guitarScraperStrandberg struct {
    gScraper guitarScraper
}

type callBacksStrandberg struct {
    funcs callBacks
}


func NewScraperStrandberg(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &guitarScraperStrandberg{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksStrandberg(logger *log.Logger) *callBacksStrandberg {
    return &callBacksStrandberg{
        callBacks{
            logger: logger,
        },
    }
}

func (g *guitarScraperStrandberg) CollectLinks(parentCtx context.Context) ([]string, error) {
    // タブごとに独立した context を作る
    tabCtx, tabCancel := chromedp.NewContext(parentCtx)
    defer tabCancel()
    // タブにだけ timeout を付ける
    ctx, cancel := context.WithTimeout(tabCtx, 20 * time.Second)
    defer cancel()

    targetLinks := []string{}

    // 詳細ページリンク収集
    doc, err := renderHTML(
        ctx,
        `https://strandbergguitars.com/en-US/guitars`,
        `div[data-sentry-component="ProductListingTypeTwo"]`,
    )
    if err != nil {
        return nil, errors.New(err.Error())
    }
    targetLinks = utils.CollectLinks(".product-card a", doc, 50)
    targetLinks = utils.GetNeedLinks(targetLinks, `/en-US/product/`, 50)
    targetLinks = utils.ToAbsLinks(targetLinks, `https://strandbergguitars.com`, 50)

    g.gScraper.urls = targetLinks
    return g.gScraper.urls, nil
}

func (g *guitarScraperStrandberg) Scrape(provider  PageProvider,
                                         parser    GuitarParser,
                                         parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

func (c *callBacksStrandberg) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`https://strandbergguitars.com/en-US/product/[a-z0-9\-]+`, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 20 * time.Second)
        defer cancel()

        var nodes []*cdp.Node

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            tryWaitVisible("body"), // 求める要素が出るまで待つ
            chromedp.Sleep(400 * time.Millisecond), // JSが動く猶予を与える
            tryWaitReady(`img[width="1200"][height="1200"]`), // 必要な要素が生成されるのを待つ
            tryWaitReady(`body div[data-sentry-component="PdpAccordion"]`),
            chromedp.Nodes(
                `//div[@data-sentry-component="PdpAccordion"]//div[@data-state="closed"]//button`,
                &nodes,
            ), // 全ｱｺｰﾃﾞｨｵﾝﾎﾞﾀﾝ取得
        )
        if err != nil {
            return "", fmt.Errorf("[Chromedp error]: %v [url]: %v\n", err, url)
        }
        // アコーディオンオープン、内部の要素を出現させる。排他的にしかオープンしないため、つどHTMLを抽出する。最後にマージ
        htmlParts := []*string{
            new(string), new(string), new(string), new(string), new(string), new(string), new(string),
        }
        for idx, node := range nodes {
            chromedp.Run(ctx,
                tryClick(node.FullXPath()),
                chromedp.Sleep(300*time.Millisecond),
                chromedp.OuterHTML(
                    `div[data-sentry-component="PdpAccordion"]`, htmlParts[idx], chromedp.ByQuery,
                ), // クリックの都度HTML抽出
            )
        }
        // HTMLを取得、マージ
        var html string
        err = chromedp.Run(ctx,
            chromedp.OuterHTML("html", &html, chromedp.ByQuery),
        )
        if err != nil {
            return "", fmt.Errorf("[Chromedp error]: %v [url]: %v\n", err, url)
        }
        for _, part := range htmlParts {
            html += *part
        }
        return html, nil
    }
}

// シリーズ名の抽出用
var regSeriesStrandberg = regexp.MustCompile(`^[A-Za-z]+\s[A-Za-z]+\b`)

var exchangeRate        = utils.GetExchangeUSDtoJPY()

func (c *callBacksStrandberg) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := []map[string]string{}
        mutex := &sync.Mutex{}

        spec    := map[string]string{}

        spec[C.Maker]            = strconv.Itoa(C.Strandberg)
        spec[C.Name]             = strings.TrimSpace(doc.Find(
                                    `div[data-sentry-component="ProductInfo"] div div h1`,
                                   ).Text())
        spec[C.Color]            = doc.Find(`h3:contains("Body finish color")`).Next().Text()
        spec[C.BodyFinish]       = doc.Find(`h3:contains("Body Finish Type")`).Next().Text()
        spec[C.BodyMaterialBack] = doc.Find(`h3:contains("Body Material")`).Next().Text()
        spec[C.BodyMaterialTop]  = doc.Find(`h3:contains("Body Top Material")`).Next().Text()
        spec[C.Bridge]           = doc.Find(`h3:contains("Bridge")`).Next().Text()
        spec[C.Controls]         = doc.Find(`h3:contains("Control Set")`).Next().Text()
        spec[C.Comment]          = ""
        spec[C.Fingerboard]      = doc.Find(`h3:contains("Fretboard Material")`).Next().Text()
        spec[C.FretCount]        = doc.Find(`h3:contains("Number of Frets")`).Next().Text()
        spec[C.Inlays]           = doc.Find(`h3:contains("Fretboard Inlays")`).Next().Text()
        spec[C.Joint]            = doc.Find(`h3:contains("Neck Construction")`).Next().Text()
        spec[C.NeckMaterial]     = doc.Find(`h3:contains("Neck Material")`).Next().Text()
        neckPickup              := doc.Find(`h3:contains("Neck pickup")`).Next().Text()
        bridgePickup            := doc.Find(`h3:contains("Bridge pickup")`).Next().Text()
        spec[C.Pickups]          = fmt.Sprintf("%v / %v", neckPickup, bridgePickup)
        priceStr                := strings.TrimSpace(doc.Find(`span:contains("Excluding vat")`).Prev().Text())
        spec[C.Price]            = utils.CalcExchangedPrice(priceStr, exchangeRate)
        spec[C.ScaleLengthMM]    = doc.Find(`h3:contains("Instrument Length Global")`).Next().Text()
        spec[C.Series]           = regSeriesStrandberg.FindString(spec[C.Name])
        proxyPath, _            := doc.Find(`img[width="1200"][height="1200"]`).Attr(`src`)
        spec[C.Src], _           = utils.ConvertRealUrl(proxyPath)
        spec[C.Weight]           = doc.Find(`h3:contains("Instrument Weight Global")`).Next().Text()

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksStrandberg) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksStrandberg) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // ありえない文字列、確実に動的ページを取得させる。
        return strings.Contains(html, "@abcd1234@")
    }
}