package scraper

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

var (
    _httpClient = &http.Client{ Timeout: 5 * time.Second }
    _regGetDate = regexp.MustCompile(`\d{4}(/|-|年)\d{1,2}(/|-|月)\d{1,2}`)
)

type Scraper[T any] interface {
	Scrape(provider PageProvider, parser ModelParser[T], ctx context.Context) []T
	CollectLinks(ctx context.Context) ([]string, error)
}

type PageProvider interface {
    IsStaticPage()                        func(html string) bool
    FetchDynamicPage(ctx context.Context) func(url string)  (string, error)
}

type ModelParser[T any] interface {
    CollectAttributes()    func(doc *goquery.Document, url string) []map[string]string
    BuildModel(url string) func(spec map[string]string) T
}

type Crawler[T any] struct {
    urls      []string
	collector *colly.Collector
    mutex     *sync.Mutex
}

type CallBacks struct {}

// スクレイピング実行のフレームワーク
func (g *Crawler[T]) scrapeFrame(provider PageProvider,
                                 parser ModelParser[T],
                                 ctx context.Context,
) []T {
    var models = make([]T, 0, 400)

    if len(g.urls) <= 0 {
        log.Println("None URL for crawling...")
        return []T{}
    }
    wg := &sync.WaitGroup{}

    for _, url := range g.urls {
        url := url

        // アクセス間隔をずらし、bot感を薄める（ランダム待ち時間 + 最低待ち時間）
        delay := time.Duration(rand.Int63n(int64(2250 * time.Millisecond))) + 250 * time.Millisecond
        time.Sleep(delay)

        // 静的/動的を判定してHTMLを取得
        html := fetchPage(url, provider.IsStaticPage(), provider.FetchDynamicPage(ctx))

        if html == "" {
            continue
        }

        wg.Add(1)
        go func(html string, url string) {
            defer wg.Done()
            doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

            if err != nil {
                log.Println("[Goquery error]:", err)
                return
            }
            funcCollectAttributes := parser.CollectAttributes()
            funcBuildModel        := parser.BuildModel(url)
            attributes            := funcCollectAttributes(doc, url) // 1ページ：N詳細ページでもOK

            for _, attribute := range attributes {
                attribute := attribute
                model     := funcBuildModel(attribute)

				models = utils.LockedAppend(g.mutex, models, model)
            }
        }(html, url)
    }
    wg.Wait()
    return models
}

// 動的、静的ページを取得（動的が優先）。funcは個々で実装の必要あり。
func fetchPage(url string,
               isStaticPage func(string)bool,
               fetchDynamicPage func(string) (string, error),
) string {
    var html string

    html = fetchStaticPage(url)

    if !isStaticPage(html) {
        var err error
        html, err = fetchDynamicPage(url)

        if err != nil {
            log.Println(err)
        }
    }
    return html
}

