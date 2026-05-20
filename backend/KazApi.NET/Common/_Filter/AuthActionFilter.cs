using CSLib.Lib;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.Extensions.Primitives;

namespace KazApi.Common._Filter
{
    /// <summary>
    /// 認証フィルター
    /// </summary>
    public class AuthActionFilter : IActionFilter
    {
        /// <summary>
        /// 実行前
        /// </summary>
        public void OnActionExecuted(ActionExecutedContext context)
        {
            // SkipAuthActionFilter付与 > 処理不要
            if (context.ActionDescriptor.EndpointMetadata.OfType<SkipAuthActionFilter>().Any())
            {
                return;
            }
      
            // 送られてきたトークンを取得
            IQueryCollection query = context.HttpContext.Request.Query;
            bool exist = query.TryGetValue("token", out StringValues token);

            if (exist)
            {
                bool isValid = UJwt.IsValidToken(token!);

                if (!isValid)
                {
                    //throw new SecurityTokenException("トークンの有効期限が切れています。ログインしてください。");
                    context.Result = new StatusCodeResult(500);
                }
            }
        }

        /// <summary>
        /// 実行後
        /// </summary>
        public void OnActionExecuting(ActionExecutingContext context)
        {
            // 処理なし
        }
    }
}
