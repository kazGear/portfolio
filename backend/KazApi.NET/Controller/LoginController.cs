using CSLib.Lib;
using KazApi.Domain._User;
using KazApi.Service;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;

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
        public ActionResult<string> FetchLoginUsers([FromBody] User? request)
        {
            // パスワード暗号化
            request.Password = UAes.AesEncrypt(request.Password);

            IEnumerable<IUser> users = _service.SelectLoginUsers(request.UserName, request.Password);

            IList<string> userNames = new List<string>();

            // ユーザー名一覧作成
            foreach (IUser user in users) 
                userNames.Add(user.UserName);

            return JsonConvert.SerializeObject(userNames);
        }
    }
}
