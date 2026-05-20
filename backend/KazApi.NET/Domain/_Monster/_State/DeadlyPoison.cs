using CSLib.Lib;
using KazApi.Common._Log;
using KazApi.Domain._Const;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._State
{
    /// <summary>
    /// 猛毒状態クラス
    /// </summary>
    public class DeadlyPoison : IState
    {
        private static readonly double POISON_DAMAGE_RATE = 0.25;
        private static readonly double ADJUST_RATE = 0.2;
        private static readonly string DEADLY_POISON_SKILL_ID = "skill046";

        /// <summary>
        /// コンストラクタ
        /// </summary>
        /// <param name="dto"></param>
        public DeadlyPoison(StateDTO dto) : base(dto)
        {
            base.Activate = true;
        }
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public DeadlyPoison(string name, string shortName, int stateType, double cancelRate)
                     : base(name, shortName, stateType, cancelRate)
        {
            base.StateType = CStateType.DEADLY_POISON.Value;
            base.Activate = true;
        }

        public override IState DeepCopy()
            => new DeadlyPoison(base.Name, base.ShortName, base.StateType, base.CancelRate);

        public override void DisabledLogging(IMonster monster, ILog<BattleMetaData> logger)
        {
            logger.Logging(new BattleMetaData(
                monster.MonsterId,
                base._disabledState,
                base.ShortName,
                $"{monster.MonsterName}の猛毒が消えたようだ。")
                );
        }

        /// <summary>
        /// 強めの毒ダメージ
        /// </summary>
        public override void Impact(IMonster monster, ILog<BattleMetaData> logger)
        {
            // 毒ダメージ算出
            int poisonDamage = (int)(monster.MaxHp * POISON_DAMAGE_RATE);
            poisonDamage = new URandom().RandomChangeInt(poisonDamage, ADJUST_RATE);

            logger.Logging(new BattleMetaData(monster.MonsterId, $"猛毒に侵されている　..."));
            logger.Logging(new BattleMetaData(
                monster.MonsterId,
                DEADLY_POISON_SKILL_ID,
                monster.Hp,
                poisonDamage,
                $"{monster.MonsterName}は{poisonDamage}ダメージを受けた。"
                ));

            // 被ダメージ
            monster.AcceptDamage(poisonDamage);
        }
    }
}
