package service

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/kazGear/portfolio/webCrawler/internal/repository"
	"github.com/kazGear/portfolio/webCrawler/internal/scraper"
	C "github.com/kazGear/portfolio/webCrawler/pkg/constants"
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
    name     string
    scraper  scraper.Scraper
    provider scraper.PageProvider
    parser   scraper.GuitarParser
    logger   *log.Logger
}

func NewMaker(name string,
              scraper  scraper.Scraper,
              provider scraper.PageProvider,
              parser   scraper.GuitarParser,
              logger *log.Logger,
) *Maker {
    return &Maker{ name, scraper, provider, parser ,logger }
}

func (g *guitarCrawlerService) RunCrawler() {
    makers := makersFactory()

    wg := &sync.WaitGroup{}
    queue := make(chan struct{}, 10) // 並列数制御

    // クロール + スクレイピング + DB保存
    for _, maker := range makers {
        maker := maker
        wg.Add(1)

        go func(maker Maker) {
            queue <- struct{}{}
            defer wg.Done()
            defer func() { <- queue }() // 次のワーカーへ

            // chromedpコンテキスト構築
            cancelAlloc, cancelParent, parentCtx := createChromedpCtx()
            defer cancelAlloc()
            defer cancelParent()

            maker.logger.Printf(C.DecoLabel, "Started crawler " + maker.name)

            startTime := time.Now() // 処理時間計測開始

            // クローラー起動
            maker.scraper.CollectLinks(parentCtx)
            guitars := maker.scraper.Scrape(maker.provider, maker.parser, parentCtx)
            okCnt, ngCnt, errs := g.repository.UpsertAll(guitars)

            maker.logger.Printf("[Upsert result]: OK %v 件, NG %v 件", okCnt, ngCnt)

            for _, err := range errs {
                maker.logger.Println(err)
            }
            maker.logger.Printf(C.DecoLabel, "Finished crawler " + maker.name)
            maker.logger.Printf("Crawler processing time: %v\n", time.Since(startTime))
        }(*maker)
    }
    wg.Wait()
}

// 各種メーカー作成
func makersFactory() map[string]*Maker {
    makers := map[string]*Maker{}

    makerName := "ESP"
    logger    := utils.NewLogger(makerName)
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperEsp(logger),
        scraper.NewCallBacksEsp(logger), // callbacksは複数のインターフェイスを実装
        scraper.NewCallBacksEsp(logger),
        logger,
    )

    // makerName = "ESP_sig"
    // logger    = utils.NewLogger(makerName)
    // makers[makerName] = NewMaker(
    //     makerName,
    //     scraper.NewScraperEspSig(logger),
    //     scraper.NewCallBacksEspSig(logger),
    //     scraper.NewCallBacksEspSig(logger),
    //     logger,
    // )

    // makerName = ".strandberg"
    // logger    = utils.NewLogger(makerName)
    // makers[makerName] = NewMaker(
    //     makerName,
    //     scraper.NewScraperStrandberg(logger),
    //     scraper.NewCallBacksStrandberg(logger),
    //     scraper.NewCallBacksStrandberg(logger),
    //     logger,
    // )

    // makerName = "Gibson"
    // logger    = utils.NewLogger(makerName)
    // makers[makerName] = NewMaker(
    //     makerName,
    //     scraper.NewScraperGibson(logger),
    //     scraper.NewCallBacksGibson(logger),
    //     scraper.NewCallBacksGibson(logger),
    //     logger,
    // )

    // makerName = "Ibanez"
    // logger    = utils.NewLogger(makerName)
    // makers[makerName] = NewMaker(
    //     makerName,
    //     scraper.NewScraperIbanez(logger),
    //     scraper.NewCallBacksIbanez(logger),
    //     scraper.NewCallBacksIbanez(logger),
    //     logger,
    // )

    // makerName = "PRS"
    // logger    = utils.NewLogger(makerName)
    // makers[makerName] = NewMaker(
    //     makerName,
    //     scraper.NewScraperPRS(logger),
    //     scraper.NewCallBacksPRS(logger),
    //     scraper.NewCallBacksPRS(logger),
    //     logger,
    // )

    return makers
}

// chromedp環境構築
func createChromedpCtx() (cancelAlloc context.CancelFunc,
                          cancelParent context.CancelFunc,
                          parentCtx context.Context,
) {
    // 動的ページ取得のためのchromedpコンテキスト構築
    allocCtx, allocCancel := chromedp.NewExecAllocator(
        context.Background(),
        chromedp.DefaultExecAllocatorOptions[:]...,
    )
    parentCtx, parentCancel := chromedp.NewContext(allocCtx)

    return allocCancel, parentCancel, parentCtx
}

// chromedp環境構築（実際にchromeを起動して挙動を確認できる。clickされているか？等）
func createChromedpCtxDebug() (cancelAlloc context.CancelFunc,
                               cancelParent context.CancelFunc,
                               parentCtx context.Context,
) {
    // 動的ページ取得のためのchromedpコンテキスト構築
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", false),
    )

    allocCtx, allocCancel := chromedp.NewExecAllocator(
        context.Background(),
        opts...
    )
    parentCtx, parentCancel := chromedp.NewContext(allocCtx)

    return allocCancel, parentCancel, parentCtx
}