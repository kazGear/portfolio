using CSLib.Lib;
using CSLib.Middleware;
using Microsoft.AspNetCore.Mvc;
using PrivateApi.Domain._User;
using PrivateApi.Service;

namespace PrivateApi.Controller
{
    [Route("api/[controller]")]
    [ApiController]
    public class LoginController : ControllerBase
    {
        private readonly LoginService _service;

        public LoginController(IConfiguration configuration)
        {
            _service = new LoginService(configuration);
        }

        // ユーザ一覧を取得する
        [HttpPost("FetchLoginUsers")]
        public async Task<IActionResult> FetchLoginUsers([FromBody] User? request)
        {
            if (request == null) return StatusCode(HttpStatus.BadRequest);

            // パスワード暗号化
            request.Password = Aes.AesEncrypt(request.Password);

            IEnumerable<IUser> users =
                await _service.SelectLoginUsers(request.UserName, request.Password);

            IList<string> userNames = new List<string>();

            // ユーザー名一覧作成
            foreach (IUser user in users)
                userNames.Add(user.UserName);

            return StatusCode(HttpStatus.OK, userNames);
        }
    }
}
