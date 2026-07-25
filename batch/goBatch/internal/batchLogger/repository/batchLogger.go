package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/model"
	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/repository/sql"
)

type BatchLoggerRepository struct {
    db *sqlx.DB
}

func NewBatchLoggerRepository(db *sqlx.DB) *BatchLoggerRepository {
    return &BatchLoggerRepository{ db: db }
}

func (b *BatchLoggerRepository) InsertStartLog(batchName string) (*model.BatchConfig, error) {
    params := &model.BatchLoggerParam{
        BatchName: batchName,
    }
    // 開始ログ挿入
    _, err := b.db.NamedExec(sql.InsertStartLog(), params)

    if err != nil {
        return nil, err
    }
    // バッチ情報を取得
    rows, err := b.db.NamedQuery(sql.SelectBatchConfig(), params)

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

func (b *BatchLoggerRepository) UpdateError(config *model.BatchConfig, err error) error {
    param := &model.BatchLoggerParam{
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

func (b *BatchLoggerRepository) UpdateStatus(config *model.BatchConfig, timeSpan *time.Duration) error {
    params := &model.BatchLoggerParam{
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