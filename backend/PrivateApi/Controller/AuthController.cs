using CSLib.Lib;
using PrivateApi.Common;
using PrivateApi.Domain.DTO;
using PrivateApi.Service;
using Microsoft.AspNetCore.Mvc;

namespace PrivateApi.Controller
{
    public class AuthController : ControllerBase
    {
        private readonly IConfiguration _configuration;
        private readonly AuthService _service;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public AuthController(IConfiguration configuration)
        {
            _configuration = configuration;
            _service = new AuthService(configuration);
        }

        /// <summary>
        /// ログイン実行
        /// </summary>
        [HttpPost("api/auth/login")]
        public IActionResult Login([FromBody] ReqLogin req)
        {
            try
            {
                string loginId  = req.loginId.Trim();
                string password = req.password.Trim();

                // ユーザの認証
                UserDTO? user = _service.AuthenticateUser(loginId, password);

                // 認証失敗
                if (user == null) return StatusCode(HttpStatus.Unauthorized);

                // トークン発行
                string token = Jwt.GenerateJwtToken(user.LoginId, _configuration);
                user.Token = token;

                return StatusCode(HttpStatus.OK, user);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// トークンが有効か確認する
        /// </summary>
        [HttpPost("api/auth/checkToken")]
        public IActionResult IsValidToken()
        {
            string? token = Request.Headers["Authorization"].ToString();

            if (string.IsNullOrEmpty(token))
            {
                return StatusCode(HttpStatus.Unauthorized);
            }
            bool isValid = Jwt.IsValidToken(token);

            if (isValid)
            {
                return StatusCode(HttpStatus.OK);
            }
            else
            {
                return StatusCode(HttpStatus.Unauthorized);
            }
        }
    }
}
