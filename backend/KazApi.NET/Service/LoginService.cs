using CSLib.Lib;
using PrivateApi.Domain._User;
using Repository.Repository;

namespace PrivateApi.Service
{
    public class LoginService
    {
        private readonly IDatabase _posgre;

        public LoginService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(configuration));
        }

        // ユーザの一覧を取得
        public IEnumerable<IUser> SelectLoginUsers(string name, string pass)
        {
            object parameters = new
            {
                username = name,
                password = pass
            };

            string select = @"
                SELECT username
                  FROM users
                 WHERE username   = @username
	               AND password   = @password
	               AND is_invalid = FALSE;
                ";

            return _posgre.Select<IUser>(select, parameters);
        }
    }
}
