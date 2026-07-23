package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/model"
	batchLoggerService "github.com/kazGear/portfolio/goBatch/internal/batchLogger/service"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/repository"
	crawlerService "github.com/kazGear/portfolio/goBatch/internal/crawler/service"
	"github.com/kazGear/portfolio/goBatch/pkg/db"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

func init() {
	log.Println("Start guitar crawler.")

	envFile := os.Getenv("ENV_FILE")

	if envFile == "" { // ローカル環境なら空
		envFile = `C:\repository\portfolio\batch\goBatch\.env.dev`
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Skip env file: %v", envFile)
		log.Println("Use environment of compose.yaml.")
	} else {
		log.Printf("Loaded env file: %v", envFile)
	}
}

func main() {
	stopWatch := time.Now()

	// DBセットアップ
	database := db.Connect()
	defer database.Close()
	repository := repository.NewGuitarRepository(database)

	// DBロガー
	dbLogger := batchLoggerService.NewBatchLogger(database)
	config, err := dbLogger.InsertStartLog("GuitarCrawler")

	defer func(config *model.BatchConfig) {
		if r := recover(); r != nil {
			panic := fmt.Errorf("panic: %v\n", r)
			dbLogger.UpdateError(config, panic)
		}
	}(config)

	if err != nil {
		dbLogger.UpdateError(config, err)
		return
	}

	// クローラー起動
	crawler := crawlerService.NewGuitarCrawlerService(repository)
	crawler.RunCrawler()

	// 過去ログの整理
	logPath 		 := os.Getenv("LOGS_PATH")
	logsKeepCount, _ := strconv.Atoi(os.Getenv("LOGS_KEEP_COUNT"))
	utils.CleanupLogs(
		logPath,
		logsKeepCount,
	)

	timeSpan := time.Since(stopWatch)
	dbLogger.UpdateStatus(config, &timeSpan)

	log.Println("Finished guitar crawler.")
}