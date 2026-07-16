using Microsoft.AspNetCore.Http;

namespace CSLib.Notify
{
    public interface INotify
    {
        public Task NotifyAsync(string message);

        public string CreateErrorMessage(HttpContext context, Exception ex);
    }
}
