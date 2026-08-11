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
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

type CrawlerFreelanceStart struct {
    jScraper Crawler[*model.Job]
}

type CallBacksFreelanceStart struct {
    funcs CallBacks
}

func NewScraperFreelanceStart(logger *log.Logger) Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1, // URL収集漏れが発生するため5に制限
	})
    return &CrawlerFreelanceStart{
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksFreelanceStart(logger *log.Logger) *CallBacksFreelanceStart {
    return &CallBacksFreelanceStart{
        CallBacks{
            logger: logger,
        },
    }
}

// CollectAttributesへ
var _parentCtxFreelanceStart context.Context

func (c *CrawlerFreelanceStart) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector              := c.jScraper.collector
    _parentCtxFreelanceStart = parentCtx

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(collector ,crawlStats, c.jScraper.logger)

    mutex   := &sync.Mutex{}

    // URL生成の設定
    pageIdFrom, pageIdTo := loadPageIdFromTo("PAGE_ID_FROM_FREELANCE_START", "PAGE_ID_TO_FREELANCE_START")
    visited              := make(map[string]struct{}, pageIdTo - pageIdFrom)

    validatePageIdFromTo(pageIdFrom, pageIdTo)

    // URL生成
    for pageId := pageIdFrom; pageId <= pageIdTo; pageId++ {
        url := fmt.Sprintf("https://freelance-start.com/jobs/detail/%v", pageId)
        isFirstVisit(mutex, url, visited)
    }

    loggingCrawlStats(crawlStats, c.jScraper.logger)

    c.jScraper.urls = utils.MapToSliceUrl(visited)
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
        ctx, cancel := context.WithTimeout(tabCtx, 3 * time.Second)
        defer cancel()

        // 404ページに対する対応
        isNotFount := isNotFountPage("職務内容", ctx)

        if isNotFount { return "", fmt.Errorf(C.This404page, url)}

        // クロームで対応
        var html string

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.WaitReady(`.job-title`, chromedp.ByQuery), // 求める要素が出るまで待つ
            chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )

        if err != nil {
            log.Printf("Chromedp error: %v", err)
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
        data[C.CompanyName] = ""
        data[C.Location]    = salvageLocation(normalizedDescription)

        minPrice, maxPrice      := getJobPrice(doc.Find(`.salary`).Text())

        data[C.MinSalaryAtMonth] = strconv.Itoa(minPrice)
        data[C.MaxSalaryAtMonth] = strconv.Itoa(maxPrice)

        data[C.Description]    = description
        data[C.EmploymentType] = salvageEmploymentType(normalizedDescription)
        data[C.WorkPlace]      = salvageWorkPlace(normalizedDescription)

        data[C.IsActive]       = isActiveFreelanceStart(doc)
        // data[C.SimilarityScore] =
        data[C.SourceSite]     = C.FreelanceStart
        data[C.UpdatedAt]      = ""

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

func isActiveFreelanceStart(doc *goquery.Document) string {
    isActive := "invalid"

    errorPage := doc.Find(`.error-page-hero`).Text()

    if strings.Contains(errorPage, "404") {
        isActive = "false"
    }
    buttonText := doc.Find(`.action-buttons`).Text()

    if strings.Contains(buttonText, "案件の話") {
        isActive = "true"
    } else if strings.Contains(buttonText, "募集終了") {
        isActive = "false"
    }

    return isActive
}

// 必要な情報を抽出する
func salvageFeaturesFreelanceStart(normalizedText string) []*model.JobFeature {
    return salvageJobData(normalizedText, "")
}

func (c *CallBacksFreelanceStart) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data, c.funcs.logger)
    }
}

func (c *CallBacksFreelanceStart) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "job-title")
    }
}