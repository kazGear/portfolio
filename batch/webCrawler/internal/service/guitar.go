package service

import (
	"log"
	"time"

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

func (s *guitarCrawlerService) RunCrawler() {
    utils.LoggerInit("esp")
    log.Printf(constants.DecoLabel, "Started crawler.")

    // 処理時間計測
    startTime := time.Now()
    defer func() { log.Printf("Crawler processing time: %v\n", time.Since(startTime)) }()
    // メーカーが増えたら追加
    scrapers := []scraper.Scraper{
        scraper.NewESPGuitarScraper(),
    }
    // chrome ターミネート
    for i := len(scrapers)-1; i >= 0; i-- {
        defer scrapers[i].Cancel()
    }
    // スクレイピング + DB保存
    for _, scraper := range scrapers {
        scraper.CollectLinks()
        guitars, err := scraper.Scrape()

        if err != nil {
            log.Println(err)
        }

        _ = s.repository.UpsertAll(guitars) // エラー処理はupsert内で。
    }
    log.Printf(constants.DecoLabel, "Finished crawler.")
}
