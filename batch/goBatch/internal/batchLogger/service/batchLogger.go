package service

import (
	"time"

	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/model"
	"github.com/kazGear/portfolio/goBatch/internal/batchLogger/repository"
)

type BatchLogger struct {
    repository repository.BatchLoggerRepository
}

func NewBatchLogger(repository repository.BatchLoggerRepository) *BatchLogger {
    return &BatchLogger{ repository: repository }
}

func (b *BatchLogger) InsertStartLog(batchName string) (*model.BatchConfig, error) {
    config, err := b.repository.InsertStartLog(batchName)

    if err != nil {
        return nil, err
    }
    return config, nil
}

func (b *BatchLogger) UpdateError(config *model.BatchConfig, err error) error {
    err = b.repository.UpdateError(config, err)

    if err != nil {
        return err
    }
    return nil
}

func (b *BatchLogger) UpdateStatus(config *model.BatchConfig, timeSpan *time.Duration) error {
    err := b.repository.UpdateStatus(config, timeSpan)

    if err != nil {
        return err
    }
    return nil
}