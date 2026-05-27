using CSLib.Lib;
using KazApi.Common;
using KazApi.Domain._User;
using KazApi.Service;
using Microsoft.AspNetCore.Mvc;

namespace KazApi.Controller
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
        public IActionResult FetchLoginUsers([FromBody] User? request)
        {
            if (request == null) return StatusCode(HttpStatus.BadRequest);

            try
            {
                // パスワード暗号化
                request.Password = Aes.AesEncrypt(request.Password);

                IEnumerable<IUser> users = _service.SelectLoginUsers(request.UserName, request.Password);

                IList<string> userNames = new List<string>();

                // ユーザー名一覧作成
                foreach (IUser user in users)
                    userNames.Add(user.UserName);

                return StatusCode(HttpStatus.OK, userNames);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }
    }
}
