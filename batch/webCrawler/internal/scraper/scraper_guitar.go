package scraper

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type Scraper interface {
	Scrape(GuitarCallbacks) ([]model.Guitar, error)
	CollectLinks()                []string
    Cancel()
}

type GuitarCallbacks interface {
    FetchDynamicPage() func(url string) string
    CollectSpec()      func(doc *goquery.Document) map[string]string
    BuildGuitar()      func(spec map[string]string) *model.Guitar
}

type guitarScraper struct {
    urls      []string
	collector *colly.Collector
    mutex     *sync.Mutex
    cancel    context.CancelFunc
}

type callBacks struct {
    ctx context.Context
}

func NewCallBacks(ctx context.Context) GuitarCallbacks {
    return &callBacks{
        ctx: ctx,
    }
}

// スクレイピング実行
func (e *guitarScraper) scrapeFrame(funcs GuitarCallbacks) ([]model.Guitar, error) {
    var guitars []model.Guitar

    if len(e.urls) <= 0 {
        return []model.Guitar{}, errors.New("巡回用URLがありません。")
    }
    for _, url := range e.urls {
        // 静的/動的を判定して HTML を取得
        html := fetchPage(url, isStaticPage, funcs.FetchDynamicPage())
        // goquery >>> DOM化
        doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
        if err != nil {
            log.Println("goquery error:", err)
            continue
        }
        collectSpec := funcs.CollectSpec()
        buildGuitar := funcs.BuildGuitar()
        spec        := collectSpec(doc)
        guitar      := buildGuitar(spec)

        if len(guitar.Name) <= 0 || len(guitar.Color) <= 0 { continue  }
        guitars = utils.LockedAppend(e.mutex, guitars, *guitar)
    }
    return guitars, nil
}


// ギター構造体の構築
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
	guitar.BodyMaterialFront = utils.SearchWoodCode(spec["BodyMaterialFront"])
    guitar.Bridge            = spec["Bridge"]
	guitar.Color             = spec["Color"]
    guitar.ColorCd           = constants.InvalidNumber //strconv.Atoi(spec["ColorCd"])
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
    weight, _    := strconv.Atoi(spec["Weight"])
	guitar.Weight = float64(weight)

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

// HTMLを取得
func fetchStaticPage(url string) string {
    var html string
    c := colly.NewCollector()

    c.OnHTML("html", func(e *colly.HTMLElement) {
        html, _ = e.DOM.Html()
    })
    c.Visit(url)
    return html
}

// WaitReady を実行し、失敗しても無視するフォールバック
func tryWaitReady(elem string) chromedp.ActionFunc {
    return func(ctx context.Context) error {
        _ = chromedp.Run(ctx,
			chromedp.Sleep(200 * time.Millisecond),
            chromedp.WaitReady(elem, chromedp.ByQuery),
        )
        return nil
    }
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