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
    ctx, cancel := context.WithTimeout(tabCtx, 600 * time.Second)
    defer cancel()

    modelLinks, err := collectLinksModelView(ctx)

    if err != nil {
        return []string{}, err
    }
    detailLinks, err := collectLinksDetailView(ctx, modelLinks)

    if err != nil {
        return []string{}, err
    }
    utils.LogCollectedLinks(detailLinks, g.gScraper.logger)
    g.gScraper.urls = detailLinks
    return g.gScraper.urls, nil
}

// モデル一覧へのリンク収集
func collectLinksModelView(ctx context.Context) ([]string, error) {
    var nodes []*cdp.Node
    var modelLinks = make([]string, 0, 450)

    // クリック対象のノード取得
    err := chromedp.Run(ctx,
        chromedp.Navigate(`https://www.ibanez.com/jp/`),
        tryWaitVisible(".idx-product-tabs-wrap"),
        chromedp.Nodes(
            `//div[contains(@class, "idx-product-tabs")]//div[contains(@class, "js-tab-parent-in")]`,
            &nodes,
        ),
    )
    if err != nil {
        return []string{}, fmt.Errorf("[Chromedp error]: %w\n", err)
    }
    if len(nodes) == 0 {
        return []string{}, fmt.Errorf(`[WARN] nodes is empty, but continuing`)
    }

    var htmlBuilder strings.Builder
    var htmlParts = make([]*string, 0, len(nodes))
    for i := 0; i < len(nodes); i++ {
        htmlParts = append(htmlParts, new(string))
    }
    // クリックの都度HTML抽出
    // 左から４ノードのみ必要。[nodes]: ELECTRIC GUITARS | BASSES | HOLLOW BODIES | ACOUSTICS | ELECTRONICS | ACCESSORIES
    for i := 0; i < 4; i++ {
        chromedp.Run(ctx,
            tryClick(nodes[i].FullXPath()),
            chromedp.Sleep(200 * time.Millisecond),
            chromedp.OuterHTML(`.idx-product-tabs-wrap`, htmlParts[i], chromedp.ByQuery),
        )
        htmlBuilder.WriteString(*htmlParts[i])
    }
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBuilder.String()))
    if err != nil {
        return []string{}, fmt.Errorf(`[Html read error(goquery)]: %w`, err)
    }
    modelLinks = collectLinks(".idx-product-tabs-wrap a.rt_cf_pm_href", doc, 450)
    modelLinks = toAbsLinks(modelLinks, `https://www.ibanez.com`, 450)
    return modelLinks, nil
}

// 詳細ページのリンク収集
func collectLinksDetailView(ctx context.Context, modelLinks []string) ([]string, error) {
    var htmlBuilder strings.Builder
    htmlParts := make([]*string, 0, len(modelLinks))

    for i := 0; i < len(modelLinks); i++ {
        htmlParts = append(htmlParts, new(string))
    }
    // モデル一覧から詳細ページへのリンクを収集
    for idx, link := range modelLinks {
        chromedp.Run(ctx,
            chromedp.Navigate(link),
            tryWaitVisible(".products-model-series-lineup-list a"),
            chromedp.OuterHTML(
                "main", htmlParts[idx], chromedp.ByQuery, // 必要なリンクが散らばっているので大きく取得 main
            ),
        )
        htmlBuilder.WriteString(*htmlParts[idx])
    }
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBuilder.String()))

    if err != nil {
        return []string{}, fmt.Errorf(`[Html read error(goquery)]: %w`, err)
    }
    detailLinks := collectLinks(".products-model-series-lineup-list a", doc, 2050)
    detailLinks  = toAbsLinks(detailLinks, `https://www.ibanez.com`, 2050)
    return detailLinks, nil
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
