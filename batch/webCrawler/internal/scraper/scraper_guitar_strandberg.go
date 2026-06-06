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
		Parallelism: 10,
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

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (e *callBacksStrandberg) FetchDynamicPage(parentCtx context.Context) func(url string) string {
    return func(url string) string {
        if !isDetailPage(`^https://strandbergguitars.com/en-US/product/[a-z0-9\-]+`, url) {
            return ""
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
               chromedp.Sleep(200 * time.Millisecond), // JSが動く猶予を与える
               tryWaitReady(`img[width="1200"][height="1200"]`), // 必要な要素が生成されるのを待つ
               tryWaitReady(`body div[data-sentry-component="PdpAccordion"]`),
               chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )
        if err != nil {
            log.Printf("[chromedp error]: %v [url]: %v\n", err, url)
            return ""
        }
        return html
    }
}

func (e *callBacksStrandberg) CollectSpec() func(doc *goquery.Document) *[]map[string]string {
    return func(doc *goquery.Document) *[]map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec  := map[string]string{}

        spec["Maker"]   = strconv.Itoa(constants.Strandberg)
        spec["Name"]    = strings.TrimSpace(doc.Find(`div[data-sentry-component="ProductInfo"] div div h1`).Text())
        spec["Color"]   = "tmpColor"//strings.TrimSpace(doc.Find(`h3:contains("Body finish color")`).Next().Text())
        spec["BodyFinish"] = strings.TrimSpace(doc.Find(`h3:contains("Body Finish Type")`).Next().Text())
        spec["BodyMaterialBack"] = strings.TrimSpace(doc.Find(`h3:contains("Body Material")`).Next().Text())
        spec["BodyMaterialFront"] = strings.TrimSpace(doc.Find(`h3:contains("Body Top Material")`).Next().Text())
        spec["BodyMaterial"] = spec["BodyMaterialFront"] + " " + spec["BodyMaterialBack"]
        spec["Bridge"] = strings.TrimSpace(doc.Find(`h3:contains("Bridge")`).Next().Text())
        spec["Controls"] = strings.TrimSpace(doc.Find(`h3:contains("Control Set")`).Next().Text())
        spec["Comment"] = strings.TrimSpace(doc.Find(``).Text())
        spec["Fingerboard"] = strings.TrimSpace(doc.Find(`h3:contains("Fretboard Material")`).Next().Text())
        spec["FretCount"] = strings.TrimSpace(doc.Find(`h3:contains("Number of Frets")`).Next().Text())
        spec["Inlays"] = strings.TrimSpace(doc.Find(`h3:contains("Fretboard Inlays")`).Next().Text())
        spec["Joint"] = strings.TrimSpace(doc.Find(`h3:contains("Neck Construction")`).Next().Text())
        spec["NeckMaterial"] = strings.TrimSpace(doc.Find(`h3:contains("Neck Material")`).Next().Text())
        neckPickup := strings.TrimSpace(doc.Find(`h3:contains("Neck pickup")`).Next().Text())
        bridgePickup := strings.TrimSpace(doc.Find(`h3:contains("Bridge pickup")`).Next().Text())
        spec["Pickups"] = fmt.Sprintf(`(Neck) %v (Bridge) %v`, neckPickup, bridgePickup)
        // TODO $ 1 149 形式でもできるよう改良する util
        spec["Price"]   = strings.TrimSpace(doc.Find(`span:contains("Excluding vat")`).Prev().Text())
        spec["ScaleLengthMM"] = strings.TrimSpace(doc.Find(`h3:contains("Instrument Length Global")`).Next().Text())
        spec["Series"] = strings.TrimSpace(doc.Find(`h3:contains("Body Shape")`).Next().Text())
        // TODO srcはダウンロード方式を作成
        src, _         := doc.Find(`img[title="English"]`).Attr(`src`)
        spec["Src"]     = strings.TrimSpace(src)
        // TODO Kg単位の数値を抜き出す処理追加 util
        spec["Weight"] = strings.TrimSpace(doc.Find(`h3:contains("Instrument Weight Global")`).Next().Text())
log.Println(spec)
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
        return strings.Contains(html, "Included Accessories")
    }
}

// key: ESPの項目名, value: 構造体フィールド名
var strandbergMap = map[string]string{
	"BODY":         "BodyMaterial",
	"NECK":         "NeckMaterial",
	"FINGERBOARD":  "Fingerboard",
	"BRIDGE":       "Bridge",
	"PICKUPS":      "Pickups",
	"CONTROLS":     "Controls",
	"Price":        "Price",
	"SCALE":        "ScaleLengthMM",
	"FRET":         "FretCount",
	"INLAY":        "Inlays",
	"CONSTRUCTION": "Joint",
}

// サイトの項目名をフィールド名に変換
// func convertLabelStrandberg(label string) string {
//     return strandbergMap[label]
// }
