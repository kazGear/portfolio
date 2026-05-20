using KazApi.Domain._Monster._State;

namespace KazApi.Domain._Const
{
    public class CTarget : Enumeration<int>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CTarget(int id, string name) : base(id, name) { }

        /// <summary>
        /// 無し
        /// </summary>
        public static readonly CTarget NONE = new(0, "NONE");
        /// <summary>
        /// 敵ランダム
        /// </summary>
        public static readonly CTarget ENEMY_RANDOM = new(1, "ENEMY_RANDOM");
        /// <summary>
        /// 敵全体
        /// </summary>
        public static readonly CTarget ENEMY_ALL = new(2, "ENEMY_ALL");
        /// <summary>
        /// 敵ランダム・敵全体
        /// </summary>
        public static readonly CTarget ENEMY_RANDOM_OR_ALL = new(3, "ENEMY_RANDOM_OR_ALL");
        /// <summary>
        /// 敵ランダム・複数回
        /// </summary>
        public static readonly CTarget ENEMY_RANDOM_SOME_TIMES = new(4, "ENEMY_RANDOM_SOME_TIMES");
        /// <summary>
        /// 自身
        /// </summary>
        public static readonly CTarget ME = new(5, "ME");
        /// <summary>
        /// 味方ランダム
        /// </summary>
        public static readonly CTarget FRIEND_RANDOM = new(6, "FRIEND_RANDOM");
        /// <summary>
        /// 味方全体
        /// </summary>
        public static readonly CTarget FRIEND_ALL = new(7, "FRIEND_ALL");
        /// <summary>
        /// 味方ランダム・味方全体
        /// </summary>
        public static readonly CTarget FRIEND_RANDOM_OR_ALL = new(8, "FRIEND_RANDOM_OR_ALL");

        /// <summary>
        /// フィールド情報の一覧を取得
        /// </summary>
        public static IReadOnlyCollection<int> GetValues()
        {
            IReadOnlyCollection<int> values = new HashSet<int>()
            {
                NONE.Value,
                ENEMY_RANDOM.Value,
                ENEMY_ALL.Value,
                ENEMY_RANDOM_OR_ALL.Value,
                ENEMY_RANDOM_SOME_TIMES.Value,
                ME.Value,
                FRIEND_RANDOM.Value,
                FRIEND_ALL.Value,
                FRIEND_RANDOM_OR_ALL.Value
            };
            return values;
        }
    }
}
