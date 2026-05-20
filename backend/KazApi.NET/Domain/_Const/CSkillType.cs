namespace KazApi.Domain._Const
{
    /// <summary>
    /// スキルタイプ定数
    /// </summary>
    public class CSkillType : Enumeration<int>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CSkillType(int id, string name) : base(id, name) { }

        /// <summary>
        /// 何も起こらない
        /// </summary>
        public static readonly CSkillType NONE = new(0, "NONE");
        /// <summary>
        /// 打撃攻撃
        /// </summary>
        public static readonly CSkillType BLOW = new(1, "BLOW");
        /// <summary>
        /// 斬撃攻撃
        /// </summary>
        public static readonly CSkillType SLASH = new(2, "SLASH");
        /// <summary>
        /// 魔法攻撃
        /// </summary>
        public static readonly CSkillType ATTACK_MAGIC = new(3, "ATTACK_MAGIC");
        /// <summary>
        /// 割合攻撃
        /// </summary>
        public static readonly CSkillType ATTACK_RATE = new(4, "ATTACK_RATE");
        /// <summary>
        /// 即死攻撃
        /// </summary>
        public static readonly CSkillType DEAD = new(5, "DEAD");
        /// <summary>
        /// 回復スキル
        /// </summary>
        public static readonly CSkillType HEAL = new(6, "HEAL");
        /// <summary>
        /// 状態系
        /// </summary>
        public static readonly CSkillType STATE = new(7, "STATE");
        /// <summary>
        /// 何もしない
        /// </summary>
        public static readonly CSkillType NOT_MOVE = new(8, "NOT_MOVE");

        /// <summary>
        /// フィールド情報の一覧を取得
        /// </summary>
        public static IReadOnlyCollection<int> GetValues()
        {
            IReadOnlyCollection<int> values = new HashSet<int>()
            {
                NONE.Value,
                BLOW.Value,
                SLASH.Value,
                ATTACK_MAGIC.Value,
                ATTACK_RATE.Value,
                DEAD.Value,
                HEAL.Value,
                STATE.Value,
                NOT_MOVE.Value
            };
            return values;
        }
    }
}
