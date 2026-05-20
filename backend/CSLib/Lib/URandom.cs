namespace CSLib.Lib
{
    /// <summary>
    /// 乱数ユーティリティ
    /// </summary>
    public class URandom
    {
        //private readonly int _seed = Environment.TickCount;
        private readonly Random _random;

        /// <summary>
        /// コンストラクタ
        /// 乱数シードが毎回変わる
        /// </summary>
        public URandom()
        {
            DateTime dt = DateTime.Now;
            int seed = dt.Second + dt.Millisecond + dt.Nanosecond;
            _random = new Random(seed);
        }

        /// <summary>
        /// int型の乱数を取得（範囲指定）
        /// minValue: 下限を含む
        /// maxValue: 上限値 + 1（この数値は範囲に含まない）
        /// </summary>
        public int RandomInt(int minValue, int maxValue)
        {
            return _random.Next(minValue, maxValue);
        }

        /// <summary>
        /// double型の乱数を取得（範囲指定、小数点以下切り捨て） 
        /// minValue: 下限を含む
        /// maxValue: 上限値 + 1（この数値は範囲に含まない）
        /// </summary>
        public double RandomDouble(double minValue, double maxValue)
        {
            return _random.NextDouble() * (maxValue - minValue) + minValue;
        }

        /// <summary>
        /// 対象の数値を rate(%) に応じて増減させる
        /// rateは小数点で表現すること（0.0 <= rate <= 1.0）
        /// </summary>
        public int RandomChangeInt(int target, double rate)
        {
            if (rate >= 1.01) throw new ArgumentException("比率は1.0以下の数値で表現してください。");

            double impactFromRate = target * RandomDouble(0, rate);
            bool randBool = RandomBool();

            // 負数に変換
            if (!randBool) impactFromRate *= -1;

            return target + (int)impactFromRate;
        }

        /// <summary>
        /// true / false のいずれかを生成 
        /// </summary>
        public bool RandomBool()
        {
            return _random.Next(2) == 1 ? true : false;
        }
    }
}
