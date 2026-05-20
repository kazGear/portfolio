namespace KazApi.Repository
{
    /// <summary>
    /// DBインターフェイス
    /// </summary>
    public interface IDatabase
    {
        /// <summary>
        /// データ取得（パラメータバインド）
        /// </summary>
        public IEnumerable<T> Select<T>(string query, object parameters);
        /// <summary>
        /// データ取得
        /// </summary>
        public IEnumerable<T> Select<T>(string query);
        /// <summary>
        /// 更新操作
        /// </summary>
        public void Execute(string query);
        /// <summary>
        /// 更新操作（パラメータバインド）
        /// </summary>
        public void Execute(string query, object parameters);
                
    }
}
