package sql

func InsertStartLog() string {
	return `
   INSERT INTO
               t_batch_execution
               (batch_name, exec_date, status, start_at, end_at, message)
        VALUES
               (
                :batch_name,
                CURRENT_DATE,
                'STARTED',
                CURRENT_TIMESTAMP,
                NULL ,
                NULL
               )
             ;
    `
}

func SelectBatchConfig() string {
	return `
        SELECT
               exe.log_id                       AS log_id,
			   exe.start_at                     AS start_at,
			   config.batch_name                AS batch_name,
			   config.scheduled_time            AS scheduled_time,
			   config.start_delay_minutes       AS start_delay_minutes,
			   config.expected_duration_minutes AS expected_duration_minutes,
			   config.timeout_minutes           AS timeout_minutes
          FROM
               m_batch_config AS config
    INNER JOIN
               t_batch_execution AS exe
            ON exe.batch_name = config.batch_name

         WHERE
               exe.log_id = (SELECT max(log_id) FROM t_batch_execution)
           AND exe.batch_name = :batch_name
             ;
    `
}

func UpdateStatus() string {
	return `
        UPDATE
               t_batch_execution
           SET
               end_at  = CURRENT_TIMESTAMP,
               status  = :status,
               message = :message
         WHERE
               log_id = :log_id
             ;
    `
}