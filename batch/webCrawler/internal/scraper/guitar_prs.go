package scraper

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"maps"
	"net/http"
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
	C "github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type guitarScraperPRS struct {
    gScraper guitarScraper
}

type callBacksPRS struct {
    funcs callBacks
}

func NewScraperPRS(logger *log.Logger) *guitarScraperPRS {
    collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &guitarScraperPRS{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksPRS(logger *log.Logger) *callBacksPRS {
    return &callBacksPRS{
        callBacks{
            logger: logger,
        },
    }
}

type priceData struct {
    name  string
    price string
}
var regItems = regexp.MustCompile(`(製|Product)[\s\S]+?csvUrl`)
var regPrice = regexp.MustCompile(`,(\d{3})`)
var regItemAndPrice = regexp.MustCompile(`(¥,?\d{3,7})+([A-Za-z]{1})`) // 例 ¥200,000Archon1x12ClosedBack

func getPrices(ctx context.Context) map[string]string {
    var baseUrl = `https://www.prsguitars.jp/products/plicelist`

    sheets := getPrsPriceSheets(ctx, baseUrl)

    // 価格データ収集
    var priceSet = map[string]string{}
    for _, node := range sheets {
        if node.NodeName != "IFRAME" {
            continue
        }
        csvStr := getRoughPriceData(baseUrl, node)

        if len(csvStr) <= 0 { continue }

        // データクレンジング
        csvStr = cleanData(csvStr)

        // csv変換
        loadCsv  := convertCsv(csvStr)
        flatData := createFlatData(loadCsv)

        // 価格表構築
        prices := searchGuitarPrices(flatData)
        maps.Copy(priceSet, prices)
    }
    return priceSet
}

func getPrsPriceSheets(ctx context.Context, baseUrl string) []*cdp.Node {
    var nodes []*cdp.Node

    // 価格表に繋がるデータを複数取得。直接はとれない
    err := chromedp.Run(ctx,
        chromedp.Navigate(baseUrl),
        autoScroll(),
        chromedp.WaitReady("body"),
        chromedp.Sleep(2000 * time.Millisecond),
        chromedp.Nodes(`iframe`, &nodes),
    )
    if err != nil {
        log.Printf("[getPrice chromedp error]: %+v", err)
    }
    return nodes
}

func getRoughPriceData(baseUrl string, node *cdp.Node) string {
    // アクセス先の設定
    iframeSrc := utils.GetAttr(node, "src")
    req, _    := http.NewRequest("GET", iframeSrc, nil)
    req.Header.Set("Referer", baseUrl) // 無いとforbiddenで弾かれる

    // 雑な価格データ取得
    client    := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        log.Printf("client do res error: %+v\n", err)
        return ""
    }
    body, _ := io.ReadAll(resp.Body)
    csvStr  := regItems.FindString(string(body))
    resp.Body.Close()

    return csvStr
}

func cleanData(csvStr string) string {
    csvStr = strings.ReplaceAll(csvStr, "\n", "")
    csvStr = strings.ReplaceAll(csvStr, "\\n", "")
    csvStr = strings.ReplaceAll(csvStr, "\t", "")
    csvStr = strings.ReplaceAll(csvStr, "\\t", "")
    csvStr = strings.ReplaceAll(csvStr, "\"", "")
    csvStr = strings.ReplaceAll(csvStr, "\\", "")
    csvStr = strings.ReplaceAll(csvStr, " ", "")
    csvStr = strings.ReplaceAll(csvStr, "  ", "")
    csvStr = strings.ReplaceAll(csvStr, "csvUrl", "")
    csvStr = strings.ReplaceAll(csvStr, "{", ",{")
    csvStr = regPrice.ReplaceAllString(csvStr, "$1") // (\d{3}) 部分
    csvStr = regItemAndPrice.ReplaceAllString(csvStr, "$1,$2") // (¥,?\d{3,7}), ([A-Za-z]{1}) 部分
    csvStr = strings.ReplaceAll(csvStr, ",,", ",")

    return csvStr
}

func convertCsv(csvStr string) [][]string {
    csv := csv.NewReader(strings.NewReader(csvStr))
    csv.LazyQuotes       = true
    csv.TrimLeadingSpace = true

    loadCsv, err := csv.ReadAll()

    if err != nil {
        log.Printf("[csv read error]: %+v", err)
    }
    return loadCsv
}

func createFlatData(loadCsv [][]string) []string {
    var flatData []string

    for _, outer := range loadCsv {
        for _, csv := range outer {
            flatData = append(flatData, csv)
        }
    }
    return flatData
}

func searchGuitarPrices(flatData []string) map[string]string {
    var prices = map[string]string{}

    // まず価格を探し、そこから商品名を探す
    for idx, data := range flatData {
        if !strings.HasPrefix(data, "¥") {
            continue
        }
        var price   = priceData{}
        price.price = data

        // 商品名を探す
        for backIdx := idx - 1; idx - backIdx <= 6; backIdx-- { // ６つ前まで走査
            if strings.Contains(flatData[backIdx], ":") {
                if strings.Contains(flatData[backIdx], "text") {
                    price.name = utils.NormalizeString(
                        strings.ReplaceAll(flatData[backIdx], "text:", ""),
                    )
                    break
                }
            } else {
                price.name = utils.NormalizeString(flatData[backIdx])
                break
            }
        }
        prices[price.name] = price.price
    }
    return prices
}

var regNeedPatterPrs = regexp.MustCompile(`https://www.prsguitars.jp/products/.+/.+`)

func (g *guitarScraperPRS) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(c ,crawlStats, g.gScraper.logger)

    // URL収集、クロール
    visited := make(map[string]struct{}, 150)
    mutex   := &sync.Mutex{}

    c.OnHTML("fluid-columns-repeater a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.Visit("https://www.prsguitars.jp/products")
    c.Wait()

    loggingCrawlStats(crawlStats, g.gScraper.logger)

    g.gScraper.urls = utils.MapToSliceUrl(visited)
    g.gScraper.urls = utils.GetNeedLinks(g.gScraper.urls, regNeedPatterPrs, 120)
    return g.gScraper.urls, nil
}

func (g *guitarScraperPRS) Scrape(provider  PageProvider,
                                  parser    GuitarParser,
                                  parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)

    ctx, cancel := context.WithTimeout(parentCtx, 60 * time.Second)
    defer cancel()

    priceSet := getPrices(ctx)

    mergePrice(guitars, priceSet)
    return guitars
}

