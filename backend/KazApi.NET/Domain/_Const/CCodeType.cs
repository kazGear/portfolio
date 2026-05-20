namespace KazApi.Domain._Const
{
    /// <summary>
    /// コードタイプ定数
    /// </summary>
    public class CCodeType : Enumeration<string>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CCodeType(string id, string name) : base(id, name) { }

        /// <summary>
        /// 自然属性
        /// </summary>
        public static readonly CCodeType ELEMENT = new("code001", "ELEMENT");
        ///// <summary>
        ///// 状態
        ///// </summary>
        public static readonly CCodeType STATE = new("code002", "STATE");
        ///// <summary>
        ///// 対象
        ///// </summary>
        public static readonly CCodeType TARGET = new("code003", "TARGET");
        ///// <summary>
        ///// スキル
        ///// </summary>
        public static readonly CCodeType SKILL = new("code004", "SKILL");
        ///// <summary>
        ///// システム補正率
        ///// </summary>
        public static readonly CCodeType SYS_RATE = new("code005", "SYS_RATE");
        ///// <summary>
        ///// モンスター
        ///// </summary>
        public static readonly CCodeType MONSTER = new("code006", "MONSTER");
        ///// <summary>
        ///// ロール
        ///// </summary>
        public static readonly CCodeType ROLE = new("code007", "ROLE");
        ///// <summary>
        ///// 設定種類
        ///// </summary>
        public static readonly CCodeType EDIT = new("code008", "EDIT");
    }
}
