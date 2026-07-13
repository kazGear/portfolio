using CSLib.Const;
using CSLib.Lib;
using Dapper;
using PrivateApi.Domain.DTO;
using Repository.Repository;
using Repository.Repository.sql;

namespace PrivateApi.Service
{
    public class BattleReportService
    {
        private readonly IDatabase _posgre;
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public BattleReportService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(configuration));
        }

        /// <summary>
        /// モンスター種族を取得
        /// </summary>
        /// <returns></returns>
        public async Task<IEnumerable<MonsterTypeDTO>> SelectMonsterTypes()
        {
            try
            {
                object parameter = new
                {
                    code_id = CCodeType.MONSTER.Value
                };
                string sql = ReportSQL.SelectMonsterTypes();

                return await _posgre.Select<MonsterTypeDTO>(sql, parameter); ;
            }
            catch
            {
                throw;
            }
        }

        /// <summary>
        /// モンスター毎のレポートを取得
        /// </summary>
        /// <returns></returns>
        public async Task<IEnumerable<MonsterReportDTO>> SelectMonsterReport(int monsterTypeId,
                                                                             int sortType,
                                                                             bool isAscOrder)
        {
            try
            {
                DynamicParameters param = new DynamicParameters();
                param.Add("monster_type", monsterTypeId);
                param.Add("sort_type", sortType);
                param.Add("is_asc_order", isAscOrder);

                string SQL = ReportSQL.SelectMonsterReport(monsterTypeId, sortType, isAscOrder);

                IEnumerable<MonsterReportDTO> report =
                    await _posgre.Select<MonsterReportDTO>(SQL, param);

                return report;
            }
            catch (Exception)
            {
                throw;
            }
        }

        public async Task<IEnumerable<BattleReportDTO>> SelectBattleReport(int battleScale,
                                                                           DateTime? dateFrom,
                                                                           DateTime? dateTo)
        {
            try
            {
                DynamicParameters param = new DynamicParameters();
                param.Add("battle_scale", battleScale);
                param.Add("from", dateFrom);
                param.Add("to", dateTo);

                string SQL = ReportSQL.SelectBattleReport(battleScale, dateFrom, dateTo);

                IEnumerable <BattleReportDTO> report
                    = await _posgre.Select<BattleReportDTO>(SQL, param);

                return report;
            }
            catch (Exception)
            {
                throw;
            }
        }
    }
}
