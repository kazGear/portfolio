using Dapper;
using Npgsql;
using System.Data;

namespace Repository.Repository
{
    public class PostgreSQL : IDisposable, IDatabase
    {
        private IDbConnection Connection { get; set; }
        private IDbTransaction? Transaction { get; set; }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public PostgreSQL(string connectionString)
        {
            Connection = new NpgsqlConnection(connectionString);
        }
        
        /// <summary>
        /// 接続状態
        /// </summary>
        public bool IsConnected()
        {
            return Connection.State == ConnectionState.Open;
        }

        /// <summary>
        /// 後処理
        /// </summary>
        public void Dispose()
        {
            throw new NotImplementedException();
        }

        /// <summary>
        /// コネクションオープン
        /// </summary>
        public void ConnectionOpen()
        {
            if (!IsConnected())
            {
                Connection.Open();
            }
        }

        /// <summary>
        /// コネクションクローズ
        /// </summary>
        public void ConnectionClose()
        {
            if (IsConnected())
            {
                Connection.Close();
            }
        }

        /// <summary>
        /// トランザクション開始
        /// </summary>
        public void BeginTransaction()
        {
            Transaction = Connection.BeginTransaction(IsolationLevel.Serializable);
        }

        /// <summary>
        /// コミット
        /// </summary>
        public void CommitTransaction()
        {
            if (Transaction != null && Transaction.Connection != null)
            {
                Transaction.Commit();
            }
        }

        /// <summary>
        /// ロールバック
        /// </summary>
        public void RollBackTransaction()
        {
            if (Transaction != null && Transaction.Connection != null)
            {
                Transaction.Rollback();
            }
        }

        public async Task<IEnumerable<T>> Select<T>(string query, object parameters)
        {
            return await Connection.QueryAsync<T>(query, parameters);
        }

        public async Task<IEnumerable<T>> Select<T>(string query)
        {
            return await Connection.QueryAsync<T>(query);
        }

        public async Task Execute(string query)
        {
            await Connection.ExecuteAsync(query);
        }

        public async Task Execute(string query, object parameters)
        {
            await Connection.ExecuteAsync(query, parameters);
        }
    }
}
