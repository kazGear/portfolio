package scraper

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	C "github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type Scraper interface {
	Scrape(provider PageProvider, parser GuitarParser, ctx context.Context) []*model.Guitar
	CollectLinks(ctx context.Context)                                       ([]string, error)
}

type PageProvider interface {
    IsStaticPage() func(html string) bool
    FetchDynamicPage(ctx context.Context) func(url string) (string, error)
}

type GuitarParser interface {
    CollectSpec()           func(doc *goquery.Document)  []map[string]string
    BuildGuitar(url string) func(spec map[string]string) *model.Guitar
}

type guitarScraper struct {
    urls      []string
	collector *colly.Collector
    mutex     *sync.Mutex
    logger    *log.Logger
}

type callBacks struct {
    logger *log.Logger
}

// スクレイピング実行のフレームワーク
func (g *guitarScraper) scrapeFrame(provider PageProvider,
                                    parser GuitarParser,
                                    ctx context.Context,
) []*model.Guitar {
    var guitars = make([]*model.Guitar, 0, 400)

    if len(g.urls) <= 0 {
        g.logger.Println("None URL for crawling...")
        return []*model.Guitar{}
    }
    utils.LoggingCollectedLinks(g.urls, g.logger)
    g.logger.Printf("[Urls count]: %v 件\n", len(g.urls))

    wg := &sync.WaitGroup{}

    for _, url := range g.urls {
        url := url
        // 静的/動的を判定してHTMLを取得
        html := g.fetchPage(url, provider.IsStaticPage(), provider.FetchDynamicPage(ctx))

        wg.Add(1)
        go func(html string, url string) {
            defer wg.Done()
            doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

            if err != nil {
                g.logger.Println("[Goquery error]:", err)
                return
            }
            collectSpec := parser.CollectSpec()
            buildGuitar := parser.BuildGuitar(url)
            specs       := collectSpec(doc) // 1ページ：N詳細ページでもOK

            for _, spec := range specs {
                spec := spec
                guitar := buildGuitar(spec)

                if len(guitar.Name) <= 0 || len(guitar.Color) <= 0 {
                    g.logger.Printf("[Skip guitar]: %v\n", guitar)
                    continue
                }
                guitars = utils.LockedAppend(g.mutex, guitars, guitar)
            }
        }(html, url)
    }
    wg.Wait()
    return guitars
}

// ギター構造体の構築フレームワーク
func buildGuitarFrame(spec map[string]string, url string, logger *log.Logger) (*model.Guitar) {
	guitar := model.Guitar{}
    trim   := utils.TrimSpace()

    var errMaker error
	guitar.Maker, errMaker = strconv.Atoi(spec[C.Maker])
	guitar.Name            = trim(spec[C.Name])

    if errMaker != nil {
        logger.Printf("[Maker convert error]: %v", errMaker)
        return &model.Guitar{}
	}
	guitar.BodyFinish        = trim(spec[C.BodyFinish])
	guitar.BodyMaterial      = trim(spec[C.BodyMaterialTop]) + " / " + trim(spec[C.BodyMaterialBack])
    guitar.BodyMaterialBack  = searchWoodCode(spec[C.BodyMaterialBack])
	guitar.BodyMaterialTop   = searchWoodCode(spec[C.BodyMaterialTop])
    guitar.Bridge            = trim(spec[C.Bridge])
	guitar.Color             = trim(spec[C.Color])
    guitar.ColorCd           = utils.ConvertColorCd(guitar.Color)
	guitar.Comment           = trim(spec[C.Comment])
	guitar.Controls          = trim(spec[C.Controls])
    guitar.Fingerboard       = searchWoodCode(spec[C.Fingerboard])

	var errFretCount error
	fretCount                     := trim(spec[C.FretCount])
    guitar.FretCount, errFretCount = utils.GetFretCount(fretCount)
    if errFretCount != nil {
        // logger.Println(errFretCount)
    }
	guitar.Inlays       = trim(spec[C.Inlays])
	guitar.Joint        = trim(spec[C.Joint])
    guitar.NeckMaterial = searchWoodCode(spec[C.NeckMaterial])
    guitar.Pickups      = trim(spec[C.Pickups])

    var errPrice error
	guitar.Price, errPrice = utils.ParsePrice(spec[C.Price])

    if errPrice != nil {
        // logger.Println(errPrice)
    }
    scaleLengthMM       := trim(spec[C.ScaleLengthMM])
    guitar.ScaleLengthMM = int(utils.ParseScale(scaleLengthMM))
	guitar.Series        = trim(spec[C.Series])

    guitar.Src           = trim(spec[C.Src])
    // 画像の相対パスをフルパスへ
    if strings.HasPrefix(guitar.Src, "/") {
        fullPass, err := utils.CreateImagePath(url, guitar.Src)

        if err != nil {
            logger.Println(err)
        }
        guitar.Src = fullPass
    }

    var errWeight error
    guitar.Weight, errWeight = utils.ParseWight(trim(spec[C.Weight]))

    if errWeight != nil {
        // logger.Println(errWeight)
    }
	return &guitar
}

