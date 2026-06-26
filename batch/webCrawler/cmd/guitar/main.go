package main

import (
	"github.com/joho/godotenv"
	"github.com/kazGear/portfolio/webCrawler/internal/repository"
	"github.com/kazGear/portfolio/webCrawler/internal/service"
	"github.com/kazGear/portfolio/webCrawler/pkg/db"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

func init() {
	godotenv.Load()
}

func main() {
	// DBセットアップ
	database := db.Connect()
	defer database.Close()
	repository := repository.NewGuitarRepository(database)
	// // クローラー起動
	service := service.NewGuitarCrawlerService(repository)
	service.RunCrawler()

	utils.CleanupLogs(`../../logs/`, 5)
}