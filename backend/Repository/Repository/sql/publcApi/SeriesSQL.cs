namespace Repository.Repository.sql.publicApi;

/// <summary>
/// SQL文格納クラス
/// </summary>
public static class SeriesSQL
{
    public static string GetSeries()
    {
        string SQL = @"
                SELECT
                       'series'   AS Category,
                       max(maker) AS Code,
                       series     AS Name
                  FROM
                       t_guitars
                 WHERE
                       maker = @maker
              GROUP BY
                       series
              ORDER BY
                       series ASC
                     ;
        ";
        return SQL;
    }
}
