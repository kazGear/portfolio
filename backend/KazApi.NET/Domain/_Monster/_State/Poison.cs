using CSLib.Lib;
using KazApi.Common._Log;
using KazApi.Domain._Const;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._State
{
    /// <summary>
    /// 毒状態クラス
    /// </summary>
    public class Poison : IState
    {
        private static readonly double POISON_DAMAGE_RATE = 0.1;
        private static readonly double ADJUST_RATE = 0.4;
        private static readonly string POISON_SKILL_ID = "skill044";

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public Poison(StateDTO dto) : base(dto) 
        {
            base.Activate = true;
        }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public Poison(string name, string shortName, int stateType, double cancelRate)
               : base(name, shortName, stateType, cancelRate)
        {
            base.StateType = CStateType.POISON.Value;
            base.Activate = true;
        }

        public override IState DeepCopy()
            => new Poison(base.Name, base.ShortName, base.StateType, base.CancelRate);

        public override void DisabledLogging(IMonster monster, ILog<BattleMetaData> logger)
        {
            logger.Logging(new BattleMetaData(
                monster.MonsterId,
                base._disabledState,
                base.ShortName,
                $"{monster.MonsterName}の毒が消えたようだ。")
                );
        }

        /// <summary>
        /// 毒ダメージを受ける
        /// </summary>
        public override void Impact(IMonster monster, ILog<BattleMetaData> logger)
        {
            // 毒ダメージ算出
            int poisonDamage = (int)(monster.MaxHp * POISON_DAMAGE_RATE);
            poisonDamage = new URandom().RandomChangeInt(poisonDamage, ADJUST_RATE);

            logger.Logging(new BattleMetaData(monster.MonsterId, $"毒がまわってきた。。。"));
            logger.Logging(new BattleMetaData(
                monster.MonsterId,
                POISON_SKILL_ID,
                monster.Hp,
                poisonDamage,
                $"{monster.MonsterName}は{poisonDamage}ダメージを受けた。"
                ));

            // 被ダメージ
            monster.AcceptDamage(poisonDamage);
        }
    }
}
