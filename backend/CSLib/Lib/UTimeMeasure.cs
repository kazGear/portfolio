using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CSLib.Lib
{
    /// <summary>
    /// タイマーユーティリティ
    /// </summary>
    public class UTimeMeasure
    {
        private static Stopwatch _stopWatch = new();

        /// <summary>
        /// コンストラクタ
        /// </summary>
        private UTimeMeasure()
        {

        }
        /// <summary>
        /// 時間計測開始
        /// </summary>
        public static void Start()
        {
            _stopWatch = new();
            _stopWatch.Start();
        }
        /// <summary>
        /// 計測終了
        /// </summary>
        public static string Stop()
        {
            _stopWatch.Stop();
            TimeSpan ts = _stopWatch.Elapsed;

            return 
                $"{ts.Hours}:".ToString().PadLeft(3, '0')
                + $"{ts.Minutes}:".ToString().PadLeft(3, '0')
                + $"{ts.Seconds}.".ToString().PadLeft(3, '0')
                + $"{ts.Milliseconds}".ToString().PadLeft(3, '0');
        }
    }
}
