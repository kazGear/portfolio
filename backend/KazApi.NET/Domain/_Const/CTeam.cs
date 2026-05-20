namespace KazApi.Domain._Const
{
    /// <summary>
    /// チーム定数
    /// </summary>
    public class CTeam : Enumeration<int>
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        private CTeam(int id, string name) : base(id, name) { }

        public static readonly CTeam UNKNOWN = new(0, "UNKNOWN");
        public static readonly CTeam A = new(1, "A");
        public static readonly CTeam B = new(2, "B");
        public static readonly CTeam C = new(3, "C");
        public static readonly CTeam D = new(4, "D");
        public static readonly CTeam E = new(5, "E");
        public static readonly CTeam F = new(6, "F");
    }
}
