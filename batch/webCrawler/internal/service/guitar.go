package service

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/kazGear/portfolio/webCrawler/internal/repository"
	"github.com/kazGear/portfolio/webCrawler/internal/scraper"
	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type CrawlerService interface {
    RunCrawler()
}

type guitarCrawlerService struct {
    repository repository.Repository
}

func NewGuitarCrawlerService(repository repository.Repository) CrawlerService {
    return &guitarCrawlerService{ repository: repository }
}

type maker struct {
    name    string
    scraper scraper.Scraper
    funcs   scraper.GuitarCallbacks
}

func NewMaker(name string, scraper scraper.Scraper, funcs scraper.GuitarCallbacks) maker {
    return maker{ name, scraper, funcs}
}

func (s *guitarCrawlerService) RunCrawler() {
    // chromedpコンテキスト構築
    allocCtx, allocCancel := chromedp.NewExecAllocator(
        context.Background(),
        chromedp.DefaultExecAllocatorOptions[:]...,
    )
    defer allocCancel()
    parentCtx, parentCancel := chromedp.NewContext(allocCtx)
    defer parentCancel()

    // メーカーが増えたら追加
    makers := []maker {
        // NewMaker("ESP", scraper.NewScraperEsp(), scraper.NewCallBacksEsp()),
        // NewMaker("ESP_sig", scraper.NewScraperEspSig(), scraper.NewCallBacksEspSig()),
        NewMaker(".strandberg", scraper.NewScraperStrandberg(), scraper.NewCallBacksStrandberg()),
    }

    // スクレイピング + DB保存
    for _, maker := range makers {
        // ログ設定
        utils.LoggerInit(maker.name)
        log.Printf(constants.DecoLabel, "Started crawler " + maker.name)
        // 処理時間計測開始
        startTime := time.Now()

        maker.scraper.CollectLinks(parentCtx)
        guitars, err := maker.scraper.Scrape(maker.funcs, parentCtx)

        if err != nil { log.Println(err) }

        s.repository.UpsertAll(*guitars)

        log.Printf(constants.DecoLabel, "Finished crawler " + maker.name)
        log.Printf("Crawler processing time: %v\n", time.Since(startTime))
    }
}
