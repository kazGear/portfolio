using Dapper;
using Npgsql;
using System.Data;
using System.Text.Json;

namespace KazApi.Repository
{
    public class PostgreSQL : IDisposable, IDatabase
    {
        private string _dbUser { get; set; }
        private string _dbPassword { get; set; }
        private string _dbName { get; set; }
        private string _dbHost { get; set; }
        private string _dbPort { get; set; }

        private IDbConnection Connection { get; set; }
        private IDbTransaction Transaction { get; set; }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public PostgreSQL(IConfiguration configuration)
        {
            Connection = CreateConnection();
        }
        
        /// <summary>
        /// コンストラクタ autoBattle batch 用
        /// </summary>
        public PostgreSQL()
        {
            Connection = CreateConnection();
        }

        /// <summary>
        /// 接続文字列の作成
        /// </summary>
        private IDbConnection CreateConnection()
        {
            string connectionString = string.Empty;
            bool onTheDocker = Environment.GetEnvironmentVariable("DOTNET_RUNNING_IN_CONTAINER") == "true";

            if (onTheDocker)
            {
                _dbUser     = Environment.GetEnvironmentVariable("DB_USER")!;
                _dbPassword = Environment.GetEnvironmentVariable("DB_PASSWORD")!;
                _dbName     = Environment.GetEnvironmentVariable("DB_NAME")!;
                _dbHost     = Environment.GetEnvironmentVariable("DB_HOST")!;
                _dbPort     = Environment.GetEnvironmentVariable("DB_PORT")!;

                connectionString = $"Server={_dbHost};Port={_dbPort};User Id={_dbUser};Password={_dbPassword};Database={_dbName}";
            }
            else // windowsデバック
            {
                string json      = File.ReadAllText("appsettings.Development.json");
                JsonDocument doc = JsonDocument.Parse(json);
                connectionString = doc.RootElement
                                      .GetProperty("ConnectionStrings")
                                      .GetProperty("DefaultConnection")
                                      .GetString()!;
            }
            return new NpgsqlConnection(connectionString);
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

        public IEnumerable<T> Select<T>(string query, object parameters)
        {
            return Connection.Query<T>(query, parameters).ToList();
        }

        public IEnumerable<T> Select<T>(string query)
        {
            return Connection.Query<T>(query).ToList();
        }

        public void Execute(string query)
        {
            Connection.Execute(query);
        }

        public void Execute(string query, object parameters)
        {
            Connection.Execute(query, parameters);
        }
    }
}
