namespace KazApi.Domain._Const
{
    /// <summary>
    /// モンスタータイプ定数
    /// </summary>
    public class CMonsterType : Enumeration<string>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CMonsterType(string value, string name) : base(value, name) { }

        public static readonly CMonsterType キラービー = new("monsterType001", "キラービー");
        public static readonly CMonsterType カーミラ = new("monsterType002", "カーミラ");
        public static readonly CMonsterType デーモン = new("monsterType003", "デーモン");
        public static readonly CMonsterType ゴブリン = new("monsterType004", "ゴブリン");
        public static readonly CMonsterType マシンゴーレム = new("monsterType005", "マシンゴーレム");

        public static readonly CMonsterType ハーピー = new("monsterType006", "ハーピー");
        public static readonly CMonsterType アーマーナイト = new("monsterType007", "アーマーナイト");
        public static readonly CMonsterType マジシャン = new("monsterType008", "マジシャン");
        public static readonly CMonsterType マイコニド = new("monsterType009", "マイコニド");
        public static readonly CMonsterType ニードルバード = new("monsterType010", "ニードルバード");

        public static readonly CMonsterType プチドラゴン = new("monsterType011", "プチドラゴン");
        public static readonly CMonsterType ポト = new("monsterType012", "ポト");
        public static readonly CMonsterType プリースト = new("monsterType013", "プリースト");
        public static readonly CMonsterType ラビ = new("monsterType014", "ラビ");
        public static readonly CMonsterType グリーンスライム = new("monsterType015", "グリーンスライム");

        public static readonly CMonsterType イビルソード = new("monsterType016", "イビルソード");
        public static readonly CMonsterType ウルフ = new("monsterType017", "ウルフ");
        public static readonly CMonsterType ダック = new("monsterType018", "ダック");
        public static readonly CMonsterType モールベア = new("monsterType019", "モールベア");
        public static readonly CMonsterType ギャルビー = new("monsterType020", "ギャルビー");

        public static readonly CMonsterType サハギン = new("monsterType021", "サハギン");
        public static readonly CMonsterType クロウラー = new("monsterType022", "クロウラー");
        public static readonly CMonsterType パックン = new("monsterType023", "パックン");
        public static readonly CMonsterType チビデビル = new("monsterType024", "チビデビル");
        public static readonly CMonsterType オーガボックス = new("monsterType025", "オーガボックス");

        public static readonly CMonsterType バレッテ = new("monsterType026", "バレッテ");
        public static readonly CMonsterType バシリスク = new("monsterType027", "バシリスク");
        public static readonly CMonsterType スペクター = new("monsterType028", "スペクター");
        public static readonly CMonsterType ユニコーンヘッド = new("monsterType029", "ユニコーンヘッド");
        public static readonly CMonsterType シェイプシフター = new("monsterType030", "シェイプシフター");

        public static readonly CMonsterType ボルダー = new("monsterType031", "ボルダー");
        public static readonly CMonsterType パンプキンボム = new("monsterType032", "パンプキンボム");

        /// <summary>
        /// ユーザ登録初期から登場するモンスターたち 
        /// </summary>
        public static readonly IList<CMonsterType> START_UP = new List<CMonsterType>
        {
            ラビ,
            カーミラ,
            キラービー,
            ゴブリン,
            マシンゴーレム,
            ハーピー,
            アーマーナイト,
            マジシャン,
            マイコニド,
            ニードルバード
        };

        /// <summary>
        /// フィールド情報の一覧を取得
        /// </summary>
        public static IReadOnlyCollection<string> GetValues()
        {
            IReadOnlyCollection<string> values = new HashSet<string>()
            {
                キラービー.Value,
                カーミラ.Value,
                デーモン.Value,
                ゴブリン.Value,
                マシンゴーレム.Value,

                ハーピー.Value,
                アーマーナイト.Value,
                マジシャン.Value,
                マイコニド.Value,
                ニードルバード.Value,

                プチドラゴン.Value,
                ポト.Value,
                プリースト.Value,
                ラビ.Value,
                グリーンスライム.Value,

                イビルソード.Value,
                ウルフ.Value,
                ダック.Value,
                モールベア.Value,
                ギャルビー.Value,

                サハギン.Value,
                クロウラー.Value,
                パックン.Value,
                チビデビル.Value,
                オーガボックス.Value,

                バレッテ.Value,
                バシリスク.Value,
                スペクター.Value,
                ユニコーンヘッド.Value,
                シェイプシフター.Value,

                ボルダー.Value,
                パンプキンボム.Value,
            };
            return values;
        }

    }
}
