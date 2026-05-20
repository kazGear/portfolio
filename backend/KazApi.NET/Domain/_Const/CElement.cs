using System.Reflection;
using System.Security.Policy;

namespace KazApi.Domain._Const
{
    /// <summary>
    /// 自然属性定数
    /// </summary>
    public class CElement : Enumeration<int>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CElement(int id, string name) : base(id, name) { }

        /// <summary>
        /// 無
        /// </summary>
        public static readonly CElement NONE = new(0, "NONE");
        /// <summary>
        /// 炎
        /// </summary>
        public static readonly CElement FIRE = new(1, "FIRE");
        /// <summary>
        /// 雷
        /// </summary>
        public static readonly CElement THUNDER = new(2, "THUNDER");
        /// <summary>
        /// 氷
        /// </summary>
        public static readonly CElement ICE = new(3, "ICE");
        /// <summary>
        /// 土
        /// </summary>
        public static readonly CElement EARTH = new(4, "EARTH");
        /// <summary>
        /// 聖
        /// </summary>
        public static readonly CElement HOLY = new(5, "HOLY");
        /// <summary>
        /// 闇
        /// </summary>
        public static readonly CElement DARK = new(6, "DARK");

        /// <summary>
        /// フィールド情報の一覧を取得
        /// </summary>
        public static IReadOnlyCollection<int> GetValues()
        {
            IReadOnlyCollection<int> values = new HashSet<int>()
            {
                NONE.Value,
                FIRE.Value,
                THUNDER.Value,
                ICE.Value,
                EARTH.Value,
                HOLY.Value,
                DARK.Value
            };
            return values;
        }
    }
}
