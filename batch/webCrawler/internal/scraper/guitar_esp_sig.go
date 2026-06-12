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

type guitarScraperEspSig struct {
    gScraper guitarScraper
}

type callBacksEspSig struct {
    funcs callBacks
}

func NewScraperEspSig(logger *log.Logger) Scraper {
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
			logger:    logger,
        },
    }
}

func NewCallBacksEspSig() GuitarCallbacks {
    return &callBacksEspSig{
        callBacks{},
    }
}

func (g *guitarScraperEspSig) CollectLinks(parentCtx context.Context) ([]string, error) {
    c       := g.gScraper.collector
    mutex   := &sync.Mutex{}
    visited := make(map[string]struct{}, 100)

    // URL収集、クロール
    c.OnHTML(`.searchResultBlock.gallery_item .searchResultBlock_item a[href*="/artists/"]`,
			 func(html *colly.HTMLElement,
	) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.Visit("https://espguitars.co.jp/signatureseries/")
    c.Wait()

    g.gScraper.urls = mapToSliceUrl(visited)
    return g.gScraper.urls, nil
}

func (g *guitarScraperEspSig) Scrape(funcs GuitarCallbacks,
									 parentCtx context.Context,
) []*model.Guitar {
	guitars := g.gScraper.scrapeFrame(funcs, parentCtx)
    return guitars
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksEspSig) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://espguitars\.co\.jp/artists/\d{4,}/?$`, url) {
            return "", nil
        }
		// タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 8*time.Second)
        defer cancel()

        var html string

        err := chromedp.Run(ctx,
               chromedp.Navigate(url),
               chromedp.WaitVisible("#main", chromedp.ByQuery), // 求める要素が出るまで待つ
               chromedp.Sleep(200 * time.Millisecond), // JSが動く猶予を与える
			   chromedp.Poll(`() => document.querySelectorAll("section.tab_detail").length >= 7`,
			   				 nil, chromedp.WithPollingInterval(500*time.Millisecond)), // 必要な要素が生成されるのを待つ
			   chromedp.Poll(`() => document.querySelectorAll("section.tab_detail .signatures_brand_logo").length >= 7`,
							 nil, chromedp.WithPollingInterval(500*time.Millisecond)),
			   chromedp.Poll(`() => document.querySelectorAll("section.tab_detail .content_spec-detail").length >= 7`,
							 nil, chromedp.WithPollingInterval(500*time.Millisecond)),
               chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )
        if err != nil {
            return "", fmt.Errorf("[chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

func (c *callBacksEspSig) CollectSpec() func(doc *goquery.Document) []map[string]string {
	return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 10)
        mutex := &sync.Mutex{}

		doc.Find("#main section.tab_detail").Each(func(idx int, selector1 *goquery.Selection) {
			spec := map[string]string{}

			spec["Maker"]   = strconv.Itoa(constants.EspSignature)
			spec["Name"]    = strings.TrimSpace(selector1.Find(".product_series_logo_name").Text())
			src, _         := selector1.Find("img.main_image").Attr("src")
			spec["Src"]     = strings.TrimSpace(src)
			spec["Comment"] = strings.TrimSpace(selector1.Find(".content_spec-detail em strong").Text())
			spec["Price"]   = strings.TrimSpace(selector1.Find(
				".content_borderline.text-center p, .content_spec-detail div p.text-center",
			).Text())

			spec["Series"] 	= strings.TrimSpace(doc.Find("div.pd30 h1.text-center span").Text())

			selector1.Find(".tbl_spec tr").Each(func(idx int, selector2 *goquery.Selection) {
				th      := strings.TrimSpace(selector2.Find("th").Text())
				td      := strings.TrimSpace(selector2.Find("td").Text())
				th       = utils.ConvertLabel(th, fieldMapEspSig)
				spec[th] = td
			})
			specs = utils.LockedAppend(mutex, specs, spec)
		})
        return specs
    }
}

func (c *callBacksEspSig) BuildGuitar() func(spec map[string]string) *model.Guitar {
	return func(spec map[string]string) *model.Guitar {
		return buildGuitarFrame(spec)
    }
}

func (c *callBacksEspSig) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "tbl_spec")
    }
}

// key: ESPの項目名, value: 構造体フィールド名
var fieldMapEspSig = map[string]string{
	"BODY":         "BodyMaterial",
	"NECK":         "NeckMaterial",
	"FINGERBOARD":  "Fingerboard",
	"BRIDGE":       "Bridge",
	"PICKUPS":      "Pickups",
	"CONTROLS":     "Controls",
	"Price":        "Price",
	"SCALE":        "ScaleLengthMM",
	"FRET":         "FretCount",
	"FRETS":        "FretCount",
	"INLAY":        "Inlays",
	"CONSTRUCTION": "Joint",
	"COLOR":		"Color",
}