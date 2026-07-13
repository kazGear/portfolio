using CSLib.Lib;
using PrivateApi.Domain.DTO;
using Repository.Repository;
using Repository.Repository.sql;

namespace PrivateApi.Service
{
    public class AuthService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public AuthService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(configuration));
        }

        /// <summary>
        /// 認証実行 ユーザー検索・パスワード検証
        /// </summary>

        internal async Task<UserDTO?> AuthenticateUser(string loginId, string password)
        {
            string encryptPass = Aes.AesEncrypt(password);

            var param = new
            {
                login_id = loginId,
                login_pass = encryptPass
            };

            // ユーザー検索
            IEnumerable<UserDTO?> users = 
                await _posgre.Select<UserDTO>(AuthSQL.SelectLoginUser(), param);

            return users.SingleOrDefault();
        }
    }
}
