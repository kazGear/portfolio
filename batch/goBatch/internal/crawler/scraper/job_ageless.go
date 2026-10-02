package scraper

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/repository"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/kazGear/portfolio/goBatch/pkg/db"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

type CrawlerAgeless struct {
    name     string
    jScraper Crawler[*model.Job]
}

type CallBacksAgeless struct {
    funcs CallBacks
}

func NewScraperAgeless() Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(3),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
        Delay:       500 * time.Millisecond,
        RandomDelay: 1500 * time.Millisecond,
	})
    return &CrawlerAgeless{
        "AGELESS",
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksAgeless() *CallBacksAgeless {
    return &CallBacksAgeless{
        CallBacks{},
    }
}

func (c *CrawlerAgeless) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector := c.jScraper.collector
    mutex     := c.jScraper.mutex
    visited   := make(map[string]struct{}, 1000)

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(collector ,crawlStats)

    // 詳細ページ収集
    collector.OnHTML(`.card-project a[href*="/projects/"]`, func(html *colly.HTMLElement) {
        link := html.Request.AbsoluteURL(html.Attr("href"))
        utils.LockedAddSet(mutex, visited, link)
    })

    // ページネーション
    for i := 1; i < C.PaginationLimit; i++ {
        collector.Visit(fmt.Sprintf(`https://freelance.ageless.co.jp/projects/search?page=%v`, i))
    }
    collector.Wait()

    // 保存済ページID取得
    repository := repository.NewJobRepository(db.GetInstance())
    savedUrls  := repository.Select(c.name)
    log.Printf("%v savedUrls: %v件\n", c.name, len(savedUrls))

    // 保存されていない案件urlだけを残す
    needUrls := make(map[string]struct{}, 1000)
    for url := range visited {
        if _, exist := savedUrls[url]; !exist {
            utils.LockedAddSet(mutex, needUrls, url)
        }
    }

    loggingCrawlStats(c.name, crawlStats)

    c.jScraper.urls = utils.MapToSliceUrl(needUrls)

    return c.jScraper.urls, nil
}

func (c *CrawlerAgeless) Scrape(provider  PageProvider,
                                parser    ModelParser[*model.Job],
                                parentCtx context.Context,
) []*model.Job {
    jobs := c.jScraper.scrapeFrame(provider, parser, parentCtx)
    return jobs
}

func (c *CallBacksAgeless) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        // 動的ページを取得しない場合、引数のパターンは記載しないで良い
        if !isDetailPage(``, url) {
            return "", nil
        }
        // 無駄なchromedpの起動を回避
        if err := checkHttpStatusOK(_httpClient, url); err != nil {
            return "", err
        }

        // タブごとに独立した context を作る
        tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        defer tabCancel()
        // // タブにだけ timeout を付ける
        ctx, cancel := context.WithTimeout(tabCtx, 10 * time.Second)
        defer cancel()

        var html string

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.WaitReady(".job-btn", chromedp.ByQuery), // 求める要素が出るまで待つ
            chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )

        if err != nil {
            log.Printf("Chromedp error: %v", err)
            return "", err
        }
        return html, nil
    }
}

var _regDeleteDescriptionAgeLess = regexp.MustCompile(`\{.*\}`)

func (c *CallBacksAgeless) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)

        description           := doc.Find("main").Text()
        description            = _regDeleteDescriptionAgeLess.ReplaceAllString(description, "")
        normalizedDescription := normalizeForSearchFeatures(description)

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesAgeless(normalizedDescription)
        // 保存するべき案件か
        if len(features) <= 0 {
            return []map[string]string{}
        }
        repository.InjectionJobFeatures(features, url)

        // 案件のオプションを収集し、repositoryへ（このサイトは無し）


        data := map[string]string{}

        data[C.Url]         = url
        data[C.Title]       = doc.Find(".project-card-ttl").Text()
        data[C.Location]    = salvageLocation(normalizedDescription)

        data[C.MinSalaryAtMonth] = doc.Find(".income-num").Text()
        data[C.MaxSalaryAtMonth] = doc.Find(".income-num").Text()

        data[C.Description]    = description
        data[C.EmploymentType] = salvageEmploymentType(normalizedDescription)
        data[C.WorkPlace]      = salvageWorkPlace(normalizedDescription)
        data[C.SourceSite]     = C.AGELESS

        data[C.UpdatedAt] = ""

        dataset = append(dataset, data)
        return dataset
    }
}

// 必要な情報を抽出する
func salvageFeaturesAgeless(normalizedText string) []*model.JobFeature {
    return salvageJobData(normalizedText, "")
}

func (c *CallBacksAgeless) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data)
    }
}

func (c *CallBacksAgeless) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, "body")
    }
}