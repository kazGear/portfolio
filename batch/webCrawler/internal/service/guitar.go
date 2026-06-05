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

    // 処理時間計測
    startTime := time.Now()
    defer func() { log.Printf("Crawler processing time: %v\n", time.Since(startTime)) }()
    // メーカーが増えたら追加
    makers := []maker {
        // NewMaker("ESP", scraper.NewEspScraper(), scraper.NewCallBacksEsp()),
        NewMaker("ESP_sig", scraper.NewEspSigScraper(), scraper.NewCallBacksEspSig()),
    }
    // スクレイピング + DB保存
    for _, maker := range makers {
        utils.LoggerInit(maker.name)
        log.Printf(constants.DecoLabel, "Started crawler.")

        maker.scraper.CollectLinks()
        guitars, err := maker.scraper.Scrape(maker.funcs, parentCtx)

        if err != nil { log.Println(err) }

        _ = s.repository.UpsertAll(*guitars) // エラー処理はupsert内で。
    }
    log.Printf(constants.DecoLabel, "Finished crawler.")
}
