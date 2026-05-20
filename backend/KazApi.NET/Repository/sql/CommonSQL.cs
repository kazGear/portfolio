using KazApi.Domain._Const;

namespace KazApi.Repository.sql
{
    /// <summary>
    /// SQL文格納クラス
    /// </summary>
    public static class CommonSQL
    {
        public static string UpdateUserImage()
        {
            string SQL = @"
                UPDATE
                       m_user
                   SET
                       user_image = @image
                 WHERE
                       login_id = @login_id ;
            ";
            return SQL;
        }

        public static string FetchElementCode()
        {
            string SQL = $@"
                SELECT
                       code_id AS CodeId
                     , value   AS VALUE
                     , name    AS Name
                     , param1  AS Param1
                     , param2  AS Param2
                     , param3  AS Param3
                     , remarks AS Remarks
                  FROM
                       m_code
                 WHERE
                       code_id = '{CCodeType.ELEMENT.Value}' ;
            ";
            return SQL;
        }
    }
}
