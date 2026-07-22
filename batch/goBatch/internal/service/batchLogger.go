package service

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kazGear/portfolio/goBatch/internal/model"
	"github.com/kazGear/portfolio/goBatch/internal/repository"
)

type BatchLogger struct {
    db *sqlx.DB
}

func NewBatchLogger(db *sqlx.DB) *BatchLogger {
    return &BatchLogger{ db: db }
}

func (b *BatchLogger) InsertStartLog(batchName string) (*model.BatchConfig, error) {
    repository := repository.NewBatchLoggerRepository(b.db)
    config, err := repository.InsertStartLog(batchName)

    if err != nil {
        return nil, err
    }
    return config, nil
}

func (b *BatchLogger) UpdateError(config *model.BatchConfig, err error) error {
    repository := repository.NewBatchLoggerRepository(b.db)
    err = repository.UpdateError(config, err)

    if err != nil {
        return err
    }
    return nil
}

func (b *BatchLogger) UpdateStatus(config *model.BatchConfig, timeSpan *time.Duration) error {
    repository := repository.NewBatchLoggerRepository(b.db)
    err := repository.UpdateStatus(config, timeSpan)

    if err != nil {
        return err
    }
    return nil
}