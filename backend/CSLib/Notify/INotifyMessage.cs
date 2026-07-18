using Microsoft.AspNetCore.Http;

namespace CSLib.Notify
{
    public interface INotifyMessage
    {
        public string CreateMessage(HttpContext context, Exception ex);
    }
}
