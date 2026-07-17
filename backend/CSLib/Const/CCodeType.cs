namespace CSLib.Const
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
        public static readonly CCodeType ELEMENT = new("battle_elements", "ELEMENT");
        ///// <summary>
        ///// 状態
        ///// </summary>
        public static readonly CCodeType STATE = new("battle_status", "STATE");
        ///// <summary>
        ///// 対象
        ///// </summary>
        public static readonly CCodeType TARGET = new("battle_target_types", "TARGET");
        ///// <summary>
        ///// スキル
        ///// </summary>
        public static readonly CCodeType SKILL = new("battle_skill_types", "SKILL");
        ///// <summary>
        ///// システム補正率
        ///// </summary>
        public static readonly CCodeType SYS_RATE = new("battle_system_rates", "SYS_RATE");
        ///// <summary>
        ///// モンスター
        ///// </summary>
        public static readonly CCodeType MONSTER = new("battle_monster_types", "MONSTER");
        ///// <summary>
        ///// ロール
        ///// </summary>
        public static readonly CCodeType ROLE = new("user_rolls", "ROLE");
        ///// <summary>
        ///// 設定種類
        ///// </summary>
        public static readonly CCodeType EDIT = new("battle_configs", "EDIT");
    }
}
