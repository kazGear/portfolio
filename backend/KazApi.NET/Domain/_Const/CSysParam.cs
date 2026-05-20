namespace KazApi.Domain._Const
{
    /// <summary>
    /// 環境独自の設定クラス
    /// </summary>
    public class CEnvironment : Enumeration<bool>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CEnvironment(bool value, string name) : base(value, name) { }
       
        /// <summary>
        /// 開発、ローカル環境
        /// </summary>
        public static readonly CEnvironment DEVELOPMENT = new(false, "DEVELOPMENT");
        /// <summary>
        /// リモートサーバ、公開、本番
        /// </summary>
        public static readonly CEnvironment PRODUCTION = new(true, "PRODUCTION");

        /// <summary> ////////////////////////////////////////////////////////////
        /// 現環境を決定 false: 開発, true: 本番
        /// デプロイ前に確認
        /// </summary> ///////////////////////////////////////////////////////////
        public static readonly CEnvironment THIS_ENVIRONMENT = DEVELOPMENT;
        //public static readonly CEnvironment THIS_ENVIRONMENT = PRODUCTION;
    }

    public class CFilePath : Enumeration<string>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CFilePath(string value, string name) : base(value, name) { }
        /// <summary>
        /// ユーザーイメージ保管場所
        /// </summary>
        public static readonly CFilePath USER_IMAGES = new("Domain/_User/_Images", "USER_IMAGES");
    }

    public class CPrefix : Enumeration<string>
    {
        //data:image/jpeg;base64,
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CPrefix(string value, string name) : base(value, name) { }
        /// <summary>
        /// ユーザーイメージ保管場所
        /// </summary>
        public static readonly CPrefix BASE64 = new("data:image/jpeg;base64,", "BASE64");
    }

    /// <summary>
    /// 補正率クラス
    /// </summary>
    public class CSysRate : Enumeration<double>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CSysRate(double value, string name) : base(value, name) { }

        /// <summary>
        /// クリティカル補正率
        /// </summary>
        public static readonly CSysRate CRITICAL_DAMAGE = new(1.8, "CRITICAL_DAMAGE");
        /// <summary>
        /// 行動順補正率
        /// </summary>
        public static readonly CSysRate MOVE_SPEED = new(0.3, "MOVE_SPEED");
        /// <summary>
        /// 弱点属性ダメージ
        /// </summary>
        public static readonly CSysRate WEEK_DAMAGE = new(1.8, "WEEK_DAMAGE");
        /// <summary>
        /// 物理スキルダメージ補正率
        /// </summary>
        public static readonly CSysRate PHYSICAL_SKILL_DAMAGE = new(0.10, "PHYSICAL_SKILL_DAMAGE");
        /// <summary>
        /// 魔法スキルダメージ補正率
        /// </summary>
        public static readonly CSysRate MAGIC_SKILL_DAMAGE = new(0.05, "MAGIC_SKILL_DAMAGE");
        /// <summary>
        /// 全体攻撃調整・下限
        /// </summary>
        public static readonly CSysRate ALL_ATTACK_ADJUST_MIN = new(0.45, "ALL_ATTACK_ADJUST_MIN");
        /// <summary>
        /// 全体攻撃調整・上限
        /// </summary>
        public static readonly CSysRate ALL_ATTACK_ADJUST_MAX = new(0.65, "ALL_ATTACK_ADJUST_MAX");
    }
}