package scraper

import (
	"strconv"
	"sync"

	"log"

	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
)

type espGuitarScraper struct {
	collector *colly.Collector
	html      *colly.HTMLElement
}

func NewESPGuitarScraper() Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 10,
	})
	return &espGuitarScraper{
		collector: collector,
	}
}

// スクレイピング実行
func (e *espGuitarScraper) Scrape() ([]model.Guitar, error) {
	guitars := make([]model.Guitar, 0, 500)
    mutex   := &sync.Mutex{}

	// リンククロール
	e.collector.OnHTML("#item", navigateLinks(".item_guitar")) // ギターシリーズ一覧
	e.collector.OnHTML("div.items div .fig", navigateLinks(".wrap")) // ギター一覧
	e.collector.OnHTML("div.items", navigateLinks(".fig .figcap")) // ギター一覧
	// ギター詳細、データ収集
	e.collector.OnHTML("#main", func(html *colly.HTMLElement) {
		guitar := buildGuitar(collectSpec(html))

        mutex.Lock()
        guitars = append(guitars, *guitar)
        mutex.Unlock()
	})

	e.collector.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	e.collector.Visit("https://espguitars.co.jp/products/esp")
	e.collector.Wait()

	return guitars, nil
}

// ギターのスペック情報を収集
func collectSpec(html *colly.HTMLElement) map[string]string {
    const skip = "99"

    spec := map[string]string{
        "Maker":             strconv.Itoa(constants.ESP),
        "Name":              html.ChildText("div.overlay_bgr_header div.header_content h1.header_title"),
        "BodyFinish":        "",
        "BodyMaterial":      html.ChildText("#specifications table:nth-child(2) tr:nth-child(1) td"),
        "BodyMaterialFront": skip,
        "BodyMaterialBack":  skip,
        "Bridge":            html.ChildText("#specifications table:nth-child(2) tr:nth-child(12) td"),
        "Color":             "",
        "Controls":          html.ChildText("#specifications table:nth-child(2) tr:nth-child(15) td"),
        "Comment":           html.ChildText("#specialfeatures div:nth-child(2) .content p"),
        "Fingerboard":       html.ChildText("#specifications table:nth-child(2) tr:nth-child(4) td"),
        "FretCount":         html.ChildText("#specifications table:nth-child(2) tr:nth-child(9) td"),
        "Inlays":            html.ChildText("#specifications table:nth-child(2) tr:nth-child(8) td"),
        "Joint":             html.ChildText("#specifications table:nth-child(2) tr:nth-child(10) td"),
        "NeckMaterial":      html.ChildText("#specifications table:nth-child(2) tr:nth-child(2) td"),
        "Pickups":           html.ChildText("#specifications table:nth-child(2) tr:nth-child(13) td"),
        "Price":             html.ChildText("#specifications div:nth-child(2) p.detail_price"),
        "ScaleLengthMM":     html.ChildText("#specifications table:nth-child(2) tr:nth-child(6) td"),
        "Series":            html.ChildText(".header div:nth-child(1) div:nth-child(1) h2.subltitle"),
        "Src":               html.ChildAttr(".header .header_content div:nth-child(1) img", "src"),
        "Weight":            strconv.Itoa(constants.InvalidNumber),
    }
    return spec
}
