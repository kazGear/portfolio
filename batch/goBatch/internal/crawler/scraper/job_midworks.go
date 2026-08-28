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

type CrawlerMidworks struct {
    name     string
    jScraper Crawler[*model.Job]
}

type CallBacksMidworks struct {
    funcs CallBacks
}

func NewScraperMidworks() Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
        Delay:       500 * time.Millisecond,
        RandomDelay: 500 * time.Millisecond,
	})
    return &CrawlerMidworks{
        "Midworks",
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksMidworks() *CallBacksMidworks {
    return &CallBacksMidworks{
        CallBacks{},
    }
}

// CollectAttributesへ
var _parentCtxMidworks context.Context

func (c *CrawlerMidworks) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector              := c.jScraper.collector
    _parentCtxMidworks = parentCtx

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(collector ,crawlStats)

    mutex   := &sync.Mutex{}

    // URL生成の設定
    pageIdFrom, pageIdTo := loadPageIdFromTo("PAGE_ID_FROM_MIDWORKS", "PAGE_ID_TO_MIDWORKS")
    visited              := make(map[string]struct{}, pageIdTo - pageIdFrom)

    validatePageIdFromTo(pageIdFrom, pageIdTo)

    // 保存済ページID取得
    repository   := repository.NewJobRepository(db.GetInstance())
    savedPageIds := repository.Select(c.name)

    log.Printf("%v savedPageIds: %v件\n", c.name, len(savedPageIds))

    // URL生成
    for pageId := pageIdFrom; pageId <= pageIdTo; pageId++ {
        if _, exist := savedPageIds[pageId]; exist {
            continue
        }
        url := fmt.Sprintf("https://mid-works.com/projects/%v", pageId)
        isFirstVisit(mutex, url, visited)
    }
    loggingCrawlStats(c.name, crawlStats)

    c.jScraper.urls = utils.MapToSliceUrl(visited)
    log.Printf("%v visit urls: %v件\n", c.name, len(c.jScraper.urls))

    return c.jScraper.urls, nil
}

func (c *CrawlerMidworks) Scrape(provider  PageProvider,
                                 parser    ModelParser[*model.Job],
                                 parentCtx context.Context,
) []*model.Job {
    jobs := c.jScraper.scrapeFrame(provider, parser, parentCtx)
    return jobs
}

func (c *CallBacksMidworks) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`https://mid-works.com/projects/\d+`, url) {
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
        isNotFount := isNotFountPage("業務内容", ctx)

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

func (c *CallBacksMidworks) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)

        description           := collectTextMidworks(doc)
        normalizedDescription := normalizeForSearchFeatures(description)

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesMidworks(doc)
        // 保存するべき案件か
        if len(features) <= 0 {
            return []map[string]string{}
        }
        repository.InjectionJobFeatures(features, url)

        // 案件のオプションを収集し、repositoryへ（このサイトは無し）


        data := map[string]string{}

        data[C.Url]         = url
        data[C.Title]       = doc.Find(`.p-jobBoard__title`).Text()
        data[C.Location]    = salvageLocation(normalizedDescription)

        minPrice, maxPrice      := getJobPrice(doc.Find(`.p-jobBoard__salary`).Text())
        data[C.MinSalaryAtMonth] = strconv.Itoa(minPrice)
        data[C.MaxSalaryAtMonth] = strconv.Itoa(maxPrice)

        data[C.Description]    = description
        data[C.EmploymentType] = salvageEmploymentType(normalizedDescription)
        data[C.WorkPlace]      = salvageWorkPlace(normalizedDescription)
        data[C.SourceSite]     = C.Midworks

        data[C.UpdatedAt] = getOpenDateMidworks(doc)

        dataset = append(dataset, data)
        return dataset
    }
}

func collectTextMidworks(doc *goquery.Document) string {
    builder := &strings.Builder{}

    builder.WriteString(doc.Find(`.p-jobBoard__title`).Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`.p-jobBoard__salary`).Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`dt:contains("勤務地")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`h2:contains("業務内容")`).Parent().Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`h2:contains("求めるスキル")`).Parent().Next().Text())
    builder.WriteString("\n")

    remote := doc.Find(`#remote_working`).Text()
    builder.WriteString(strings.ReplaceAll(remote, "可", "一部リモート"))
    builder.WriteString("\n")

    return builder.String()
}

func getOpenDateMidworks(doc *goquery.Document) string {
    return ""
}

func isActiveMidworks(doc *goquery.Document) string {
    isActive := "invalid"

    errorText1 := doc.Find(`p:contains("お探しのページが")`).Text()
    errorText2 := doc.Find(`p:contains("URLにお間違いがないか")`).Text()

    if errorText1 != "" || errorText2 != "" {
        isActive = "false"
    }
    return isActive
}

func salvageFeaturesMidworks(doc *goquery.Document) []*model.JobFeature {
    skillText := doc.Find(`h2:contains("求めるスキル")`).Parent().Next().Text()
    skillText  = normalizeForSearchFeatures(skillText)
    skills    := salvageJobData(skillText, "")

    return skills
}

func (c *CallBacksMidworks) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data)
    }
}

func (c *CallBacksMidworks) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "body")
    }
}