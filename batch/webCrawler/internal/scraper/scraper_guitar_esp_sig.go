package scraper

import (
	"context"
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

type guitarScraperEspSig struct {
    gScraper guitarScraper
}

type callBacksEspSig struct {
    funcs callBacks
}

func NewEspSigScraper() Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5,
	})
    return &guitarScraperEspSig{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksEspSig() GuitarCallbacks {
    return &callBacksEsp{
        callBacks{},
    }
}

func (e *guitarScraperEspSig) CollectLinks() *[]string {
    c       := e.gScraper.collector
    mutex   := &sync.Mutex{}
    visited := make(map[string]struct{}, 100)

    // URL収集、クロール
    c.OnHTML(`.searchResultBlock.gallery_item .searchResultBlock_item a[href*="/artists/"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.Visit("https://espguitars.co.jp/signatureseries/")
    c.Wait()

    e.gScraper.urls = getDistinctUrls(visited)
    return &e.gScraper.urls
}

func (e *guitarScraperEspSig) Scrape(funcs GuitarCallbacks, ctx context.Context) (*[]model.Guitar, error) {
    guitars, _ := e.gScraper.scrapeFrame(funcs, ctx)
    return guitars, nil
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (e *guitarScraperEspSig) FetchDynamicPage(parentCtx context.Context) func(url string) string {
    return func(url string) string {
        if !isDetailPage(`^https://espguitars\.co\.jp/product/\d{4,}/?$`, url) {
            return ""
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 15*time.Second)
        defer cancel()

        var html string

        // 大本となるHTMLを取得
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
            log.Printf("[chromedp error]: %v [url]: %v\n", err, url)
            return ""
        }
        return html
    }
}

func (e *callBacksEspSig) CollectSpec() func(doc *goquery.Document) *[]map[string]string {
    return func(doc *goquery.Document) *[]map[string]string {
        specs := make([]map[string]string, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec["Maker"]   = strconv.Itoa(constants.Esp)
        spec["Name"]    = strings.TrimSpace(doc.Find("h1.header_title").Text())
        src, _         := doc.Find("#main .header_content img.transform-5").Attr("src")
        spec["Src"]     = strings.TrimSpace(src)
        spec["Color"]   = strings.TrimSpace(doc.Find(".header_content h3.clr_name").Text())
        spec["Comment"] = strings.TrimSpace(doc.Find("#specialfeatures .container_small p").Text())
        spec["Price"]   = strings.TrimSpace(doc.Find("p.detail_price").Text())

        doc.Find("#specifications table.tbl_spec tr").Each(func(i int, selector *goquery.Selection) {
            th      := strings.TrimSpace(selector.Find("th").Text())
            td      := strings.TrimSpace(selector.Find("td").Text())
            th       = convertLabelEspSig(th)
            spec[th] = td
        })
        specs = utils.LockedAppend(mutex, specs, spec)
        return &specs
    }
}

func (e *callBacksEspSig) BuildGuitar() func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec)
    }
}

func (e *callBacksEspSig) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "tbl_spec")
    }
}

// key: ESPの項目名, value: 構造体フィールド名
var espSigFieldMap = map[string]string{
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
func convertLabelEspSig(label string) string {
    return espFieldMap[label]
}
