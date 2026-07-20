using Microsoft.AspNetCore.Http;
using System.Text;

namespace CSLib.Notify
{
    public class MessageApiError : INotifyMessage
    {
        public string CreateMessage(HttpContext context, Exception ex)
        {
            StringBuilder sb = new StringBuilder();

            sb.AppendLine("🚨 portfolio API Error\n");

            sb.AppendLine($"Time: {DateTime.Now:yyyy-MM-dd HH:mm}\n");
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
