namespace KazApi.Domain._Const
{
    public class CRoleType : Enumeration<int>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CRoleType(int id, string name) : base(id, name) { }

        /// <summary>
        /// 一般
        /// </summary>
        public static readonly CRoleType NORMAL = new(1, "一般");
        /// <summary>
        /// 優良
        /// </summary>
        public static readonly CRoleType EXCELLENT = new(2, "優良");
        /// <summary>
        /// カモ
        /// </summary>
        public static readonly CRoleType KAMO = new(3, "カモ");
        /// <summary>
        /// 要注意
        /// </summary>
        public static readonly CRoleType DANGER = new(4, "要注意");
        /// <summary>
        /// ブラックリスト
        /// </summary>
        public static readonly CRoleType BLACKLIST = new (5, "ブラックリスト");
        /// <summary>
        /// 管理者
        /// </summary>
        public static readonly CRoleType ADMIN = new (90, "管理者");
        /// <summary>
        /// 統括管理者
        /// </summary>
        public static readonly CRoleType MASTER_ADMIN = new (91, "統括管理者");
    }
}
