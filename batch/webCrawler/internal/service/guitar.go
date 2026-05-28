package service

import (
	"log"

	"github.com/kazGear/portfolio/webCrawler/internal/repository"
	"github.com/kazGear/portfolio/webCrawler/internal/scraper"
	"github.com/kazGear/portfolio/webCrawler/pkg/db"
)

func RunGuitarCrawler() {
    log.Println("ギタークローラー開始")

    guitars, err := scraper.ScrapeGuitars()
    if err != nil {
        log.Println("スクレイピング失敗:", err)
        return
    }

    database := db.Connect()
    defer database.Close()

    err = repository.SaveGuitars(database, guitars)
    if err != nil {
        log.Println("DB保存失敗:", err)
        return
    }

    log.Println("ギタークローラー完了")
}
