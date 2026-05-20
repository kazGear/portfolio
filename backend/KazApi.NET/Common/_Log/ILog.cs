namespace KazApi.Common._Log
{
    /// <summary>
    /// ログインターフェイス
    /// </summary>
    public interface ILog<T>
    {
        /// <summary>
        /// ログを記録
        /// </summary>
        public void Logging(T log);

        /// <summary>
        /// ログを出力
        /// </summary>
        public IList<T> DumpMemory();

        /// <summary>
        /// ログを全て出力
        /// </summary>
        public IList<T> DumpAll();
    }
}
