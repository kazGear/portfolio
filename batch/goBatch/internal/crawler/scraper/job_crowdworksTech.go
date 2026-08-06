package scraper

import (
	"context"
	"encoding/json"
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
var _parentCtxCrowdworksTech context.Context

func (c *CrawlerCrowdworksTech) CollectLinks(parentCtx context.Context) ([]string, error) {
    collector              := c.jScraper.collector
    _parentCtxCrowdworksTech = parentCtx

    // クロールログ収集
    crawlStats := &crawlStats{}
    statsCrawlLogs(collector ,crawlStats, c.jScraper.logger)

    mutex   := &sync.Mutex{}

    // URL生成の設定
    pageIdFrom, pageIdTo := loadPageIdFromTo("PAGE_ID_FROM_CROWDWORKS_TECH", "PAGE_ID_TO_CROWDWORKS_TECH")
    visited              := make(map[string]struct{}, pageIdTo - pageIdFrom)

    validatePageIdFromTo(pageIdFrom, pageIdTo)

    // URL生成
    for pageId := pageIdFrom; pageId <= pageIdTo; pageId++ {
        url := fmt.Sprintf("https://tech.crowdworks.jp/job_offers/%v", pageId)
        isFirstVisit(mutex, url, visited)
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

var _regCrowdworksTechPageId   = regexp.MustCompile(`\d{1,6}`)
var _regCrowdworksTechMaxPrice = regexp.MustCompile(`\d{0,3},*\d{0,3},\d{0,3}`)

func (c *CallBacksCrowdworksTech) CollectAttributes() func(doc *goquery.Document, url string) []map[string]string {
    return func(doc *goquery.Document, url string) []map[string]string {
        dataset := make([]map[string]string, 0, 1)

        // apiからJSON取得 > struct化
        pageId   := _regCrowdworksTechPageId.FindString(url)
        res, err := fetchApiData(
            fmt.Sprintf("https://tech.crowdworks.jp/api/v1/users/job_offers/%v/detail", pageId),
        )
        if err != nil {
            log.Println(err)
            return []map[string]string{}
        }
        defer res.Body.Close()

        var jsonModel model.ApiResponseCrowdworksTech

        if err := json.NewDecoder(res.Body).Decode(&jsonModel); err != nil {
            log.Printf(C.JsonDecodeError, err)
            return []map[string]string{}
        }

        data := map[string]string{}

        data[C.Url]         = url
        data[C.Title]       = jsonModel.DetailedTitle
        data[C.CompanyName] = jsonModel.ClientName // なぜか会社名だけ取得できない
        data[C.Location]    = "" // 別関数で抽出

        maxSalaryAtMonth, _     := doc.Find(`meta[name="description"]`).Attr("content")
        maxSalaryAtMonth         = _regCrowdworksTechMaxPrice.FindString(maxSalaryAtMonth)
        data[C.MinSalaryAtMonth] = maxSalaryAtMonth
        data[C.MaxSalaryAtMonth] = maxSalaryAtMonth

        data[C.Description]    = jsonModel.DetailedTitle + "\n" +
                                 jsonModel.SpecificWorkContent + "\n" +
                                 jsonModel.RelatedServicesProducts
        data[C.EmploymentType] = "" // 別関数で抽出
        data[C.WorkPlace]      = "" // 別関数で抽出
        // data[C.IsActive]       = isActiveCrowdworksTech(doc)
        // data[C.SimilarityScore] =
        data[C.SourceSite]     = C.CrowdWorksTech

        // 案件の特徴を収集し、repositoryへ
        features := salvageFeaturesCrowdWorksTech(data, data[C.Description])
        repository.InjectionJobFeatures(features, url)

        // 案件のオプションを収集し、repositoryへ
        options := make([]*model.JobOption, 0, 20)

        doc.Find(".job-point").Each(func(idx int, selector *goquery.Selection) {
            option := &model.JobOption{
                JobId: -1,
                Option: selector.Text(),
            }
            options = append(options, option)
        })
        repository.InjectionJobOptions(options, url)

        dataset = append(dataset, data)
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

func (c *CallBacksCrowdworksTech) BuildModel(url string) func(data map[string]string) *model.Job {
    return func(data map[string]string) *model.Job {
        return buildJobFrame(data, c.funcs.logger)
    }
}

func (c *CallBacksCrowdworksTech) IsStaticPage() func(html string) bool {
    return func(html string) bool {
        return strings.Contains(html, "body")
    }
}