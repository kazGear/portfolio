using System.Diagnostics;

namespace CSLib.Lib
{
    /// <summary>
    /// タイマーユーティリティ
    /// </summary>
    public class TimeMeasure
    {
        private Stopwatch _stopWatch;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public TimeMeasure()
        {
            _stopWatch = new();
        }
        /// <summary>
        /// 時間計測開始
        /// </summary>
        public void Start()
        {
            _stopWatch.Start();
        }
        /// <summary>
        /// 計測終了
        /// </summary>
        public TimeSpan Stop()
        {
            _stopWatch.Stop();

            TimeSpan ts = _stopWatch.Elapsed;
            return ts;
        }
    }
}
