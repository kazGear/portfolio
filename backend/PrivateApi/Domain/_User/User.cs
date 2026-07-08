using System.Text.Json.Serialization;

namespace PrivateApi.Domain._User
{
    public class User
    {
        [JsonPropertyName("userName")]
        public string UserName { get; set; }
        [JsonPropertyName("password")]
        public string Password { get; set; }
    }
}
