using CSLib.Lib;
using KazApi.Domain.DTO;
using KazApi.Repository;
using KazApi.Repository.sql;

namespace KazApi.Service
{
    public class AuthService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public AuthService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(configuration);
        }

        /// <summary>
        /// 認証実行 ユーザー検索・パスワード検証
        /// </summary>

        internal UserDTO? AuthenticateUser(string loginId, string password)
        {
            string encryptPass = UAes.AesEncrypt(password);

            var param = new
            {
                login_id = loginId,
                login_pass = encryptPass
            };

            // ユーザー検索
            UserDTO? user = _posgre.Select<UserDTO>(AuthSQL.SelectLoginUser(), param)
                                   .SingleOrDefault();
            return user;
        }

    }
}
