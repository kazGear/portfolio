using Microsoft.Extensions.Hosting;
using Serilog;


namespace CSLib.Logging
{
    public static class LoggerConfig
    {
        public static ILogger Create(string logPath)
        {
            return 
                new LoggerConfiguration().MinimumLevel.Information()
                                         .Enrich.FromLogContext()
                                         .WriteTo.File(path: logPath,
                                                       rollingInterval: RollingInterval.Day,
                                                       retainedFileCountLimit: 30,
                                                       outputTemplate: "{Timestamp:yyyy-MM-dd HH:mm:ss.fff} [{Level:u3}] {Message:lj}{NewLine}{Exception}")
#if DEBUG
                                         .WriteTo.Console()
#endif
                                         .CreateLogger();
        }
    }
}
