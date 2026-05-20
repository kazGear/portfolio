using KazApi.Common._Log;
using KazApi.Domain._Const;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._State
{
    /// <summary>
    /// 状態異常なしクラス
    /// </summary>
    public class None : IState
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public None(StateDTO dto) : base(dto) { }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public None(string name, string shortName, int stateType, double cancelRate)
             : base(name, shortName, stateType, cancelRate)
        {
            StateType = CStateType.NONE.Value;
        }


        public override IState DeepCopy()
            => new None(base.Name, base.ShortName, base.StateType, base.CancelRate);

        public override void DisabledLogging(IMonster monster, ILog<BattleMetaData> logger)
            => throw new NotImplementedException();

        public override void Impact(IMonster monster, ILog<BattleMetaData> logger)
        {
            throw new NotImplementedException();
        }
    }
}
