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
)

type CrawlerFreelanceJob struct {
    name     string
    jScraper Crawler[*model.Job]
}

type CallBacksFreelanceJob struct {
    funcs CallBacks
}

func NewScraperFreelanceJob() Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1, // URL収集漏れが発生するため5に制限
	})
    return &CrawlerFreelanceJob{
        "フリーランスジョブ",
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksFreelanceJob() *CallBacksFreelanceJob {
    return &CallBacksFreelanceJob{
        CallBacks{},
    }
}

// CollectAttributesへ
var _parentCtxFreelanceJob context.Context

func (c *CrawlerFreelanceJob) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector             := c.jScraper.collector
    _parentCtxFreelanceJob = parentCtx

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(collector ,crawlStats)

    mutex   := &sync.Mutex{}

    // URL生成の設定
    pageIdFrom, pageIdTo := loadPageIdFromTo("PAGE_ID_FROM_FREELANCE_JOB", "PAGE_ID_TO_FREELANCE_JOB")
    visited              := make(map[string]struct{}, pageIdTo - pageIdFrom)

    validatePageIdFromTo(pageIdFrom, pageIdTo)

    // URL生成
    for pageId := pageIdFrom; pageId <= pageIdTo; pageId++ {
        url := fmt.Sprintf("https://freelance-job.com/job/detail/%v", pageId)
        isFirstVisit(mutex, url, visited)
    }

    loggingCrawlStats(c.name, crawlStats)

    c.jScraper.urls = utils.MapToSliceUrl(visited)
    return c.jScraper.urls, nil
}

func (c *CrawlerFreelanceJob) Scrape(provider  PageProvider,
                                     parser    ModelParser[*model.Job],
                                     parentCtx context.Context,
) []*model.Job {
    jobs := c.jScraper.scrapeFrame(provider, parser, parentCtx)
    return jobs
}

func (c *CallBacksFreelanceJob) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
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
        ctx, cancel := context.WithTimeout(tabCtx, 2 * time.Second)
        defer cancel()

        // 404ページに対する対応
        isNotFount := isNotFountPage("案件詳細", ctx)

        if isNotFount { return "", fmt.Errorf(C.This404page, url)}

        // クロームで対応
        var html string

        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        )

        if err != nil {
            log.Printf("Chromedp error: %v", err)
            return "", err
        }
        return html, nil
    }
}

var _regFindPriceFreelanceJob = regexp.MustCompile(`(\d{0,3},?\d{0,3},\d{0,3})?\s*~\s*(\d{0,3},?\d{0,3},\d{0,3})?`)

func (c *CallBacksFreelanceJob) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)

        description           := collectTextFreelanceJob(doc)
        normalizedDescription := normalizeForSearchFeatures(description)

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesFreelanceJob(normalizedDescription)
        // 保存するべき案件か
        if len(features) <= 0 {
            return []map[string]string{}
        }
        repository.InjectionJobFeatures(features, url)

        // 案件のオプションを収集し、repositoryへ（このサイトは無し）


        data := map[string]string{}

        data[C.Url]         = url
        data[C.Title]       = doc.Find(`p:contains("業務委託")`).Parent().Parent().Prev().Children().Next().Text()
        data[C.Location]    = salvageLocation(normalizedDescription)

        priceText               := doc.Find(`p:contains("業務委託")`).Parent().Next().Text()
        minPrice, maxPrice      := getJobPrice(_regFindPriceFreelanceJob.FindString(priceText))
        data[C.MinSalaryAtMonth] = strconv.Itoa(minPrice)
        data[C.MaxSalaryAtMonth] = strconv.Itoa(maxPrice)

        data[C.Description]    = description
        data[C.EmploymentType] = salvageEmploymentType(normalizedDescription)
        data[C.WorkPlace]      = salvageWorkPlace(normalizedDescription)
        data[C.SourceSite]     = C.FreelanceJob

        data[C.UpdatedAt] = getOpenDateFreelanceJob(doc)

        dataset = append(dataset, data)
        return dataset
    }
}

func collectTextFreelanceJob(doc *goquery.Document) string {
    builder := &strings.Builder{}

    builder.WriteString(doc.Find(`dt:contains("案件詳細")`).Next().Text())

    doc.Find(`dt:contains("開発言語")`).Next().Children().Children().Each(
        func(idx int, selector *goquery.Selection) {
            builder.WriteString(selector.Text())
            builder.WriteString(" ")
    })
    builder.WriteString(doc.Find(`dt:contains("必須スキル・経験")`).Next().Text())
    builder.WriteString(doc.Find(`dt:contains("尚可スキル・経験")`).Next().Text())

    doc.Find(`dt:contains("職種・ポジション")`).Next().Children().Children().Each(
        func(idx int, selector *goquery.Selection) {
            builder.WriteString(selector.Text())
            builder.WriteString(" ")
    })
    builder.WriteString(doc.Find(`dt:contains("業界")`).Next().Text())
    builder.WriteString(doc.Find(`dt:contains("案件詳細")`).Next().Text())
    builder.WriteString(doc.Find(`dt:contains("募集背景")`).Next().Text())
    builder.WriteString(doc.Find(`dt:contains("開発環境")`).Next().Text())
    builder.WriteString(doc.Find(`dt:contains("出社頻度")`).Next().Text())
    builder.WriteString(doc.Find(`dt:contains("案件公開日時")`).Next().Text())
    builder.WriteString(doc.Find(`span:contains("おすすめポイント")`).Parent().Next().Next().Text())
    builder.WriteString(doc.Find(`p:contains("業務委託")`).Next().Text())

    return builder.String()
}

func getOpenDateFreelanceJob(doc *goquery.Document) string {
    dateText := doc.Find(`dt:contains("案件公開日時")`).Next().Text()
    dateText  = _regGetDate.FindString(dateText)

    return strings.ReplaceAll(dateText, "/", "-")
}

// 必要な情報を抽出する
func salvageFeaturesFreelanceJob(normalizedText string) []*model.JobFeature {
    return salvageJobData(normalizedText, "")
}

func (c *CallBacksFreelanceJob) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data)
    }
}

func (c *CallBacksFreelanceJob) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "body")
    }
}