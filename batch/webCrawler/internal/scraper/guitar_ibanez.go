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
    ctx, cancel := context.WithTimeout(tabCtx, 300 * time.Second)
    defer cancel()

    modelLinks, err := collectLinksModelView(ctx)

    if err != nil {
        return []string{}, err
    }
    detailLinks, err := collectLinksDetailView(ctx, modelLinks)

    if err != nil {
        return []string{}, err
    }
    g.gScraper.urls = utils.GetDistinctLinks(detailLinks)
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
    modelLinks = collectLinks(".idx-product-tabs-wrap a.rt_cf_pm_href", doc, 500)
    modelLinks = toAbsLinks(modelLinks, `https://www.ibanez.com`, 500)
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
    detailLinks := collectLinks(".products-model-series-lineup-list a", doc, 2100)
    detailLinks  = toAbsLinks(detailLinks, `https://www.ibanez.com`, 2100)
    return detailLinks, nil
}

func (g *guitarScraperIbanez) Scrape(funcs GuitarCallbacks,
                                     parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(funcs, parentCtx)
    return guitars
}

func (c *callBacksIbanez) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`https://www/ibanez/com/jp/products/detail/[a-z]+\d+`, url) {
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
            tryWaitVisible(".products-spec-table-li"), // 求める要素が出るまで待つ
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

func (c *callBacksIbanez) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        getElem       := utils.GetElem(doc)

        // 詳細urlにギター、ベース以外が紛れてしまうのでフィルタリング
        factoryTuning := getElem(`.rt_cf_p_data_factory_tuning`)
        if len(factoryTuning) <= 0 {
            return []map[string]string{} // upsertで弾かれる
        }

        specs := []map[string]string{}
        spec  := map[string]string{}
        mutex := &sync.Mutex{}

        spec["Maker"]            = strconv.Itoa(constants.Ibanez)
        spec["Name"]             = getElem(`.rt_cf_p_cm_product_code`)
        spec["Color"]            = getElem(`.rt_cf_pcl_color_name_jp_1, .rt_cf_pcl_color_name_ag_jp_1`)
        spec["Comment"]          = ""
        spec["BodyFinish"]       = ""
        spec["BodyMaterialBack"] = getElem(`.rt_cf_p_data_body_material, .rt_cf_p_ag_side_material`)
        spec["BodyMaterialTop"]  = getElem(`.rt_cf_p_data_body_top_material`)
        spec["BodyMaterial"]     = spec["BodyMaterialTop"] + " / " + spec["BodyMaterialBack"]
        spec["Bridge"]           = getElem(`.rt_cf_p_data_bridge`)
        spec["Controls"]         = ""
        spec["Comment"]          = ""
        spec["Fingerboard"]      = getElem(`.rt_cf_p_data_fretboard`)
        spec["FretCount"]        = getElem(`.rt_cf_p_data_number_fret`)

        inlays                  := getElem(`.rt_cf_p_ag_face_inlay`)
        if len(inlays) > 0 {
            spec["Inlays"] = inlays
        } else {
            spec["Inlays"] = getElem(`.rt_cf_p_data_in`)
        }

        spec["Joint"]            = ""
        spec["NeckMaterial"]     = getElem(`.rt_cf_p_data_neck_material`)

        neckPickup              := getElem(`.rt_cf_p_data_neck_pickup`)
        middlePickup            := getElem(`.rt_cf_p_data_middle_pickup`)
        bridgePickup            := getElem(`.rt_cf_p_data_bridge_pickup`)
        spec["Pickups"]          = fmt.Sprintf(`%v / %v / %v `, neckPickup, middlePickup, bridgePickup)

        spec["Price"]            = getElem(`.rt_cf_p_cm_price`)
        src, _                  := doc.Find(`.products-detail-main-modal-img`).Attr(`src`)
        spec["Src"]              = strings.TrimSpace(src)
        spec["Series"]           = strings.TrimSpace(doc.Find(
                                    `ul a:contains("` + spec["Name"] + `")`,
                                         ).Parent().Parent().Prev().Children().Text())
        spec["ScaleLengthMM"]    = getElem(`.rt_cf_p_data_scale_mm`)
        spec["Weight"]           = strconv.Itoa(constants.InvalidNumber)

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
        return strings.Contains(html, "products-spec-table-li")
    }
}