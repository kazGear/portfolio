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
	"github.com/kazGear/portfolio/goBatch/internal/model"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
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
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &guitarScraperIbanez{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksIbanez(logger *log.Logger) *callBacksIbanez {
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
    modelLinks = utils.CollectLinks(".idx-product-tabs-wrap a.rt_cf_pm_href", doc, 500)
    modelLinks = utils.ToAbsLinks(modelLinks, `https://www.ibanez.com`, 500)
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
    detailLinks := utils.CollectLinks(".products-model-series-lineup-list a", doc, 2100)
    detailLinks  = utils.ToAbsLinks(detailLinks, `https://www.ibanez.com`, 2100)
    return detailLinks, nil
}

func (g *guitarScraperIbanez) Scrape(provider  PageProvider,
                                     parser    GuitarParser,
                                     parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
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
        // 詳細urlにギター、ベース以外が紛れてしまうのでフィルタリング
        factoryTuning := doc.Find(`.rt_cf_p_data_factory_tuning`).Text()
        if len(factoryTuning) <= 0 {
            return []map[string]string{} // upsertで弾かれる
        }

        specs := []map[string]string{}
        spec  := map[string]string{}
        mutex := &sync.Mutex{}

        spec[C.Maker]            = strconv.Itoa(C.Ibanez)
        spec[C.Name]             = doc.Find(`.rt_cf_p_cm_product_code`).Text()
        spec[C.Color]            = doc.Find(`.rt_cf_pcl_color_name_jp_1, .rt_cf_pcl_color_name_ag_jp_1`).Text()
        spec[C.BodyFinish]       = ""
        spec[C.BodyMaterialBack] = doc.Find(`.rt_cf_p_data_body_material, .rt_cf_p_ag_side_material`).Text()
        spec[C.BodyMaterialTop]  = doc.Find(`.rt_cf_p_data_body_top_material`).Text()
        spec[C.Bridge]           = doc.Find(`.rt_cf_p_data_bridge`).Text()
        spec[C.Controls]         = ""

        // コメント収集
        var comment = strings.Builder{}
        doc.Find(`#products_detail_features section section .fl_left`).Each(func(idx int, selector *goquery.Selection) {
            title  := selector.Find(`h3`).Text()
            detail := selector.Find(`p`).Text()
            comment.WriteString(fmt.Sprintf("%v\n%v\n", title, detail))
        })
        spec[C.Comment] = comment.String()
        // spec[C.Comment]          = ""

        spec[C.Fingerboard] = doc.Find(`.rt_cf_p_data_fretboard`).Text()
        spec[C.FretCount]   = doc.Find(`.rt_cf_p_data_number_fret`).Text()

        inlays             := doc.Find(`.rt_cf_p_ag_face_inlay`).Text()

        if len(inlays) > 0 {
            spec[C.Inlays] = inlays
        } else {
            spec[C.Inlays] = doc.Find(`.rt_cf_p_data_in`).Text()
        }

        spec[C.Joint]        = ""
        spec[C.NeckMaterial] = doc.Find(`.rt_cf_p_data_neck_material`).Text()
        spec[C.NeckPickup]   = doc.Find(`.rt_cf_p_data_neck_pickup`).Text()
        spec[C.CenterPickup] = doc.Find(`.rt_cf_p_data_middle_pickup`).Text()
        spec[C.BridgePickup] = doc.Find(`.rt_cf_p_data_bridge_pickup`).Text()
        spec[C.Price]  = doc.Find(`.rt_cf_p_cm_price`).Text()
        src, _        := doc.Find(`.products-detail-main-modal-img`).Attr(`src`)
        spec[C.Src]    = src
        spec[C.Series] = doc.Find(`ul a:contains("` + spec[C.Name] + `")`).
                            Parent().Parent().Prev().Children().Text()

        spec[C.ScaleLengthMM] = doc.Find(`.rt_cf_p_data_scale_mm`).Text()
        spec[C.Weight]        = strconv.Itoa(C.InvalidNumber)

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksIbanez) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksIbanez) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "products-spec-table-li")
    }
}