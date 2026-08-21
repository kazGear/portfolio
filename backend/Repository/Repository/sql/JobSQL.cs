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
                       f.category          = @category
                   AND NOW() - updated_at <= '3 month'

              GROUP BY
                       f.feature_name
              ORDER BY
                       feature_name DESC
                ;
            ";
            return SQL;
        }

        public static string SelectWorkPlaceByPrefecture()
        {
            string SQL = @"
                SELECT
                       location                                                   AS Location,
                       sum(CASE work_place WHEN 'フルリモート' THEN 1 ELSE 0 END ) AS FullRemote,
                       sum(CASE work_place WHEN 'ハイブリッド' THEN 1 ELSE 0 END ) AS Hybrid,
                       sum(CASE work_place WHEN '常駐' THEN 1 ELSE 0 END)         AS OnSite,
                       CASE location
                            WHEN '北海道' THEN 1
                            WHEN '青森県' THEN 2
                            WHEN '岩手県' THEN 3
                            WHEN '宮城県' THEN 4
                            WHEN '秋田県' THEN 5
                            WHEN '山形県' THEN 6
                            WHEN '福島県' THEN 7
                            WHEN '茨城県' THEN 8
                            WHEN '栃木県' THEN 9
                            WHEN '群馬県' THEN 10
                            WHEN '埼玉県' THEN 11
                            WHEN '千葉県' THEN 12
                            WHEN '東京都' THEN 13
                            WHEN '神奈川県' THEN 14
                            WHEN '新潟県' THEN 15
                            WHEN '富山県' THEN 16
                            WHEN '石川県' THEN 17
                            WHEN '福井県' THEN 18
                            WHEN '山梨県' THEN 19
                            WHEN '長野県' THEN 20
                            WHEN '岐阜県' THEN 21
                            WHEN '静岡県' THEN 22
                            WHEN '愛知県' THEN 23
                            WHEN '三重県' THEN 24
                            WHEN '滋賀県' THEN 25
                            WHEN '京都府' THEN 26
                            WHEN '大阪府' THEN 27
                            WHEN '兵庫県' THEN 28
                            WHEN '奈良県' THEN 29
                            WHEN '和歌山県' THEN 30
                            WHEN '鳥取県' THEN 31
                            WHEN '島根県' THEN 32
                            WHEN '岡山県' THEN 33
                            WHEN '広島県' THEN 34
                            WHEN '山口県' THEN 35
                            WHEN '徳島県' THEN 36
                            WHEN '香川県' THEN 37
                            WHEN '愛媛県' THEN 38
                            WHEN '高知県' THEN 39
                            WHEN '福岡県' THEN 40
                            WHEN '佐賀県' THEN 41
                            WHEN '長崎県' THEN 42
                            WHEN '熊本県' THEN 43
                            WHEN '大分県' THEN 44
                            WHEN '宮崎県' THEN 45
                            WHEN '鹿児島県' THEN 46
                            WHEN '沖縄県' THEN 47
                            WHEN '' THEN 98
                            ELSE 99 END AS sort_key
                  FROM
                       t_jobs
                 WHERE
                       NOW() - updated_at <= '3 month'
              GROUP BY
                       location
              ORDER BY
                       sort_key DESC
                     ;
            ";
            return SQL;
        }

        public static string SelectSalaryRangeByFeature()
        {
            string SQL = @$"
                SELECT
                       f.feature_name AS FeatureName,
                      (percentile_cont(0.2) WITHIN GROUP (ORDER BY j.min_salary_at_month) +
                       percentile_cont(0.2) WITHIN GROUP (ORDER BY j.max_salary_at_month)) / 2 AS SalaryLower,
                      (percentile_cont(0.5) WITHIN GROUP (ORDER BY j.min_salary_at_month) +
                       percentile_cont(0.5) WITHIN GROUP (ORDER BY j.max_salary_at_month)) / 2 AS SalaryMedian,
                      (percentile_cont(0.8) WITHIN GROUP (ORDER BY j.min_salary_at_month) +
                       percentile_cont(0.8) WITHIN GROUP (ORDER BY j.max_salary_at_month)) / 2 AS SalaryHigher
                  FROM
                       t_job_features AS f
            INNER JOIN
                       t_jobs AS j
                    ON j.id = f.job_id

                 WHERE
                       f.category           = @category
                   AND NOW() - updated_at  <= '3 month'
                   AND min_salary_at_month >  0 AND max_salary_at_month > 0

              GROUP BY
                       f.feature_name
              ORDER BY
                       f.feature_name DESC
                     ;
            ";
            return SQL;
        }
    }
}
