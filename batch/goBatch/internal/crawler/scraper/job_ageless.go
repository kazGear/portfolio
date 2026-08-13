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

    mutex   := &sync.Mutex{}

    // URL生成の設定
    pageIdFrom, pageIdTo := loadPageIdFromTo("PAGE_ID_FROM_AGELESS", "PAGE_ID_TO_AGELESS")
    visited              := make(map[string]struct{}, pageIdTo - pageIdFrom)

    validatePageIdFromTo(pageIdFrom, pageIdTo)

    // URL生成
    for pageId := pageIdFrom; pageId <= pageIdTo; pageId++ {
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
        return buildJobFrame(data, c.funcs.logger)
    }
}

func (c *CallBacksAgeless) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "body")
    }
}