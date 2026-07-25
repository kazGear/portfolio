package model

import "time"

type BatchConfig struct {
	LogId                   int64     `db:"log_id"                    json:"log_id"`
	StartAt                 time.Time `db:"start_at"                  json:"start_at"`
	BatchName               string    `db:"batch_name"                json:"batch_name"`
	ScheduledTime           time.Time `db:"scheduled_time"            json:"scheduled_time"`
	StartDelayMinutes       int       `db:"start_delay_minutes"       json:"start_delay_minutes"`
	ExpectedDurationMinutes int       `db:"expected_duration_minutes" json:"expected_duration_minutes"`
	TimeoutMinutes          int       `db:"timeout_minutes"           json:"timeout_minutes"`
}

type BatchLoggerParam struct {
	LogId     int64  `db:"log_id"`
	BatchName string `db:"batch_name"`
	Status    string `db:"status"`
	Message   string `db:"message"`
}