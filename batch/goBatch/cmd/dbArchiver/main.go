package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/model"
	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/repository"
	batchLoggerService "github.com/kazGear/portfolio/goBatch/internal/batchLogger/service"
	dbArchiverService "github.com/kazGear/portfolio/goBatch/internal/dbArchiver/service"
	"github.com/kazGear/portfolio/goBatch/pkg/db"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

func init() {
	log.Println("Start DB archiver.")
	utils.LoadEnv()
}

func main() {
	stopWatch := time.Now()

	// DBセットアップ
	database := db.Connect()
	defer database.Close()

	// リポジトリ作成
	repository := repository.NewBatchLoggerRepository(database)

	// DBロガー
	dbLogger    := batchLoggerService.NewBatchLogger(*repository)
	config, err := dbLogger.InsertStartLog("DbArchiver")

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
	dbArchiverService := dbArchiverService.NewDbArchiverService()
	err = dbArchiverService.Archive()

	if err != nil {
		dbLogger.UpdateError(config, err)
		return
	}

	timeSpan := time.Since(stopWatch)
	dbLogger.UpdateStatus(config, &timeSpan)

	log.Println("Finished DB archiver.")
}