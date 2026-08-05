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
	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/repository"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
	"golang.org/x/text/width"
)

type CrawlerSesJobLink struct {
    jScraper Crawler[*model.Job]
}

type CallBacksSesJobLink struct {
    funcs CallBacks
}

func NewScraperSesJobLink(logger *log.Logger) Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1, // URL収集漏れが発生するため5に制限
	})
    return &CrawlerSesJobLink{
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksSesJobLink(logger *log.Logger) *CallBacksSesJobLink {
    return &CallBacksSesJobLink{
        CallBacks{
            logger: logger,
        },
    }
}

// CollectAttributesへ
var _parentCtxSesJobLink context.Context

func (c *CrawlerSesJobLink) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector              := c.jScraper.collector
    _parentCtxSesJobLink = parentCtx

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(collector ,crawlStats, c.jScraper.logger)

    // URL収集、クロール
    visited := make(map[string]struct{}, 60000)
    mutex   := &sync.Mutex{}

    for pageId := 1; pageId <= 60000; pageId++ {
        url := fmt.Sprintf("https://ses-job-link.com/projects/%v", pageId)
        isFirstVisit(mutex, url, visited)
    }

    loggingCrawlStats(crawlStats, c.jScraper.logger)

    c.jScraper.urls = utils.MapToSliceUrl(visited)
    return c.jScraper.urls, nil
}

func (c *CrawlerSesJobLink) Scrape(provider  PageProvider,
                                   parser    ModelParser[*model.Job],
                                   parentCtx context.Context,
) []*model.Job {
    jobs := c.jScraper.scrapeFrame(provider, parser, parentCtx)
    return jobs
}

func (c *CallBacksSesJobLink) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://freelance.SesJobLink.co.jp/projects/\d+`, url) {
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

func (c *CallBacksSesJobLink) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)

        data := map[string]string{}

        data[C.Url]         = url
        data[C.Title]       = doc.Find(".project-title").Text()
        data[C.CompanyName] = ""
        data[C.Location]    = "" // 別関数で抽出

        description := doc.Find("#sys_project_detail_page").Text()
        minPrice, maxPrice      := getJobPrice(description)
        data[C.MinSalaryAtMonth] = strconv.Itoa(minPrice)
        data[C.MaxSalaryAtMonth] = strconv.Itoa(maxPrice)

        data[C.Description]    = description
        data[C.EmploymentType] = "" // 別関数で抽出
        data[C.WorkPlace]      = "" // 別関数で抽出
        data[C.IsActive]       = isActiveSesJobLink()
        // data[C.SimilarityScore] =
        data[C.SourceSite]     = C.SES_JOB_LINK
        data[C.UpdatedAt]      = getUpdatedAtSesJobLink(description)

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesSesJobLink(data, data[C.Description])
        repository.InjectionJobFeatures(features, url)

        // 案件のオプションを収集し、repositoryへ（このサイトは無し）

        dataset = append(dataset, data)
        return dataset
    }
}

func isActiveSesJobLink() string {
    isActive := "invalid"
    return isActive
}

var _regUpdatedAtSesJobLink = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

func getUpdatedAtSesJobLink(text string) string {
    text = width.Narrow.String(text)
    text = strings.ReplaceAll(text, "\r\n", "")
    text = strings.ReplaceAll(text, "\n", "")
    text = strings.ReplaceAll(text, " ", "")
    text = strings.ReplaceAll(text, "年", "-")
    text = strings.ReplaceAll(text, "月", "-")

    updatedAt := _regUpdatedAtSesJobLink.FindString(text)

    return updatedAt
}

// 必要な情報を抽出する
func salvageFeaturesSesJobLink(data map[string]string, target string) []*model.JobFeature {
    normalizedTarget := normalizeForSearchFeature(target)

    data[C.Location]       = salvageLocation(normalizedTarget)
    data[C.WorkPlace]      = salvageWorkPlace(normalizedTarget)
    data[C.EmploymentType] = salvageEmploymentType(normalizedTarget)

    return salvageJobData(normalizedTarget)
}

func (c *CallBacksSesJobLink) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data, c.funcs.logger)
    }
}

func (c *CallBacksSesJobLink) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "body")
    }
}