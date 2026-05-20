namespace KazApi.Domain._Const
{
    /// <summary>
    /// モンスター定数
    /// </summary>
    public class CMonster : Enumeration<string>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CMonster(string id, string name) : base(id, name) { }

        public static readonly CMonster キラービー = new("monster001", "キラービー");
        public static readonly CMonster アサシンバグ = new("monster002", "アサシンバグ");
        public static readonly CMonster ラスターバグ = new("monster003", "ラスターバグ");
        public static readonly CMonster カーミラ = new("monster004", "カーミラ");
        public static readonly CMonster カーミラクイーン = new("monster005", "カーミラクイーン)");

        public static readonly CMonster デーモン = new("monster006", "デーモン");
        public static readonly CMonster グレートデーモン = new("monster007", "グレートデーモン");
        public static readonly CMonster ゴブリン = new("monster008", "ゴブリン");
        public static readonly CMonster ゴブリンガード = new("monster009", "ゴブリンガード");
        public static readonly CMonster ゴブリンロード = new("monster010", "ゴブリンロード");
        
        public static readonly CMonster マシンゴーレム = new("monster011", "マシンゴーレム");
        public static readonly CMonster ガーディアン = new("monster012", "ガーディアン");
        public static readonly CMonster デスマシン = new("monster013", "デスマシン");
        public static readonly CMonster ハーピー = new("monster014", "ハーピー");
        public static readonly CMonster セイレーン = new("monster015", "セイレーン");

        public static readonly CMonster アーマーナイト = new("monster016", "アーマーナイト");
        public static readonly CMonster ダークナイト = new("monster017", "ダークナイト");
        public static readonly CMonster ターミネーター = new("monster018", "ターミネーター");
        public static readonly CMonster マジシャン = new("monster019", "マジシャン");
        public static readonly CMonster ウィザード = new("monster020", "ウィザード");

        public static readonly CMonster ハイウィザード = new("monster021", "ハイウィザード");
        public static readonly CMonster マイコニド = new("monster022", "マイコニド");
        public static readonly CMonster ダースマタンゴ = new("monster023", "ダースマタンゴ");
        public static readonly CMonster ニードルバード = new("monster024", "ニードルバード");
        public static readonly CMonster コカトバード = new("monster025", "コカトバード");

        public static readonly CMonster プチドラゴン = new("monster026", "プチドラゴン");
        public static readonly CMonster プチドラゾンビ = new("monster027", "プチドラゾンビ");
        public static readonly CMonster フロストドラゴン = new("monster028", "フロストドラゴン");
        public static readonly CMonster プチティアマット = new("monster029", "プチティアマット");
        public static readonly CMonster ポト = new("monster030", "ポト");

        public static readonly CMonster マーマポト = new("monster031", "マーマポト");
        public static readonly CMonster パーパポト = new("monster032", "パーパポト");
        public static readonly CMonster プリースト = new("monster033", "プリースト");
        public static readonly CMonster カオスソーサラー = new("monster034", "カオスソーサラー");
        public static readonly CMonster イビルシャーマン = new("monster035", "イビルシャーマン");

        public static readonly CMonster ラビ = new("monster036", "ラビ");
        public static readonly CMonster ラビリオン = new("monster037", "ラビリオン");
        public static readonly CMonster キングラビ = new("monster038", "キングラビ");
        public static readonly CMonster グレートラビ = new("monster039", "グレートラビ");
        public static readonly CMonster グリーンスライム = new("monster040", "グリーンスライム");

        public static readonly CMonster ブルーババロア = new("monster041", "ブルーババロア");
        public static readonly CMonster レッドマシュマロ = new("monster042", "レッドマシュマロ");
        public static readonly CMonster イビルソード = new("monster043", "イビルソード");
        public static readonly CMonster イビルウェポン = new("monster044", "イビルウェポン");
        public static readonly CMonster エレメントソード = new("monster045", "エレメントソード");

        public static readonly CMonster ケルベロス = new("monster046", "ケルベロス");
        public static readonly CMonster バウンドウルフ = new("monster047", "バウンドウルフ");
        public static readonly CMonster ジャッカル = new("monster048", "ジャッカル");
        public static readonly CMonster ダックソルジャー = new("monster049", "ダックソルジャー");
        public static readonly CMonster ダックジェネラル = new("monster050", "ダックジェネラル");

        public static readonly CMonster モールベア = new("monster051", "モールベア");
        public static readonly CMonster ニードリオン = new("monster052", "ニードリオン");
        public static readonly CMonster ギャルビー = new("monster053", "ギャルビー");
        public static readonly CMonster レディビー = new("monster054", "レディビー");
        public static readonly CMonster クインビー = new("monster055", "クインビー");

        public static readonly CMonster サハギン = new("monster056", "サハギン");
        public static readonly CMonster プチポセイドン = new("monster057", "プチポセイドン");
        public static readonly CMonster クロウラー = new("monster058", "クロウラー");
        public static readonly CMonster メガクロウラー = new("monster059", "メガクロウラー");
        public static readonly CMonster ギガクロウラー = new("monster060", "ギガクロウラー");

        public static readonly CMonster ぱっくんオタマ = new("monster061", "ぱっくんオタマ");
        public static readonly CMonster ぱっくりオタマ = new("monster062", "ぱっくりオタマ");
        public static readonly CMonster ぱっくんトカゲ = new("monster063", "ぱっくんトカゲ");
        public static readonly CMonster ぱっくんドラゴン = new("monster064", "ぱっくんドラゴン");
        public static readonly CMonster チビデビル = new("monster065", "チビデビル");

        public static readonly CMonster グレムリン = new("monster066", "グレムリン");
        public static readonly CMonster オーガボックス = new("monster067", "オーガボックス");
        public static readonly CMonster カイザーミミック = new("monster068", "カイザーミミック");
        public static readonly CMonster バレッテ = new("monster069", "バレッテ");
        public static readonly CMonster ゴールドバレッテ = new("monster070", "ゴールドバレッテ");

        public static readonly CMonster バシリスク = new("monster071", "バシリスク");
        public static readonly CMonster ファイアドレイク = new("monster072", "ファイアドレイク");
        public static readonly CMonster スペクター = new("monster073", "スペクター");
        public static readonly CMonster ゴースト = new("monster074", "ゴースト");
        public static readonly CMonster ユニコーンヘッド = new("monster075", "ユニコーンヘッド");

        public static readonly CMonster ゴールドユニコ = new("monster076", "ゴールドユニコ");
        public static readonly CMonster シェイプシフター = new("monster077", "シェイプシフター");
        public static readonly CMonster シャドウゼロ = new("monster078", "シャドウゼロ");
        public static readonly CMonster シャドウゼロワン = new("monster079", "シャドウゼロワン");
        public static readonly CMonster ボルダー = new("monster080", "ボルダー");

        public static readonly CMonster パワーボルダー = new("monster081", "パワーボルダー");
        public static readonly CMonster デスボルダー = new("monster082", "デスボルダー");
        public static readonly CMonster パンプキンボム = new("monster083", "パンプキンボム");
        public static readonly CMonster グレネードボム = new("monster084", "グレネードボム");
    }
}
