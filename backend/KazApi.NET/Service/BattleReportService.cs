using KazApi.Domain._Const;
using KazApi.Domain.DTO;
using KazApi.Repository;
using KazApi.Repository.sql;

namespace KazApi.Service
{
    public class BattleReportService
    {
        private readonly IDatabase _posgre;
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public BattleReportService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(configuration);
        }

        /// <summary>
        /// モンスター種族を取得
        /// </summary>
        /// <returns></returns>
        public IEnumerable<MonsterTypeDTO> SelectMonsterTypes()
        {
            try
            {
                object parameter = new
                {
                    code_id = CCodeType.MONSTER.Value
                };
                string sql = ReportSQL.SelectMonsterTypes();

                return _posgre.Select<MonsterTypeDTO>(sql, parameter); ;
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
        public IEnumerable<MonsterReportDTO> SelectMonsterReport(
            int monsterTypeId,
            int sortType,
            bool isAscOrder)
        {
            try
            {
                var param = new
                {
                    monster_type = monsterTypeId,
                    sort_type = sortType,
                    is_asc_order = isAscOrder
                };
                IEnumerable<MonsterReportDTO> report
                    = _posgre.Select<MonsterReportDTO>(
                        ReportSQL.SelectMonsterReport(param), param);

                return report;
            }
            catch (Exception)
            {
                throw;
            }
        }

        public IEnumerable<BattleReportDTO> SelectBattleReport(
            int battleScale, DateTime? dateFrom, DateTime? dateTo
        )
        {
            try
            {
                var param = new
                {
                    battle_scale = battleScale,
                    from = dateFrom,
                    to = dateTo
                };

                IEnumerable<BattleReportDTO> report
                    = _posgre.Select<BattleReportDTO>(
                        ReportSQL.SelectBattleReport(
                            param.battle_scale,
                            param.from,
                            param.to
                            ), param);

                return report;
            }
            catch (Exception)
            {
                throw;
            }
        }
    }
}
