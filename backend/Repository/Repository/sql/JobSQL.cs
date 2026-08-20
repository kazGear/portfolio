namespace Repository.Repository.sql
{
    /// <summary>
    /// SQL文格納クラス
    /// </summary>
    public static class JobSQL
    {
        public static string SelectJobs(string conditions, string featureConditions)
        {
            // 検索結果の案件カード用
            string SQL = @$"
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
                         {conditions}

       /* 動的検索条件 AND title            iLIKE '%' || @title || '%'
                      AND location              = @location
                      AND min_salary_at_month  >= @min_salary_at_month_specified_min
                      AND min_salary_at_month  <= @min_salary_at_month_specified_max
                      AND max_salary_at_month  >= @max_salary_at_month_specified_min
                      AND max_salary_at_month  <= @max_salary_at_month_specified_max
                      AND work_place            = @work_place
                      AND source_site           = @source_site 
                      AND NOW() - j.updated_at <= '3 month' */

                 GROUP BY
                          j.url
                  ) AS v

                 WHERE TRUE
                      {featureConditions}

    /* 動的検索条件 AND EXISTS (
                           SELECT 1
                             FROM t_job_features AS f
                            WHERE v.id = f.job_id
                              AND f.feature_name = @feature_name
                     )
                     ... */

                 LIMIT
                       @page_size
                OFFSET
                       (@page - 1) * @page_size -- ページネーション
                     ;
            ";
            return SQL;
        }

        public static string GetTotalCount(string conditions, string featureConditions)
        {
            string SQL = @$"
                SELECT
                       count(*)
                  FROM
                       t_jobs AS v -- サブクエリのテーブル名に合わせておく
                 WHERE
                       TRUE
                      {conditions}

       /* 動的検索条件 AND title            iLIKE '%' || @title || '%'
                      AND location              = @location
                      AND min_salary_at_month  >= @min_salary_at_month_specified_min
                      AND min_salary_at_month  <= @min_salary_at_month_specified_max
                      AND max_salary_at_month  >= @max_salary_at_month_specified_min
                      AND max_salary_at_month  <= @max_salary_at_month_specified_max
                      AND work_place            = @work_place
                      AND source_site           = @source_site 
                      AND NOW() - j.updated_at <= '3 month' */

                      {featureConditions}
    /* 動的検索条件 AND EXISTS (
                           SELECT 1
                             FROM t_job_features AS f
                            WHERE v.id = f.job_id
                              AND f.feature_name = @feature_name
                     )
                     ... */
                     ;
            ";
            return SQL;
        }

        public static string GetFeatures()
        {
            string SQL = @$"
                SELECT
                       feature_name
                  FROM
                       t_job_features
                 WHERE
                       category = @category
              GROUP BY
                       feature_name
              ORDER BY
                       feature_name ASC
                     ;
            ";
            return SQL;
        }

        public static string SelectProjectUsageByFeature()
        {
            string SQL = @$"
                SELECT
                       f.feature_name                                   AS FeatureName,
                       count(*)                                         AS FeatureCount,
                       trunc(count(*) / sum(count(*)) OVER (), 4) * 100 AS Ratio
                  FROM
                       t_job_features AS f
            INNER JOIN
                       t_jobs AS j
                    ON j.id = f.job_id

                 WHERE
                       f.category = @category -- 'LANGUAGE' 'FRAMEWORK_LIBRARY' 'ROLE' 'INFRASTRUCTURE' 'DATABASE' 'CLOUD'
              GROUP BY
                       f.feature_name
              ORDER BY
                       feature_name DESC
                ;
            ";
            return SQL;
        }
    }
}
