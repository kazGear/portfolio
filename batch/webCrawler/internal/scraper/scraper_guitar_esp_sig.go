package scraper

// import (
// 	"context"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"

// 	"log"

// 	"github.com/PuerkitoBio/goquery"
// 	"github.com/chromedp/chromedp"
// 	"github.com/gocolly/colly/v2"
// 	"github.com/kazGear/portfolio/webCrawler/internal/model"
// 	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
// )

// type guitarScraperEspSig struct {
//     gScraper guitarScraper
// }

// type callBacksEspSig struct {
//     funcs callBacks
// }

// func NewEspSigScraper(ctx context.Context) Scraper {
// 	collector := colly.NewCollector(
// 		colly.Async(true),
// 		colly.MaxDepth(1),
// 	)
// 	collector.Limit(&colly.LimitRule{
// 		DomainGlob:  "*",
// 		Parallelism: 5,
// 	})
//     return &guitarScraperEspSig{
//         guitarScraper{
//             collector: collector,
//             mutex:     &sync.Mutex{},
//             ctx:       ctx,
//         },
//     }
// }

// func NewCallBacksEspSig(ctx context.Context) GuitarCallbacks {
//     return &callBacksEsp{
//         callBacks{
//             ctx: ctx,
//         },
//     }
// }

// func (e *guitarScraperEspSig) CollectLinks() *[]string {
//     c       := e.gScraper.collector
//     mutex   := &sync.Mutex{}
//     visited := make(map[string]struct{}, 100)

//     // URL収集、クロール
//     c.OnHTML(`.searchResultBlock.gallery_item .searchResultBlock_item a[href*="/artists/"]`, func(html *colly.HTMLElement) {
//         link := html.Request.AbsoluteURL(html.Attr("href"))
//         if isFirstVisit(mutex, link, visited) {
//             c.Visit(link)
//         }
//     })
//     c.Visit("https://espguitars.co.jp/signatureseries/")
//     c.Wait()

//     e.gScraper.urls = getDistinctUrls(visited)
//     return &e.gScraper.urls
// }

// func (e *guitarScraperEspSig) Scrape(funcs GuitarCallbacks) (*[]model.Guitar, error) {
//     guitars, _ := e.gScraper.scrapeFrame(funcs)
//     return guitars, nil
// }

// func (e *callBacksEspSig) FetchDynamicPage() func(url string) string {
//     return func(url string) string {
//         // タイムアウト（JS が遅いページ対策）
//         // e.funcs.ctx, _ = context.WithTimeout(e.funcs.ctx, 30*time.Second)

//         // 大本となるHTMLを取得
//         err := chromedp.Run(e.funcs.ctx,
//                chromedp.Navigate(url),
//                chromedp.WaitVisible("#main", chromedp.ByQuery), // 求める要素が出るまで待つ
//                chromedp.Sleep(400 * time.Millisecond), // JSが動く猶予を与える
//         )
//         if err != nil {
//             log.Println("chromedp error:", err)
//             return ""
//         }
//         // 必要な要素が生成されるのを待つ
//         _ = chromedp.Run(e.funcs.ctx,
//             tryWaitReady(".tab_detail"),

//         )
// //         chromedp.Poll(`
// //   () => document.querySelectorAll("table.tbl_spec tr").length >= 20
// // `, nil, chromedp.WithPollingInterval(100*time.Millisecond))

//         // 最終的なHTML出力
//         var html string
//         err = chromedp.Run(e.funcs.ctx,
//               chromedp.OuterHTML("html", &html, chromedp.ByQuery),
//         )
//         return html
//     }
// }

// func (e *callBacksEspSig) CollectSpec() func(doc *goquery.Document) map[string]string {
//     return func(doc *goquery.Document) map[string]string {
//         spec := map[string]string{}

//         spec["Maker"]   = strconv.Itoa(constants.EspSignature)
//         spec["Name"]    = strings.TrimSpace(doc.Find("h1.header_title").Text())
//         src, _         := doc.Find("#main .header_content img.transform-5").Attr("src")
//         spec["Src"]     = strings.TrimSpace(src)
//         spec["Color"]   = strings.TrimSpace(doc.Find(".header_content h3.clr_name").Text())
//         spec["Comment"] = strings.TrimSpace(doc.Find("#specialfeatures .container_small p").Text())
//         spec["Price"]   = strings.TrimSpace(doc.Find("p.detail_price").Text())

//         doc.Find("#specifications table.tbl_spec tr").Each(func(i int, s *goquery.Selection) {
//             th      := strings.TrimSpace(s.Find("th").Text())
//             td      := strings.TrimSpace(s.Find("td").Text())
//             th       = convertLabel(th)
//             spec[th] = td
//         })
//         return spec
//     }
// }
// // func SplitModels(doc *goquery.Document) ([]*goquery.Document, error) {
// //     var docs []*goquery.Document

// //     doc.Find("section.artist-model").Each(func(i int, s *goquery.Selection) {
// //         if len(s.Nodes) > 0 {
// //             sub := goquery.NewDocumentFromNode(s.Nodes[0])
// //             docs = append(docs, sub)
// //         }
// //     })

// //     return docs, nil
// // }
// func (e *callBacksEspSig) BuildGuitar() func(spec map[string]string) *model.Guitar {
//     return func(spec map[string]string) *model.Guitar {
//         return buildGuitarFrame(spec)
//     }
// }

// func (e *callBacksEspSig) IsStaticPage() func(html string) bool {
//     return func(html string) bool {
//         return strings.Contains(html, "tbl_spec")
//     }
// }

// // key: ESPの項目名, value: 構造体フィールド名
// var espSigFieldMap = map[string]string{
// 	"BODY":         "BodyMaterial",
// 	"NECK":         "NeckMaterial",
// 	"FINGERBOARD":  "Fingerboard",
// 	"BRIDGE":       "Bridge",
// 	"PICKUPS":      "Pickups",
// 	"CONTROLS":     "Controls",
// 	"Price":        "Price",
// 	"SCALE":        "ScaleLengthMM",
// 	"FRET":         "FretCount",
// 	"INLAY":        "Inlays",
// 	"CONSTRUCTION": "Joint",
// }

// // サイトの項目名をフィールド名に変換
// func convertLabel(label string) string {
//     return espSigFieldMap[label]
// }