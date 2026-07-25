package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/model"
	batchLoggerRepository "github.com/kazGear/portfolio/goBatch/internal/batchLogger/repository"
	batchLoggerService "github.com/kazGear/portfolio/goBatch/internal/batchLogger/service"
	batchMonitorRepository "github.com/kazGear/portfolio/goBatch/internal/batchMonitor/repository"
	batchMonitorService "github.com/kazGear/portfolio/goBatch/internal/batchMonitor/service"
	"github.com/kazGear/portfolio/goBatch/pkg/db"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

func init() {
	log.Println("Start batch monitor.")
	utils.LoadEnv()
}

func main() {
	stopWatch := time.Now()

	// DBセットアップ
	database := db.Connect()
	defer database.Close()

	// リポジトリ作成
	batchMonitorRepository := batchMonitorRepository.NewBatchMonitorRepository(database)
	batchLoggerRepository  := batchLoggerRepository.NewBatchLoggerRepository(database)

	// DBロガー
	dbLogger    := batchLoggerService.NewBatchLogger(*batchLoggerRepository)
	config, err := dbLogger.InsertStartLog("BatchMonitor")

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

	// サービス作成・実行
	discordWebHook := os.Getenv("DISCORD_WEBHOOK_URL")
	batchMonitor   := batchMonitorService.NewBatchMonitorService(batchMonitorRepository)
	batchMonitor.Notify(discordWebHook)

	timeSpan := time.Since(stopWatch)
	dbLogger.UpdateStatus(config, &timeSpan)

	log.Println("Finished batch monitor.")
}