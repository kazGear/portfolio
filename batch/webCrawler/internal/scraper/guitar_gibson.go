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

type guitarScraperGibson struct {
    gScraper guitarScraper
}

type callBacksGibson struct {
    funcs callBacks
}


func NewScraperGibson() Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5,
	})
    return &guitarScraperGibson{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksGibson() GuitarCallbacks {
    return &callBacksGibson{
        callBacks{},
    }
}

func (e *guitarScraperGibson) CollectLinks(parentCtx context.Context) []string {
    // タブごとに独立した context を作る
    tabCtx, tabCancel := chromedp.NewContext(parentCtx)
    defer tabCancel()
    // タブにだけ timeout を付ける
    ctx, cancel := context.WithTimeout(tabCtx, 20 * time.Second)
    defer cancel()

    targetLinks := []string{}

    // シリーズ一覧リンク収集
    doc := renderHTML(
        ctx,
        // `https://jp.fender.com/collections/electric-guitars`
        `https://jp.fender.com/collections/electric-guitars`,
        // `div[data-component-id*="section"] div[role="region"] .swiper-container`,
        `div.footer-title`,
    )
    targetLinks = collectLinks(`a[class*="Title"]`, doc, 10)
    utils.LogCollectedLinks(targetLinks)
    // targetLinks = getNeedLinks(targetLinks, `/en-US/product/`, 50)
    // targetLinks = toAbsLinks(targetLinks, `https://strandbergguitars.com`, 50)

    e.gScraper.urls = targetLinks
    return e.gScraper.urls
}

func (e *guitarScraperGibson) Scrape(funcs GuitarCallbacks,
                                     parentCtx context.Context,
) ([]*model.Guitar, error) {
    guitars, _ := e.gScraper.scrapeFrame(funcs, parentCtx)
    return guitars, nil
}

func (e *callBacksGibson) FetchDynamicPage(parentCtx context.Context) func(url string) string {
    return func(url string) string {
        if !isDetailPage(`https://strandbergguitars.com/en-US/product/[a-z0-9\-]+`, url) {
            return ""
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

func (e *callBacksGibson) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := []map[string]string{}
        mutex := &sync.Mutex{}

        spec    := map[string]string{}
        getElem := utils.GetElemNextToLabel(doc)

        spec["Maker"]            = strconv.Itoa(constants.Gibson)
        spec["Name"]             = strings.TrimSpace(doc.Find(
                                    `div[data-sentry-component="ProductInfo"] div div h1`,
                                   ).Text())
        spec["Color"]            = getElem(`h3:contains("Body finish color")`)
        spec["BodyFinish"]       = getElem(`h3:contains("Body Finish Type")`)
        spec["BodyMaterialBack"] = getElem(`h3:contains("Body Material")`)
        spec["BodyMaterialTop"]  = getElem(`h3:contains("Body Top Material")`)
        spec["BodyMaterial"]     = spec["BodyMaterialTop"] + " " + spec["BodyMaterialBack"]
        spec["Bridge"]           = getElem(`h3:contains("Bridge")`)
        spec["Controls"]         = getElem(`h3:contains("Control Set")`)
        spec["Comment"]          = strings.TrimSpace(doc.Find(``).Text())
        spec["Fingerboard"]      = getElem(`h3:contains("Fretboard Material")`)
        spec["FretCount"]        = getElem(`h3:contains("Number of Frets")`)
        spec["Inlays"]           = getElem(`h3:contains("Fretboard Inlays")`)
        spec["Joint"]            = getElem(`h3:contains("Neck Construction")`)
        spec["NeckMaterial"]     = getElem(`h3:contains("Neck Material")`)
        neckPickup              := getElem(`h3:contains("Neck pickup")`)
        bridgePickup            := getElem(`h3:contains("Bridge pickup")`)
        spec["Pickups"]          = fmt.Sprintf(constants.PickupsFormat, neckPickup, bridgePickup)
        spec["Price"]            = strings.TrimSpace(doc.Find(`span:contains("Excluding vat")`).Prev().Text())
        spec["ScaleLengthMM"]    = getElem(`h3:contains("Instrument Length Global")`)
        spec["Series"]           = regSeries.FindString(spec["Name"])

        // 画像保存、保存場所の記録
        proxyUrl, _             := doc.Find(`img[width="1200"][height="1200"]`).Attr(`src`)
        realUrl                 := utils.ConvertRealUrl(proxyUrl)
        savePath                := utils.CreateImageSavePath("images/strandberg", realUrl)
        utils.DownloadImage(realUrl, savePath)

        spec["Src"]              = strings.TrimSpace(savePath)
        spec["Weight"]           = getElem(`h3:contains("Instrument Weight Global")`)

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (e *callBacksGibson) BuildGuitar() func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec)
    }
}

func (e *callBacksGibson) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // ありえない文字列、確実に動的ページを取得させる。
        return strings.Contains(html, "@abcd1234@")
    }
}