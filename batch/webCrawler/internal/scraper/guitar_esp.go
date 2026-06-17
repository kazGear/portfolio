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

type guitarScraperEsp struct {
    gScraper guitarScraper
}

type callBacksEsp struct {
    funcs callBacks
}


func NewScraperEsp(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(4),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 20,
	})
    return &guitarScraperEsp{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksEsp(logger *log.Logger) GuitarCallbacks {
    return &callBacksEsp{
        callBacks{
            logger: logger,
        },
    }
}

func (g *guitarScraperEsp) CollectLinks(parentCtx context.Context) ([]string, error) {
    c       := g.gScraper.collector
    visited := make(map[string]struct{}, 500)
    mutex   := &sync.Mutex{}

    // URL収集、クロール
    c.OnHTML("#item .figcap a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML("#inner_content .figcap a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML("section.color_variation a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.Visit("https://espguitars.co.jp/products/esp")
    c.Wait()

    g.gScraper.urls = utils.MapToSliceUrl(visited)
    return g.gScraper.urls, nil
}

func (g *guitarScraperEsp) Scrape(funcs GuitarCallbacks,
                                  parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(funcs, parentCtx)
    return guitars
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksEsp) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://espguitars\.co\.jp/product/\d{4,}/?$`, url) {
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
               chromedp.WaitVisible("#main", chromedp.ByQuery), // 求める要素が出るまで待つ
               chromedp.Sleep(300 * time.Millisecond), // JSが動く猶予を与える
               tryWaitReady("h1.header_title"), // 必要な要素が生成されるのを待つ
               tryWaitReady(".tbl_spec"),
               tryWaitReady("p.detail_price"),
               chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )
        if err != nil {
            return "", fmt.Errorf("[chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

var regSeries = regexp.MustCompile(`^[\w-]+`)

func (c *callBacksEsp) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec["Maker"]   = strconv.Itoa(constants.Esp)
        spec["Name"]    = strings.TrimSpace(doc.Find("h1.header_title").Text())
        spec["Color"]   = strings.TrimSpace(doc.Find(".header_content h3.clr_name").Text())
        spec["Comment"] = strings.TrimSpace(doc.Find("#specialfeatures .container_small p").Text())
        spec["Price"]   = strings.TrimSpace(doc.Find("p.detail_price").Text())
        src, _         := doc.Find("#main .header_content img.transform-5").Attr("src")
        spec["Src"]     = strings.TrimSpace(src)
        spec["Series"]  = regSeries.FindString(spec["Name"])

        doc.Find("#specifications table.tbl_spec tr").Each(func(idx int, selector *goquery.Selection) {
            th      := strings.TrimSpace(selector.Find("th").Text())
            td      := strings.TrimSpace(selector.Find("td").Text())
            th       = utils.ConvertLabel(th, fieldMapEsp)
            spec[th] = td
        })

        bodyMaterial := spec["BodyMaterial"]
        materials    := strings.Split(bodyMaterial, ",")

        if len(materials) == 1 {
            spec["BodyMaterialBack"] = materials[0]
        } else if len(materials) == 2 {
            spec["BodyMaterialTop"]  = materials[0]
            spec["BodyMaterialBack"] = materials[1]
        } else {
            spec["BodyMaterialBack"] = materials[0]
        }

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksEsp) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksEsp) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "tbl_spec")
    }
}

// key: ESPの項目名, value: 構造体フィールド名
var fieldMapEsp = map[string]string{
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
