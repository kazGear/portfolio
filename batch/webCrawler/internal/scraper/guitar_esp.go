package scraper

import (
	"context"
	"errors"
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
type espGuitarScraper struct {
	collector *colly.Collector
    urls      []string
    mutex     *sync.Mutex
    ctx       context.Context
    cancel    context.CancelFunc
}

func NewESPGuitarScraper() Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(4),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 10,
	})
    ctx, cancel := chromedp.NewContext(context.Background())

    return &espGuitarScraper{
		collector: collector,
        mutex:     &sync.Mutex{},
        ctx:       ctx,
        cancel:    cancel,
	}
}

// chromeプロセスのターミネート
func (e *espGuitarScraper) Cancel() {
    e.cancel()
}

// スクレイプ対象のURLを収集
func (e *espGuitarScraper) CollectURLs() /*[]string*/ {
    c       := e.collector
    visited := make(map[string]struct{}, 500)
    mutex   := &sync.Mutex{}

    // URL収集、クロール
    c.OnHTML("#item .figcap a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML("#inner_content .figcap a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML("section.color_variation a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.Visit("https://espguitars.co.jp/products/esp")
    c.Wait()

    e.urls = getDistinctUrls(visited)
}

// スクレイピング実行
func (e *espGuitarScraper) Scrape() ([]model.Guitar, error) {
    var guitars []model.Guitar

    if len(e.urls) <= 0 {
        return []model.Guitar{}, errors.New("巡回用URLがありません。")
    }
    for _, url := range e.urls {
        // 静的/動的を判定して HTML を取得
        html := fetchPage(url, isStaticPage, e.fetchDynamicPage)
        // goquery >>> DOM化
        doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
        if err != nil {
            log.Println("goquery error:", err)
            continue
        }
        spec   := collectSpec(doc)
        guitar := buildGuitar(spec)

        if len(guitar.Name) <= 0 || len(guitar.Color) <= 0 { continue  }
        guitars = utils.LockedAppend(e.mutex, guitars, *guitar)
    }
    return guitars, nil
}

// 動的ページ（JS実行後、処理重め）のHTMLを取得
// 必要に応じて、scraper.go TryWaitReadyを組み込む
func (e *espGuitarScraper) fetchDynamicPage(url string) string {
    // タイムアウト（JS が遅いページ対策）
    e.ctx, _ = context.WithTimeout(e.ctx, 30*time.Second)

    // 大本となるHTMLを取得
    err := chromedp.Run(e.ctx,
           chromedp.Navigate(url),
           chromedp.WaitVisible("#main", chromedp.ByQuery), // 求める要素が出るまで待つ
	       chromedp.Sleep(250 * time.Millisecond), // JSが動く猶予を与える
    )
    if err != nil {
        log.Println("chromedp error:", err)
		return ""
    }
    // 必要な要素が生成されるのを待つ
    _ = chromedp.Run(e.ctx,
        TryWaitReady("h1.header_title"),
        TryWaitReady(".tbl_spec"),
        TryWaitReady("p.detail_price"),
    )
    // 最終的なHTML出力
    var html string
    err = chromedp.Run(e.ctx,
          chromedp.OuterHTML("html", &html, chromedp.ByQuery),
    )
	return html
}

// ギターのスペック情報を収集
func collectSpec(doc *goquery.Document) map[string]string {
    spec := map[string]string{}

    spec["Maker"]   = strconv.Itoa(constants.ESP)
    spec["Name"]    = strings.TrimSpace(doc.Find("h1.header_title").Text())
    src, _         := doc.Find("#main .header_content img.transform-5").Attr("src")
    spec["Src"]     = strings.TrimSpace(src)
    spec["Color"]   = strings.TrimSpace(doc.Find(".header_content h3.clr_name").Text())
    spec["Comment"] = strings.TrimSpace(doc.Find("#specialfeatures .container_small p").Text())
    spec["Price"]   = strings.TrimSpace(doc.Find("p.detail_price").Text())

    doc.Find("#specifications table.tbl_spec tr").Each(func(i int, s *goquery.Selection) {
        th      := strings.TrimSpace(s.Find("th").Text())
        td      := strings.TrimSpace(s.Find("td").Text())
        th       = convertLabel(th)
        spec[th] = td
    })
    return spec
}

// 静的ページであるか判断
func isStaticPage(html string) bool {
    return strings.Contains(html, "tbl_spec")
}

// key: ESPの項目名, value: 構造体フィールド名
var espFieldMap = map[string]string{
	"BODY":         "BodyMaterial",
	"NECK":         "NeckMaterial",
	"FINGERBOARD":  "Fingerboard",
	"BRIDGE":       "Bridge",
	"PICKUPS":      "Pickups",
	"CONTROLS":     "Controls",
	"Price":        "Price",
	"SCALE":        "ScaleLengthMM",
	"FRET":         "FretCount",
	"INLAY":        "Inlays",
	"CONSTRUCTION": "Joint",
}

// サイトの項目名をフィールド名に変換
func convertLabel(label string) string {
    return espFieldMap[label]
}