using Microsoft.Extensions.Configuration;
using System.Net.Http.Json;

namespace CSLib.Notify
{
    public class NotifyDiscord : INotify
    {
        private readonly HttpClient _httpClient;
        private readonly string? _webhookUrl;

        public NotifyDiscord(HttpClient httpClient, IConfiguration configuration)
        {
            _httpClient = httpClient;
            _webhookUrl = configuration["Discord:WebhookUrl"];

            if (string.IsNullOrWhiteSpace(_webhookUrl))
            {
                throw new InvalidOperationException("Not set Discord webhookUrl.");
            }
        }

        public async Task NotifyAsync(string message)
        {
            var body = new
            {
                content = message
            };

            HttpResponseMessage? response = await _httpClient.PostAsJsonAsync(_webhookUrl, body);
            response.EnsureSuccessStatusCode();
        }
    }
}
