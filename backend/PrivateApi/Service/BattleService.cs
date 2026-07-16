using CSLib.Const;
using CSLib.Lib;
using PrivateApi.Domain._Monster;
using PrivateApi.Domain.DTO;
using Repository.Repository;
using Repository.Repository.sql;
using System.Transactions;


namespace PrivateApi.Service
{
    public class BattleService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public BattleService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(configuration));
        }

        /// <summary>
        /// モンスターデータ取得
        /// </summary>
        public async Task<IEnumerable<MonsterDTO>> SelectMonsters(string? loginId = null)
        {
            var param = new { login_id = loginId };
            return await _posgre.Select<MonsterDTO>(MonsterSQL.SelectMonsters(), param);
        }

        /// <summary>
        /// スキルーデータの読込み
        /// </summary>
        public async Task<IEnumerable<SkillDTO>> SelectSkills()
            => await _posgre.Select<SkillDTO>(SkillSQL.SelectSkill());

        /// <summary>
        /// スキルマップデータの読込み
        /// </summary>
        public async Task<IEnumerable<MonsterSkillDTO>> SelectMonsterSkills()
            => await _posgre.Select<MonsterSkillDTO>(MonsterSQL.SelectMonsterSkill());

        public async Task<IEnumerable<CodeDTO>> SelectStateCode()
        {
            IEnumerable<CodeDTO> codes = await _posgre.Select<CodeDTO>(CodeSQL.SelectCode());
            
            return codes.Where(e => e.CodeId == CCodeType.STATE.Value);

        }

        /// <summary>
        /// 勝敗結果を記録（モンスター）
        /// </summary>
        public async Task InsertBattleResult(IEnumerable<MonsterDTO> monsters,
                                             DateTime endDate, 
                                             TimeSpan endTime)
        {
            using (var transaciton = new TransactionScope(TransactionScopeAsyncFlowOption.Enabled))
            {
                int seq = 1;
                foreach (MonsterDTO m in monsters)
                {
                    object parameters = new
                    {
                        battle_end_date = endDate,
                        battle_end_time = endTime,
                        serial = seq,
                        monster_id = m.MonsterId,
                        is_win = m.Hp > 0 ? true : false
                    };

                    string sql = BattleSQL.InsertBattleResult();
                    await _posgre.Execute(sql, parameters);

                    seq++;
                }
                transaciton.Complete();
            }
        }
    }
}
