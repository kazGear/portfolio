package sql

func SelectBatchExecution() string {
	return `
        SELECT
               log_id,
               batch_name,
               exec_date,
               status,
               start_at,
               end_at,
               message
          FROM
               t_batch_execution
      ORDER BY
               start_at DESC
         LIMIT
               :record_amount  -- 各バッチが1日1回実行される前提で、直近のバッチ実行ログをバッチ数分取得する
    `
}