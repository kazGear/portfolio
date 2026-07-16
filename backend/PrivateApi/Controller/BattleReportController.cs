using Microsoft.AspNetCore.Mvc;
using PrivateApi.Common._Filter;
using PrivateApi.Domain.DTO;
using PrivateApi.Service;
using CSLib.Lib;
using PrivateApi.Common;

namespace PrivateApi.Controller
{
    [SkipAuthActionFilter]
    [ApiController]
    public class BattleReportController : ControllerBase
    {
        private readonly BattleReportService _service;

        public BattleReportController(IConfiguration configuration)
        {
            _service = new BattleReportService(configuration);
        }

        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpGet("api/battleReport/init")]
        public async Task<IActionResult> Init()
        {
            IEnumerable<MonsterTypeDTO> monsterTypes = await _service.SelectMonsterTypes();
            return StatusCode(HttpStatus.OK, monsterTypes);
        }

        /// <summary>
        /// モンスターのレポートを取得
        /// </summary>
        [HttpPost("api/battleReport/monsterReport")]
        public async Task<IActionResult> SelectMonsterReport([FromBody] ReqMonsterReport req)
        {
            IEnumerable<MonsterReportDTO> report 
                = await _service.SelectMonsterReport(req.monsterTypeId, req.sortType, req.isAscOrder);

            // 勝率を算出
            IEnumerable<MonsterReportDTO> editedReport = report.Select(e => new MonsterReportDTO
            {
                MonsterId   = e.MonsterId,
                MonsterName = e.MonsterName,
                BattleCount = e.BattleCount,
                Wins        = e.Wins,
                WinRate     = (e.Wins / (double)e.BattleCount * 100).ToString("N2") + "%"
            });

            return StatusCode(HttpStatus.OK, editedReport);
        }

        /// <summary>
        /// 戦闘のレポートを取得
        /// </summary>
        [HttpPost("api/battleReport/battleReport")]
        public async Task<IActionResult> SelectBattleReport([FromBody] ReqBattleReport req)
        {
            DateTime? dateFrom = string.IsNullOrEmpty(req.from) ? null : DateTime.Parse(req.from);
            DateTime? dateTo   = string.IsNullOrEmpty(req.to) ? null : DateTime.Parse(req.to);

            IEnumerable<BattleReportDTO> battleReports
                = await _service.SelectBattleReport(req.battleScale, dateFrom, dateTo);
                
            IEnumerable<BattleReportDTO> editedReport = battleReports.Select(e => new BattleReportDTO
            {
                BattleId         = e.BattleId,
                BattleEndDateStr = e.BattleEndDate.ToString().Substring(0, 10),
                BattleEndTimeStr = e.BattleEndTime.ToString().Substring(0, 8),
                Serial           = e.Serial,
                MonsterId        = e.MonsterId,
                MonsterName      = e.MonsterName,
                IsWin            = e.IsWin
            });

            return StatusCode(HttpStatus.OK, editedReport);
        }
    }
}
