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
	"github.com/kazGear/portfolio/goBatch/pkg/db"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

type CrawlerEngineerFactory struct {
    name     string
    jScraper Crawler[*model.Job]
}

type CallBacksEngineerFactory struct {
    funcs CallBacks
}

func NewScraperEngineerFactory() Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
	})
    return &CrawlerEngineerFactory{
        "エンジニアファクトリー",
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksEngineerFactory() *CallBacksEngineerFactory {
    return &CallBacksEngineerFactory{
        CallBacks{},
    }
}

// CollectAttributesへ
var _parentCtxEngineerFactory context.Context

func (c *CrawlerEngineerFactory) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector              := c.jScraper.collector
    _parentCtxEngineerFactory = parentCtx

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(collector ,crawlStats)

    mutex   := &sync.Mutex{}

    // URL生成の設定
    pageIdFrom, pageIdTo := loadPageIdFromTo("PAGE_ID_FROM_ENGINEER_FACTORY", "PAGE_ID_TO_ENGINEER_FACTORY")
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
        url := fmt.Sprintf("https://www.engineer-factory.com/freelance/jobs/%v", pageId)
        isFirstVisit(mutex, url, visited)
    }
    loggingCrawlStats(c.name, crawlStats)

    c.jScraper.urls = utils.MapToSliceUrl(visited)
    log.Printf("%v visit urls: %v件\n", c.name, len(c.jScraper.urls))

    return c.jScraper.urls, nil
}

func (c *CrawlerEngineerFactory) Scrape(provider  PageProvider,
                                        parser    ModelParser[*model.Job],
                                        parentCtx context.Context,
) []*model.Job {
    jobs := c.jScraper.scrapeFrame(provider, parser, parentCtx)
    return jobs
}

func (c *CallBacksEngineerFactory) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
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

func (c *CallBacksEngineerFactory) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)

        description           := collectTextEngineerFactory(doc)
        normalizedDescription := normalizeForSearchFeatures(description)

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesEngineerFactory(doc)
        // 保存するべき案件か
        if len(features) <= 0 {
            return []map[string]string{}
        }
        repository.InjectionJobFeatures(features, url)

        // 案件のオプションを収集し、repositoryへ（このサイトは無し）


        data := map[string]string{}

        data[C.Url]         = url
        data[C.Title]       = doc.Find(`.modJobBlock__detailTitle`).Text()
        data[C.Location]    = salvageLocation(normalizedDescription)

        minPrice, maxPrice      := getJobPrice(normalizedDescription)
        data[C.MinSalaryAtMonth] = strconv.Itoa(minPrice)
        data[C.MaxSalaryAtMonth] = strconv.Itoa(maxPrice)

        data[C.Description]    = description
        data[C.EmploymentType] = salvageEmploymentType(normalizedDescription)
        data[C.WorkPlace]      = salvageWorkPlace(normalizedDescription)
        data[C.SourceSite]     = C.EngineerFactory

        data[C.UpdatedAt] = getOpenDateEngineerFactory(doc)

        dataset = append(dataset, data)
        return dataset
    }
}

func collectTextEngineerFactory(doc *goquery.Document) string {
    builder := &strings.Builder{}

    builder.WriteString(doc.Find(`.modJobBlock__detailTitle`).Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`.modJobBlock__update`).Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`dt:contains("単価")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`h3:contains("都道府県")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`h3:contains("最寄駅")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`h3:contains("契約形態")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`h3:contains("必須スキル")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`h3:contains("歓迎スキル")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`h3:contains("業務内容")`).Next().Text())
    builder.WriteString("\n")

    return builder.String()
}

func getOpenDateEngineerFactory(doc *goquery.Document) string {
    dateText := doc.Find(`.modJobBlock__update`).Text()
    dateText  = _regGetDate.FindString(dateText)
    dateText  = strings.ReplaceAll(dateText, "/", "-")

    return dateText
}

func isActiveEngineerFactory(doc *goquery.Document) string {
    isActive := "invalid"

    errorText1 := doc.Find(`h1:contains("404")`).Text()
    errorText2 := doc.Find(`p:contains("お探しのページは見つかりません")`).Text()
    errorText3 := doc.Find(`p:contains("案件は終了しました")`).Text()

    if errorText1 != "" || errorText2 != "" || errorText3 != "" {
        isActive = "false"
    } else {
        isActive = "true"
    }
    return isActive
}

func salvageFeaturesEngineerFactory(doc *goquery.Document) []*model.JobFeature {
    requiredSkillText := doc.Find(`h3:contains("必須スキル")`).Next().Text()
    optionalSkillText := doc.Find(`h3:contains("歓迎スキル")`).Next().Text()

    requiredSkillText = normalizeForSearchFeatures(requiredSkillText)
    optionalSkillText = normalizeForSearchFeatures(optionalSkillText)

    requiredSkills := salvageJobData(requiredSkillText, C.Required)
    optionalSkills := salvageJobData(optionalSkillText, C.Optional)

    return slices.Concat(requiredSkills, optionalSkills)
}

func (c *CallBacksEngineerFactory) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data)
    }
}

func (c *CallBacksEngineerFactory) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "body")
    }
}