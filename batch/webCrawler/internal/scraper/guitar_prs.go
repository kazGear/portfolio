package scraper

import (
	"context"
	"fmt"
	"maps"
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

type guitarScraperPRS struct {
    gScraper guitarScraper
}

type callBacksPRS struct {
    funcs callBacks
}

func NewScraperPRS(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 20,
	})
    return &guitarScraperPRS{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksPRS(logger *log.Logger) *callBacksPRS {
    return &callBacksPRS{
        callBacks{
            logger: logger,
        },
    }
}

var regNotNeed = regexp.MustCompile(`(privatestock|amplifiers|pedals|accessories|contents/color)`)

func (g *guitarScraperPRS) CollectLinks(parentCtx context.Context) ([]string, error) {
    c       := g.gScraper.collector
    visited := make(map[string]struct{}, 150)
    mutex   := &sync.Mutex{}

    c.OnHTML("fluid-columns-repeater a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.Visit("https://www.prsguitars.jp/products")
    c.Wait()

    g.gScraper.urls = utils.MapToSliceUrl(visited)
    g.gScraper.urls = utils.RemoveNotNeedLinks(g.gScraper.urls, regNotNeed)
    return g.gScraper.urls, nil
}

func (g *guitarScraperPRS) Scrape(provider  PageProvider,
                                  parser    GuitarParser,
                                  parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksPRS) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://www.prsguitars.jp/products/[\w-]+/[\w-]+`, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 10 * time.Second)
        defer cancel()

        var html string

        err := chromedp.Run(ctx,
               chromedp.Navigate(url),
               chromedp.WaitVisible(`//span[contains("Tuning")]`, chromedp.ByQuery), // 求める要素が出るまで待つ
               chromedp.Sleep(200 * time.Millisecond), // JSが動く猶予を与える
               chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )
        if err != nil {
            return "", fmt.Errorf("[chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

func (c *callBacksPRS) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec[C.Maker]   = strconv.Itoa(C.PRS)
        spec[C.Name]    = doc.Find(`h1 span span span span span`).First().Text()
        spec[C.Comment] = ""
        spec[C.Series]  = strings.SplitN(spec[C.Name], " ", 2)[0]

        /*
        スペック表の構造例
        <span>
            Body Construction : Solidbody<br>
            Body Wood : Poplar<br>
            Top Carve : Flat Top
        </span>
        */
        bodySection        := doc.Find(`p span:contains("Top Wood"), p span:contains("Body Wood")`).Text()
        neckSection        := doc.Find(`p span:contains("Fretboard")`).Text()
        jointSection       := doc.Find(`p span:contains("Assembly")`).Text()
        finishSection      := doc.Find(`p span:contains("Finish Type")`).Text()
        hardwareSection    := doc.Find(`p span:contains("Bridge")`).Text()
        electronicsSection := doc.Find(`p span:contains("Pickup")`).Text()

        spec = parseSpec(bodySection, spec)
        spec = parseSpec(neckSection, spec)
        spec = parseSpec(jointSection, spec)
        spec = parseSpec(finishSection, spec)
        spec = parseSpec(hardwareSection, spec)
        spec = parseSpec(electronicsSection, spec)

        if len(spec["TreblePickup"]) <= 0 {
            spec["TreblePickup"] = spec["BassPickup"]
        }
        spec[C.Pickups] = fmt.Sprintf("%v / %v / %v", spec["BassPickup"], spec["MiddlePickup"], spec["TreblePickup"])

        // 画像、カラー取得
        doc.Find(`span:contains("COLORS")`).
            Closest("div").Parent(). // 直近の親divの親
            Children().Next().Children().Children().Children(). // 各色のギターカード
                Each(func(idx int, selector *goquery.Selection) {
                    nextSpec := make(map[string]string)
                    maps.Copy(nextSpec, spec)

                    nextSpec[C.Src], _ = selector.Find("img").Attr("src")
                    nextSpec[C.Color]  = strings.TrimSpace(selector.Text())

                    specs = utils.LockedAppend(mutex, specs, nextSpec)
                })
        return specs
    }
}

func parseSpec(specSection string, spec map[string]string) map[string]string {
    specSection     = strings.ReplaceAll(specSection, "\r\n", "\n") // 改行正規化
    specSection     = "\n" + specSection + "\n"
    splitedSection := strings.Split(specSection, "\n")

    for _, elem := range splitedSection {
        if len(elem) == 0 {
            continue
        }
        labelAndSpec := strings.SplitN(elem, ":", 2)
        specLabel   := strings.TrimSpace(labelAndSpec[0])
        key, exist := utils.ConvertLabel(specLabel, fieldMapPRS)

        if exist && len(labelAndSpec) > 1 {
            spec[key] = labelAndSpec[1]
        }
    }
    return spec
}

func (c *callBacksPRS) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksPRS) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "Tuning")
    }
}

// key: PRSの項目名, value: 構造体フィールド名

var fieldMapPRS = map[string]string{
	"Finish Type":            "BodyFinish",
    "Top Wood":               "BodyMaterialTop",
	"Body Wood":              "BodyMaterialBack",
    "Back Wood":              "BodyMaterialBack",
	"Bridge":                 "Bridge",
	"Controls":               "Controls",
	"Fretboard Wood":         "Fingerboard",
	"Number of Frets":        "FretCount",
	"Fretboard Inlay":        "Inlays",
	"Neck/Body Assembly Type":"Joint",
	"Neck Wood":              "NeckMaterial",
	"Scale Length":           "ScaleLengthMM",
    "Treble Pickup":          "TreblePickup",
    "Middle Pickup":          "MiddlePickup",
    "Bass Pickup":            "BassPickup",
}