package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/repository"
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
		colly.MaxDepth(1),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1, // URL収集漏れが発生するため5に制限
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

    for pageId := 94800; pageId <= 95000; pageId++ {
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

var regCrowdworksTechPageId   = regexp.MustCompile(`\d{1,6}`)
var regCrowdworksTechMaxPrice = regexp.MustCompile(`\d{0,3},*\d{0,3},\d{0,3}`)

func (c *CallBacksCrowdworksTech) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)
        mutex := &sync.Mutex{}

        // apiからJSON取得 > struct化
        pageId   := regCrowdworksTechPageId.FindString(url)
        res, err := fetchApiData(
            fmt.Sprintf("https://tech.crowdworks.jp/api/v1/users/job_offers/%v/detail", pageId),
        )
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

        data[C.Url]         = url
        data[C.Title]       = jsonModel.DetailedTitle
        data[C.CompanyName] = jsonModel.ClientName // なぜか会社名だけ取得できない
        data[C.Location]    = ""

        data[C.MinSalaryAtHour]  = ""
        data[C.MinSalaryAtMonth] = ""
        data[C.MaxSalaryAtHour]  = ""
        maxSalaryAtMonth, _     := doc.Find(`meta[name="description"]`).Attr("content")
        maxSalaryAtMonth         = regCrowdworksTechMaxPrice.FindString(maxSalaryAtMonth)
        data[C.MaxSalaryAtMonth] = maxSalaryAtMonth

        data[C.Description]    = jsonModel.DetailedTitle + "\n" +
                                 jsonModel.SpecificWorkContent + "\n" +
                                 jsonModel.RelatedServicesProducts
        data[C.EmploymentType] = ""
        data[C.WorkPlace]      = ""
        // data[C.IsActive]       = "true"
        // data[C.SimilarityScore] =
        data[C.SourceSite]     = "CrowdWorks Tech"

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesCrowdWorksTech(data, data[C.Description])
        repository.InjectionJobFeaturesCrowdWorksTech(features, url)

        dataset = utils.LockedAppend(mutex, dataset, data)
        return dataset
    }
}

// 必要な情報を抽出する
func salvageFeaturesCrowdWorksTech(data map[string]string, target string) []*model.JobFeature {
    normalizedTarget := normalizeForSearchFeature(target)

    data[C.Location]       = salvageLocation(normalizedTarget)
    data[C.WorkPlace]      = salvageWorkPlace(normalizedTarget)
    data[C.EmploymentType] = salvageEmploymentType(normalizedTarget)

    return salvageJobData(normalizedTarget)
}

func salvageLocation(target string) string {
    for _, location := range locationDictionary {
        for _, locationName := range location.Keywords {
            if strings.Contains(target, locationName) {
                return location.Name
            }
        }
    }
    return ""
}

func salvageWorkPlace(target string) string {
    for _, workPlace := range workPlaces {
        if strings.Contains(target, workPlace) {
            return workPlace
        }
    }
    return ""
}

func salvageEmploymentType(target string) string {
    for _, employmentType := range employmentTypes {
        if strings.Contains(target, employmentType) {
            return employmentType
        }
    }
    return ""
}

func salvageJobData(target string) []*model.JobFeature {
    jobFeatures := make([]*model.JobFeature, 0, 10)

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

    jobFeatures = append(jobFeatures, languages...)
    jobFeatures = append(jobFeatures, frameworkLibraries...)
    jobFeatures = append(jobFeatures, databases...)
    jobFeatures = append(jobFeatures, clouds...)
    jobFeatures = append(jobFeatures, infrastructures...)
    jobFeatures = append(jobFeatures, tools...)
    jobFeatures = append(jobFeatures, tests...)
    jobFeatures = append(jobFeatures, architectures...)
    jobFeatures = append(jobFeatures, methodologies...)
    jobFeatures = append(jobFeatures, roles...)
    jobFeatures = append(jobFeatures, ais...)

    return jobFeatures
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