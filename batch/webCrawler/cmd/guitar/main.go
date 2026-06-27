package main

import (
	"os"
	"strconv"

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

	logPath 		 := os.Getenv("LOGS_PATH")
	logsKeepCount, _ := strconv.Atoi(os.Getenv("LOGS_KEEP_COUNT"))
	utils.CleanupLogs(
		logPath,
		logsKeepCount,
	)
}