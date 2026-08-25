package service

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/repository"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/scraper"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
)

type guitarCrawlerService struct {
    repository repository.GuitarRepository
}

func NewGuitarCrawlerService(repository repository.GuitarRepository) CrawlerService {
    return &guitarCrawlerService{ repository: repository }
}

type Maker struct {
    name     string
    scraper  scraper.Scraper[*model.Guitar]
    provider scraper.PageProvider
    parser   scraper.ModelParser[*model.Guitar]
}

func NewMaker(name     string,
              scraper  scraper.Scraper[*model.Guitar],
              provider scraper.PageProvider,
              parser   scraper.ModelParser[*model.Guitar],
) *Maker {
    return &Maker{ name, scraper, provider, parser }
}

func (g *guitarCrawlerService) RunCrawler() {
    envName            := "PARALLEL_COUNT_GUITAR"
    parallelCount, err := strconv.Atoi(os.Getenv(envName))

    if err != nil || parallelCount <= 0 {
        log.Panicf("%v must be a positive integer.", envName)
    }

    queue := make(chan struct{}, parallelCount) // 並列数制御

    makers := makersFactory()
    wg     := &sync.WaitGroup{}

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

            log.Printf(C.DecoLabel, "Started guitar crawler " + maker.name)

            startTime := time.Now() // 処理時間計測開始

            // クローラー起動
            maker.scraper.CollectLinks(parentCtx)
            guitars := maker.scraper.Scrape(maker.provider, maker.parser, parentCtx)
            okCnt, ngCnt, errs := g.repository.Save(guitars)

            // ログ
            log.Printf("[Upsert result %v]: OK %v 件, NG %v 件", maker.name, okCnt, ngCnt)

            for _, err := range errs {
                log.Println(err)
            }
            log.Printf(C.DecoLabel, "Finished guitar crawler " + maker.name)
            log.Printf("%v crawler processing time: %v\n", maker.name, time.Since(startTime))
        }(*maker)
    }
    wg.Wait()
}

// 各種メーカー作成
func makersFactory() map[string]*Maker {
    makers := map[string]*Maker{}

    makerName := "Momose"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperMomose(),
        scraper.NewCallBacksMomose(),
        scraper.NewCallBacksMomose(),
    )

    makerName = "ESP-sig"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperEspSig(),
        scraper.NewCallBacksEspSig(), // callbacksは複数のインターフェイスを実装
        scraper.NewCallBacksEspSig(),
    )

    makerName = "ESP"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperEsp(),
        scraper.NewCallBacksEsp(),
        scraper.NewCallBacksEsp(),
    )

    makerName = "Gibson"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperGibson(),
        scraper.NewCallBacksGibson(),
        scraper.NewCallBacksGibson(),
    )

    makerName = ".strandberg"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperStrandberg(),
        scraper.NewCallBacksStrandberg(),
        scraper.NewCallBacksStrandberg(),
    )

    makerName = "Ibanez"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperIbanez(),
        scraper.NewCallBacksIbanez(),
        scraper.NewCallBacksIbanez(),
    )

    makerName = "PRS"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperPRS(),
        scraper.NewCallBacksPRS(),
        scraper.NewCallBacksPRS(),
    )

    makerName = "SCHECTER"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperSchecter(),
        scraper.NewCallBacksSchecter(),
        scraper.NewCallBacksSchecter(),
    )

    makerName = "ZEMAITIS"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperZemaitis(),
        scraper.NewCallBacksZemaitis(),
        scraper.NewCallBacksZemaitis(),
    )

    makerName = "MusicMan"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperMusicMan(),
        scraper.NewCallBacksMusicMan(),
        scraper.NewCallBacksMusicMan(),
    )

    makerName = "Fender"
    makers[makerName] = NewMaker(
        makerName,
        scraper.NewScraperFender(),
        scraper.NewCallBacksFender(),
        scraper.NewCallBacksFender(),
    )

    return makers
}