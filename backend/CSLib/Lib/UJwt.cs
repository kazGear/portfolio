using Microsoft.Extensions.Configuration;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;


namespace CSLib.Lib
{
    /// <summary>
    /// Json Web Token Util
    /// </summary>
    public class UJwt
    {
        /// <summary>
        /// Json Web Token の発行
        /// </summary>
        public static string GenerateJwtToken(string userName, IConfiguration configuration)
        {
            SymmetricSecurityKey securityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(configuration["Jwt:Key"]!));
            SigningCredentials credentials = new SigningCredentials(securityKey, SecurityAlgorithms.HmacSha256);

            DateTime now = DateTime.UtcNow; // 基準時刻

            var claims = new[] {
                new Claim(JwtRegisteredClaimNames.Sid, userName),
                new Claim(JwtRegisteredClaimNames.Sub, "user authentication for kazApp."),
                new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString())
            };

            JwtSecurityToken jwtPayload = new JwtSecurityToken(
                issuer: configuration["Jwt:Issuer"],
                audience: configuration["Jwt:Audience"],
                claims: claims,
                notBefore: now,
                expires: DateTime.UtcNow.AddDays(Convert.ToDouble(configuration["Jwt:ExpireDays"])),
                signingCredentials: credentials
                );
            return new JwtSecurityTokenHandler().WriteToken(jwtPayload);
        }

        /// <summary>
        /// TODO: 本メソッド削除、AuthActionFilterクラスを削除
        /// 認証確認したいエンドポイントに [Authorize] を付与
        /// サービス起動時のコードに認証設定をしている
        /// 仕組みの概要
        //
        //- AddAuthenticationとAddJwtBearerなどを通じて認証の設定をアプリケーション全体に適用します。
        //- [Authorize] 属性をつけると、そのエンドポイントまたはコントローラーで認証が強制されます。
        //- 有効なJWTが提供されない場合は、自動的にアクセスが拒否されます（例えば、HTTPステータスコード401 Unauthorizedが返されます）。
        //
        /// 
        /// トークンが有効か確認
        /// true: 有効, false: 無効
        /// </summary>
        public static bool IsValidToken(string token)
        {
            JwtSecurityTokenHandler handler = new JwtSecurityTokenHandler();

            JwtSecurityToken? jwtToken;
            try
            {
                // トークンデコード。改ざんされていれば例外発生
                jwtToken = handler.ReadToken(token) as JwtSecurityToken;
                if (jwtToken == null) return false;
            }
            catch (Exception)
            {
                return false;
            }

            // 有効期限
            DateTime limit = jwtToken!.ValidTo;
            DateTime now = DateTime.UtcNow;

            // 期限比較
            if (limit < now) return false;
            return true;
        }
    }
}
