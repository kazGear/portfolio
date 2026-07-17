namespace Repository.Repository.sql.publicApi;

/// <summary>
/// SQL文格納クラス
/// </summary>
public static class MakersSQL
{
    public static string GetMakers()
    {
        string SQL = @"
                SELECT 
                       'maker' AS Category,
                       value   AS Code,
                       name    AS Name
                  FROM
                       m_code
                 WHERE
                       code_id = 'guitar_makers'
              ORDER BY
                       code ASC
                     ;
        ";
        return SQL;
    }
}
