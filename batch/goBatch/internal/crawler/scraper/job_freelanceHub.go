package scraper

import (
	"context"
	"fmt"
	"slices"
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

type CrawlerFreelanceHub struct {
    jScraper Crawler[*model.Job]
}

type CallBacksFreelanceHub struct {
    funcs CallBacks
}

func NewScraperFreelanceHub(logger *log.Logger) Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
	})
    return &CrawlerFreelanceHub{
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksFreelanceHub(logger *log.Logger) *CallBacksFreelanceHub {
    return &CallBacksFreelanceHub{
        CallBacks{
            logger: logger,
        },
    }
}

// CollectAttributesへ
var _parentCtxFreelanceHub context.Context

func (c *CrawlerFreelanceHub) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector              := c.jScraper.collector
    _parentCtxFreelanceHub = parentCtx

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(collector ,crawlStats, c.jScraper.logger)

    mutex   := &sync.Mutex{}

    // URL生成の設定
    pageIdFrom, pageIdTo := loadPageIdFromTo("PAGE_ID_FROM_FREELANCE_HUB", "PAGE_ID_TO_FREELANCE_HUB")
    visited              := make(map[string]struct{}, pageIdTo - pageIdFrom)

    validatePageIdFromTo(pageIdFrom, pageIdTo)

    // URL生成
    for pageId := pageIdFrom; pageId <= pageIdTo; pageId++ {
        url := fmt.Sprintf("https://freelance-hub.jp/project/detail/%v", pageId)
        isFirstVisit(mutex, url, visited)
    }

    loggingCrawlStats(crawlStats, c.jScraper.logger)

    c.jScraper.urls = utils.MapToSliceUrl(visited)
    return c.jScraper.urls, nil
}

func (c *CrawlerFreelanceHub) Scrape(provider  PageProvider,
                                     parser    ModelParser[*model.Job],
                                     parentCtx context.Context,
) []*model.Job {
    jobs := c.jScraper.scrapeFrame(provider, parser, parentCtx)
    return jobs
}

func (c *CallBacksFreelanceHub) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
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

func (c *CallBacksFreelanceHub) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)

        description           := collectTextFreelanceHub(doc)
        normalizedDescription := normalizeForSearchFeatures(description)

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesFreelanceHub(doc)
        // 保存するべき案件か
        if len(features) <= 0 {
            return []map[string]string{}
        }
        repository.InjectionJobFeatures(features, url)

        // 案件のオプションを収集し、repositoryへ
        options := make([]*model.JobOption, 0, 20)

        doc.Find(`h3:contains("案件特徴")`).Next().Children().Each(func(idx int, selector *goquery.Selection) {
            option := &model.JobOption{
                JobId: -1,
                Option: selector.Text(),
            }
            options = append(options, option)
        })
        repository.InjectionJobOptions(options, url)

        data := map[string]string{}

        data[C.Url]         = url
        data[C.Title]       = doc.Find(`h2.Detail_Title`).Text()
        data[C.CompanyName] = ""
        data[C.Location]    = salvageLocation(normalizedDescription)

        minPrice, maxPrice      := getJobPrice(doc.Find(`.Detail_SummaryItem--money`).Text())
        data[C.MinSalaryAtMonth] = strconv.Itoa(minPrice)
        data[C.MaxSalaryAtMonth] = strconv.Itoa(maxPrice)

        data[C.Description]    = description
        data[C.EmploymentType] = salvageEmploymentType(normalizedDescription)
        data[C.WorkPlace]      = salvageWorkPlace(normalizedDescription)

        data[C.IsActive]       = isActiveFreelanceHub(doc)
        // data[C.SimilarityScore] =
        data[C.SourceSite]     = C.FreelanceHub
        data[C.UpdatedAt]      = getOpenDateFreelanceHub(doc)

        dataset = append(dataset, data)
        return dataset
    }
}

func collectTextFreelanceHub(doc *goquery.Document) string {
    builder := &strings.Builder{}

    builder.WriteString(doc.Find(`.Detail_SummaryItem--money`).Text())
    builder.WriteString(" ")
    builder.WriteString(doc.Find(`.Detail_SummaryItem--contract`).Text())
    builder.WriteString(" ")
    builder.WriteString(doc.Find(`.Detail_SummaryItem--location`).Text())
    builder.WriteString(" ")

    doc.Find(`h3:contains("職種・ポジション")`).Next().Children().Each(
        func(idx int, selector *goquery.Selection) {
            builder.WriteString(selector.Text())
            builder.WriteString(" ")
    })
    builder.WriteString(doc.Find(`h3:contains("作業内容")`).Next().Text())

    doc.Find(`h3:contains("開発環境")`).Next().Children().Each(
        func(idx int, selector *goquery.Selection) {
            builder.WriteString(selector.Text())
            builder.WriteString(" ")
    })
    builder.WriteString(doc.Find(`h3:contains("必須スキル")`).Next().Text())
    builder.WriteString(doc.Find(`h3:contains("歓迎スキル")`).Next().Text())

    doc.Find(`h3:contains("案件特徴")`).Next().Children().Each(
        func(idx int, selector *goquery.Selection) {
            builder.WriteString(selector.Text())
            builder.WriteString(" ")
    })
    builder.WriteString(doc.Find(`h3:contains("最終更新日")`).Next().Text())

    return builder.String()
}

func getOpenDateFreelanceHub(doc *goquery.Document) string {
    dateText := doc.Find(`h3:contains("最終更新日")`).Next().Text()

    dateText  = _regGetDate.FindString(dateText)

    dateText  = strings.ReplaceAll(dateText, "年", "-")
    dateText  = strings.ReplaceAll(dateText, "月", "-")

    return dateText
}

func isActiveFreelanceHub(doc *goquery.Document) string {
    isActive := "invalid"

    errorText := doc.Find(`a:contains("エージェントに相談する")`).Text()

    if errorText != ""  {
        isActive = "false"
    } else {
        isActive = "true"
    }
    return isActive
}

func salvageFeaturesFreelanceHub(doc *goquery.Document) []*model.JobFeature {
    requiredSkillText := doc.Find(`h3:contains("必須スキル")`).Next().Text()
    optionalSkillText := doc.Find(`h3:contains("歓迎スキル")`).Next().Text()

    requiredSkillText = normalizeForSearchFeatures(requiredSkillText)
    optionalSkillText = normalizeForSearchFeatures(optionalSkillText)

    requiredSkills := salvageJobData(requiredSkillText, C.Required)
    optionalSkills := salvageJobData(optionalSkillText, C.Optional)

    return slices.Concat(requiredSkills, optionalSkills)
}

func (c *CallBacksFreelanceHub) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data, c.funcs.logger)
    }
}

func (c *CallBacksFreelanceHub) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "body")
    }
}