// 動的、静的ページを取得（動的が優先）。funcは個々で実装の必要あり。
func (g *guitarScraper) fetchPage(url string,
                                  isStaticPage func(string)bool,
                                  fetchDynamicPage func(string) (string, error),
) string {
    var html string
    html = g.fetchStaticPage(url)

    if !isStaticPage(html) {
        var err error
        html, err = fetchDynamicPage(url)

        if err != nil {
            g.logger.Println(err)
        }
    }
    return html
}

// 静的HTMLを取得
func (g *guitarScraper) fetchStaticPage(url string) string {
    var html string
    c := colly.NewCollector()

    c.OnHTML("html", func(e *colly.HTMLElement) {
        var err error
        html, err = e.DOM.Html()

        if err != nil {
            g.logger.Printf("[fetchStaticPage failed]: %v", err)
        }
    })
    c.Visit(url)
    return html
}

// 動的ページ取得用ヘルパー
// WaitVisible を実行し、失敗しても無視するフォールバック
func tryWaitVisible(selector string) chromedp.Action {
    return chromedp.ActionFunc(func(ctx context.Context) error {
        err := chromedp.WaitVisible(selector, chromedp.ByQuery).Do(ctx)
        if err != nil {
            log.Printf("[TryWaitVisible fallback]: selector=%s err=%v\n", selector, err)
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
        err := chromedp.WaitReady(elem, chromedp.ByQuery).Do(ctx)

        if err != nil {
            log.Printf("[TryWaitReady fallback]: elem=%v err=%v\n", elem, err)
        }
        return nil
    })
}

// ブラウザクリックのフォールバック版
func tryClick(path string) chromedp.Action {
    return chromedp.ActionFunc(func(ctx context.Context) error {
        err := chromedp.Click(path, chromedp.NodeVisible).Do(ctx)
        if err != nil {
            log.Printf("[TryClick fallback]: selector=%s err=%v\n", path, err)
            return nil
        }
        return nil
    })
}

// URLセットに追加（重複なし）
// true: 初visit, false: visit済
func (g *guitarScraper) isFirstVisit(mutex *sync.Mutex, url string, visited map[string]struct{}) bool {
    mutex.Lock()
    defer mutex.Unlock()

    _, exists := visited[url]

    if exists {
        return false
    }
    visited[url] = struct{}{} // struct{} = use memory 0
    return true
}

// 詳細データが載っているページであるか判定
func isDetailPage(pattern string, url string) bool {
    matched, _ := regexp.MatchString(pattern, url)
    return matched
}

// 動的ページのレンダー(CSR/SSRに影響を受けない)
func renderHTML(ctx context.Context, startURL string, waitElem string,
) (*goquery.Document, error) {

    var html string

    // 一覧ページをレンダリング
    err := chromedp.Run(ctx,
        chromedp.Navigate(startURL),
        tryWaitVisible(waitElem), // 商品一覧の親
        chromedp.Sleep(2000 * time.Millisecond), // JS描画待
        chromedp.OuterHTML("html", &html),
    )
    if err != nil {
        return nil, fmt.Errorf("[Chromedp error]: %v %v\n", err, waitElem)
    }
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

    if err != nil {
        return nil, fmt.Errorf("[Document read error]: %v %v\n", err, waitElem)
    }
    return doc, nil
}

// htmlを自動でスクロールさせる
func autoScroll() chromedp.Action {
    return chromedp.ActionFunc(func(ctx context.Context) error {
        var lastHeight int
        var scrollY int
        var innerHeight int

        for i := 0; i < 50; i++ {
            chromedp.Run(ctx,
                chromedp.Evaluate(`document.body.scrollHeight`, &lastHeight), // ページ全体の高さ
            )
            chromedp.Run(ctx,
                chromedp.Evaluate(`window.scrollBy(0, 800)`, nil),
                chromedp.Sleep(300 * time.Millisecond),
                chromedp.Evaluate(`window.scrollY`, &scrollY), // 画面上端のスクロール位置
                chromedp.Evaluate(`window.innerHeight`, &innerHeight),
            )
            if scrollY + innerHeight >= lastHeight - 50 {
                break // 最後までスクロール済
            }
            time.Sleep(5 * time.Second)
        }
        return nil
    })
}

type crawlStats struct {
    requests  atomic.Int64
    responses atomic.Int64
    errors    atomic.Int64
}

// クロールの req,res,err の数を集計する
func statsCrawlLogs (c *colly.Collector,
                     stats *crawlStats,
                     logger *log.Logger,
) {
    c.OnRequest(func(c *colly.Request) {
        stats.requests.Add(1)
    })
    c.OnResponse(func(c *colly.Response) {
        stats.responses.Add(1)
    })
    c.OnError(func(c *colly.Response, err error) {
        stats.errors.Add(1)

        logger.Printf(
            "[Crawl error]: status=%d url=%s err=%v\n",
            c.StatusCode,
            c.Request.URL,
            err,
        )
    })
}

// クロールのロギング
func loggingCrawlStats(stats *crawlStats, logger *log.Logger) {
    logger.Printf(
        "[Crawl stats]: requests=%d responses=%d errors=%d\n",
        stats.requests.Load(),
        stats.responses.Load(),
        stats.errors.Load(),
    )
}

var specFieldMap = map[string]string{
	"Top":                     C.BodyMaterialTop,
	"Top Wood":                C.BodyMaterialTop,
    "Body Top":                C.BodyMaterialTop,

    "Back Wood":               C.BodyMaterialBack,
	"Body":                    C.BodyMaterialBack,
	"BODY":                    C.BodyMaterialBack,
	"Body Material":           C.BodyMaterialBack,
	"Body Wood":               C.BodyMaterialBack,
    "Body Back":               C.BodyMaterialBack,
    "Back & Sides":            C.BodyMaterialBack,

	"Finish":                  C.BodyFinish,
	"Finish Type":             C.BodyFinish,
    "Body Finish":             C.BodyFinish,

	"Bridge":                  C.Bridge,
	"BRIDGE":                  C.Bridge,

	"COLOR":                   C.Color,
    "Color":                   C.Color,
    "Body Color":              C.Color,

	"Controls":                C.Controls,
	"CONTROLS":                C.Controls,
    "CONTROL":                 C.Controls,

	"Fingerboard Material":    C.Fingerboard,
	"FINGERBOARD":             C.Fingerboard,
    "FINGER BOARD":            C.Fingerboard,
	"Fretboard Wood":          C.Fingerboard,
    "Fingerboard":             C.Fingerboard,
    "Fingerboard & Bridge":    C.Fingerboard,

	"FRET":                    C.FretCount,
	"FRETS":                   C.FretCount,
    "Frets":                   C.FretCount,
	"Number Of Frets":         C.FretCount,
	"Number of Frets":         C.FretCount,

	"INLAY":                   C.Inlays,
	"Inlays":                  C.Inlays,
	"Fretboard Inlay":         C.Inlays,
    "Position Inlays":         C.Inlays,
    "Fret Markers":            C.Inlays,

	"CONSTRUCTION":            C.Joint,
	"Neck Joint":              C.Joint,
    "Joint":                   C.Joint,
    "JOINT":                   C.Joint,
	"Neck/Body Assembly Type": C.Joint,

	"Material":                C.NeckMaterial,
	"NECK":                    C.NeckMaterial,
    "Neck":                    C.NeckMaterial,
	"Neck Wood":               C.NeckMaterial,
    "Neck Material":           C.NeckMaterial,

	"PICKUPS":                 C.Pickups,
    "Pickups":                 C.Pickups,
    "Pickup":                  C.Pickups,
    "Pickup(Neck, Middle, Bridge)":C.Pickups,
	"Bass Pickup":             C.NeckPickup,
    "Neck Pickup":             C.NeckPickup,
    "Neck P ickup":            C.NeckPickup,
	"Middle Pickup":           C.CenterPickup,
	"Treble Pickup":           C.BridgePickup,
    "Bridge Pickup":           C.BridgePickup,

	"Price":                   C.Price,
    "PRICE":                   C.Price,

	"SCALE":                   C.ScaleLengthMM,
    "Scale":                   C.ScaleLengthMM,
	"Scale Length":            C.ScaleLengthMM,

    "Series":                  C.Series,
}

var regWood = regexp.MustCompile(`\s+`)
// 木材コードを探しだす
func searchWoodCode(s string) int {
	trimed := regWood.ReplaceAllString(s, "")

	for _, wood := range GetWoods() {
		if strings.Contains(strings.ToLower(trimed), strings.ToLower(wood.Name)) {
			return wood.Code
		}
	}
	return 0 // 該当なし
}
type wood struct {
	Name string
	Code int
}

// wood materials
func GetWoods() []wood {
	woods := []wood{
		{"Unknown", 0},
		{"HardMaple", 1},
		{"FlameMaple", 2},
		{"FlamedMaple", 2},
		{"QuiltedMaple", 3},
		{"BirdseyeMaple", 4},
		{"RoastedMaple", 5},
		{"Maple", 6},
		{"HonduranMahogany", 7},
		{"Mahogany", 8},
		{"Sapele", 9},
		{"Korina", 10},
		{"WhiteKorina", 11},
		{"Alder", 12},
		{"Ash", 13},
		{"Basswood", 14},
		{"Linden", 14},
		{"Poplar", 15},
		{"Spruce", 16},
		{"Cedar", 17},
		{"IndianRosewood", 18},
		{"BrazilianRosewood", 19},
		{"Rosewood", 20},
		{"PauFerro", 21},
		{"Ovangkol", 22},
		{"Ebony", 23},
		{"Walnut", 24},
		{"Padauk", 25},
		{"Koa", 26},
		{"Nato", 27},
		{"Agathis", 28},
		{"Bubinga", 29},
		{"Wenge", 30},
		{"Purpleheart", 31},
		{"Zebrawood", 32},
		{"Okoume", 33},
		{"Meranti", 34},
		{"Sakura", 35},
		{"Tochi", 36},
	}
	return woods
}