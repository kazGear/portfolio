package service

import (
	"context"
	"log"
	"sync"
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

type Maker struct {
    name    string
    scraper scraper.Scraper
    funcs   scraper.GuitarCallbacks
    logger  *log.Logger
}

func NewMaker(name string, scraper scraper.Scraper, funcs scraper.GuitarCallbacks) *Maker {
    return &Maker{
        name,
        scraper,
        funcs,
        utils.NewLogger(name),
    }
}

func (s *guitarCrawlerService) RunCrawler() {
    // メーカーが増えたら追加
    makers := []*Maker {
        NewMaker("ESP", scraper.NewScraperEsp(), scraper.NewCallBacksEsp()),
        NewMaker("ESP_sig", scraper.NewScraperEspSig(), scraper.NewCallBacksEspSig()),
        // NewMaker(".strandberg", scraper.NewScraperStrandberg(), scraper.NewCallBacksStrandberg()),
        // NewMaker("Gibson", scraper.NewScraperGibson(), scraper.NewCallBacksGibson()),
    }
    wg := &sync.WaitGroup{}
    queue := make(chan struct{}, 5) // 並列数制御

    // スクレイピング + DB保存
    for _, maker := range makers {
        maker := maker
        wg.Add(1)
        go func(maker Maker) {
            queue <- struct{}{}
            defer wg.Done()
            defer func() { <- queue }() // 次のワーカーへ

            // chromedpコンテキスト構築
            allocCtx, allocCancel := chromedp.NewExecAllocator(
                context.Background(),
                chromedp.DefaultExecAllocatorOptions[:]...,
            )
            defer allocCancel()
            parentCtx, parentCancel := chromedp.NewContext(allocCtx)
            defer parentCancel()

            maker.logger.Printf(constants.DecoLabel, "Started crawler " + maker.name)

            startTime := time.Now() // 処理時間計測開始

            maker.scraper.CollectLinks(parentCtx)
            guitars := maker.scraper.Scrape(maker.funcs, parentCtx)

            s.repository.UpsertAll(guitars)

            maker.logger.Printf(constants.DecoLabel, "Finished crawler " + maker.name)
            maker.logger.Printf("Crawler processing time: %v\n", time.Since(startTime))
        }(*maker)
    }
    wg.Wait()
}
