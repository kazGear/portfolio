namespace Repository.Repository.sql.publicApi;

/// <summary>
/// SQL文格納クラス
/// </summary>
public static class ColorsSQL
{
    public static string GetColors()
    {
        string SQL = @"
                SELECT 
                       'color' AS Category,
                       value   AS Code,
                       name    AS Name
                  FROM
                       m_code
                 WHERE
                       code_id = 'code011'
              ORDER BY
                       code ASC
                     ;
        ";
        return SQL;
    }
}
