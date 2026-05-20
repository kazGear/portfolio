using CSLib.Lib;
using KazApi.Common._Log;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._State
{
    /// <summary>
    /// 自動回復状態クラス
    /// </summary>
    public class AutoHeal : IState
    {
        private static readonly double HEAL_RATE = 0.15;
        private static readonly double ADJUST_RATE = 0.25;
        private static readonly string HEAL_SKILL_ID = "skill039";

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public AutoHeal(StateDTO dto) : base(dto) { }
 
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public AutoHeal(string name, string shortName, int stateType, double cancelRate)
                 : base(name, shortName, stateType, cancelRate) { }

        public override IState DeepCopy()
            => new AutoHeal(base.Name, base.ShortName, base.StateType, base.CancelRate);

        public override void DisabledLogging(IMonster me, ILog<BattleMetaData> logger)
        {
            logger.Logging(new BattleMetaData(
                me.MonsterId,
                base._disabledState,
                base.ShortName,
                $"{me.MonsterName}の自然治癒力がなくなった。")
                );
        }

        /// <summary>
        /// 自ターンに自動回復する
        /// </summary>
        public override void Impact(IMonster me, ILog<BattleMetaData> logger)
        {
            int healPoint1 = (int)(me.MaxHp * HEAL_RATE);
            int healPoint2 = new URandom().RandomChangeInt(healPoint1, ADJUST_RATE);
            int healLimit = me.MaxHp - me.Hp;
            int healPointFix = healPoint2 >= healLimit ? healLimit : healPoint2;

            logger.Logging(new BattleMetaData(me.MonsterId, $"{me.MonsterName}の自然治癒！"));
            logger.Logging(new BattleMetaData(
                me.MonsterId,
                HEAL_SKILL_ID,
                me.Hp,
                healPointFix * -1,
                $"{me.MonsterName}のHPが{healPointFix}回復した！")
                );

            me.AcceptDamage(healPointFix * -1); // マイナスダメージを加えてHP増加
        }

    }
}
