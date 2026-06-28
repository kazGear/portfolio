using Microsoft.Extensions.Configuration;

namespace CSLib.Lib
{
    public class ConnectionString
    {
        public static string Get(IConfiguration configuration)
        {
            bool onTheDocker = Environment.GetEnvironmentVariable("DOTNET_RUNNING_IN_CONTAINER") == "true";

            if (onTheDocker)
            {
                string dbUser     = Environment.GetEnvironmentVariable("DB_USER")!;
                string dbPassword = Environment.GetEnvironmentVariable("DB_PASSWORD")!;
                string dbName     = Environment.GetEnvironmentVariable("DB_NAME")!;
                string dbHost     = Environment.GetEnvironmentVariable("DB_HOST")!;
                string dbPort     = Environment.GetEnvironmentVariable("DB_PORT")!;

                return $"Server={dbHost};Port={dbPort};User Id={dbUser};Password={dbPassword};Database={dbName}";
            }
            return configuration.GetConnectionString("DefaultConnection")!;
        }
    }
}
