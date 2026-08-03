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
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

type CrawlerAgeless struct {
    jScraper Crawler[*model.Job]
}

type CallBacksAgeless struct {
    funcs CallBacks
}

func NewScraperAgeless(logger *log.Logger) Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1, // URL収集漏れが発生するため5に制限
	})
    return &CrawlerAgeless{
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksAgeless(logger *log.Logger) *CallBacksAgeless {
    return &CallBacksAgeless{
        CallBacks{
            logger: logger,
        },
    }
}

// CollectAttributesへ
var _parentCtxAgeless context.Context

func (c *CrawlerAgeless) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector              := c.jScraper.collector
    _parentCtxAgeless = parentCtx

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(collector ,crawlStats, c.jScraper.logger)

    // URL収集、クロール
    visited := make(map[string]struct{}, 10000)
    mutex   := &sync.Mutex{}

    for pageId := 1; pageId <= 10000; pageId++ {
        url := fmt.Sprintf("https://freelance.ageless.co.jp/projects/%v", pageId)
        isFirstVisit(mutex, url, visited)
    }

    loggingCrawlStats(crawlStats, c.jScraper.logger)

    c.jScraper.urls = utils.MapToSliceUrl(visited)
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
        if !isDetailPage(`^https://freelance.ageless.co.jp/projects/\d+`, url) {
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

        data := map[string]string{}

        data[C.Url]         = url
        data[C.Title]       = doc.Find(".project-card-ttl").Text()
        data[C.CompanyName] = ""
        data[C.Location]    = "" // 別関数で抽出

        data[C.MinSalaryAtMonth] = doc.Find(".income-num").Text()
        data[C.MaxSalaryAtMonth] = doc.Find(".income-num").Text()

        description           := doc.Find("main").Text()
        data[C.Description]    = _regDeleteDescriptionAgeLess.ReplaceAllString(description, "")
        data[C.EmploymentType] = "" // 別関数で抽出
        data[C.WorkPlace]      = "" // 別関数で抽出
        data[C.IsActive]       = isActiveAgeless(doc)
        // data[C.SimilarityScore] =
        data[C.SourceSite]     = C.AGELESS

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesAgeless(data, data[C.Description])
        repository.InjectionJobFeatures(features, url)

        // 案件のオプションを収集し、repositoryへ（このサイトは無し）

        dataset = append(dataset, data)
        return dataset
    }
}

func isActiveAgeless(doc *goquery.Document) string {
    isActive := "invalid"

    title := doc.Find("title").Text()

    if strings.Contains(title, "404") {
        isActive = "false"
    } else if strings.Contains(doc.Find(".btn-disabled").Text(), "募集終了") {
        isActive = "false"
    } else {
        isActive = "true"
    }
    return isActive
}

// 必要な情報を抽出する
func salvageFeaturesAgeless(data map[string]string, target string) []*model.JobFeature {
    normalizedTarget := normalizeForSearchFeature(target)

    data[C.Location]       = salvageLocation(normalizedTarget)
    data[C.WorkPlace]      = salvageWorkPlace(normalizedTarget)
    data[C.EmploymentType] = salvageEmploymentType(normalizedTarget)

    return salvageJobData(normalizedTarget)
}

func (c *CallBacksAgeless) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data, c.funcs.logger)
    }
}

func (c *CallBacksAgeless) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "body")
    }
}