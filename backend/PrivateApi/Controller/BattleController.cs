using Microsoft.AspNetCore.Mvc;
using PrivateApi.Common._Log;
using PrivateApi.Domain._ViewModel;
using PrivateApi.Domain._Factory;
using PrivateApi.Domain._GameSystem;
using PrivateApi.Domain._Monster;
using PrivateApi.Domain.DTO;
using Microsoft.CodeAnalysis;
using PrivateApi.Service;
using CSLib.Lib;
using PrivateApi.Common;
using System.Transactions;
using CSLib.Const;

namespace PrivateApi.Controller
{
    [ApiController]
    public class BattleController : ControllerBase
    {
        private readonly ILog<BattleMetaData> _logger;
        private readonly BattleService _service;
        private readonly MonsterFactory _monsterFactory;

        public BattleController(IConfiguration configuration)
        {
            _logger = new BattleLogger();
            _service = new BattleService(configuration);
            _monsterFactory = new MonsterFactory();
        }

        /// <summary>
        /// モンスター情報
        /// </summary>
        [HttpPost("api/battle/monstersInfo")]
        public async Task<IActionResult> MonstersInfo([FromBody] string? loginId)
        {
            if (string.IsNullOrEmpty(loginId)) return StatusCode(HttpStatus.BadRequest);

            IEnumerable<MonsterDTO> monsters = await _service.SelectMonsters(loginId);
            return StatusCode(HttpStatus.OK, monsters);
        }

        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpPost("api/battle/init")]
        public async Task<IActionResult> Init([FromBody] ReqBattleInit req)
        {
            // モンスターデータ等の読込み
            IEnumerable<MonsterDTO> monstersFromDB          = await _service.SelectMonsters(req.loginId);
            IEnumerable<SkillDTO> skillsFromDB              = await _service.SelectSkills();
            IEnumerable<MonsterSkillDTO> monsterSkillFromDB = await _service.SelectMonsterSkills();

            // モンスターDTO構築
            IEnumerable<MonsterDTO> monstersDTO =
                _monsterFactory.MappingToMonsterDTO(monstersFromDB, skillsFromDB, monsterSkillFromDB);

            // 参加モンスター（ランダム）
            IEnumerable<MonsterDTO> battleMonsters =
                BattleSystem.MonsterSelector(monstersDTO, int.Parse(req.selectMonstersCount));

            // 賭けレート算出
            BattleSystem.CalcBetRate(battleMonsters);

            // テスト用モンスターで対戦
            //battleMonsters = UseTestMonsters(monstersDTO);

            return StatusCode(HttpStatus.OK, battleMonsters);
        }

        /// <summary>
        /// テストしたいモンスターを指定する
        /// </summary>
        /// <param name="monstersDTO"></param>
        /// <returns></returns>
        private IEnumerable<MonsterDTO> UseTestMonsters(IEnumerable<MonsterDTO> monstersDTO)
        {
            IEnumerable<MonsterDTO> testMonsters = new List<MonsterDTO>()
            {
                //monstersDTO.Where(e => e.MonsterId == CMonster.カーミラクイーン.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.マーマポト.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.パーパポト.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.カーミラ.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.キラービー.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.アサシンバグ.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.ラスターバグ.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.イビルウェポン.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.クロウラー.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.ダースマタンゴ.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.ゴブリン.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.ゴブリンガード.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.ゴブリンロード.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.サハギン.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.プチポセイドン.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.パンプキンボム.VALUE).Single(),
                //monstersDTO.Where(e => e.MonsterId == CMonster.グレネードボム.VALUE).Single(),
                monstersDTO.Where(e => e.MonsterId == CMonster.シェイプシフター.Value).Single(),
                monstersDTO.Where(e => e.MonsterId == CMonster.シャドウゼロ.Value).Single(),
                monstersDTO.Where(e => e.MonsterId == CMonster.シャドウゼロワン.Value).Single(),

            };
            return testMonsters;
        }

        /// <summary>
        /// モンスターたちの行動
        /// </summary>
        [HttpPost("api/battle/nextTurn")]
        public async Task<IActionResult> NextTurn([FromBody] IEnumerable<MonsterDTO>? monsters)
        {
            if (monsters == null || monsters.Count() == 0) return StatusCode(HttpStatus.BadRequest); 

            // 戦闘用モンスターを構築
            IEnumerable<CodeDTO> codes = await _service.SelectStateCode();
            IEnumerable<IMonster> battleMonsters = _monsterFactory.CreateBattleMonsters(monsters, codes);

            // TODO 未実装 チーム決め
            ((List<IMonster>)battleMonsters).ForEach(e => e.DefineTeam(CTeam.A.Value));
            if (battleMonsters.Where(e => e.Team == CTeam.UNKNOWN.Value).Count() > 0)
            {
                throw new Exception("チーム決めが完了していません。");
            }
            // 行動順決め
            IEnumerable<IMonster> orderedMonsters = BattleSystem.ActionOrder(battleMonsters);

            // モンスターの行動
            foreach (IMonster me in orderedMonsters)
            {
                if (me.Hp <= 0) continue;

                // 誰のターンか
                MessageInfo.WhoseTurn(me, _logger);

                // 状態異常の影響
                me.StateImpact(_logger);

                // モンスターの行動
                IList<IMonster> otherMonsters = orderedMonsters.Where(e => e.MonsterId != me.MonsterId)
                                                                .ToList();
                if (me.IsMoveAble()) me.Move(otherMonsters, _logger);

                // 状態異常解除
                if (me.CurrentStatus().Count() > 0) me.RefreshStatus(_logger);

                _logger.Logging(new BattleMetaData());
            }

            // 勝敗判定
            MessageInfo.BattleResult(battleMonsters, _logger);

            // DTOへ変換
            IEnumerable<MonsterDTO> monstersDTO = _monsterFactory.ConvertToDTO(battleMonsters);

            BattleViewModel model = new BattleViewModel();
            model.Monsters = monstersDTO;
            model.BattleLog = _logger.DumpMemory();

            return StatusCode(HttpStatus.OK, model);
        }

        /// <summary>
        /// 勝敗結果を記録（モンスター）
        /// </summary>
        [HttpPost("api/battle/recordResults")]
        public async Task<IActionResult> RecordResults([FromBody] IEnumerable<MonsterDTO>? monsters)
        {
            if (monsters == null || monsters.Count() == 0) return StatusCode(HttpStatus.BadRequest);

            using (var transaction = new TransactionScope(TransactionScopeAsyncFlowOption.Enabled))
            { 
                DateTime endDate = DateTime.Now;
                TimeSpan endTime = new TimeSpan(endDate.Ticks);

                bool result = await _service.InsertBattleResult(monsters!, endDate, endTime);
                transaction.Complete();

                return StatusCode(HttpStatus.OK, result);
            }
        }
    }
}
