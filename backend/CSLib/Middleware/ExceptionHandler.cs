using CSLib.Notify;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Logging;

namespace CSLib.Middleware
{
    public class ExceptionMiddleware
    {
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
                Console.WriteLine($"middleware exception handler: {ex}");
                _logger.LogError(ex, "Unhandled Exception");

                context.Response.StatusCode  = 500;
                context.Response.ContentType = "application/json; charset=utf-8";

                await context.Response.WriteAsJsonAsync(new
                {
                    message = $"Internal Server Error.\n{ex}"
                });

                INotifyMessage message = new MessageApiError();
                await _notify.NotifyAsync(message.CreateMessage(context, ex));

                throw;
            }
        }
    }
}
