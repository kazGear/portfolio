package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/internal/repository/sql"
)

type batchLoggerRepository struct {
    db *sqlx.DB
}

type batchLoggerParam struct {
	LogId     int64  `db:"log_id"`
	BatchName string `db:"batch_name"`
	Status    string `db:"status"`
	Message   string `db:"message"`
}

func NewBatchLoggerRepository(db *sqlx.DB) *batchLoggerRepository {
    return &batchLoggerRepository{ db: db }
}

func (b *batchLoggerRepository) InsertStartLog(batchName string) (*model.BatchConfig, error) {
    params := &batchLoggerParam{
        BatchName: batchName,
    }
    // 開始ログ挿入
    _, err := b.db.NamedExec(sql.InsertStartLog(), params)

    if err != nil {
        return nil, err
    }
    // バッチ情報を取得
    rows, err := b.db.NamedQuery(sql.GetConfig(), params)

    if err != nil {
        return nil, err
    }
    defer rows.Close()
    rows.Next()

    var config model.BatchConfig
    err = rows.StructScan(&config)

    if err != nil {
        return nil, err
    }
    return &config, nil
}

func (b *batchLoggerRepository) UpdateError(config *model.BatchConfig, err error) error {
    param := &batchLoggerParam{
        LogId:   config.LogId,
        Status:  "ERROR",
        Message: err.Error(),
    }
    _, err = b.db.NamedExec(sql.UpdateStatus(), param)

    if err != nil {
        return err
    }
    return nil
}

func (b *batchLoggerRepository) UpdateStatus(config *model.BatchConfig, timeSpan *time.Duration) error {
    params := &batchLoggerParam{
        LogId:   config.LogId,
        Message: "",
    }

    if config.TimeoutMinutes < int(timeSpan.Minutes()) {
        params.Status = "TIMEOUT"
    } else if config.ExpectedDurationMinutes < int(timeSpan.Minutes()) {
        params.Status = "SLOW"
    } else {
        params.Status = "SUCCESS"
    }
    _, err := b.db.NamedExec(sql.UpdateStatus(), params)

    if err != nil {
        return err
    }
    return nil
}