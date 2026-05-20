using KazApi.Common._Log;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._State
{
    /// <summary>
    /// 状態異常インターフェイス
    /// </summary>
    public abstract class IState
    {
        protected readonly bool _disabledState = true;

        public string Name { get; protected set; }
        public string ShortName { get; protected set; }
        public int StateType { get; protected set; }
        public double CancelRate { get; protected set; } // ステータス解除率

        /// <summary>
        /// 非アクティブ時には解除されない
        /// true: 解除可能性あり、false: 解除されない
        /// </summary>
        public bool Activate { get; protected set; } = false; 

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public IState(string name, string shortName, int stateType, double cancelRate)
        {
            Name = name;
            ShortName = shortName;
            StateType = stateType;
            CancelRate = cancelRate;
        }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public IState(StateDTO dto)
        {
            Name = dto.Name;
            ShortName = dto.ShortName;
            StateType = dto.StateType;
            CancelRate = dto.CancelRate;
            Activate = dto.Activate;
        }

        /// <summary>
        /// 状態有効化、状態解除の可能性が出現
        /// </summary>
        public void Activation() => Activate = true;

        /// <summary>
        /// 状態タイプを取得
        /// </summary>
        public int GetStateType() => StateType;

        /// <summary>
        /// DTOへ変換
        /// </summary>
        public static IEnumerable<StateDTO> ModelToDTO(IEnumerable<IState> models)
        {
            IList<StateDTO> result = [];

            foreach (IState model in models)
                result.Add(new StateDTO(model));

            return result;
        }

        /// <summary>
        /// 新しくオブジェクトを生成
        /// </summary>
        public abstract IState DeepCopy();

        /// <summary>
        /// 状態解除時のログ
        /// </summary>
        public abstract void DisabledLogging(IMonster monster, ILog<BattleMetaData> logger);

        /// <summary>
        /// 状態の影響をモンスターに与える
        /// </summary>
        public abstract void Impact(IMonster monster, ILog<BattleMetaData> logger);
    }
}
