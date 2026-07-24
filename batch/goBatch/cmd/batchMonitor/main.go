package main

import (
	"log"
	"os"

	"github.com/kazGear/portfolio/goBatch/internal/batchMonitor/repository"
	"github.com/kazGear/portfolio/goBatch/internal/batchMonitor/service"
	"github.com/kazGear/portfolio/goBatch/pkg/db"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

func init() {
	log.Println("Start batch monitor.")
	utils.LoadEnv()
}

func main() {
	// DBセットアップ
	database := db.Connect()
	defer database.Close()

	// リポジトリ作成
	repository := repository.NewBatchMonitorRepository(database)

	// サービス作成・実行
	discordWebHook := os.Getenv("DISCORD_WEBHOOK_URL")
	service        := service.NewBatchMonitorService(repository)
	service.Notify(discordWebHook)

	log.Println("Finished batch monitor.")
}