using KazApi.Domain._Const;

namespace KazApi.Repository.sql
{
    /// <summary>
    /// SQL文格納クラス
    /// </summary>
    public static class ReportSQL
    {
        public static string SelectMonsterTypes()
        {
            string SQL = $@"
                SELECT
                       value AS MonsterTypeId
                     , name  AS MonsterTypeName
                  FROM
                       m_code
                 WHERE
                       code_id = @code_id
              ORDER BY
                       MonsterTypeName ASC
            ";
            return SQL;
        }

        public static string SelectMonsterReport(dynamic param)
        {
            string WHERE = PartialWhereMonsterReport(param.monster_type);
            string ORDER_BY = PartialOrderByMonsterReport(param.sort_type , param.is_asc_order);

            string SQL = $@"
                SELECT
                       m.monster_id          AS MonsterId
                     , max(m.monster_name)   AS MonsterName
                     , count(*)              AS BattleCount
                     , sum(CASE WHEN is_win = TRUE
                                THEN 1
                                ELSE 0 END ) AS Wins
                  FROM
                       m_monster AS m
            INNER JOIN
                       t_battle_result AS br
                    ON m.monster_id = br.monster_id

                {WHERE}

              GROUP BY
                       m.monster_id
             {ORDER_BY} ;
            ";

            return SQL;
        }
        public static string PartialWhereMonsterReport(int monsterType)
        {
            return monsterType > 0
                ? $"WHERE monster_type = @monster_type"
                : "";
        }
        public static string PartialOrderByMonsterReport(int sortKey, bool isAscOrder)
        {
            string result = string.Empty;

            if (sortKey == CReportSortType.MONSTER_NAME.Value)
            {
                result = $@" ORDER BY MonsterName {IsAscOrder(isAscOrder)} ";
            }
            else if (sortKey == CReportSortType.WIN_COUNT.Value)
            {
                result = $@" ORDER BY Wins {IsAscOrder(isAscOrder)}
                                 , MonsterName {IsAscOrder(isAscOrder)}
                ";
            }
            else if (sortKey == CReportSortType.BATTLE_COUNT.Value)
            {
                result = $@" ORDER BY BattleCount {IsAscOrder(isAscOrder)}
                                    , MonsterName {IsAscOrder(isAscOrder)}
                ";
            }
            else if (sortKey == CReportSortType.WINS_RATE.Value)
            {
                result = $@" ORDER BY sum(CASE WHEN is_win = TRUE
                                               THEN 1
                                               ELSE 0 END ) / count(*)::real
                                                  {IsAscOrder(isAscOrder)}
                                    , MonsterName {IsAscOrder(isAscOrder)}
                ";
            } 
            else // 初期表示のソート
            {
                result = " ORDER BY MonsterName ASC" ;
            }
            return result;
        }
        public static string IsAscOrder(bool isAscOrder)
            => isAscOrder ? "ASC" : "DESC";

        public static string SelectBattleReport(int battleScale, DateTime? from, DateTime? to)
        {
            string HAVING = PartialBattleReportHaving(battleScale);
            string AND_fromTo = PartialBattleReportFromTo(from, to);

            string SQL = $@"
                SELECT
                       DENSE_RANK() OVER (
                           ORDER BY battle_end_date ASC, battle_end_time ASC
                       )                 AS BattleId
                     , b.battle_end_date AS BattleEndDate 
                     , b.battle_end_time AS BattleEndTime
                     , b.serial          AS Serial
                     , b.monster_id      AS MonsterId
                     , m.monster_name    AS MonsterName 
                     , b.is_win          AS IsWin
                  FROM
                       t_battle_result AS b
            INNER JOIN
                       m_monster AS m
                    ON b.monster_id = m.monster_id

                 WHERE
                       EXISTS 
                      (
                        SELECT
                               1
                          FROM
                               t_battle_result AS br
                         WHERE
                               b.battle_end_date = br.battle_end_date
                           AND b.battle_end_time = br.battle_end_time

                      GROUP BY
                               battle_end_date
                             , battle_end_time 
                       {HAVING}
                      )
                  {AND_fromTo}

              ORDER BY
                       BattleEndDate ASC 
                     , BattleEndTime ASC 
                     , Serial ASC ;
            ";

            return SQL;
        }
        public static string PartialBattleReportHaving(int battleScale)
        {
            return battleScale != 0 ? "HAVING count(*) = @battle_scale"
                                    : "";
        }
        public static string PartialBattleReportFromTo(DateTime? from, DateTime? to)
        {
            string fromTo = string.Empty;
            if (from != null && to != null)
            {
                fromTo = " AND battle_end_date >= @from "
                       + " AND battle_end_date <= @to ";
            }
            else if (from != null)
            {
                fromTo = " AND battle_end_date >= @from ";
            }
            else if (to != null)
            {
                fromTo = " AND battle_end_date <= @to ";
            }
            return fromTo;
        }
    }
}
