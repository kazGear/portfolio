package service

import (
	"fmt"
	"slices"
	"time"

	"github.com/kazGear/portfolio/goBatch/internal/batchMonitor/model"
	"github.com/kazGear/portfolio/goBatch/internal/batchMonitor/repository"
	"github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/kazGear/portfolio/goBatch/pkg/notify"
)

type BatchMonitor struct {
	Repository *repository.BatchMonitorRepository
}

func NewBatchMonitorService(repository *repository.BatchMonitorRepository) *BatchMonitor {
	return &BatchMonitor{Repository: repository}
}

func (b *BatchMonitor) Notify(discordWebHook string) error {
	batchLogs, err := b.Repository.SelectBatchExecution()

	if err != nil {
		return err
	}
	checkedLogs := make([]string, 0, len(constants.BatchNames))
	notifier    := notify.NewNotify(discordWebHook)
	notifyError := "通知の送信が失敗しました。\n%v"

	// 正常終了していなければ通知する
	checkedLogs, err = allBatchCheck(notifier, batchLogs, checkedLogs, notifyError)

	if err != nil {
		return err
	}
	// 処理されていないバッチを探す、あれば通知
	err = checkNotStarted(notifier, checkedLogs, notifyError)

	if err != nil {
		return err
	}
	return nil
}

func allBatchCheck(notifier    notify.Notify,
				   batchLogs   []*model.BatchExecution,
				   checkedLogs []string,
				   notifyError string,
) ([]string, error) {
	// 各バッチが正常に完了しているか確認
	for _, batchLog := range batchLogs {
		// 正常終了以外の状態である
		if slices.Contains(constants.BatchBadStatus, batchLog.Status) {
			// エラーメッセージ作成、通知
			message := notify.NewMessageDiscord()
			err 	:= notifier.Notify(message.CreateMessage(batchLog))

			if err != nil {
				return nil, fmt.Errorf(notifyError, err)
			}
		}
		checkedLogs = append(checkedLogs, batchLog.BatchName)
	}
	return checkedLogs, nil
}

func checkNotStarted(notifier 	 notify.Notify,
				     checkedLogs []string,
					 notifyError string,
) error {
	// NOT_STARTED: ログレコードが無い = 未処理
	for _, batchName := range constants.BatchNames {
		// 各バッチは確認済か
		if !slices.Contains(checkedLogs, batchName) {
			// エラーメッセージ作成
			m      := "バッチ処理が開始されていません。"
			source := &model.BatchExecution{
				ExecDate:  time.Now(),
				BatchName: batchName,
				Status:	   "NOT_STARTED",
				LogId: 	   -1,
				Message:   &m,
			}
			message := notify.NewMessageDiscord()

			// 通知
			err := notifier.Notify(message.CreateMessage(source))

			if err != nil {
				return fmt.Errorf(notifyError, err)
			}
		}
	}
	return nil
}