namespace Repository.Repository
{
    /// <summary>
    /// DBインターフェイス
    /// </summary>
    public interface IDatabase
    {
        /// <summary>
        /// データ取得（パラメータバインド）
        /// </summary>
        public Task<IEnumerable<T>> Select<T>(string query, object parameters);
        /// <summary>
        /// データ取得
        /// </summary>
        public Task<IEnumerable<T>> Select<T>(string query);
        /// <summary>
        /// 更新操作
        /// </summary>
        public Task Execute(string query);
        /// <summary>
        /// 更新操作（パラメータバインド）
        /// </summary>
        public Task Execute(string query, object parameters);
    }
}
