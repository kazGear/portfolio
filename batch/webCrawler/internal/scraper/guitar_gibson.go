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
	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type guitarScraperGibson struct {
    gScraper guitarScraper
}

type callBacksGibson struct {
    funcs callBacks
}


func NewScraperGibson(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 30,
	})
    return &guitarScraperGibson{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksGibson() GuitarCallbacks {
    return &callBacksGibson{
        callBacks{},
    }
}

func (e *guitarScraperGibson) CollectLinks(parentCtx context.Context) []string {
    c       := e.gScraper.collector
    visited := make(map[string]struct{}, 600)
    mutex   := &sync.Mutex{}

    c.OnHTML(".body-types a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML(".category-wrapper .model-card a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.Visit("https://gibson.jp/")
    c.Wait()

    e.gScraper.urls = mapToSliceUrl(visited)
    return e.gScraper.urls
}

func (e *guitarScraperGibson) Scrape(funcs GuitarCallbacks,
                                     parentCtx context.Context,
) []*model.Guitar {
    guitars, _ := e.gScraper.scrapeFrame(funcs, parentCtx)
    utils.AutoDownLoader(guitars, "images/gibson")
    return guitars
}

func (e *callBacksGibson) FetchDynamicPage(parentCtx context.Context) func(url string) string {
    return func(url string) string {
        if !isDetailPage(`https://gibson.jp/(electric|acoustic)/[a-z0-9\-]+`, url) {
            return ""
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 20 * time.Second)
        defer cancel()

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            tryWaitVisible("#body-wrap"), // 求める要素が出るまで待つ
            chromedp.Sleep(200 * time.Millisecond), // JSが動く猶予を与える
        )
        if err != nil {
            log.Printf("[Chromedp error]: %v [url]: %v\n", err, url)
            return ""
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
        return html
    }
}

// シリーズ名の抽出用
var regSeriesGibson = regexp.MustCompile(
    `(Les Paul|SG|ES-\d+|Flying V|Explorer|Firebird|Hummingbird|J\-\d+)+\s[A-Za-z]+\b`,
)

func (e *callBacksGibson) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := []map[string]string{}
        mutex := &sync.Mutex{}

        spec        := map[string]string{}
        getElem     := utils.GetElem(doc)
        getElemNext := utils.GetElemNextToLabel(doc)

        doc.Find(`#cart-options h2.marketing-headline small`).Remove() // Nameからノイズを除去

        spec["Maker"]            = strconv.Itoa(constants.Gibson)
        spec["Name"]             = getElem(`h2.marketing-headline`)
        spec["Color"]            = getElem(`div#displayed-finish`)
        spec["Comment"]          = getElem(`#cart-options .marketing-copy p`)
        neckPickup              := getElemNext(`.spec-item div:contains("Neck pickup")`)
        bridgePickup            := getElemNext(`.spec-item div:contains("Bridge pickup")`)
        spec["Pickups"]          = fmt.Sprintf(constants.PickupsFormat, neckPickup, bridgePickup)
        spec["Price"]            = strconv.Itoa(constants.InvalidNumber)
        src, _                  := doc.Find(`img#gallery-front`).Attr(`src`)
        spec["Src"]              = strings.TrimSpace(src)
        spec["Series"]           = regSeriesGibson.FindString(spec["Name"])
        spec["Weight"]           = strconv.Itoa(constants.InvalidNumber)

        doc.Find(`#product-overview .spec-section .spec-item`).Each(func(idx int, selector *goquery.Selection) {
            label      := strings.TrimSpace(selector.Find(`div:nth-child(1)`).Text())
            elem       := strings.TrimSpace(selector.Find(`div:nth-child(2)`).Text())
            field      := utils.ConvertLabel(label, fieldMapGibson)
            spec[field] = elem
        })
        spec["BodyMaterial"] = spec["BodyMaterialTop"] + " " + spec["BodyMaterialBack"]

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
        return strings.Contains(html, "product-overview")
    }
}

var fieldMapGibson = map[string]string{
    "Finish":               "BodyFinish",
    "Top":                  "BodyMaterialTop",
    "Body Material":        "BodyMaterialBack",
    "Body":                 "BodyMaterialBack",
    "Bridge":               "Bridge",
    "Controls":             "Controls",
    "Fingerboard Material": "Fingerboard",
    "Number Of Frets":      "FretCount",
    "Inlays":               "Inlays",
    "Joint":                "Joint",
    "Material":             "NeckMaterial",
    "Scale Length":         "ScaleLengthMM",
}
