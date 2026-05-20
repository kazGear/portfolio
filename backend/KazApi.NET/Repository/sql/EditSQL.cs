using KazApi.Domain._Const;

namespace KazApi.Repository.sql
{
    /// <summary>
    /// SQL文格納クラス
    /// </summary>
    public static class EditSQL
    {
        public static string FetchDropDown()
        {
            string SQL = @$"
                SELECT
                       value   AS VALUE
                     , name    AS Name
                  FROM
                       m_code
                 WHERE
                       code_id = '{CCodeType.EDIT.Value}'
            ";
            return SQL;
        }

        public static string FetchEditMonsters()
        {
            string SQL = @$"
                SELECT
                       monster.monster_id   AS MonsterId
                     , monster.monster_name AS MonsterName
                     , monster.monster_type AS MonsterType
                     , monster_type.name    AS MonsterTypeName 
                     , monster.hp           AS Hp
                     , monster.attack       AS Attack
                     , monster.speed        AS Speed
                     , monster.week         AS Week
                     , week.name            AS WeekName
                     , item.is_disabled     AS IsDisabled
                  FROM
                       m_user AS ""user""
            INNER JOIN
                       t_my_item AS item
                    ON item.login_id = ""user"".login_id

            INNER JOIN
                       m_monster AS monster
                    ON item.item_id = 'monsterType' || lpad(monster.monster_type::text, 3, '0')

            INNER JOIN
                       m_code AS monster_type
                    ON monster_type.code_id = 'code006'
                   AND monster_type.value = monster.monster_type

            INNER JOIN
                       m_code AS week
                    ON week.code_id = 'code001'
                   AND week.value = monster.week

                 WHERE 
                       ""user"".login_id = @login_id
              ORDER BY
                       monster.monster_id ASC ;
            ";
            return SQL;
        }

        public static string UpdateMonsterStatus()
        {
            string SQL = @"
                UPDATE
                       m_monster
                   SET 
                       monster_name = @monster_name 
                     , hp           = @hp
                     , attack       = @attack
                     , speed        = @speed
                     , week         = @week
                 WHERE
                       monster_id = @monster_id ;
            ";
            return SQL;
        }

        public static string InitAllMonsterStatus()
        {
            string SQL = @"
                   TRUNCATE
                            m_monster;
                INSERT INTO
                            m_monster 
                          (
                            SELECT
                                   * 
                              FROM
                                   m_monster_origin
                          ) ;
            ";
            return SQL;
        }

        public static string InitAllMonsterSkills()
        {
            string SQL = @"
                   TRUNCATE
                            m_monster_skills;
                INSERT INTO
                            m_monster_skills 
                           (
                            SELECT
                                   *
                              FROM
                                   m_monster_skills_origin
                           ) ;
            ";
            return SQL;
        }

        public static string FecthEditSkills()
        {
            string SQL = @$"
                SELECT item.item_id                 AS ItemId
                     , monster.monster_id           AS MonsterId
                     , monster.monster_name         AS MonsterName
                     , monster.hp                   AS Hp
                     , monster.attack               AS MonsterAttack
                     , monster.speed                AS Speed
                     , week.name                    AS WeekName
                     , monster_skills.myskill_id    AS MySkillId
                     , skill.skill_id               AS SkillId
                     , skill.skill_name             AS SkillName
                     , skill.attack                 AS SkillAttack
                     , skill_element.name           AS SkillElementName
                  FROM t_my_item AS item 
            INNER JOIN m_monster AS monster
                    ON item.item_id = 'monsterType' || lpad(monster.monster_type::text, 3, '0')
            INNER JOIN m_monster_skills AS monster_skills
                    ON monster_skills.monster_id = monster.monster_id
            INNER JOIN m_skill AS skill
                    ON skill.skill_id = monster_skills.skill_id
            INNER JOIN m_code AS week
                    ON week.code_id = 'code001'
                   AND week.value = monster.week
            INNER JOIN m_code AS skill_element
                    ON skill_element.code_id = 'code001'
                   AND skill_element.value = skill.element_type 
                 WHERE login_id = @login_id
              ORDER BY monster.monster_id ASC
                     , skill.skill_id ASC ;
            ";
            return SQL;
        }

        public static string FetchAllSkills()
        {
            string SQL = @"
                SELECT skill.skill_id AS SkillId
                     , skill.skill_name AS SkillName
                     , skill_type.name AS SkillTypeName
                     , element.name AS ElementName
                     , skill.attack AS Attack
                  FROM m_skill AS skill
            INNER JOIN m_code AS skill_type
                    ON skill_type.code_id = 'code004'
                   AND skill_type.value = skill.skill_type 
            INNER JOIN m_code AS element
                    ON element.code_id = 'code001'
                   AND element.value = skill.element_type
              ORDER BY SkillName ASC ;
                ";
            return SQL;
        }

        public static string UpdateMonsterSkills()
        {
            string SQL = @"
                UPDATE m_monster_skills 
                   SET skill_id = @skill_id 
                 WHERE myskill_id = @myskill_id ;
            ";
            return SQL;
        }
    }
}
