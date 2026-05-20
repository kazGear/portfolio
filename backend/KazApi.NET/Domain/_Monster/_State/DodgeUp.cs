using KazApi.Common._Log;
using KazApi.Domain._Monster._Skill;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._State
{
    /// <summary>
    /// 回避率アップクラス
    /// </summary>
    public class DodgeUp : IState, IPositiveSkill
    {
        private static readonly double DODGE_GAIN = 2.0;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public DodgeUp(StateDTO dto) : base(dto) { }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public DodgeUp(string name, string shortName, int stateType, double cancelRate)
                : base(name, shortName, stateType, cancelRate) { }

        public override IState DeepCopy()
            => new DodgeUp(base.Name, base.ShortName, base.StateType, base.CancelRate);

        public override void DisabledLogging(IMonster me, ILog<BattleMetaData> logger)
        {
            me.InitDodge();

            logger.Logging(new BattleMetaData(
                me.MonsterId,
                base._disabledState,
                base.ShortName,
                $"{me.MonsterName}の回避力が元に戻った。")
                );
        }

        /// <summary>
        /// 回避力を上昇させる
        /// </summary>
        public override void Impact(IMonster me, ILog<BattleMetaData> logger)
        {
            if (me.Dodge == me.DefaultDodge)
            {
                double gainerDodge = me.Dodge * DODGE_GAIN;
                me.SetDodge(gainerDodge);

                logger.Logging(new BattleMetaData(
                    me.MonsterId,
                    $"{me.MonsterName}の回避力が上昇した！")
                    );
            }
        }
    }
}
