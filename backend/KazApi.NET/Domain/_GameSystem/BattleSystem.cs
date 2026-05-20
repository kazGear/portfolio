using CSLib.Lib;
using KazApi.Common._Log;
using KazApi.Domain._Monster;
using KazApi.Domain._Monster._State;
using KazApi.Domain.DTO;

namespace KazApi.Domain._GameSystem
{
    /// <summary>
    /// 戦闘システムクラス
    /// </summary>
    public class BattleSystem
    {
        /// <summary>
        /// 敵を選択する
        /// </summary>
        public static IMonster SelectOneEnemy(IEnumerable<IMonster> monsters)
        {
            IEnumerable<IMonster> enemies = monsters.Where(e => e.Hp > 0);

            int enemyIndex = new URandom().RandomInt(0, enemies.Count());
            return enemies.ElementAt(enemyIndex);
        }

        /// <summary>
        /// 戦闘モンスターをランダムに選出する
        /// </summary>
        public static IEnumerable<T> MonsterSelector<T>(IEnumerable<T> monsters, int needAmount)
        {
            if (monsters.Count() < 2) throw new Exception("バトルは２体以上必要です。");

            IList<T> result = [];
            IList<int> usedMonsterId = [];

            // 必要数のモンスタを用意
            for (int i = 0; i < needAmount; i++)
            {
                int monsterId = new URandom().RandomInt(0, monsters.Count());

                // 同じモンスターは選べない
                while (usedMonsterId.Contains(monsterId))
                    // ランダムに選出
                    monsterId = new URandom().RandomInt(0, monsters.Count());

                usedMonsterId.Add(monsterId);

                T monster = monsters.ElementAt(monsterId);
                result.Add(monster);
            }
            return result;
        }

        /// <summary>
        /// 行動順を決定する
        /// </summary>
        public static IEnumerable<IMonster> ActionOrder(IEnumerable<IMonster> monsters)
        {
            // スピードを乱数調整した上で順番決め
            IList<IMonster> result =
                monsters.Where(e => e.Hp > 0)
                        .OrderByDescending(
                            e => new URandom().RandomChangeInt(e.Speed, 0.4))
                        .ToList();

            return result;
        }
        /// <summary>
        /// 状態異常が解除されたか判定する
        /// true: 解除・無効, false: 続行・有効
        /// </summary>
        public static bool StateIsDisabled(IState state)
        {
            bool result = false;
            double randomNumber = new URandom().RandomDouble(0.0, 1.0);

            if (randomNumber < state.CancelRate) result = true;
            return result;
        }

        /// <summary>
        /// 賭け金レートを算出
        /// </summary>
        /// <param name="monsters"></param>
        public static void CalcBetRate(IEnumerable<MonsterDTO> monsters)
        {
            int monsterCount = monsters.Count() - 1; // モンスター数が多いほど倍率UP
            double maxScore = monsters.Max(e => e.BetScore);

            foreach (MonsterDTO monster in monsters)
            {
                if (maxScore == monster.BetScore)
                {
                    monster.BetRate = double.Parse(
                        (maxScore / monster.BetScore * monsterCount * 1.15).ToString("F2")
                        );
                }
                else
                {
                    monster.BetRate = double.Parse(
                        (maxScore / monster.BetScore * monsterCount).ToString("F2")
                        );
                }
            }
        }
    }
}
