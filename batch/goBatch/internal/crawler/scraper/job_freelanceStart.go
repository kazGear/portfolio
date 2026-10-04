package scraper

import (
	"context"
	"fmt"
	"strconv"
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

type CrawlerFreelanceStart struct {
    name     string
    jScraper Crawler[*model.Job]
}

type CallBacksFreelanceStart struct {
    funcs CallBacks
}

func NewScraperFreelanceStart() Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
        Delay:       500 * time.Millisecond,
        RandomDelay: 1500 * time.Millisecond,
	})
    return &CrawlerFreelanceStart{
        "フリーランススタート",
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksFreelanceStart() *CallBacksFreelanceStart {
    return &CallBacksFreelanceStart{
        CallBacks{},
    }
}

func (c *CrawlerFreelanceStart) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector := c.jScraper.collector
    mutex     := c.jScraper.mutex
    visited   := make(map[string]struct{}, 2500)

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(collector ,crawlStats)

    // pagination urls を作成
    paginationUrls := make(map[string]struct{}, C.PaginationLimit)
    for i := 1; i <= C.PaginationLimit; i++ {
        paginationUrls[fmt.Sprintf(`https://freelance-start.com/jobs?page=%v`, i)] = struct{}{}
    }

    // 詳細ページの url を取得
    for url := range paginationUrls {
        pageCtx, pageCtxCancel := chromedp.NewContext(parentCtx) // 独立したコンテキストの作成
        ctx, ctxCancel         := context.WithTimeout(pageCtx, 60 * time.Second)

        var html string
        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.WaitVisible(`#job-list`, chromedp.ByQuery),
            chromedp.OuterHTML(`#job-list`, &html, chromedp.ByQueryAll),
        )
        ctxCancel()
        pageCtxCancel()
        if err != nil {
            log.Println(err)
            break
        }

        doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
        if err != nil {
            log.Println(err)
            break
        }

        // 案件詳細への url を取得
        doc.Find(`a[href^="/jobs/detail/"]`).Each(func(_ int, selector *goquery.Selection) {
            url, _ := selector.Attr("href")
            utils.LockedAddSet(mutex, visited, "https://freelance-start.com" + url)
        })

        // JS が動くための猶予
        time.Sleep(1 * time.Second)
    }

    // 保存済ページID取得
    repository := repository.NewJobRepository(db.GetInstance())
    savedUrls  := repository.Select(c.name)
    log.Printf("%v savedUrls: %v件\n", c.name, len(savedUrls))

    // 保存されていない案件urlだけを残す
    needUrls := make(map[string]struct{}, 2500)
    for url := range visited {
        if _, exist := savedUrls[url]; !exist {
            utils.LockedAddSet(mutex, needUrls, url)
        }
    }

    loggingCrawlStats(c.name, crawlStats)

    c.jScraper.urls = utils.MapToSliceUrl(needUrls)

    return c.jScraper.urls, nil
}

func (c *CrawlerFreelanceStart) Scrape(provider  PageProvider,
                                       parser    ModelParser[*model.Job],
                                       parentCtx context.Context,
) []*model.Job {
    jobs := c.jScraper.scrapeFrame(provider, parser, parentCtx)
    return jobs
}

func (c *CallBacksFreelanceStart) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        // 動的ページを取得しない場合、引数のパターンは記載しないで良い
        if !isDetailPage(`^https://freelance-start.com/jobs/detail/\d+`, url) {
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
        ctx, cancel := context.WithTimeout(tabCtx, 20 * time.Second)
        defer cancel()

        var html string
        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.WaitVisible(`.job-title`, chromedp.ByQuery),
            chromedp.OuterHTML("html", &html, chromedp.ByQuery),
        )
        if err != nil {
            log.Printf("Chromedp error %v: %v, %v", "フリーランススタート", err, url)
            return "", err
        }

        return html, nil
    }
}

func (c *CallBacksFreelanceStart) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)

        description           := collectTextFreelanceStart(doc)
        normalizedDescription := normalizeForSearchFeatures(description)

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesFreelanceStart(normalizedDescription)

        // 保存するべき案件か
        if len(features) <= 0 {
            return []map[string]string{}
        }
        repository.InjectionJobFeatures(features, url)

        // 案件のオプションを収集し、repositoryへ
        options := collectOptionsFreelanceStart(doc)
        repository.InjectionJobOptions(options, url)

        data := map[string]string{}

        data[C.Url]         = url
        data[C.Title]       = doc.Find(`title`).Text()
        data[C.Location]    = salvageLocation(normalizedDescription)

        minPrice, maxPrice      := getJobPrice(doc.Find(`.salary`).Text())
        data[C.MinSalaryAtMonth] = strconv.Itoa(minPrice)
        data[C.MaxSalaryAtMonth] = strconv.Itoa(maxPrice)

        data[C.Description]    = normalizedDescription
        data[C.EmploymentType] = salvageEmploymentType(normalizedDescription)
        data[C.WorkPlace]      = salvageWorkPlace(normalizedDescription)
        data[C.SourceSite]     = C.FreelanceStart

        data[C.UpdatedAt] = ""

        dataset = append(dataset, data)
        return dataset
    }
}

func collectTextFreelanceStart(doc *goquery.Document) string {
    builder := &strings.Builder{}

    builder.WriteString(doc.Find(`.tech-stack`).Text())
    builder.WriteString(doc.Find(`.left-column`).Text())
    builder.WriteString(doc.Find(`.right-column`).Text())

    return builder.String()
}

func collectOptionsFreelanceStart(doc *goquery.Document) []*model.JobOption {
    optionsArr := make([]*model.JobOption, 0, 20)

    optionsText := doc.Find(`.detail-label:contains("特徴")`).Next().Text()

    if options := strings.Split(optionsText, "/"); len(options) > 1 {
        for _, opt := range options {
            option := &model.JobOption{
                JobId: -1,
                Option: opt,
            }
            optionsArr = append(optionsArr, option)
        }
    }
    return optionsArr
}

// 必要な情報を抽出する
func salvageFeaturesFreelanceStart(normalizedText string) []*model.JobFeature {
    return salvageJobData(normalizedText, "")
}

func (c *CallBacksFreelanceStart) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data)
    }
}

func (c *CallBacksFreelanceStart) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        // 静的ソースのみからデータを取得する場合、必ず存在する bodyタグ(bodyの文字列)を指定しておく
        return strings.Contains(html, "job-title")
    }
}