package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/model"
	loggerRepository "github.com/kazGear/portfolio/goBatch/internal/batchLogger/repository"
	batchLoggerService "github.com/kazGear/portfolio/goBatch/internal/batchLogger/service"
	guitarRepository "github.com/kazGear/portfolio/goBatch/internal/crawler/repository"
	crawlerService "github.com/kazGear/portfolio/goBatch/internal/crawler/service"
	"github.com/kazGear/portfolio/goBatch/pkg/db"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

func init() {
	log.Println("Start guitar crawler.")
	utils.LoadEnv()
}

func main() {
	stopWatch := time.Now()

	// DBセットアップ
	database := db.Connect()
	defer database.Close()

	// リポジトリ作成
	guitarRepository := guitarRepository.NewGuitarRepository(database)
	loggerRepository := loggerRepository.NewBatchLoggerRepository(database)

	// DBロガー
	dbLogger    := batchLoggerService.NewBatchLogger(*loggerRepository)
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

	// サービス作成・起動
	crawler := crawlerService.NewGuitarCrawlerService(guitarRepository)
	crawler.RunCrawler()

	// 過去ログの整理
	logPath 		 := os.Getenv("LOGS_PATH_GUITAR")
	logsKeepCount, _ := strconv.Atoi(os.Getenv("LOGS_KEEP_COUNT"))
	utils.CleanupLogs(
		logPath,
		logsKeepCount,
	)

	timeSpan := time.Since(stopWatch)
	dbLogger.UpdateStatus(config, &timeSpan)

	log.Println("Finished guitar crawler.")
}