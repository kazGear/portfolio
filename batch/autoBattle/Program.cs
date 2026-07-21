using CSLib.Const;
using CSLib.Lib;
using Microsoft.Extensions.Configuration;
using PrivateApi.Common._Log;
using PrivateApi.Domain._Factory;
using PrivateApi.Domain._GameSystem;
using PrivateApi.Domain._Monster;
using PrivateApi.Domain._Monster._State;
using PrivateApi.Domain.DTO;
using PrivateApi.Service;
using Repository.Repository;
using System.Text;

Console.WriteLine("Auto battle start...");

// 時間計測
TimeMeasure stopWatch = new TimeMeasure();
stopWatch.Start();

// 接続先情報
string environment = Environment.GetEnvironmentVariable("ASPNETCORE_ENVIRONMENT") ?? "Development";

var configuration = new ConfigurationBuilder()
    .SetBasePath(AppContext.BaseDirectory)
    .AddJsonFile("appsettings.json", optional: false)
    .AddJsonFile($"appsettings.{environment}.json", optional: true)
    .Build();

// batch log
string batchName            = "AutoBattle";
BatchLogger _batchLogger    = new BatchLogger(configuration);
BatchConfigDTO _batchConfig = await _batchLogger.InsertStartLog(batchName);

// services
IDatabase _posgre              = new PostgreSQL(ConnectionString.Get(configuration));
BattleService _service         = new BattleService(configuration);
MonsterFactory _monsterFactory = new MonsterFactory();
Randoms _random                = new Randoms();
ILog<BattleMetaData> _logger   = new BattleLogger();

int battleTimes = 10; // 戦闘回数

for (int i = 0; i < battleTimes; i++)
{
    try
    {
        Console.OutputEncoding = Encoding.UTF8;

        /**
         * モンスタ－用意
         */

        // モンスターデータ等の読込み
        IEnumerable<MonsterDTO> monstersFromDB          = await _service.SelectMonsters("admin");
        IEnumerable<SkillDTO> skillsFromDB              = await _service.SelectSkills();
        IEnumerable<MonsterSkillDTO> monsterSkillFromDB = await _service.SelectMonsterSkills();

        // モンスターDTO構築
        IEnumerable<MonsterDTO> monstersDTO =
            _monsterFactory.MappingToMonsterDTO(monstersFromDB, skillsFromDB, monsterSkillFromDB);

        // 参加モンスター（モンスター数はランダム）
        IEnumerable<MonsterDTO> battleMonstersDTO
            = BattleSystem.MonsterSelector(monstersDTO, _random.RandomInt(2, 7));

        // 戦闘用モンスターを構築
        IEnumerable<CodeDTO> stateCodeFromDB = await _service.SelectStateCode();
        IEnumerable<IMonster> battleMonsters 
            = _monsterFactory.CreateBattleMonsters(battleMonstersDTO, stateCodeFromDB);

        // TODO 未実装 チーム決め
        ((List<IMonster>)battleMonsters).ForEach(e => e.DefineTeam(CTeam.A.Value));

        if (battleMonsters.Where(e => e.Team == (CTeam.UNKNOWN.Value)).Count() > 0)
        {
            throw new Exception("チーム決めが完了していません。");
        }

        IEnumerable<IMonster> alives = []; // 生き残りモンスター

        /**
         * 戦闘不能が1人以下になるまで戦う
         */
        do
        {
            // 行動順決め
            IEnumerable<IMonster> orderedMonsters = BattleSystem.ActionOrder(battleMonsters);

            // モンスタ達のーの行動
            foreach (IMonster me in orderedMonsters)
            {
                // 行動不可
                if (me.Hp <= 0) continue;

                // 状態異常の効果
                me.StateImpact(_logger);

                // モンスターの行動
                IList<IMonster> otherMonsters = orderedMonsters.Where(e => e.MonsterId != me.MonsterId).ToList();
                if (me.IsMoveAble()) me.Move(otherMonsters, _logger);

                // 状態異常解除
                IEnumerable<IState> currentStatus = me.CurrentStatus();
                ISet<IState> changedStatus = new HashSet<IState>();
                foreach (IState state in currentStatus)
                {
                    if (!BattleSystem.StateIsDisabled(state))
                        changedStatus.Add(state);
                    else
                        state.DisabledLogging(me, _logger);
                }
                me.UpdateStatus(changedStatus);

                // HP 現状確認
                foreach (IMonster monster in battleMonsters)
                {
                    int hp = monster.Hp > 0 ? monster.Hp : 0;
                }

                // 勝敗判定
                alives = battleMonsters.Where(e => e.Hp > 0);
                IMonster? alive = alives.Count() == 1 ? alives.Single() : null;
            }
        } while (alives.Count() > 1);

        /**
         * 戦績の記録
         */
        IList<MonsterDTO> records = [];
        foreach (IMonster monster in battleMonsters)
        {
            records.Add(new MonsterDTO(monster));
        }
        DateTime endDate = DateTime.Now;
        TimeSpan endTime = new TimeSpan(endDate.Ticks);

        await _service.InsertBattleResult(records, endDate, endTime);

        // ログ
        Console.WriteLine($"{i + 1}戦目 終了.");
        Console.WriteLine("-- 参戦モンスター --");
        foreach (IMonster monster in battleMonsters)
        {
            Console.WriteLine($"{monster.MonsterName}");
        }
        Console.WriteLine("");

        // 間隔を空け再選（最終回は待たない）
        if (i < battleTimes - 1)
        {
            // 再戦待ち...(5秒)");
            await Task.Delay(5000);
        }
    }
    catch (Exception e)
    {
        await _batchLogger.UpdateError(e, _batchConfig);

        Console.WriteLine("batch [AutoBattle] が異常終了しました。");
        Console.WriteLine(e);

        return;
    }
}
await _batchLogger.UpdateStatus(_batchConfig, stopWatch.Stop());
Console.WriteLine($"Finished auto battle.（{battleTimes}戦）");