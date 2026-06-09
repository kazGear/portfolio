package scraper

import (
	"context"
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type Scraper interface {
	Scrape(funcs GuitarCallbacks, ctx context.Context) (*[]model.Guitar, error)
	CollectLinks(ctx context.Context)                  []string
}

type GuitarCallbacks interface {
    IsStaticPage() func(html string) bool
    FetchDynamicPage(ctx context.Context) func(url string) string
    CollectSpec()  func(doc *goquery.Document)  *[]map[string]string
    BuildGuitar()  func(spec map[string]string) *model.Guitar
}

type guitarScraper struct {
    urls      []string
	collector *colly.Collector
    mutex     *sync.Mutex
}

type callBacks struct {}

// スクレイピング実行のフレームワーク
func (e *guitarScraper) scrapeFrame(funcs GuitarCallbacks, ctx context.Context) (*[]model.Guitar, error) {
    var guitars = make([]model.Guitar, 0, 400)

    if len(e.urls) <= 0 {
        return &[]model.Guitar{}, errors.New("巡回用URLがありません。")
    }
    for _, url := range e.urls {
        // 静的/動的を判定して HTML を取得、DOM化
        html := fetchPage(url, funcs.IsStaticPage(), funcs.FetchDynamicPage(ctx))
        doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

        if err != nil {
            log.Println("goquery error:", err)
            continue
        }
        if len(html) <= 0 { continue }

        collectSpec := funcs.CollectSpec()
        buildGuitar := funcs.BuildGuitar()
        specs       := collectSpec(doc) // 1ページ：N詳細ページでもOK

        for _, spec := range *specs {
            guitar := buildGuitar(spec)

            if len(guitar.Name) <= 0 || len(guitar.Color) <= 0 { continue }
            guitars = utils.LockedAppend(e.mutex, guitars, *guitar)
        }
    }
    return &guitars, nil
}

// ギター構造体の構築フレームワーク
func buildGuitarFrame(spec map[string]string) (*model.Guitar) {

	guitar := model.Guitar{}

    var errMaker error
	guitar.Maker, errMaker = strconv.Atoi(spec["Maker"])
	guitar.Name            = spec["Name"]

    if errMaker != nil {
        return &model.Guitar{}
	}

	guitar.BodyFinish        = spec["BodyFinish"]
	guitar.BodyMaterial      = spec["BodyMaterial"]
    guitar.BodyMaterialBack  = utils.SearchWoodCode(spec["BodyMaterialBack"])
	guitar.BodyMaterialTop = utils.SearchWoodCode(spec["BodyMaterialTop"])
    guitar.Bridge            = spec["Bridge"]
	guitar.Color             = spec["Color"]
    guitar.ColorCd           = utils.ConvertColorCd(guitar.Color)
	guitar.Comment           = spec["Comment"]
	guitar.Controls          = spec["Controls"]
    guitar.Fingerboard       = utils.SearchWoodCode(spec["Fingerboard"])

	var errFretCount error
	guitar.FretCount, errFretCount = utils.GetFretCount(spec["FretCount"])

    if errFretCount != nil {
        log.Println(errFretCount)
    }

	guitar.Inlays       = spec["Inlays"]
	guitar.Joint        = spec["Joint"]
    guitar.NeckMaterial = utils.SearchWoodCode(spec["NeckMaterial"])
    guitar.Pickups      = spec["Pickups"]

    var errPrice error
	guitar.Price, errPrice = utils.ParsePrice(spec["Price"])

    if errPrice != nil {
        log.Println(errPrice)
    }

    guitar.ScaleLengthMM = utils.TrimScaleUnit(spec["ScaleLengthMM"])
	guitar.Series        = spec["Series"]
	guitar.Src           = spec["Src"]

    if len(guitar.Src) <= 0 {
        return &model.Guitar{}
    }
    guitar.Weight = utils.ParseWight(spec["Weight"])

	return &guitar
}

// 動的、静的ページを取得（動的が優先）。funcは個々で実装の必要あり。
func fetchPage(
    url string,
    isStaticPage func(string)bool,
    fetchDynamicPage func(string)string,
) string {

    html := fetchStaticPage(url)

    if !isStaticPage(html) {
        html = fetchDynamicPage(url)
    }
    return html
}

// 静的HTMLを取得
func fetchStaticPage(url string) string {
    var html string
    c := colly.NewCollector()

    c.OnHTML("html", func(e *colly.HTMLElement) {
        html, _ = e.DOM.Html()
    })
    c.Visit(url)
    return html
}

// 動的ページ取得用ヘルパー
// WaitVisible を実行し、失敗しても無視するフォールバック
func tryWaitVisible(sel string) chromedp.Action {
    return chromedp.ActionFunc(func(ctx context.Context) error {
        err := chromedp.WaitVisible(sel, chromedp.ByQuery).Do(ctx)
        if err != nil {
            log.Printf("[TryWaitVisible fallback] selector=%s err=%v\n", sel, err)
            return nil
        }
        return nil
    })
}

// 動的ページ取得用ヘルパー
// WaitReady を実行し、失敗しても無視するフォールバック
func tryWaitReady(elem string) chromedp.ActionFunc {
  return chromedp.ActionFunc(func(ctx context.Context) error {
        // 失敗しても止めない
        _ = chromedp.WaitReady(elem, chromedp.ByQuery).Do(ctx)
        return nil
    })
}

// ブラウザクリックのフォールバック版
func tryClick(path string) chromedp.Action {
    return chromedp.ActionFunc(func(ctx context.Context) error {
        err := chromedp.Click(path, chromedp.NodeVisible).Do(ctx)
        if err != nil {
            log.Printf("[TryClick fallback] selector=%s err=%v\n", path, err)
            return nil
        }
        return nil
    })
}

// URLセットに追加（重複なし）
// true: 初visit, false: visit済
func isFirstVisit(mutex *sync.Mutex, url string, visited map[string]struct{}) bool {
    mutex.Lock()
    defer mutex.Unlock()

    _, exists := visited[url]

    if exists {
        return false
    }
    visited[url] = struct{}{} // struct{} = use memory 0
    log.Printf("visited: %v", url)

    return true
}

// 重複なしのURL配列を取得
func getDistinctUrls(visited map[string]struct{}) []string {
    urls  := make([]string, 0, 500)
    mutex := &sync.Mutex{}

    for k, _ := range visited {
        urls = utils.LockedAppend(mutex, urls, k)
    }
    return urls
}

// 詳細データが載っているページであるか判定
func isDetailPage(pattern string, url string) bool {
    matched, _ := regexp.MatchString(pattern, url)
    return matched
}

// 動的ページのレンダー(CSR/SSRに影響を受けない)
func renderHTML(ctx context.Context, startURL string, waitElem string) *goquery.Document {
    var html string

    // 一覧ページをレンダリング
    err := chromedp.Run(ctx,
        chromedp.Navigate(startURL),
        tryWaitVisible(waitElem), // 商品一覧の親
        chromedp.Sleep(1500 * time.Millisecond),    // JS描画待
        chromedp.OuterHTML("html", &html),
    )
    if err != nil {
        log.Printf("[Chromedp error]: %v", err)
    }
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

    if err != nil {
        log.Printf("[Document read error]: %v", err)
    }
    return doc
}

// link収集
func collectLinks(eachSelector string, doc *goquery.Document, cap int) []string {
    var links = make([]string, 0, cap)

    // 複数リンクを収集
    doc.Find(eachSelector).Each(func(idx int, selector *goquery.Selection) {
        link, _ := selector.Attr("href")
        if link != "" {
            links = append(links, link)
        }
    })
    return links
}

// 必要なリンクだけ取得
func getNeedLinks(links []string, needPattern string, cap int) []string {
    needLinks := make([]string, 0, cap)

    for _, link := range links {
        if strings.Contains(link, needPattern) {
            needLinks = append(needLinks, link)
        }
    }
    return needLinks
}

// 必要なリンクだけ取得
func toAbsLinks(links []string, prefix string, cap int) []string {
    absLinks := make([]string, 0, cap)

    for _, link := range links {
        absLinks = append(absLinks, prefix + link)
    }
    return absLinks
}