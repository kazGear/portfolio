namespace Repository.Repository.sql
{
    /// <summary>
    /// SQL文格納クラス
    /// </summary>
    public static class BatchLogSQL
    {
        public static string InsertStartLog()
        {
            string SQL = @"
                INSERT INTO
                            t_batch_execution 
                            (batch_name, exec_date, status, start_at, end_at, message)
                     VALUES 
                            (
                                @batch_name,
                                CURRENT_DATE,
                                'STARTED',
                                CURRENT_TIMESTAMP,
                                NULL ,
                                NULL 
                            )
                          ;
            ";
            return SQL;
        }

        public static string GetConfig()
        {
            string SQL = @"
                SELECT
                       exe.log_id                       AS LogId,
                       exe.start_at                     AS StartAt,
                       config.batch_name                AS BatchName,
                       config.scheduled_time            AS ScheduledTime,
                       config.start_delay_minutes       AS StartDelayMinutes,
                       config.expected_duration_minutes AS ExpectedDurationMinutes,
                       config.timeout_minutes           AS TimeoutMinutes 
                  FROM
                       m_batch_config AS config
            INNER JOIN
                       t_batch_execution AS exe
                    ON exe.batch_name = config.batch_name

                 WHERE
                       exe.log_id = (SELECT max(log_id) FROM t_batch_execution)
                   AND exe.batch_name = @batch_name 
                     ;
            ";
            return SQL;
        }

        public static string UpdateStatus()
        {
            string SQL = @"
                UPDATE
                       t_batch_execution
                   SET
                       end_at  = CURRENT_TIMESTAMP,
                       status  = @status,
                       message = @message
                 WHERE
                       log_id = @log_id
                     ;
            ";
            return SQL;
        }
    }
}
