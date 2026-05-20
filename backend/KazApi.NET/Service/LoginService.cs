using KazApi.Domain._User;
using KazApi.Repository;

namespace KazApi.Service
{
    public class LoginService
    {
        private readonly IDatabase _posgre;

        public LoginService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(configuration);
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
