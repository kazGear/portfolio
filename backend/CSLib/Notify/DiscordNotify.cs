using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Configuration;
using System.Net.Http.Json;
using System.Text;

namespace CSLib.Notify
{
    public class DiscordNotify : INotify
    {
        private readonly HttpClient _httpClient;
        private readonly string _webhookUrl;

        public DiscordNotify(HttpClient httpClient, IConfiguration configuration)
        {
            _httpClient = httpClient;
            _webhookUrl = configuration["Discord:WebhookUrl"]!;
        }

        public async Task NotifyAsync(string message)
        {
            var body = new
            {
                content = message
            };

            await _httpClient.PostAsJsonAsync(_webhookUrl, body);
        }

        public string CreateErrorMessage(HttpContext context, Exception ex)
        {
            StringBuilder sb = new StringBuilder();

            sb.AppendLine("🚨 portfolio API Error\n");

            sb.AppendLine($"Time: {DateTime.Now:yyyy-MM-dd HH:mm:ss}\n");
            sb.AppendLine($"Path: {context.Request.Method} {context.Request.Path}\n");
            sb.AppendLine($"Exception: {ex.GetType().Name}\n");
            sb.AppendLine($"Message: {ex.Message}\n");

            var stack = ex.StackTrace?.Split(Environment.NewLine)
                                      .Take(3); // 短いスタックトレース

            sb.AppendLine($"StackTrace: {string.Join(Environment.NewLine, stack ?? [])}\n");

            return sb.ToString();
        }
    }
}
