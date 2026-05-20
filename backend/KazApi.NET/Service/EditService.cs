using KazApi.Domain.DTO;
using KazApi.Repository;
using KazApi.Repository.sql;

namespace KazApi.Service
{
    public class EditService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        internal EditService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(configuration);
        }

        internal IEnumerable<CodeDTO> FetchDropDown()
            => _posgre.Select<CodeDTO>(EditSQL.FetchDropDown());

        /// <summary>
        /// 編集用モンスターデータ
        /// </summary>
        internal IEnumerable<EditMonsterDTO> FetchEditMonsters(string loginId)
        {
            var param = new { login_id = loginId };
            return _posgre.Select<EditMonsterDTO>(EditSQL.FetchEditMonsters(), param);
        }

        /// <summary>
        /// モンスターのステータスを設定する
        /// </summary>
        internal void UpdateMonsterStatus(IEnumerable<EditMonsterDTO> changeMonsters)
        {
            foreach (EditMonsterDTO monster in changeMonsters)
            {
                var param = new
                {
                    monster_id = monster.MonsterId,
                    monster_name = monster.AfterName,
                    hp = monster.AfterHp,
                    attack = monster.AfterAttack,
                    speed = monster.AfterSpeed,
                    week = monster.AfterWeek
                };
                _posgre.Execute(EditSQL.UpdateMonsterStatus(), param);
            }
        }

        /// <summary>
        /// 全モンスターのステータスを初期化する
        /// </summary>
        internal void InitAllMonsterStatus()
            => _posgre.Execute(EditSQL.InitAllMonsterStatus());

        /// <summary>
        /// 全モンスターのスキルを初期化する
        /// </summary>
        internal void InitAllMonsterSkills()
            => _posgre.Execute(EditSQL.InitAllMonsterSkills());

        /// <summary>
        /// 編集用モンスターデータ（スキル付き）を取得
        /// </summary>
        internal IEnumerable<EditSkillsDTO> FecthEditSkills(string loginId)
        {
            var param = new { login_id = loginId };
            return _posgre.Select<EditSkillsDTO>(EditSQL.FecthEditSkills(), param);
        }

        internal IEnumerable<AllSkillDTO> FetchAllSkills()
            => _posgre.Select<AllSkillDTO>(EditSQL.FetchAllSkills());

        /// <summary>
        /// モンスターのスキルを変更する
        /// </summary>
        internal void UpdateMonsterSkills(IEnumerable<EditSkillsDTO> skills)
        {
            foreach (EditSkillsDTO skill in skills)
            {
                for (int i = 0; i < skill.MySkillIds.Count(); i++)
                {
                    var param = new
                    {
                        myskill_id = skill.MySkillIds[i],
                        skill_id = skill.SkillIds[i],
                    };
                    _posgre.Execute(EditSQL.UpdateMonsterSkills(), param);
                }
            }
        }
    }
}
