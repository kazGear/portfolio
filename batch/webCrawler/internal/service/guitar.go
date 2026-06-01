package service

import (
	"log"

	"github.com/kazGear/portfolio/webCrawler/internal/repository"
	"github.com/kazGear/portfolio/webCrawler/internal/scraper"
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

    // メーカーが増えたら追加
    scrapers := []scraper.Scraper{
        scraper.NewESPGuitarScraper(),
    }
    // スクレイピング + DB保存
    for _, scraper := range scrapers {
        guitars, err := scraper.Scrape()

        if err != nil {
            log.Println(err)
        }

        _ = s.repository.UpsertAll(guitars) // エラー処理はupsert内で。
    }
}
