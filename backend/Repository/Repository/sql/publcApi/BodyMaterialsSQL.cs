namespace Repository.Repository.sql.publicApi;

/// <summary>
/// SQL文格納クラス
/// </summary>
public static class BodyMaterialsSQL
{
    public static string GetBodyMaterials()
    {
        string SQL = @"
                SELECT 
                       'bodyMaterial' AS Category,
                       value          AS Code,
                       name           AS Name
                  FROM
                       m_code
                 WHERE
                       code_id = 'code010'
              ORDER BY
                       code ASC
                     ;
        ";
        return SQL;
    }
}
