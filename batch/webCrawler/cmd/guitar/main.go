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
	log.Println("Start guitar crawler.")

	envFile := os.Getenv("ENV_FILE")

	if envFile == "" { // ローカル環境なら空
		envFile = ".env.dev"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Skip env file: %v", envFile)
		log.Println("Use environment of compose.yaml.")
	} else {
		log.Printf("Loaded env file: %v", envFile)
	}
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
	log.Println("Finished guitar crawler.")
}