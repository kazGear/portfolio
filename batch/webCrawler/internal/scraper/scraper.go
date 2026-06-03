package scraper

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type Scraper interface {
	Scrape()       ([]model.Guitar, error)
	CollectLinks() []string
    Cancel()
}

// ギター構造体の構築
func buildGuitar(spec map[string]string) (*model.Guitar) {

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
func fetchPage(url string,
			   isStaticPage func(string)bool,
			   fetchDynamicPage func(string)string) string {

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
func TryWaitReady(elem string) chromedp.ActionFunc {
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