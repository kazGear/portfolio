namespace KazApi.Repository.sql
{
    /// <summary>
    /// SQL文格納クラス
    /// </summary>
    public static class BattleSQL
    {
        public static string InsertBattleResult()
        {
            string SQL = $@"
                INSERT INTO
                            t_battle_result
                     VALUES 
                          (                  
                              @battle_end_date
                            , @battle_end_time
                            , @serial
                            , @monster_id
                            , @is_win
                          );
            ";
            return SQL;
        }
    }
}
