using CSLib.Lib;
using PrivateApi.Domain.DTO;
using Repository.Repository;
using Repository.Repository.sql;

namespace PrivateApi.Service
{
    public class EditService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        internal EditService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(configuration));
        }

        internal async Task<IEnumerable<CodeDTO>> FetchDropDown()
            => await _posgre.Select<CodeDTO>(EditSQL.FetchDropDown());

        /// <summary>
        /// 編集用モンスターデータ
        /// </summary>
        internal async Task<IEnumerable<EditMonsterDTO>> FetchEditMonsters(string loginId)
        {
            var param = new { login_id = loginId };
            return await _posgre.Select<EditMonsterDTO>(EditSQL.FetchEditMonsters(), param);
        }

        /// <summary>
        /// モンスターのステータスを設定する
        /// </summary>
        internal async Task UpdateMonsterStatus(IEnumerable<EditMonsterDTO> changeMonsters)
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
                await _posgre.Execute(EditSQL.UpdateMonsterStatus(), param);
            }
        }

        /// <summary>
        /// 全モンスターのステータスを初期化する
        /// </summary>
        internal async Task InitAllMonsterStatus()
            => await _posgre.Execute(EditSQL.InitAllMonsterStatus());

        /// <summary>
        /// 全モンスターのスキルを初期化する
        /// </summary>
        internal async Task InitAllMonsterSkills()
            => await _posgre.Execute(EditSQL.InitAllMonsterSkills());

        /// <summary>
        /// 編集用モンスターデータ（スキル付き）を取得
        /// </summary>
        internal async Task<IEnumerable<EditSkillsDTO>> FecthEditSkills(string loginId)
        {
            var param = new { login_id = loginId };
            return await _posgre.Select<EditSkillsDTO>(EditSQL.FecthEditSkills(), param);
        }

        internal async Task<IEnumerable<AllSkillDTO>> FetchAllSkills()
            => await _posgre.Select<AllSkillDTO>(EditSQL.FetchAllSkills());

        /// <summary>
        /// モンスターのスキルを変更する
        /// </summary>
        internal async Task UpdateMonsterSkills(IEnumerable<EditSkillsDTO> skills)
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
                    await _posgre.Execute(EditSQL.UpdateMonsterSkills(), param);
                }
            }
        }
    }
}
