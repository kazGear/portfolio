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

type CrawlerTechReach struct {
    name     string
    jScraper Crawler[*model.Job]
}

type CallBacksTechReach struct {
    funcs CallBacks
}

func NewScraperTechReach() Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
	})
    return &CrawlerTechReach{
        "TechReach",
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
        },
    }
}

func NewCallBacksTechReach() *CallBacksTechReach {
    return &CallBacksTechReach{
        CallBacks{},
    }
}

// CollectAttributesへ
var _parentCtxTechReach context.Context

func (c *CrawlerTechReach) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector              := c.jScraper.collector
    _parentCtxTechReach = parentCtx

    // クロールログ収集
    crawlStats := &crawlStats{}
    collectStatsCrawl(collector ,crawlStats)

    mutex   := &sync.Mutex{}

    // URL生成の設定
    pageIdFrom, pageIdTo := loadPageIdFromTo("PAGE_ID_FROM_TECH_REACH", "PAGE_ID_TO_TECH_REACH")
    visited              := make(map[string]struct{}, pageIdTo - pageIdFrom)

    validatePageIdFromTo(pageIdFrom, pageIdTo)

    // URL生成
    for pageId := pageIdFrom; pageId <= pageIdTo; pageId++ {
        url := fmt.Sprintf("https://tech-reach.jp/jobs/%v", pageId)
        isFirstVisit(mutex, url, visited)
    }

    loggingCrawlStats(c.name, crawlStats)

    c.jScraper.urls = utils.MapToSliceUrl(visited)
    return c.jScraper.urls, nil
}

func (c *CrawlerTechReach) Scrape(provider  PageProvider,
                                  parser    ModelParser[*model.Job],
                                  parentCtx context.Context,
) []*model.Job {
    jobs := c.jScraper.scrapeFrame(provider, parser, parentCtx)
    return jobs
}

func (c *CallBacksTechReach) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
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

func (c *CallBacksTechReach) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)

        description           := collectTextTechReach(doc)
        normalizedDescription := normalizeForSearchFeatures(description)

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesTechReach(doc)
        // 保存するべき案件か
        if len(features) <= 0 {
            return []map[string]string{}
        }
        repository.InjectionJobFeatures(features, url)

        // 案件のオプションを収集し、repositoryへ（このサイトは無し）


        data := map[string]string{}

        data[C.Url]         = url
        data[C.Title]       = doc.Find(`.m-info-head__ttl`).Text()
        data[C.Location]    = salvageLocation(normalizedDescription)

        minPrice, maxPrice      := getJobPrice(doc.Find(`dt:contains("単価")`).Next().Text())
        data[C.MinSalaryAtMonth] = strconv.Itoa(minPrice)
        data[C.MaxSalaryAtMonth] = strconv.Itoa(maxPrice)

        data[C.Description]    = description
        data[C.EmploymentType] = salvageEmploymentType(normalizedDescription)
        data[C.WorkPlace]      = salvageWorkPlace(normalizedDescription)
        data[C.SourceSite]     = C.TechReach

        data[C.UpdatedAt] = getOpenDateTechReach(doc)

        dataset = append(dataset, data)
        return dataset
    }
}

func collectTextTechReach(doc *goquery.Document) string {
    builder := &strings.Builder{}

    builder.WriteString(doc.Find(`.m-info-head__ttl`).Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`.detail-update`).Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`dt:contains("単価")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`dt:contains("業界")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`dt:contains("特徴")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`dt:contains("職務内容")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`dt:contains("ポジション")`).Next().Text())
    builder.WriteString("\n")

    doc.Find(`dt:contains("スキル")`).
        Next().Children().Children().Each(func(idx int, selector *goquery.Selection) {
            builder.WriteString(selector.Text())
            builder.WriteString("\n")
    })
    builder.WriteString(doc.Find(`dt:contains("勤務地")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`dt:contains("雇用形態")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`dt:contains("必須(MUST)")`).Next().Text())
    builder.WriteString("\n")
    builder.WriteString(doc.Find(`dt:contains("尚可(WANT)")`).Next().Text())
    builder.WriteString("\n")

    return builder.String()
}

func getOpenDateTechReach(doc *goquery.Document) string {
    dateText := doc.Find(`.detail-update`).Text()
    dateText  = _regGetDate.FindString(dateText)

    return dateText
}

func salvageFeaturesTechReach(doc *goquery.Document) []*model.JobFeature {
    requiredSkillText := doc.Find(`dt:contains("必須(MUST)")`).Next().Text()
    optionalSkillText := doc.Find(`dt:contains("尚可(WANT)")`).Next().Text()

    requiredSkillText = normalizeForSearchFeatures(requiredSkillText)
    optionalSkillText = normalizeForSearchFeatures(optionalSkillText)

    requiredSkills := salvageJobData(requiredSkillText, C.Required)
    optionalSkills := salvageJobData(optionalSkillText, C.Optional)

    return slices.Concat(requiredSkills, optionalSkills)
}

func (c *CallBacksTechReach) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data)
    }
}

func (c *CallBacksTechReach) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "body")
    }
}