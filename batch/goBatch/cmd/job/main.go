package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/model"
	loggerRepository "github.com/kazGear/portfolio/goBatch/internal/batchLogger/repository"
	batchLoggerService "github.com/kazGear/portfolio/goBatch/internal/batchLogger/service"
	jobRepository "github.com/kazGear/portfolio/goBatch/internal/crawler/repository"
	crawlerService "github.com/kazGear/portfolio/goBatch/internal/crawler/service"
	"github.com/kazGear/portfolio/goBatch/pkg/db"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

func init() {
	log.Println("Start job crawler.")
	utils.LoadEnv()
}

func main() {
	stopWatch := time.Now()

	// DBセットアップ
	database := db.Connect()
	defer database.Close()

	// リポジトリ作成
	jobRepository := jobRepository.NewJobRepository(database)
	loggerRepository := loggerRepository.NewBatchLoggerRepository(database)

	// DBロガー
	dbLogger    := batchLoggerService.NewBatchLogger(*loggerRepository)
	config, err := dbLogger.InsertStartLog("JobCrawler")

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

	// サービス作成・起動
	crawler := crawlerService.NewJobCrawlerService(jobRepository)
	crawler.RunCrawler()

	// 過去ログの整理
	// logPath 		 := os.Getenv("LOGS_PATH")
	// logsKeepCount, _ := strconv.Atoi(os.Getenv("LOGS_KEEP_COUNT"))
	// utils.CleanupLogs(
	// 	logPath,
	// 	logsKeepCount,
	// )

	timeSpan := time.Since(stopWatch)
	dbLogger.UpdateStatus(config, &timeSpan)

	log.Println("Finished job crawler.")
}