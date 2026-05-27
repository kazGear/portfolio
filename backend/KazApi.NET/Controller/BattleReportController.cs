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
        public IActionResult SelectMonsterReport([FromForm] int? monsterTypeId,
                                                 [FromForm] int? sortType,
                                                 [FromForm] bool isAscOrder = true)
        {
            if (   monsterTypeId == null
                || sortType == null) return StatusCode(HttpStatus.BadRequest);

            try
            {
                IEnumerable<MonsterReportDTO> report 
                    = _service.SelectMonsterReport((int)monsterTypeId, (int)sortType, isAscOrder);

                // 勝率を算出
                IEnumerable<MonsterReportDTO> editedReport = report.Select(e => new MonsterReportDTO
                {
                    MonsterId = e.MonsterId,
                    MonsterName = e.MonsterName,
                    BattleCount = e.BattleCount,
                    Wins = e.Wins,
                    WinRate = (e.Wins / (double)e.BattleCount * 100).ToString("N2") + "%"
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
        public IActionResult SelectBattleReport([FromForm] int? battleScale,
                                                [FromForm] string? from,
                                                [FromForm] string? to)
        {
            if (battleScale == null) return StatusCode(HttpStatus.BadRequest);

            try
            {
                DateTime? dateFrom = from == null ? null : DateTime.Parse(from);
                DateTime? dateTo = to == null ? null : DateTime.Parse(to);

                IEnumerable<BattleReportDTO> battleReports
                    = _service.SelectBattleReport((int)battleScale, dateFrom, dateTo);
                
                IEnumerable<BattleReportDTO> editedReport = battleReports.Select(e => new BattleReportDTO
                {
                    BattleId = e.BattleId,
                    BattleEndDateStr = e.BattleEndDate.ToString().Substring(0, 10),
                    BattleEndTimeStr = e.BattleEndTime.ToString().Substring(0, 8),
                    Serial = e.Serial,
                    MonsterId = e.MonsterId,
                    MonsterName = e.MonsterName,
                    IsWin = e.IsWin
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
