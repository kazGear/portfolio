using CSLib.Lib;
using KazApi.Domain.DTO;
using KazApi.Service;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;

namespace KazApi.Controller
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
        public IActionResult Login([FromForm] string? loginId,
                                   [FromForm] string? password)
        {
            loginId = loginId != null ? loginId.Trim() : null;
            if (loginId == null || password == null) return Unauthorized();

            // ユーザの認証
            UserDTO? user = _service.AuthenticateUser(loginId, password);

            // 認証失敗
            if (user == null) return Unauthorized();

            // トークン発行
            string token = Jwt.GenerateJwtToken(user.LoginId, _configuration);
            user.Token = token;

            return Ok(JsonConvert.SerializeObject(user));
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
                return StatusCode(HttpStatus.Unauthorized, new { });
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
