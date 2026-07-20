using CSLib.Lib;
using CSLib.Notify;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Logging;
using System.Collections.Concurrent;

namespace CSLib.Middleware
{
    public class ExceptionMiddleware
    {
        private static readonly IList<string> _notifiedKeys = new List<string>();
        private readonly RequestDelegate _next;
        private readonly ILogger<ExceptionMiddleware> _logger;
        private readonly INotify _notify;

        public ExceptionMiddleware(RequestDelegate next,
                                   ILogger<ExceptionMiddleware> logger,
                                   INotify notify)
        {
            _next   = next;
            _logger = logger;
            _notify = notify;
        }

        public async Task Invoke(HttpContext context)
        {
            try
            {
                await _next(context);
            }
            catch (Exception ex)
            {
                // error 通知
                try
                {
                    INotifyMessage notifyMessage = new MessageApiError();
                    string message               = notifyMessage.CreateMessage(context, ex);
                    string notifyKey =
                        $"{DateTime.Now:yyyy-MM-dd HH}{context.Request.Path}{ex.GetType().Name}{ex.Message}";

                    // react strictMode > api連続呼び出し > 重複通知を防ぐ
                    if (!_notifiedKeys.Contains(notifyKey))
                    {
                        await _notify.NotifyAsync(message);
                        _notifiedKeys.Add(notifyKey);
                    }
                }
                catch (Exception notifyEx)
                {
                    _logger.LogError(notifyEx, "Send notify error.");
                }

                _logger.LogError(ex, "Unhandled Exception from middleware.");

                context.Response.StatusCode  = 500;
                context.Response.ContentType = "application/json; charset=utf-8";

                await context.Response.WriteAsJsonAsync(new
                {
                    message = $"Internal Server Error.\n"
                });
            }
        }
    }
}
