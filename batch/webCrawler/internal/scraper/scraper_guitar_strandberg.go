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
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type guitarScraperStrandberg struct {
    gScraper guitarScraper
}

type callBacksStrandberg struct {
    funcs callBacks
}


func NewScraperStrandberg() Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5,
	})
    return &guitarScraperStrandberg{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksStrandberg() GuitarCallbacks {
    return &callBacksStrandberg{
        callBacks{},
    }
}

func (e *guitarScraperStrandberg) CollectLinks() *[]string {
    c       := e.gScraper.collector
    visited := make(map[string]struct{}, 50)
    mutex   := &sync.Mutex{}

    // ページネーション用
    c.OnHTML("nav ul li a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML("#main-plp-block div div div div .product-card .relative a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.Visit("https://strandbergguitars.com/en-US/guitars")
    c.Wait()

    e.gScraper.urls = getDistinctUrls(visited)
    return &e.gScraper.urls
}

func (e *guitarScraperStrandberg) Scrape(funcs GuitarCallbacks, ctx context.Context) (*[]model.Guitar, error) {
    guitars, _ := e.gScraper.scrapeFrame(funcs, ctx)
    return guitars, nil
}

func (e *callBacksStrandberg) FetchDynamicPage(parentCtx context.Context) func(url string) string {
    return func(url string) string {
        if !isDetailPage(`https://strandbergguitars.com/en-US/product/[a-z0-9\-]+`, url) {
            return ""
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 12 * time.Second)
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
            log.Printf("[Chromedp error]: %v [url]: %v\n", err, url)
            return ""
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
            log.Printf("[Chromedp error]: %v [url]: %v\n", err, url)
            return ""
        }
        for _, part := range htmlParts {
            html += *part
        }
        return html
    }
}

func (e *callBacksStrandberg) CollectSpec() func(doc *goquery.Document) *[]map[string]string {
    return func(doc *goquery.Document) *[]map[string]string {
        specs := []map[string]string{}
        mutex := &sync.Mutex{}

        spec    := map[string]string{}
        getElem := utils.GetElemNextToLabel(doc)

        spec["Maker"]   = strconv.Itoa(constants.Strandberg)
        spec["Name"]    = strings.TrimSpace(doc.Find(`div[data-sentry-component="ProductInfo"] div div h1`).Text())
        spec["Color"]   = getElem(`h3:contains("Body finish color")`)
        spec["BodyFinish"] = getElem(`h3:contains("Body Finish Type")`)
        spec["BodyMaterialBack"] = getElem(`h3:contains("Body Material")`)
        spec["BodyMaterialFront"] = getElem(`h3:contains("Body Top Material")`)
        spec["BodyMaterial"] = spec["BodyMaterialFront"] + " " + spec["BodyMaterialBack"]
        spec["Bridge"] = getElem(`h3:contains("Bridge")`)
        spec["Controls"] = getElem(`h3:contains("Control Set")`)
        spec["Comment"] = strings.TrimSpace(doc.Find(``).Text())
        spec["Fingerboard"] = getElem(`h3:contains("Fretboard Material")`)
        spec["FretCount"] = getElem(`h3:contains("Number of Frets")`)
        spec["Inlays"] = getElem(`h3:contains("Fretboard Inlays")`)
        spec["Joint"] = getElem(`h3:contains("Neck Construction")`)
        spec["NeckMaterial"] = getElem(`h3:contains("Neck Material")`)
        neckPickup := getElem(`h3:contains("Neck pickup")`)
        bridgePickup := getElem(`h3:contains("Bridge pickup")`)
        spec["Pickups"] = fmt.Sprintf(constants.PickupsFormat, neckPickup, bridgePickup)
        // TODO $ 1 149 形式でもできるよう改良する util
        spec["Price"]   = strings.TrimSpace(doc.Find(`span:contains("Excluding vat")`).Prev().Text())
        spec["ScaleLengthMM"] = getElem(`h3:contains("Instrument Length Global")`)
        spec["Series"] = getElem(`h3:contains("Body Shape")`)
        // TODO srcはダウンロード方式を作成
        src, _         := doc.Find(`img[title="English"]`).Attr(`src`)
        spec["Src"]     = strings.TrimSpace(src)
        // TODO Kg単位の数値を抜き出す処理追加 util
        spec["Weight"] = getElem(`h3:contains("Instrument Weight Global")`)

        specs = utils.LockedAppend(mutex, specs, spec)
        return &specs
    }
}

func (e *callBacksStrandberg) BuildGuitar() func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec)
    }
}

func (e *callBacksStrandberg) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // ありえない文字列、確実に動的ページを取得させる。
        return strings.Contains(html, "@abcd1234@")
    }
}