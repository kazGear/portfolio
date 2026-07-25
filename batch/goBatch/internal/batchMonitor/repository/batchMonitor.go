package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kazGear/portfolio/goBatch/internal/batchMonitor/model"
	"github.com/kazGear/portfolio/goBatch/internal/batchMonitor/repository/sql"
	"github.com/kazGear/portfolio/goBatch/pkg/constants"
)

type BatchMonitorRepository struct {
    db *sqlx.DB
}

func NewBatchMonitorRepository(db *sqlx.DB) *BatchMonitorRepository {
    return &BatchMonitorRepository{ db: db }
}

type batchMonitorParam struct {
    RecordAmount int `db:"record_amount"`
}

func (b *BatchMonitorRepository) SelectBatchExecution() ([]*model.BatchExecution, error) {
    // 必要な数だけレコードを取得
    params := &batchMonitorParam{
        RecordAmount: len(constants.BatchNames),
    }
    // バッチログを取得
    rows, err := b.db.NamedQuery(sql.SelectBatchExecution(), params)

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    batchLogs := make([]*model.BatchExecution, 0, len(constants.BatchNames))

    // structへ変換
    for ; rows.Next(); {
        var batchLog model.BatchExecution
        err = rows.StructScan(&batchLog)

        if err != nil {
            return nil, err
        }
        batchLogs = append(batchLogs, &batchLog)
    }
    return batchLogs, nil
}