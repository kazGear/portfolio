using KazApi.Common._Log;
using KazApi.Domain._Const;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._State
{
    /// <summary>
    /// 睡眠状態クラス
    /// </summary>
    public class Sleep : IState, IDisableMove
    {
        private static readonly string STATE_TYPE2 = "stateType2";

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public Sleep(StateDTO dto) : base(dto) 
        {
            base.Activate = true;
        }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public Sleep(string name, string shortName, int stateType, double cancelRate)
              : base(name, shortName, stateType, cancelRate)
        {
            StateType = CStateType.SLEEP.Value;
            base.Activate = true;
        }

        public override IState DeepCopy()
            => new Sleep(base.Name, base.ShortName, base.StateType, base.CancelRate);

        public override void DisabledLogging(IMonster monster, ILog<BattleMetaData> logger)
        {
            logger.Logging(new BattleMetaData(
                monster.MonsterId,
                base._disabledState,
                base.ShortName,
                $"{monster.MonsterName}は目覚めた！")
                );
        }

        /// <summary>
        /// 自ターンは行動不能
        /// </summary>
        public override void Impact(IMonster monster, ILog<BattleMetaData> logger)
        {
            int effectTime = 1400;

            logger.Logging(new BattleMetaData(
                monster.MonsterId,
                STATE_TYPE2,
                effectTime,
                $"{monster.MonsterName}は眠っている Zzz ...")
                );
        }
    }
}
