using CSLib.Lib;
using KazApi.Domain.DTO;
using KazApi.Service;
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
        public IActionResult Login([FromQuery] string? loginId, [FromQuery] string? password)
        {
            loginId = loginId != null ? loginId.Trim() : null; 
            if (loginId == null || password == null) return Unauthorized();

            // ユーザの認証
            UserDTO? user = _service.AuthenticateUser(loginId, password);
            
            // 認証失敗
            if (user == null) return Unauthorized();
            
            // トークン発行
            string token = UJwt.GenerateJwtToken(user.LoginId, _configuration);
            user.Token = token;

            return Ok(JsonConvert.SerializeObject(user));
        }

        /// <summary>
        /// トークンが有効か確認する
        /// </summary>
        [HttpPost("api/auth/checkToken")]
        public IActionResult IsValidToken([FromQuery] string token)
        {
            if (token == null) return Ok(false);

            bool isValid = UJwt.IsValidToken(token);
            return Ok(isValid);
        }
    }
}
