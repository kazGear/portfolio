namespace KazApi.Repository.sql
{
    /// <summary>
    /// SQL文格納クラス
    /// </summary>
    public static class MonsterSQL
    {
        public static string SelectMonsters()
        {
            string SQL = @"
                SELECT
                       m.monster_id   AS MonsterId 
                     , m.monster_name AS MonsterName
                     , m.monster_type AS MonsterType
                     , m.hp           AS Hp
                     , m.hp           AS MaxHp
                     , m.attack       AS Attack
                     , m.attack       AS DefaultAttack
                     , m.speed        AS Speed
                     , m.speed        AS DefaultSpeed                   
                     , m.dodge        AS Dodge
                     , m.dodge        AS DefaultDodge
                     , m.week         AS Week
                     , max(m.hp)
                         + max(m.attack) * 20
                         + max(m.speed) * 10
                         + sum(s.weight) * 20
                         + sum(s.critical * 100) AS BetScore
                  FROM
                       m_monster AS m 
            INNER JOIN
                       m_monster_skills AS ms
                    ON ms.monster_id = m.monster_id 

            INNER JOIN
                       m_skill AS s
                    ON s.skill_id = ms.skill_id

                 WHERE
                       EXISTS 
                      (
                        SELECT
                               *
                          FROM
                               t_my_item AS i
                         WHERE
                               login_id = :login_id
                           AND i.item_id = 'monsterType' || lpad(m.monster_type::text, 3, '0')
                      )
              GROUP BY
                       m.monster_id
              ORDER BY
                       m.monster_id ASC ;
            ";
            return SQL;
        }

        public static string SelectMonsterSkill()
        {
            string SQL = @"
                    SELECT
                           monster_id AS MonsterId
                         , skill_id   AS SkillId 
                         , disabled   AS Disabled
                      FROM
                           m_monster_skills ;
                ";
            return SQL;
        }
    }
}
