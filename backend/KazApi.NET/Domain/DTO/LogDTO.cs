using System.Text.Json.Serialization;
using PrivateApi.Common._Log;

namespace PrivateApi.Domain.DTO
{
    public class LogDTO
    {
        [JsonPropertyName("Log")]
        public IEnumerable<BattleMetaData> Memory { get; set; }
    }
}
