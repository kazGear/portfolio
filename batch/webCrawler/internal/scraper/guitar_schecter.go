package scraper

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	C "github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type guitarScraperSchecter struct {
    gScraper guitarScraper
}

type callBacksSchecter struct {
    funcs callBacks
}

func NewScraperSchecter(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(4),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &guitarScraperSchecter{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksSchecter(logger *log.Logger) *callBacksSchecter {
    return &callBacksSchecter{
        callBacks{
            logger: logger,
        },
    }
}

var regNeedPatterSchecter = regexp.MustCompile(`\?variation=`)

func (g *guitarScraperSchecter) CollectLinks(parentCtx context.Context) ([]string, error) {
    c := g.gScraper.collector

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(c ,crawlStats, g.gScraper.logger)

    // URL収集、クロール
    visited := make(map[string]struct{}, 130)
    mutex   := &sync.Mutex{}

    c.OnHTML("#guitar a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML("#products a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })
    c.OnHTML("#main_visual aside a", func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        if g.gScraper.isFirstVisit(mutex, link, visited) {
            c.Visit(link)
        }
    })

    c.Visit("https://schecter.co.jp/guitar/")
    c.Wait()

    loggingCrawlStats(crawlStats, g.gScraper.logger)

    g.gScraper.urls = utils.MapToSliceUrl(visited)
    g.gScraper.urls = utils.GetNeedLinks(g.gScraper.urls, regNeedPatterSchecter, 130)
    return g.gScraper.urls, nil
}

func (g *guitarScraperSchecter) Scrape(provider  PageProvider,
                                  parser    GuitarParser,
                                  parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)
    return guitars
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksSchecter) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://schecter.co.jp/[a-z]+/\d{3,5}/\?variation=`, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 4*time.Second)
        defer cancel()

        var html string

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.WaitVisible("main", chromedp.ByQuery), // 求める要素が出るまで待つ
            chromedp.Sleep(300 * time.Millisecond), // JSが動く猶予を与える
            chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )
        if err != nil {
            return "", fmt.Errorf("[chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

var regEX_IV = regexp.MustCompile(`\[EX-IV\][\n.]+`)

func (c *callBacksSchecter) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec[C.Maker]      = strconv.Itoa(C.SCHECTER)
        spec[C.Name]       = doc.Find(`h1.name.order`).Text()
        spec[C.BodyFinish] = ""
        spec[C.Comment]    = ""
        spec[C.Inlays]     = ""
        src, _            := doc.Find(`.product_main_visual_image img`).Attr(`src`)
        spec[C.Src], _     = utils.ConvertRealUrl(src)
        spec[C.Weight]     = strconv.Itoa(C.InvalidNumber)

        doc.Find(`.product_spec_tbl tr`).Each(func(idx int, selector *goquery.Selection) {
            label      := strings.TrimSpace(selector.Find(`th`).Text())
            elem       := selector.Find(`td`).Text()
            field, _   := utils.ConvertLabel(label, specFieldMap)
            spec[field] = elem
        })
        spec[C.Color] = strings.ReplaceAll(doc.Find(`.product_main_visual_title h2.color`).Text(), " ", "")

        // 不要データ排除
        if (len(spec[C.Color])) <= 4 {
            return []map[string]string{}
        }

        // コメント収集
        var comment = strings.Builder{}
        doc.Find(`div.content`).Each(func(idx int, selector *goquery.Selection) {
            title      := selector.Find(`.title`).Text()
            detail     := selector.Find(`.text`).Text()
            comment.WriteString(fmt.Sprintf("%v\n%v\n", title, detail))
        })
        spec[C.Comment] = comment.String()

        // ボデ.Children()ィ材 整形
        bodyMaterial := spec[C.BodyMaterialBack]
        if strings.Contains(bodyMaterial, "&") { // top / back の形式
            topAndBack := strings.Split(bodyMaterial, "&")
            spec[C.BodyMaterialTop]  = topAndBack[0]
            spec[C.BodyMaterialBack] = topAndBack[1]
        } else if strings.Contains(bodyMaterial, ",") { // top / back の形式
            topAndBack := strings.Split(bodyMaterial, ",")
            spec[C.BodyMaterialTop]  = topAndBack[0]
            spec[C.BodyMaterialBack] = topAndBack[1]
        } else if strings.Contains(bodyMaterial, "/") { // back / top の形式
            backAndTop := strings.Split(bodyMaterial, "/")
            spec[ C.BodyMaterialBack] = backAndTop[0]
            spec[C.BodyMaterialTop]  = backAndTop[1]
        }

        // pickup 整形
        pickup := strings.TrimSpace(spec[C.Pickups])
        pickup  = regEX_IV.ReplaceAllString(pickup, "")
        pickup  = strings.ReplaceAll(pickup, "[EX-V]\n", "")
        pickup  = strings.ReplaceAll(pickup, " ", "")
        pickups := strings.Split(pickup, "\n")

        if len(pickups) <= 1 {
            spec[C.Pickups] = pickups[0]
        } else if len(pickups) == 2 {
            spec[C.Pickups] = pickups[0] + " / " + pickups[1]
        } else {
            spec[C.Pickups] = pickups[0] + " / " + pickups[1] + " / " + pickups[2]
        }

        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksSchecter) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksSchecter) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "Notes")
    }
}