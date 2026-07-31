package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

type CrawlerCrowdworksTech struct {
    jScraper Crawler[*model.Job]
}

type CallBacksCrowdworksTech struct {
    funcs CallBacks
}

func NewScraperCrowdworksTech(logger *log.Logger) Scraper[*model.Job] {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(4),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 5, // URL収集漏れが発生するため5に制限
	})
    return &CrawlerCrowdworksTech{
        Crawler[*model.Job]{
            collector: collector,
            mutex:     &sync.Mutex{},
            logger:    logger,
        },
    }
}

func NewCallBacksCrowdworksTech(logger *log.Logger) *CallBacksCrowdworksTech {
    return &CallBacksCrowdworksTech{
        CallBacks{
            logger: logger,
        },
    }
}

// CollectAttributesへ
var parentCtxCrowdworksTech context.Context

func (c *CrawlerCrowdworksTech) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector              := c.jScraper.collector
    parentCtxCrowdworksTech = parentCtx

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(collector ,crawlStats, c.jScraper.logger)

    // URL収集、クロール
    visited := make(map[string]struct{}, 120)
    mutex   := &sync.Mutex{}

    for pageId := 94730; pageId <= 94830; pageId++ {
        isFirstVisit(mutex, fmt.Sprintf("https://tech.crowdworks.jp/job_offers/%v", pageId), visited)
    }

    loggingCrawlStats(crawlStats, c.jScraper.logger)

    c.jScraper.urls = utils.MapToSliceUrl(visited)
    return c.jScraper.urls, nil
}

func (c *CrawlerCrowdworksTech) Scrape(provider  PageProvider,
                                       parser    ModelParser[*model.Job],
                                       parentCtx context.Context,
) []*model.Job {
    jobs := c.jScraper.scrapeFrame(provider, parser, parentCtx)
    return jobs
}

func (c *CallBacksCrowdworksTech) FetchDynamicPage(parentCtx context.Context) func(url string) (string, error) {
    return func(url string) (string, error) {
        if !isDetailPage(`^https://tech.crowdworks.jp/job_offers/\d+`, url) {
            return "", nil
        }
        // タブごとに独立した context を作る
        // tabCtx, tabCancel := chromedp.NewContext(parentCtx)
        // defer tabCancel()
        // // タブにだけ timeout を付ける
        // ctx, cancel := context.WithTimeout(tabCtx, 4*time.Second)
        // defer cancel()

        var html string

        // chromedp.Run(ctx,
        //     chromedp.Navigate(url),
        //     chromedp.WaitVisible("#main", chromedp.ByQuery), // 求める要素が出るまで待つ
        //     chromedp.Sleep(300 * time.Millisecond), // JSが動く猶予を与える
        //     tryWaitReady("h1.header_title"), // 必要な要素が生成されるのを待つ
        //     tryWaitReady(".tbl_spec"),
        //     tryWaitReady("p.detail_price"),
        //     chromedp.OuterHTML("html", &html, chromedp.ByQuery), // 最終的なHTML出力
        // )
        return html, nil
    }
}

func (c *CallBacksCrowdworksTech) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        // apiからJSON取得 > struct化
        res, err := fetchApiData("https://tech.crowdworks.jp/api/v1/users/job_offers/94830/detail")

        if err != nil {
            log.Println(err)
            return []map[string]string{}
        }
        defer res.Body.Close()

        var jsonModel model.ApiResponseCrowdworksTech
        err = json.NewDecoder(res.Body).Decode(&jsonModel)

        if err != nil {
            log.Printf(C.JsonDecodeError, err)
            return []map[string]string{}
        }

        data := map[string]string{}

        data[C.Url]         = "" // BuildModel側で注入
        data[C.Title]       = jsonModel.DetailedTitle
        data[C.CompanyName] = jsonModel.ClientName
        data[C.Location]    = ""

        data[C.MinSalaryAtHour]  = ""
        data[C.MinSalaryAtMonth] = ""
        data[C.MaxSalaryAtHour]  = ""
        maxSalaryAtMonth, _ := doc.Find(`meta[name="description"]`).Attr("content")
        data[C.MaxSalaryAtMonth] = maxSalaryAtMonth

        data[C.SkillsText]          = ""
        data[C.RequiredSkillsText]  = ""
        data[C.PreferredSkillsText] = ""

        data[C.Description]    = jsonModel.SpecificWorkContent + "\n" + jsonModel.RelatedServicesProducts
        data[C.EmploymentType] = ""
        data[C.WorkPlace]      = ""
        data[C.IsActive]       = "true"
        // data[C.SimilarityScore] =
        data[C.SourceSite]     = "CrowdWorks Tech"

        salvageCrowdWorksTech(data, data[C.Description])

        dataset = utils.LockedAppend(mutex, dataset, data)
        return dataset
    }
}

// 必要な情報を抽出する
func salvageCrowdWorksTech(data map[string]string, target string) {
    normalizedTarget := normalizeForJobFeature(target)

    salvageLocation(data, normalizedTarget)
    salvageWorkPlace(data, normalizedTarget)
    jobData := salvageJobData(normalizedTarget)

    data[C.SkillsText] = jobData
}

func salvageLocation(data map[string]string, target string) {
    for _, location := range locationDictionary {
        for _, locationName := range location.Keywords {
            if strings.Contains(target, locationName) {
                data[C.Location] = location.Name
                return
            }
        }
    }
    data[C.Location] = "不明"
}

func salvageWorkPlace(data map[string]string, target string) {
    for _, workPlace := range workPlaces {
        if strings.Contains(target, workPlace) {
            data[C.WorkPlace] = workPlace
            return
        }
    }
}

func salvageJobData(target string) string {
    var builder strings.Builder

    languages          := salvageFeatures(target, languageDictionary)
    frameworkLibraries := salvageFeatures(target, frameworkLibraryDictionary)
    databases          := salvageFeatures(target, databaseDictionary)
    clouds             := salvageFeatures(target, cloudDictionary)
    infrastructures    := salvageFeatures(target, infrastructureDictionary)
    tools              := salvageFeatures(target, toolDictionary)
    tests              := salvageFeatures(target, testDictionary)
    architectures      := salvageFeatures(target, architectureDictionary)
    methodologies      := salvageFeatures(target, methodologyDictionary)
    roles              := salvageFeatures(target, roleDictionary)
    ais                := salvageFeatures(target, aiDictionary)

    builder.WriteString(languages)
    builder.WriteString(frameworkLibraries)
    builder.WriteString(databases)
    builder.WriteString(clouds)
    builder.WriteString(infrastructures)
    builder.WriteString(tools)
    builder.WriteString(tests)
    builder.WriteString(architectures)
    builder.WriteString(methodologies)
    builder.WriteString(roles)
    builder.WriteString(ais)

    return builder.String()
}

func (c *CallBacksCrowdworksTech) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data, url, c.funcs.logger)
    }
}

func (c *CallBacksCrowdworksTech) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "body")
    }
}