package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/http"
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

type guitarScraperFender struct {
    gScraper guitarScraper
}

type callBacksFender struct {
    funcs callBacks
}

func NewScraperFender(logger *log.Logger) Scraper {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &guitarScraperFender{
        guitarScraper{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksFender(logger *log.Logger) *callBacksFender {
    return &callBacksFender{
        callBacks{
            logger: logger,
        },
    }
}

type searchProductsResponse struct {
	Results []product `json:"results"`
}

type product struct {
    Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"imageUrl"`
}

var productsFender = make(map[string]*product, 550)

func (g *guitarScraperFender) CollectLinks(parentCtx context.Context) ([]string, error) {
    apiLinks := []string{
        // エレキ
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&resultsPerPage=50&bgfilter.ss_entry_type=product&bgfilter.collection_handle=electric-guitars&resultsFormat=native`,
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&resultsPerPage=50&bgfilter.ss_entry_type=product&bgfilter.collection_handle=electric-guitars&resultsFormat=native`,
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&page=2&resultsPerPage=50&bgfilter.ss_entry_type=product&bgfilter.collection_handle=electric-guitars&resultsFormat=native`,
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&page=3&resultsPerPage=50&bgfilter.ss_entry_type=product&bgfilter.collection_handle=electric-guitars&resultsFormat=native`,
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&page=4&resultsPerPage=50&bgfilter.ss_entry_type=product&bgfilter.collection_handle=electric-guitars&resultsFormat=native`,
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&page=5&resultsPerPage=50&bgfilter.ss_entry_type=product&bgfilter.collection_handle=electric-guitars&resultsFormat=native`,
        // シグネチャ
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&resultsPerPage=30&bgfilter.ss_entry_type=product&bgfilter.collection_handle=artist&resultsFormat=native`,
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&page=2&resultsPerPage=30&bgfilter.ss_entry_type=product&bgfilter.collection_handle=artist&resultsFormat=native`,
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&page=3&resultsPerPage=30&bgfilter.ss_entry_type=product&bgfilter.collection_handle=artist&resultsFormat=native`,
        // ベース
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&resultsPerPage=50&bgfilter.ss_entry_type=product&bgfilter.collection_handle=electric-basses&resultsFormat=native`,
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&page=2&resultsPerPage=50&bgfilter.ss_entry_type=product&bgfilter.collection_handle=electric-basses&resultsFormat=native`,
        // アコギ
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&resultsPerPage=30&bgfilter.ss_entry_type=product&bgfilter.collection_handle=acoustic-guitars&resultsFormat=native`,
        `https://hx9wc9.a.searchspring.io/api/search/search.json?siteId=hx9wc9&page=2&resultsPerPage=30&bgfilter.ss_entry_type=product&bgfilter.collection_handle=acoustic-guitars&resultsFormat=native`,
    }
    var productLinks = make([]string, 0, 550)

    for _, apiLink := range apiLinks {
        // 商品データ取得
        resp, err := http.Get(apiLink)

        if err != nil {
            log.Printf("[Http res error]: %v, %v\n", err, apiLink)
            continue
        }
        data, err := io.ReadAll((resp.Body))
        resp.Body.Close()

        if err != nil {
            log.Printf("[IO read error]: %v\n", err)
            continue
        }

        // 商品データから必要な部分だけ抽出
        var productsResponse searchProductsResponse
        err = json.Unmarshal(data, &productsResponse)

        if err != nil {
            log.Println("json parse error:", err)
        }
        for _, p := range productsResponse.Results {
            // リンク集積
            productLinks = append(productLinks, p.URL)
            // ギター集積 検索効率Upのため、mapを選択
            var product = map[string]*product{}
            product[utils.NormalizeString(p.Name)] = &p
            maps.Copy(productsFender, product)
        }
    }
    g.gScraper.urls = productLinks
    return g.gScraper.urls, nil
}

func (g *guitarScraperFender) Scrape(provider  PageProvider,
                                     parser    GuitarParser,
                                     parentCtx context.Context,
) []*model.Guitar {
    guitars := g.gScraper.scrapeFrame(provider, parser, parentCtx)

    // 画像URL merge
    for _, guitar := range guitars {
        product, exist := productsFender[utils.NormalizeString(guitar.Name)]
        if exist {
            guitar.Src = product.ImageURL
        }
    }
    return guitars
}

// 必要に応じて、基盤のTryWaitReadyを組み込む
func (c *callBacksFender) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://jp.fender.com/products/.+`, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 15*time.Second)
        defer cancel()

        var html string

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            // chromedp.WaitReady("#fender-react div div img", chromedp.ByQuery), // 求める要素が出るまで待つ
            chromedp.Sleep(10 * time.Second), // JSが動く猶予を与える
            chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )
        if err != nil {
            return "", fmt.Errorf("[chromedp error]: %v [url]: %v\n", err, url)
        }
        return html, nil
    }
}

func (c *callBacksFender) CollectSpec() func(doc *goquery.Document) []map[string]string {
    return func(doc *goquery.Document) []map[string]string {
        specs := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        spec := map[string]string{}

        spec[C.Maker]   = strconv.Itoa(C.Fender)
        spec[C.Name]    = doc.Find(`h1.product-title`).First().Text()
        spec[C.Color]   = doc.Find(`h4.product-specs-heading:contains("Color")`).Next().Text()
        spec[C.Comment] = doc.Find(`.product-content-description__text p`).Text()
        spec[C.Joint]   = ""
        spec[C.Price]   = doc.Find(`.price__regular .price-item`).First().Text()
        spec[C.Src]     = ""
        spec[C.Weight]  = strconv.Itoa(C.InvalidNumber)

        doc.Find(`.product-specs-wrap .product-specs-lists .product-specs-props`).Each(func(idx int, selector *goquery.Selection) {
            label      := selector.Find(`h4.product-specs-heading`).Text()
            elem       := selector.Find(`p.product-specs-content`).Text()
            field, _   := utils.ConvertLabel(label, specFieldMap)
            spec[field] = elem
        })
        specs = utils.LockedAppend(mutex, specs, spec)
        return specs
    }
}

func (c *callBacksFender) BuildGuitar(url string) func(spec map[string]string) *model.Guitar {
    return func(spec map[string]string) *model.Guitar {
        return buildGuitarFrame(spec, url, c.funcs.logger)
    }
}

func (c *callBacksFender) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "fender")
    }
}