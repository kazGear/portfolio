using Microsoft.AspNetCore.Mvc;
using KazApi.Common._Filter;
using KazApi.Domain.DTO;
using KazApi.Service;
using CSLib.Lib;
using KazApi.Common;

namespace KazApi.Controller
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
        public IActionResult Init()
        {
            try
            {
                IEnumerable<MonsterTypeDTO> monsterTypes = _service.SelectMonsterTypes();
                return StatusCode(HttpStatus.OK, monsterTypes);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// モンスターのレポートを取得
        /// </summary>
        [HttpPost("api/battleReport/monsterReport")]
        public IActionResult SelectMonsterReport([FromBody] ReqMonsterReport req)
        {
            try
            {
                IEnumerable<MonsterReportDTO> report 
                    = _service.SelectMonsterReport(req.monsterTypeId, req.sortType, req.isAscOrder);

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
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e.Message));
            }
        }

        /// <summary>
        /// 戦闘のレポートを取得
        /// </summary>
        [HttpPost("api/battleReport/battleReport")]
        public IActionResult SelectBattleReport([FromBody] ReqBattleReport req)
        {
            try
            {
                DateTime? dateFrom = string.IsNullOrEmpty(req.from) ? null : DateTime.Parse(req.from);
                DateTime? dateTo   = string.IsNullOrEmpty(req.to) ? null : DateTime.Parse(req.to);

                IEnumerable<BattleReportDTO> battleReports
                    = _service.SelectBattleReport(req.battleScale, dateFrom, dateTo);
                
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
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }
    }
}