// 静的HTMLを取得
func fetchStaticPage(url string) string {
    var html string

    c := colly.NewCollector(
        colly.UserAgent( // ブラウザからのアクセスのように振る舞う
            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
            "AppleWebKit/537.36 (KHTML, like Gecko) " +
            "Chrome/142.0.0.0 Safari/537.36",
    ),)

    c.SetRequestTimeout(10 * time.Second)

    c.OnHTML("html", func(e *colly.HTMLElement) {
        var err error
        html, err = e.DOM.Html()

        if err != nil {
            log.Printf("[fetchStaticPage failed]: %v", err)
        }
    })
    err := c.Visit(url)

    if err != nil {
        log.Printf("[Colly Visit failed] url=%s err=%v", url, err)
    }
    c.Wait()

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
func isFirstVisit(mutex *sync.Mutex, url string, visited map[string]struct{}) bool {
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
    matched, err := regexp.MatchString(pattern, url)

    if err != nil {
        return false
    }
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
func collectStatsCrawl (c *colly.Collector, stats *crawlStats) {
    c.OnRequest(func(c *colly.Request) {
        stats.requests.Add(1)
    })
    c.OnResponse(func(c *colly.Response) {
        stats.responses.Add(1)
    })
    c.OnError(func(c *colly.Response, err error) {
        stats.errors.Add(1)

        log.Printf(
            "[Crawl error]: status=%d url=%s err=%v\n",
            c.StatusCode,
            c.Request.URL,
            err,
        )
    })
}

// クロールのロギング
func loggingCrawlStats(name string, stats *crawlStats) {
    log.Printf(
        "[Crawl stats %v]: requests=%d responses=%d errors=%d\n",
        name,
        stats.requests.Load(),
        stats.responses.Load(),
        stats.errors.Load(),
    )
}

// apiから直接データを取得（関数外でClose()すること）
func fetchApiData(apiURL string) (*http.Response, error) {
    response, err := http.Get(apiURL)

    if err != nil {
        return nil, fmt.Errorf("failed to request job API: %w\n", err)
    }

    // 200系は成功扱いとする
    if response.StatusCode < 200 || response.StatusCode >= 300 {
        isShouldStopCrawler(response.StatusCode)

        return nil, fmt.Errorf(
            "Unexpected HTTP status: %d %v\n",
            response.StatusCode,
            apiURL,
        )
    }
    return response, nil
}

// return err: アクセス失敗、nil: アクセス成功
func checkHttpStatusOK(client *http.Client, url string) error {
    response, err := client.Get(url)

    if err != nil {
        return err
    }
    defer response.Body.Close()

    // 200系は成功扱いとする
    if response.StatusCode < 200 || response.StatusCode >= 300 {
        isShouldStopCrawler(response.StatusCode)

        return fmt.Errorf(
            "Unexpected HTTP status: %d, url=%s",
            response.StatusCode,
            url,
        )
    }
    return nil
}

// 場合によってはクロールを止める
func isShouldStopCrawler(httpStatus int) {
    if httpStatus == http.StatusForbidden {
        log.Panicf("Stop crawler. unexpected HTTP status: %v\n", httpStatus)
    }
}

// from: .envのPAGE_ID_FROM_..., to: .envのPAGE_ID_TO_...
func loadPageIdFromTo(envKeyFrom string, envKeyTo string) (int, int) {
    from := os.Getenv(envKeyFrom)
    to   := os.Getenv(envKeyTo)

    fromId, err:= strconv.Atoi(from)

    if err != nil {
        log.Panicf("From pageId parse error: %v", err)
    }
    toId, err := strconv.Atoi(to)

    if err != nil {
        log.Panicf("To pageId parse error: %v",err)
    }
    return fromId, toId
}

// 連番詳細ページIDの設定値が正しくなければ処理中止
func validatePageIdFromTo(fromId int, toId int) {
    if fromId > toId {
        log.Panicf(
            "連番pageIdの設定値は from <= to である必要があります。from: %v, to: %v\n",
            fromId,
            toId,
        )
    }

    if toId - fromId > 100000 {
        log.Panicf(
            "クロール対象(pageIdの範囲)は 10万件以下 に設定してください。from: %v, to: %v\n",
            fromId,
            toId,
        )
    }
}

// not found pageか調べる
func isNotFountPage(searchWord string, ctx context.Context) bool {
    // 確認は1.5秒間
    for i := 0; i < 15; i++ {
        var state int

        err := chromedp.Evaluate(fmt.Sprintf(`
            (() => {
                const body = document.body;

                if (!body) return 0;

                const notFoundWords = [
                    "404",
                    "見つかりません",
                    "募集終了",
                    "掲載終了",
                ];
                const text = body.innerText;

                if (notFoundWords.some(w => text.includes(w))) {
                    return 1;
                }

                if (text.includes(%q)) {
                    return -1;
                }
                return 0;
            })()`, searchWord), &state).Do(ctx)

        if err != nil { return false }

        switch state {
            case 1:
                return true // not found page だった
            case -1:
                return false
            case 0:
                // まだ判定できていない、ループ継続
        }
        time.Sleep(100 * time.Millisecond)
    }
    return false
}