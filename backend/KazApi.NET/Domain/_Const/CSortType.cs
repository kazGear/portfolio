namespace KazApi.Domain._Const
{
    public class CSortType
    {

    }

    /// <summary>
    /// 帳票（モンスター毎）のソート
    /// </summary>
    public class CReportSortType : Enumeration<int>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CReportSortType(int id, string name) : base(id, name) { }

        /// <summary>
        /// 初期状態ソート
        /// </summary>
        public static readonly CReportSortType DEFAULT = new(0, "DEFAULT");
        /// <summary>
        /// モンスター名
        /// </summary>
        public static readonly CReportSortType MONSTER_NAME = new(1, "MONSTER_NAME");
        /// <summary>
        /// 勝利数
        /// </summary>
        public static readonly CReportSortType WIN_COUNT = new(2, "WIN_COUNT");
        /// <summary>
        /// 対戦数
        /// </summary>
        public static readonly CReportSortType BATTLE_COUNT = new(3, "BATTLE_COUNT");
        /// <summary>
        /// 勝率
        /// </summary>
        public static readonly CReportSortType WINS_RATE = new(4, "WINS_RATE");
    }
}
