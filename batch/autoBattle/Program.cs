using CSLib.Const;
using CSLib.Lib;
using KazApi.Common._Log;
using KazApi.Domain._Factory;
using KazApi.Domain._GameSystem;
using KazApi.Domain._Monster;
using KazApi.Domain._Monster._State;
using KazApi.Domain.DTO;
using KazApi.Service;
using Microsoft.Extensions.Configuration;
using Repository.Repository;
using System.Text;

Console.WriteLine("Auto battle start...");

// 接続先情報
var configuration = new ConfigurationBuilder()
    .SetBasePath(AppContext.BaseDirectory)
    .AddJsonFile("appsettings.json", optional: false)
    .Build();

IDatabase _posgre              = new PostgreSQL(ConnectionString.Get(configuration));
BattleService _service         = new BattleService(configuration);
MonsterFactory _monsterFactory = new MonsterFactory();
Randoms _random                = new Randoms();
ILog<BattleMetaData> logger    = new BattleLogger();

int battleTimes = 5; // 戦闘回数

for (int i = 0; i < battleTimes; i++)
{
    try
    {
        Console.OutputEncoding = Encoding.UTF8;

        /**
         * モンスタ－用意
         */

        // モンスターデータ等の読込み
        IEnumerable<MonsterDTO> monstersFromDB          = _service.SelectMonsters("admin");
        IEnumerable<SkillDTO> skillsFromDB              = _service.SelectSkills();
        IEnumerable<MonsterSkillDTO> monsterSkillFromDB = _service.SelectMonsterSkills();

        // モンスターDTO構築
        IEnumerable<MonsterDTO> monstersDTO =
            _monsterFactory.MappingToMonsterDTO(monstersFromDB, skillsFromDB, monsterSkillFromDB);

        // 参加モンスター（モンスター数はランダム）
        IEnumerable<MonsterDTO> battleMonstersDTO
            = BattleSystem.MonsterSelector(monstersDTO, _random.RandomInt(2, 7));

        // 戦闘用モンスターを構築
        IEnumerable<CodeDTO> stateCodeFromDB =_service.SelectStateCode();
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
                me.StateImpact(logger);

                // モンスターの行動
                IList<IMonster> otherMonsters = orderedMonsters.Where(e => e.MonsterId != me.MonsterId).ToList();
                if (me.IsMoveAble()) me.Move(otherMonsters, logger);

                // 状態異常解除
                IEnumerable<IState> currentStatus = me.CurrentStatus();
                ISet<IState> changedStatus = new HashSet<IState>();
                foreach (IState state in currentStatus)
                {
                    if (!BattleSystem.StateIsDisabled(state))
                        changedStatus.Add(state);
                    else
                        state.DisabledLogging(me, logger);
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

        _service.InsertBattleResult(records, endDate, endTime);

        Console.WriteLine($"{i + 1}戦目 終了.");

        // 間隔を空け再選（1分ごと、最終回は待たない）
        if (i < battleTimes - 1)
        {
            // 再戦待ち...(10秒)");
            await Task.Delay(10000);
        }
    }
    catch (Exception e)
    {
        Console.WriteLine("batch [AutoBattle] が異常終了しました。");
        Console.WriteLine(e);
    }

}
Console.WriteLine($"Auto battle finish. （{battleTimes}戦）");