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
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type guitarScraperIbanez struct {
    gScraper guitarScraper
}

type callBacksIbanez struct {
    funcs callBacks
}


func NewScraperIbanez(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 30,
	})
    return &guitarScraperIbanez{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksIbanez(logger *log.Logger) GuitarCallbacks {
    return &callBacksIbanez{
        callBacks{
            logger: logger,
        },
    }
}

func (g *guitarScraperIbanez) CollectLinks(parentCtx context.Context) ([]string, error) {
    // タブごとに独立した context を作る
    tabCtx, tabCancel := chromedp.NewContext(parentCtx)
    defer tabCancel()
    // タブにだけ timeout を付ける
    ctx, cancel := context.WithTimeout(tabCtx, 20 * time.Second)
    defer cancel()

    // モデル一覧へのリンク収集

    var nodes []*cdp.Node
    var targetLinks = make([]string, 0, 450)

    chromedp.Run(ctx,
        chromedp.Navigate(`https://www.ibanez.com/jp/`),
        tryWaitVisible(".idx-product-tabs-wrap"),
        chromedp.Sleep(400 * time.Millisecond),
        chromedp.Nodes(
            `//div[contains(@class, "idx-product-tabs")]//div[contains(@class, "js-tab-parent-in")]`,
            &nodes,
        ),
    )
    if len(nodes) == 0 {
        log.Println(`[WARN] nodes is empty, but continuing`)
    }

    var html string
    var htmlParts = make([]*string, 0, len(nodes))
    for i := 0; i < len(nodes); i++ {
        htmlParts = append(htmlParts, new(string))
    }
    // クリックの都度HTML抽出
    for idx, node := range nodes {
        chromedp.Run(ctx,
            chromedp.Sleep(600 * time.Millisecond),
            tryClick(node.FullXPath()),
            chromedp.OuterHTML(`.idx-product-tabs-wrap`, htmlParts[idx], chromedp.ByQuery),
        )
        html += *htmlParts[idx]
    }
    doc, err   := goquery.NewDocumentFromReader(strings.NewReader(html))
    if err != nil {
        return nil, fmt.Errorf(`[Html read error(goquery)]: %w`, err)
    }
    targetLinks = collectLinks(".idx-product-tabs-wrap a.rt_cf_pm_href", doc, 450)
    targetLinks = toAbsLinks(targetLinks, `https://www.ibanez.com`, 450)

    // 詳細ページのリンク収集

    html = ""
    htmlParts = make([]*string, 0, len(targetLinks))

    for i := 0; i < len(targetLinks); i++ {
        htmlParts = append(htmlParts, new(string))
    }

    for idx, link := range targetLinks {
        chromedp.Run(ctx,
            chromedp.Navigate(link),
            tryWaitVisible(".products-model-series-lineup"),
            chromedp.Sleep(200 * time.Millisecond),
            autoScroll(ctx),
            chromedp.Sleep(500 * time.Millisecond),
            chromedp.OuterHTML(
                `.products-model-series-lineup-list`, htmlParts[idx], chromedp.ByQueryAll,
            ),
        )
        html += *htmlParts[idx]
    }
    doc, err    = goquery.NewDocumentFromReader(strings.NewReader(html))
    if err != nil {
        return nil, fmt.Errorf(`[Html read error(goquery)]: %w`, err)
    }
    targetLinks = collectLinks(".products-model-series-lineup-list a", doc, 1500)
    targetLinks = toAbsLinks(targetLinks, `https://www.ibanez.com`, 1500)

g.gScraper.logger.Println("[[ html ]]", html)
utils.LogCollectedLinks(targetLinks, g.gScraper.logger)
    // g.gScraper.urls = mapToSliceUrl(visited)
    return g.gScraper.urls, nil
}

func (g *guitarScraperIbanez) Scrape(funcs GuitarCallbacks,
                                     parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(funcs, parentCtx)
    utils.AutoDownLoader(guitars, "images/gibson")
    return guitars
}

func (c *callBacksIbanez) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`https://gibson.jp/(electric|acoustic)/[a-z0-9\-]+`, url) {
            return "", nil
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
            return "", fmt.Errorf("[Chromedp error]: %v [url]: %v\n", err, url)
        }
        // HTMLを取得、マージ
        var html string
        err = chromedp.Run(ctx,
            chromedp.OuterHTML("html", &html, chromedp.ByQuery),
        )
        if err != nil {
            return "", fmt.Errorf("[Chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

// シリーズ名の抽出用
var regSeriesIbanez = regexp.MustCompile(
    `(Les Paul|SG|ES-\d+|Flying V|Explorer|Firebird|Hummingbird|J\-\d+)+\s[A-Za-z]+\b`,
)

func (c *callBacksIbanez) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := []map[string]string{}
        mutex := &sync.Mutex{}

        spec        := map[string]string{}
        getElem     := utils.GetElem(doc)
        getElemNext := utils.GetElemNextToLabel(doc)

        doc.Find(`#cart-options h2.marketing-headline small`).Remove() // Nameからノイズを除去

        spec["Maker"]            = strconv.Itoa(constants.Ibanez)
        spec["Name"]             = getElem(`h2.marketing-headline`)
        spec["Color"]            = getElem(`div#displayed-finish`)
        spec["Comment"]          = getElem(`#cart-options .marketing-copy p`)
        neckPickup              := getElemNext(`.spec-item div:contains("Neck pickup")`)
        bridgePickup            := getElemNext(`.spec-item div:contains("Bridge pickup")`)
        spec["Pickups"]          = fmt.Sprintf(constants.PickupsFormat, neckPickup, bridgePickup)
        spec["Price"]            = strconv.Itoa(constants.InvalidNumber)
        src, _                  := doc.Find(`img#gallery-front`).Attr(`src`)
        spec["Src"]              = strings.TrimSpace(src)
        spec["Series"]           = regSeriesIbanez.FindString(spec["Name"])
        spec["Weight"]           = strconv.Itoa(constants.InvalidNumber)

        doc.Find(`#product-overview .spec-section .spec-item`).Each(func(idx int, selector *goquery.Selection) {
            label      := strings.TrimSpace(selector.Find(`div:nth-child(1)`).Text())
            elem       := strings.TrimSpace(selector.Find(`div:nth-child(2)`).Text())
            field      := utils.ConvertLabel(label, fieldMapIbanez)
            spec[field] = elem
        })
        spec["BodyMaterial"] = spec["BodyMaterialTop"] + " " + spec["BodyMaterialBack"]

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksIbanez) BuildGuitar() func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, c.funcs.logger)
    }
}

func (c *callBacksIbanez) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "product-overview")
    }
}

var fieldMapIbanez = map[string]string{
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
