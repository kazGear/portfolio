using CSLib.Const;
using CSLib.Lib;
using PrivateApi.Domain.DTO;
using Repository.Repository;
using Repository.Repository.sql;

namespace PrivateApi.Service
{
    public class UserService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public UserService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(configuration));
        }

        /// <summary>
        /// ユーザ情報取得
        /// </summary>
        public async Task<UserDTO?> SelectUserOne(string loginId)
        {
            var param = new { login_id = loginId };

            var users = await _posgre.Select<UserDTO>(UserSQL.SelecUserInfo(loginId), param);
            return users.SingleOrDefault();
        }

        /// <summary>
        /// 登録済ユーザーを取得
        /// </summary>
        /// <returns></returns>
        public async Task<IEnumerable<UserDTO>> RegistedSelectUsers()
            => await _posgre.Select<UserDTO>(UserSQL.SelecUserInfo());

        /// <summary>
        /// ユーザ登録
        /// </summary>

        public async Task InsertUser(string LoginId,
                                     string Password,
                                     string DispName,
                                     string DispShortName)
        {
            // 暗号化
            string encryptPass = Aes.AesEncrypt(Password);

            var param = new
            {
                login_id = LoginId,
                login_pass = encryptPass,
                disp_name = DispName,
                disp_short_name = DispShortName
            };

            await _posgre.Execute(UserSQL.InsertUser(), param);
        }

        /// <summary>
        /// ユーザが初期か使えるモンスタを設定する
        /// </summary>
        public async Task InsertStartUpMonsters(string loginId)
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
                await _posgre.Execute(UserSQL.InsertStartUpMonsters(), param);
            }
        }

        /// <summary>
        /// 自己破産（所持金初期化）
        /// </summary>
        /// <returns></returns>
        public async Task RestartAsPlayer(string loginId)
        {
            var param = new { login_id = loginId };
            await _posgre.Execute(UserSQL.RestartAsPlayer(), param);
        }

        /// <summary>
        /// 使用可能なモンスター数を取得
        /// </summary>
        public async Task<LittleDTO<int>> SelectMonsterCount(string loginId)
        {
            var param = new { login_id = loginId };

            IEnumerable<LittleDTO<int>> result =
                await _posgre.Select<LittleDTO<int>>(UserSQL.SelectMonsterCount(), param);

            return result.First();
        }

        /// <summary>
        /// 使用可能ショップを登録
        /// </summary>
        public async Task InsertUsableStore(string loginId)
        {
            var param = new { login_id = loginId };
            await _posgre.Execute(UserSQL.InsertUsableStore(), param);
        }

        /// <summary>
        /// アイテム購入に伴う所持金の更新
        /// </summary>
        public async Task Purchase(string loginId, int cashAfterPurchase)
        {
            var param = new
            {
                login_id = loginId,
                cash     = cashAfterPurchase
            };
            await _posgre.Execute(UserSQL.Purchase(), param);
        }

        /// <summary>
        /// 勝敗結果を記録（ユーザー）
        /// </summary>
        public async Task UpdateUserResults(bool hit, int betGil, decimal betRate, string loginId)
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
            await _posgre.Execute(UserSQL.InsertUserResult(), param);
        }
    }
}
