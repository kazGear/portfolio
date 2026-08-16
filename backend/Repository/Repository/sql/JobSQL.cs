namespace Repository.Repository.sql
{
    /// <summary>
    /// SQL文格納クラス
    /// </summary>
    public static class JobSQL
    {
        public static string SelectJobs()
        {
            // 検索結果の案件カード用
            string SQL = @"
                SELECT v.*
                  FROM 
                  (
                   SELECT
                          max(j.id)                                AS Id,
                          max(j.url)                               AS Url,
                          max(j.title)                             AS Title,
                          max(j.location)                          AS Location,
                          max(j.min_salary_at_month)               AS MinSalaryAtMonth,
                          max(j.max_salary_at_month)               AS MaxSalaryAtMonth,
                          max(j.employment_type)                   AS EmploymentType,
                          max(j.work_place)                        AS WorkPlace,
                          max(j.source_site)                       AS SourceSite, 
                          max(j.created_at)                        AS CreatedAt,
                          max(j.updated_at)                        AS UpdatedAt,
                          string_agg(DISTINCT f.feature_name, ',') AS FeatureNamesCSV,
                          string_agg(DISTINCT o.option, ',')       AS OptionsCSV
 
                     FROM
                          t_jobs AS j
               INNER JOIN
                          t_job_features AS f
                       ON j.id = f.job_id
          LEFT OUTER JOIN
                          t_job_options AS o
                       ON j.id = o.job_id
 
                    WHERE TRUE

       /* 動的検索条件 AND title            iLIKE '%' || @title || '%'
                      AND location             = @location
                      AND min_salary_at_month >= @min_salary_at_month_specified_min
                      AND min_salary_at_month <= @min_salary_at_month_specified_max
                      AND max_salary_at_month >= @max_salary_at_month_specified_min
                      AND max_salary_at_month <= @max_salary_at_month_specified_max
                      AND work_place           = @work_place
                      AND source_site          = @source_site */

                 GROUP BY
                          j.url
                  ) AS v

                 WHERE TRUE

    /* 動的検索条件 AND EXISTS (
                           SELECT 1
                             FROM t_job_features AS f
                            WHERE v.id = f.job_id
                              AND f.feature_name = @feature_name
                     )
                     ... */

    /* 動的検索条件 AND EXISTS (
                           SELECT 1
                             FROM t_job_options AS o
                            WHERE v.id = o.job_id
                              AND o.option = @options
                     )
                     ... */

                 LIMIT 50 --:limit
                     ;
            ";
            return SQL;
        }

        public static string GetTotalCount(string conditions)
        {
            string SQL = @$"
                SELECT
                       count(*)
                  FROM
                       t_jobs
                 WHERE
                       TRUE

       /* 動的検索条件 AND title            iLIKE '%' || @title || '%'
                      AND location             = @location
                      AND min_salary_at_month >= @min_salary_at_month_specified_min
                      AND min_salary_at_month <= @min_salary_at_month_specified_max
                      AND max_salary_at_month >= @max_salary_at_month_specified_min
                      AND max_salary_at_month <= @max_salary_at_month_specified_max
                      AND work_place           = @work_place
                      AND source_site          = @source_site */

    /* 動的検索条件 AND EXISTS (
                           SELECT 1
                             FROM t_job_features AS f
                            WHERE v.id = f.job_id
                              AND f.feature_name = @feature_name
                     )
                     ... */

    /* 動的検索条件 AND EXISTS (
                           SELECT 1
                             FROM t_job_options AS o
                            WHERE v.id = o.job_id
                              AND o.option = @options
                     )
                     ... */
                     ;
            ";
            return SQL;
        }
    }
}
