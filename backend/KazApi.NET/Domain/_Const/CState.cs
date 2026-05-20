using KazApi.Domain.DTO;
using Microsoft.CodeAnalysis.Elfie.Serialization;
using Mono.TextTemplating;
using System.Security.Policy;

namespace KazApi.Domain._Const
{
    /// <summary>
    /// 状態定数
    /// </summary>
    public class CStateType : Enumeration<int>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CStateType(int id, string name, string shortName = "")
                    : base(id, name, shortName) { }

        /// <summary>
        /// 無し
        /// </summary>
        public static readonly CStateType NONE = new(0, "通常", "通");
        /// <summary>
        /// 毒状態
        /// </summary>
        public static readonly CStateType POISON = new(1, "毒", "毒");
        /// <summary>
        /// 睡眠
        /// </summary>
        public static readonly CStateType SLEEP = new(2, "睡眠", "眠");
        /// <summary>
        /// 魅了
        /// </summary>
        public static readonly CStateType CHARM = new(3, "魅了", "魅");
        /// <summary>
        /// スロー、遅い
        /// </summary>
        public static readonly CStateType SLOW = new(4, "スロー", "遅");
        /// <summary>
        /// 攻撃力UP
        /// </summary>
        public static readonly CStateType POWERUP = new(5, "攻撃力UP", "攻↑");
        /// <summary>
        /// 回避率UP
        /// </summary>
        public static readonly CStateType DODGEUP = new(6, "回避率UP", "回↑");
        /// <summary>
        /// クリティカル率UP
        /// </summary>
        public static readonly CStateType CRITICALUP = new(7, "クリティカルUP", "ク↑");
        /// <summary>
        /// 自動回復
        /// </summary>
        public static readonly CStateType AUTOHEAL = new(8, "自動回復", "癒");
        /// <summary>
        /// 猛毒
        /// </summary>
        public static readonly CStateType DEADLY_POISON = new(9, "猛毒", "猛");

        /// <summary>
        /// フィールド情報の一覧を取得
        /// </summary>
        public static IReadOnlyCollection<int> GetValues()
        {
            IReadOnlyCollection<int> values = new HashSet<int>()
            {
                NONE.Value,
                POISON.Value,
                SLEEP.Value,
                CHARM.Value,
                SLOW.Value,
                POWERUP.Value,
                DODGEUP.Value,
                CRITICALUP.Value,
                AUTOHEAL.Value,
                DEADLY_POISON.Value
            };
            return values;
        }

        /// <summary>
        /// 状態名の詰め合わせを取得
        /// </summary>
        public static IEnumerable<string> ConvertTypeToName(IEnumerable<StateDTO> status)
        {
            IList<string> result = [];
            foreach (StateDTO state in status)
            {
                result.Add(state.ShortName);
            }
            return result;
        }
    }
}
