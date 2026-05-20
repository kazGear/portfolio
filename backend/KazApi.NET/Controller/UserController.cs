using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using KazApi.Repository;
using KazApi.Repository.sql;
using KazApi.Domain.DTO;
using System.Transactions;
using KazApi.Service;

namespace KazApi.Controller
{
    [ApiController]
    public class UserController : ControllerBase
    {
        private readonly UserService _service;
        private readonly ShopService _shopService;
        private readonly IDatabase _posgre;

        public UserController(IConfiguration configuration)
        {
            _service = new UserService(configuration);
            _shopService = new ShopService(configuration);
            _posgre = new PostgreSQL(configuration);
        }

        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpPost("api/user/init")]
        public ActionResult<string> Init()
        {
            // 検証用に登録済みユーザを取得
            IEnumerable<UserDTO> users = _service.RegistedSelectUsers();
            return JsonConvert.SerializeObject(users);
        }

        /// <summary>
        /// ユーザ情報取得
        /// </summary>
        [HttpPost("api/user/userInfo")]
        public ActionResult<string> SelectUserOne([FromQuery] string loginId)
        {
            UserDTO? user = _service.SelectUserOne(loginId);
            return JsonConvert.SerializeObject(user);
        }

        /// <summary>
        /// ユーザー登録
        /// </summary>
        [HttpPost("api/user/userRegist")]
        public ActionResult UserRegist(
            [FromQuery] string loginId,
            [FromQuery] string password,
            [FromQuery] string dispName,
            [FromQuery] string dispShortName)
        {
            using (TransactionScope transaction = new TransactionScope())
            {
                try
                {
                    // 空白除去
                    loginId = loginId.Trim();
                    password = password.Trim();
                    dispName = dispName.Trim();
                    dispShortName = dispShortName.Trim();
                    /*
                    TODO 引数検証
                    error >>> エラーページへ
                     */
                    // 初期登録
                    _service.InsertUser(loginId, password, dispName, dispShortName);
                    _service.InsertStartUpMonsters(loginId);
                    _service.InsertUsableStore(loginId);

                    // 処理完了
                    transaction.Complete();

                    string message = $"Regist user complete. LoginId: {loginId}, DispNaem: {dispName}";
                    Console.WriteLine(message);
                    return Ok(new { message = message });
                }
                catch (Exception)
                {
                    string message = $"Error regist user. LoginId: {loginId}, DispNaem: {dispName}";
                    Console.WriteLine(message);
                    return StatusCode(500, message);
                }

            }
        }

        /// <summary>
        /// ログインユーザ情報を取得
        /// </summary>
        [HttpPost("api/user/loginUser")]
        public ActionResult<string?> SelectUser([FromQuery] string? loginId)
        {
            var param = new { login_id = loginId };
            
            UserDTO? user = _posgre.Select<UserDTO>(UserSQL.SelectLoginUser(), param)
                                   .SingleOrDefault();

            return JsonConvert.SerializeObject(user);
        }

        /// <summary>
        /// 自己破産（所持金初期化）
        /// </summary>
        [HttpPost("api/user/restartAsPlayer")]
        public ActionResult<string> RestartAsPlayer([FromQuery] string loginId)
        {
            try
            {
                _service.RestartAsPlayer(loginId);

                UserDTO? user = _service.SelectUserOne(loginId);
                return JsonConvert.SerializeObject(user);
            }
            catch (Exception e)
            {
                return e.Message;
            }
            
        }

        /// <summary>
        /// 使用可能なモンスター数を取得
        /// </summary>
        [HttpPost("api/user/getMonsterCount")]
        public ActionResult<string> SelectMonsterCount([FromQuery] string loginId)
        {
            LittleDTO<int> result = _service.SelectMonsterCount(loginId);
            return JsonConvert.SerializeObject($"{result.Param1} / {result.Param2}");
        }

        /// <summary>
        /// 勝敗結果を記録（ユーザー）
        /// ショップ開放の確認
        /// </summary>
        [HttpPost("api/user/recordUserResults")]
        public ActionResult<string> RecordUserResults(
            [FromQuery] string betMonsterId,
            [FromQuery] int betGil,
            [FromQuery] decimal betRate,
            [FromQuery] string winningMonsterId,
            [FromQuery] string loginId)
        {
            using (TransactionScope transaction = new TransactionScope())
            {
                try
                {
                    if (string.IsNullOrEmpty(winningMonsterId))
                        return Ok(new { message = "No action required." });

                    // クレンジング
                    loginId = loginId.Trim();
                    betMonsterId = betMonsterId.Trim();
                    winningMonsterId = winningMonsterId.Trim();

                    // 予想的中か
                    bool hit = false;
                    if (betMonsterId == winningMonsterId) hit = true;

                    // 各種登録
                    _service.UpdateUserResults(hit, betGil, betRate, loginId);

                    IEnumerable<ShopDTO> InsertUsableShop = _shopService.ExistsUsableShop(loginId);
                    if (InsertUsableShop.Count() > 0)
                    {
                        _shopService.InsertUsableShop(loginId, InsertUsableShop);
                    }

                    // 処理完了
                    transaction.Complete();
                    return JsonConvert.SerializeObject(InsertUsableShop);
                }
                catch (Exception)
                {
                    return StatusCode(500, $"Error resist user result.");
                }
            }
        }
    }
}