func mergePrice(guitars []*model.Guitar, priceSet map[string]string) {
    for _, guitar := range guitars {
        guitarName     := utils.NormalizeString(guitar.Name)
        price          := priceSet[guitarName]
        guitar.Price, _ = utils.ParsePrice(price)
    }
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksPRS) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://www.prsguitars.jp/products/[\w-]+/[\w-]+`, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 10 * time.Second)
        defer cancel()

        var html string

        err := chromedp.Run(ctx,
               chromedp.Navigate(url),
               chromedp.WaitVisible(`//span[contains("Tuning")]`, chromedp.ByQuery), // 求める要素が出るまで待つ
               chromedp.Sleep(200 * time.Millisecond), // JSが動く猶予を与える
               chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )
        if err != nil {
            return "", fmt.Errorf("[chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

func (c *callBacksPRS) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec[C.Maker]   = strconv.Itoa(C.PRS)
        spec[C.Name]    = doc.Find(`h1 span span span span span`).First().Text()
        spec[C.Series]  = strings.SplitN(spec[C.Name], " ", 2)[0]
        spec[C.Comment] = doc.Find(`main section:nth-of-type(2)`).Text()
        /*
        スペック表の構造例
        <span>
            Body Construction : Solidbody<br>
            Body Wood : Poplar<br>
            Top Carve : Flat Top
        </span>
        */
        bodySection        := doc.Find(`p span:contains("Top Wood"), p span:contains("Body Wood")`).Text()
        neckSection        := doc.Find(`p span:contains("Fretboard")`).Text()
        jointSection       := doc.Find(`p span:contains("Assembly")`).Text()
        finishSection      := doc.Find(`p span:contains("Finish Type")`).Text()
        hardwareSection    := doc.Find(`p span:contains("Bridge")`).Text()
        electronicsSection := doc.Find(`p span:contains("Pickup")`).Text()

        spec = parseSpec(bodySection, spec)
        spec = parseSpec(neckSection, spec)
        spec = parseSpec(jointSection, spec)
        spec = parseSpec(finishSection, spec)
        spec = parseSpec(hardwareSection, spec)
        spec = parseSpec(electronicsSection, spec)

        if len(spec["TreblePickup"]) <= 0 {
            spec["TreblePickup"] = spec["BassPickup"]
        }

        trim := utils.TrimSpace()
        spec[C.Pickups] = fmt.Sprintf(
            "%v / %v / %v", trim(spec[C.NeckPickup]), trim(spec[C.CenterPickup]), trim(spec[C.BridgePickup]),
        )

        // 画像、カラー取得
        doc.Find(`span:contains("COLORS")`).
            Closest("div").Parent(). // 直近の親divの親
            Children().Next().Children().Children().Children(). // 各色のギターカード
                Each(func(idx int, selector *goquery.Selection) {
                    nextSpec := make(map[string]string)
                    maps.Copy(nextSpec, spec)

                    nextSpec[C.Src], _ = selector.Find("img").Attr("src")
                    nextSpec[C.Color]  = selector.Text()

                    specs = utils.LockedAppend(mutex, specs, nextSpec)
                })
        return specs
    }
}

func parseSpec(specSection string, spec map[string]string) map[string]string {
    specSection     = strings.ReplaceAll(specSection, "\r\n", "\n") // 改行正規化
    specSection     = "\n" + specSection + "\n"
    splitedSection := strings.Split(specSection, "\n")

    for _, elem := range splitedSection {
        if len(elem) == 0 {
            continue
        }
        labelAndSpec := strings.SplitN(elem, ":", 2)
        specLabel    := strings.TrimSpace(labelAndSpec[0])
        key, exist   := utils.ConvertLabel(specLabel, specFieldMap)

        if exist && len(labelAndSpec) > 1 {
            spec[key] = labelAndSpec[1]
        }
    }
    return spec
}

func (c *callBacksPRS) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksPRS) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "Tuning")
    }
}