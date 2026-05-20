namespace KazApi.Repository.sql
{
    /// <summary>
    /// SQL文格納クラス
    /// </summary>
    public static class AuthSQL
    {
        public static string SelectLoginUser()
        {
            string SQL = @"
                SELECT 
                       login_id          AS LoginId
                     , login_pass        AS LoginPass
                     , role              AS Role
                  FROM 
                       m_user
                 WHERE 
                       login_id          = @login_id
                   AND login_pass        = @login_pass
                   AND is_login_disabled = FALSE ;
            ";
            return SQL;
        }
    }
}
