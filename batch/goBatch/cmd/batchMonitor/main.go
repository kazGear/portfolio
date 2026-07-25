package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/model"
	batchLoggerService "github.com/kazGear/portfolio/goBatch/internal/batchLogger/service"
	"github.com/kazGear/portfolio/goBatch/internal/batchMonitor/repository"
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

	// DBロガー
	dbLogger    := batchLoggerService.NewBatchLogger(database)
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

	// リポジトリ作成
	repository := repository.NewBatchMonitorRepository(database)

	// サービス作成・実行
	discordWebHook := os.Getenv("DISCORD_WEBHOOK_URL")
	batchMonitor   := batchMonitorService.NewBatchMonitorService(repository)
	batchMonitor.Notify(discordWebHook)

	timeSpan := time.Since(stopWatch)
	dbLogger.UpdateStatus(config, &timeSpan)

	log.Println("Finished batch monitor.")
}