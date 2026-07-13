using CSLib.Lib;
using Microsoft.AspNetCore.Mvc;
using PrivateApi.Common;
using PrivateApi.Common._Filter;
using PrivateApi.Domain.DTO;
using PrivateApi.Service;
using Repository.Repository;
using Repository.Repository.sql;
using System.Transactions;

namespace PrivateApi.Controller
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
            _posgre = new PostgreSQL(ConnectionString.Get(configuration));
        }

        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpGet("api/user/init")]
        public async Task<IActionResult> Init()
        {
            // 検証用に登録済みユーザを取得
            try
            {
                IEnumerable<UserDTO> users = await _service.RegistedSelectUsers();
                return StatusCode(HttpStatus.OK, users);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// ユーザ情報取得
        /// </summary>
        [HttpPost("api/user/userInfo")]
        public async Task<IActionResult> SelectUserOne([FromBody] string loginId)
        {
            try
            {
                UserDTO? user = await _service.SelectUserOne(loginId);
                return StatusCode(HttpStatus.OK, user);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// ユーザー登録
        /// </summary>
        [HttpPut("api/user/userRegist")]
        public async Task<IActionResult> UserRegist([FromBody] ReqUserRegist req)
        {
            using (var transaction = new TransactionScope(TransactionScopeAsyncFlowOption.Enabled))
            {
                try
                {
                    Validation.LoginId(req.loginId);
                    Validation.LoginPass(req.password);
                    Validation.DispName(req.dispName);
                    Validation.DispName(req.dispShortName);
                }
                catch (Exception e)
                {
                    return StatusCode(HttpStatus.BadRequest, Message.Create(e));
                }

                string loginId       = req.loginId.Trim();
                string password      = req.password.Trim();
                string dispName      = req.dispName.Trim();
                string dispShortName = req.dispShortName.Trim();

                try
                {
                    // 初期登録
                    await _service.InsertUser(loginId, password, dispName, dispShortName);
                    await _service.InsertStartUpMonsters(loginId);
                    await _service.InsertUsableStore(loginId);

                    // 処理完了
                    transaction.Complete();
                    string message = $"Regist user complete. LoginId: {loginId}, DispNaem: {dispName}";
                    return StatusCode(HttpStatus.Created, Message.Create(message));
                }
                catch (Exception e)
                {
                    string message = $"Error regist user. LoginId: {loginId}, DispNaem: {dispName}";
                    return StatusCode(HttpStatus.InternalServerError, Message.Create(e, message));
                }
            }
        }

        /// <summary>
        /// ログインユーザ情報を取得
        /// </summary>
        [HttpPost("api/user/loginUser")]
        public async Task<IActionResult> SelectUser([FromBody] string? loginId)
        {
            
            if (string.IsNullOrEmpty(loginId)) return StatusCode(HttpStatus.BadRequest);

            try
            {
                var param = new { login_id = loginId };

                IEnumerable<UserDTO?> users =
                    await _posgre.Select<UserDTO>(UserSQL.SelectLoginUser(), param);
                                       
                return StatusCode(HttpStatus.OK, users.SingleOrDefault());
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// 自己破産（所持金初期化）
        /// </summary>
        [HttpPut("api/user/restartAsPlayer")]
        public async Task<IActionResult> RestartAsPlayer([FromBody] string loginId)
        {
            if (string.IsNullOrEmpty(loginId)) return StatusCode(HttpStatus.BadRequest);

            try
            {
                await _service.RestartAsPlayer(loginId);

                UserDTO? user = await _service.SelectUserOne(loginId);
                return StatusCode(HttpStatus.OK, user);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// 使用可能なモンスター数を取得
        /// </summary>
        [HttpPost("api/user/getMonsterCount")]
        public async Task<ActionResult<string>> SelectMonsterCount([FromBody] string? loginId)
        {
            if (string.IsNullOrEmpty(loginId)) return StatusCode(HttpStatus.BadRequest);

            try
            {
                LittleDTO<int> result = await _service.SelectMonsterCount(loginId);
                return StatusCode(HttpStatus.OK, $"{result.Param1} / {result.Param2}");
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// 勝敗結果を記録（ユーザー）
        /// ショップ開放の確認
        /// </summary>
        [HttpPut("api/user/recordUserResults")]
        public async Task<IActionResult> RecordUserResults([FromBody] ReqUserResults req)
        {
            using (var transaction = new TransactionScope(TransactionScopeAsyncFlowOption.Enabled))
            {
                try
                {
                    // クレンジング
                    string loginId          = req.loginId.Trim();
                    string betMonsterId     = req.betMonsterId.Trim();
                    string winningMonsterId = req.winningMonsterId.Trim();

                    // 予想的中か
                    bool hit = false;
                    if (betMonsterId == winningMonsterId) hit = true;

                    // 各種登録
                    await _service.UpdateUserResults(hit, req.betGil, req.betRate, loginId);

                    IEnumerable<ShopDTO> InsertUsableShop = await _shopService.ExistsUsableShop(loginId);
                    if (InsertUsableShop.Count() > 0)
                    {
                        await _shopService.InsertUsableShop(loginId, InsertUsableShop);
                    }

                    // 処理完了
                    transaction.Complete();
                    return StatusCode(HttpStatus.OK, InsertUsableShop);
                }
                catch (Exception e)
                {
                    return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
                }
            }
        }
    }
}
