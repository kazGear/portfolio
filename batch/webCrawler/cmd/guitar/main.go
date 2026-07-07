package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kazGear/portfolio/webCrawler/internal/repository"
	"github.com/kazGear/portfolio/webCrawler/internal/service"
	"github.com/kazGear/portfolio/webCrawler/pkg/db"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

func init() {
	envFile := os.Getenv("ENV_FILE")

	if envFile == "" { // ローカル環境なら空
		envFile = ".env.dev"

	}
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatalf(`godotenv load error: %v`, err)
	}
	log.Println("godotenv load OK.")
	log.Printf("loaded env file: %v\n", envFile)
}

func main() {
	// DBセットアップ
	database := db.Connect()
	defer database.Close()
	repository := repository.NewGuitarRepository(database)

	// クローラー起動
	service := service.NewGuitarCrawlerService(repository)
	service.RunCrawler()

	// 過去ログの整理
	logPath 		 := os.Getenv("LOGS_PATH")
	logsKeepCount, _ := strconv.Atoi(os.Getenv("LOGS_KEEP_COUNT"))
	utils.CleanupLogs(
		logPath,
		logsKeepCount,
	)
}