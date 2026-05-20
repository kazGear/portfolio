using KazApi.Domain._Monster;

namespace KazApi.Common._Log
{
    public class MessageInfo
    {
        /// <summary>
        /// 誰のターンか表示
        /// </summary>
        public static void WhoseTurn(IMonster monster, ILog<BattleMetaData> logger)
        {
            logger.Logging(new BattleMetaData($"\n============================================"));
            logger.Logging(new BattleMetaData($">>> {monster.MonsterName}のターン"));
            logger.Logging(new BattleMetaData($"============================================\n"));
        }

        /// <summary>
        /// 戦闘結果表示
        /// </summary>
        public static void BattleResult(IEnumerable<IMonster> monsters, ILog<BattleMetaData> logger)
        {
            bool existWinner = false;
            bool allLoser = false;

            IEnumerable<IMonster> alives = monsters.Where(e => e.Hp > 0);
            IMonster? alive = alives.Count() == 1 ? alives.Single() : null;

            if (alives.Count() == 1)
            {
                existWinner = true;

                logger.Logging(new BattleMetaData($"\n*************************************************"));
                logger.Logging(new BattleMetaData($"*************************************************"));
                logger.Logging(new BattleMetaData($"  Winner {alives.Single().MonsterName} !!"));
                logger.Logging(new BattleMetaData($"*************************************************"));
                logger.Logging(new BattleMetaData($"*************************************************\n"));
                logger.Logging(new BattleMetaData(existWinner, allLoser, alive));
            }
            else if (alives.Count() <= 0)
            {
                allLoser = true;
                logger.Logging(new BattleMetaData($"... 勝者なし。"));
                logger.Logging(new BattleMetaData(existWinner, allLoser, alive));
            }
        }
    }
}
