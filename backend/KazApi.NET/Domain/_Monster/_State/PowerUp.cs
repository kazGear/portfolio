using KazApi.Common._Log;
using KazApi.Domain._Monster._Skill;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._State
{
    /// <summary>
    /// 攻撃力UP状態クラス
    /// </summary>
    public class PowerUp : IState, IPositiveSkill
    {
        private static readonly double ATTACK_GAIN = 1.75;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public PowerUp(StateDTO dto) : base(dto) { }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public PowerUp(string name, string shortName, int stateType, double cancelRate)
                : base(name, shortName, stateType, cancelRate) { }

        public override IState DeepCopy()
            => new PowerUp(base.Name, base.ShortName, base.StateType, base.CancelRate);

        public override void DisabledLogging(IMonster me, ILog<BattleMetaData> logger)
        {
            me.InitAttack();

            logger.Logging(new BattleMetaData(
                me.MonsterId,
                base._disabledState,
                base.ShortName,
                $"{me.MonsterName}の攻撃力が元に戻った。")
                );
        }

        /// <summary>
        /// 攻撃力を上昇させる
        /// </summary>
        public override void Impact(IMonster me, ILog<BattleMetaData> logger)
        {
            if (me.Attack == me.DefaultAttack)
            {
                double gainerAttack = (double)me.Attack * ATTACK_GAIN;
                me.SetAttack((int)gainerAttack);

                logger.Logging(new BattleMetaData(
                    me.MonsterId,
                    $"{me.MonsterName}の攻撃力が上昇した！")
                    );
            }
        }
    }
}
