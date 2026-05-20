using CSLib.Lib;
using KazApi.Domain._Const;
using KazApi.Domain.DTO;
using KazApi.Repository;
using KazApi.Repository.sql;

namespace KazApi.Service
{
    public class UserService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public UserService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(configuration);
        }

        /// <summary>
        /// ユーザ情報取得
        /// </summary>
        public UserDTO? SelectUserOne(string loginId)
        {
            var param = new { login_id = loginId };

            return _posgre.Select<UserDTO>(UserSQL.SelecUserInfo(loginId), param)
                          .SingleOrDefault();
        }

        /// <summary>
        /// 登録済ユーザーを取得
        /// </summary>
        /// <returns></returns>
        public IEnumerable<UserDTO> RegistedSelectUsers()
            => _posgre.Select<UserDTO>(UserSQL.SelecUserInfo());

        /// <summary>
        /// ユーザ登録
        /// </summary>

        public void InsertUser(
            string LoginId,
            string Password,
            string DispName,
            string DispShortName)
        {
            try
            {
                // 暗号化
                string encryptPass = UAes.AesEncrypt(Password);

                var param = new
                {
                    login_id = LoginId,
                    login_pass = encryptPass,
                    disp_name = DispName,
                    disp_short_name = DispShortName
                };

                _posgre.Execute(UserSQL.InsertUser(), param);
            }
            catch (Exception)
            {
                throw;
            }
        }

        /// <summary>
        /// ユーザが初期か使えるモンスタを設定する
        /// </summary>
        public void InsertStartUpMonsters(string loginId)
        {
            try
            {
                DateTime now = DateTime.Now;

                foreach (CMonsterType monsterType in CMonsterType.START_UP)
                {
                    var param = new
                    {
                        user_id = loginId,
                        item_id = monsterType.Value,
                        not_use_this = false
                    };
                    _posgre.Execute(UserSQL.InsertStartUpMonsters(), param);
                }
            }
            catch (Exception)
            {
                throw;
            }
        }

        /// <summary>
        /// 自己破産（所持金初期化）
        /// </summary>
        /// <returns></returns>
        public void RestartAsPlayer(string loginId)
        {
            var param = new { login_id = loginId };
            _posgre.Execute(UserSQL.RestartAsPlayer(), param);
        }

        /// <summary>
        /// 使用可能なモンスター数を取得
        /// </summary>
        public LittleDTO<int> SelectMonsterCount(string loginId)
        {
            var param = new { login_id = loginId };

            LittleDTO<int> result = _posgre.Select<LittleDTO<int>>(UserSQL.SelectMonsterCount(), param)
                                          .Single();
            return result;
        }

        /// <summary>
        /// 使用可能ショップを登録
        /// </summary>
        public void InsertUsableStore(string loginId)
        {
            var param = new { login_id = loginId };
            _posgre.Execute(UserSQL.InsertUsableStore(), param);
        }

        /// <summary>
        /// アイテム購入に伴う所持金の更新
        /// </summary>
        public void Purchase(string loginId, int cashAfterPurchase)
        {
            var param = new
            {
                login_id = loginId,
                cash = cashAfterPurchase
            };
            _posgre.Execute(UserSQL.Purchase(), param);
        }

        /// <summary>
        /// 勝敗結果を記録（ユーザー）
        /// </summary>
        public bool UpdateUserResults(bool hit, int betGil, decimal betRate, string loginId)
        {
            try
            {
                var param = new
                {
                    login_id = loginId,
                    wins = hit ? 1 : 0,
                    losses = hit ? 0 : 1,
                    cash = hit ? Math.Floor(betGil * betRate) : -1 * betGil,
                    wins_get_cash = hit ? Math.Floor(betGil * betRate) : 0,
                    losses_lost_cash = hit ? 0 : betGil,
                };
                _posgre.Execute(UserSQL.InsertUserResult(), param);

                return true;
            }
            catch (Exception)
            {
                return false;
            }
        }
    }
}
