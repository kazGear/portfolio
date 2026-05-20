namespace KazApi.Domain._Const
{
    /// <summary>
    /// スキル定数
    /// </summary>
    public class CSkill : Enumeration<int>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CSkill(int id, string name) : base(id, name) { }

        public static readonly CSkill 打撃 = new(1, "打撃");
        public static readonly CSkill 正拳突き = new(2, "正拳突き");
        public static readonly CSkill ライジングドラゴン = new(3, "ライジングドラゴン");
        public static readonly CSkill リアルインパクト = new(4, "リアルインパクト,");
        public static readonly CSkill 回し蹴り = new(5, "回し蹴り");

        public static readonly CSkill ムーンサルト = new(6, "ムーンサルト");
        public static readonly CSkill ダンスマカブル = new(7, "ダンスマカブル");
        public static readonly CSkill クレイジーダンス = new(8, "クレイジーダンス");
        public static readonly CSkill 斬撃 = new(9, "斬撃");
        public static readonly CSkill 剣の舞 = new(10, "剣の舞");

        public static readonly CSkill 渾身斬り = new(11, "渾身斬り");
        public static readonly CSkill 次元斬 = new(12, "次元斬");
        public static readonly CSkill 薙ぎ払い = new(13, "薙ぎ払い");
        public static readonly CSkill 疾走居合 = new(14, "疾走居合");
        public static readonly CSkill ギガスラッシュ = new(15, "ギガスラッシュ");

        public static readonly CSkill 次元斬_絶 = new(16, "次元斬_絶");
        public static readonly CSkill ファイアボール = new(17, "ファイアボール");
        public static readonly CSkill エクスプロード = new(18, "エクスプロード");
        public static readonly CSkill ブレイズウォール = new(19, "ブレイズウォール");
        public static readonly CSkill アイススマッシュ = new(20, "アイススマッシュ");

        public static readonly CSkill メガスプラッシュ = new(21, "メガスプラッシュ");
        public static readonly CSkill コールドブレイズ = new(22, "コールドブレイズ");
        public static readonly CSkill サンダー = new(23, "サンダー");
        public static readonly CSkill サンダーボルト = new(24, "サンダーボルト");
        public static readonly CSkill サンダーストーム = new(25, "サンダーストーム");

        public static readonly CSkill ダイヤミサイル = new(26, "ダイヤミサイル");
        public static readonly CSkill アースクエイク = new(27, "アースクエイク");
        public static readonly CSkill ストーンクラウド = new(28, "ストーンクラウド");
        public static readonly CSkill ホーリーボール = new(29, "ホーリーボール");
        public static readonly CSkill セイントビーム = new(30, "セイントビーム");

        public static readonly CSkill ホーリーバースト = new(31, "ホーリーバースト");
        public static readonly CSkill イビルゲート = new(32, "イビルゲート");
        public static readonly CSkill ダークフォース = new(33, "ダークフォース");
        public static readonly CSkill ブラックレイン = new(34, "ブラックレイン");
        public static readonly CSkill グラビデ = new(35, "グラビデ");

        public static readonly CSkill グラビガ = new(36, "グラビガ");
        public static readonly CSkill グラビジャ = new(37, "グラビジャ");
        public static readonly CSkill デススペル = new(38, "デススペル");
        public static readonly CSkill ケアル = new(39, "ケアル");
        public static readonly CSkill ケアルラ = new(40, "ケアルラ");

        public static readonly CSkill ケアルガ = new(41, "ケアルガ");
        public static readonly CSkill ケアルジャ = new(42, "ケアルジャ");
        public static readonly CSkill リジェネ = new(43, "リジェネ");
        public static readonly CSkill ポイズン = new(44, "ポイズン");
        public static readonly CSkill ポイズンフラワー = new(45, "ポイズンフラワー");

        public static readonly CSkill デッドリーポイズン = new(46, "デッドリーポイズン");
        public static readonly CSkill スリプル = new(47, "スリプル");
        public static readonly CSkill スリープミスト = new(48, "スリープミスト");
        public static readonly CSkill チャーム = new(49, "チャーム");
        public static readonly CSkill スロウ = new(50, "スロウ");

        public static readonly CSkill スロウガ = new(51, "スロウガ");
        public static readonly CSkill バーサク = new(52, "バーサク");
        public static readonly CSkill ビジョン = new(53, "ビジョン");
        public static readonly CSkill エナジーボール = new(54, "エナジーボール");
        public static readonly CSkill ミス = new(55, "ミス");

        public static readonly CSkill 様子を見る = new(56, "様子を見る");
        public static readonly CSkill 余裕に構えている = new(57, "余裕に構えている");
        public static readonly CSkill 噛みつき = new(58, "噛みつき");
        public static readonly CSkill 喰いちぎり = new(59, "喰いちぎり");
        public static readonly CSkill タックル = new(60, "タックル");

        public static readonly CSkill 突撃 = new(61, "突撃");
        public static readonly CSkill 振り回す = new(62, "振り回す");
        public static readonly CSkill フルスイング = new(63, "フルスイング");
        public static readonly CSkill 突き = new(64, "突き");
        public static readonly CSkill 串刺し = new(65, "串刺し");

        public static readonly CSkill 叩きつけ = new(66, "叩きつけ");
        public static readonly CSkill 叩き潰し = new(67, "叩き潰し");
        public static readonly CSkill 引き裂く = new(68, "引き裂く");
        public static readonly CSkill 首狩り = new(69, "首狩り");
    }

}
