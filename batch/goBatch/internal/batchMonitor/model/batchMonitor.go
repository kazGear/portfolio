package model

import "time"

type BatchExecution struct {
	LogId     int64      `db:"log_id"`
	BatchName string     `db:"batch_name"`
	ExecDate  time.Time  `db:"exec_date"`
	Status    string     `db:"status"`
	StartAt   time.Time  `db:"start_at"`
	EndAt     *time.Time `db:"end_at"`
	Message   *string    `db:"message"`
}