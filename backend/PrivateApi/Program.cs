using CSLib.Logging;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.IdentityModel.Tokens;
using PrivateApi.Common._Filter;
using Serilog;

using System.Text;


public class Program
{
    public static void Main(string[] args)
    {
        Log.Logger = LoggerConfig.Create("logs/private-api-.log");
        Console.WriteLine("KazApp is starting...");
        CreateHostBuilder(args).Build().Run();
    }

    public static IHostBuilder CreateHostBuilder(string[] args) =>
        Host.CreateDefaultBuilder(args)
            .UseSerilog()
            .ConfigureWebHostDefaults(webBuilder =>
            {
                webBuilder.UseStartup<Startup>();
                webBuilder.ConfigureKestrel((context, options) => {

                });
            });
}

public class Startup
{
    public Startup(IConfiguration configuration)
    {
        Configuration = configuration;
    }

    public IConfiguration Configuration { get; }

    public void ConfigureServices(IServiceCollection services)
    {
        // Swagger
        services.AddSwaggerGen(c =>
        {
            c.SwaggerDoc("v1", new Microsoft.OpenApi.Models.OpenApiInfo
            {
                Title = "PrivateApi",
                Version = "v1"
            });
        });

        services.AddControllers(options =>
        {
            options.Filters.Add<ExceptionFilter>();
            //options.Filters.Add<AuthActionFilter>();
        })
        .AddJsonOptions(options =>
        {
            options.JsonSerializerOptions.Encoder =
                System.Text.Encodings.Web.JavaScriptEncoder.Create(
                    System.Text.Unicode.UnicodeRanges.All);
        });

        // JWT設定
        services.AddAuthentication(options =>
        {
            options.DefaultAuthenticateScheme = JwtBearerDefaults.AuthenticationScheme;
            options.DefaultChallengeScheme = JwtBearerDefaults.AuthenticationScheme;
        }).AddJwtBearer(options =>
        {
            options.TokenValidationParameters = new TokenValidationParameters
            {
                ValidateIssuer = true,
                ValidateAudience = true,
                ValidateLifetime = true,
                ValidateIssuerSigningKey = true,
                ValidIssuer = Configuration["Jwt:Issuer"],
                ValidAudience = Configuration["Jwt:Audience"],
                IssuerSigningKey = new SymmetricSecurityKey(
                    Encoding.UTF8.GetBytes(Configuration["Jwt:Key"]!))
            };
        });

        // JS: cors設定
        services.AddCors(options =>
        {
            options.AddPolicy("ApiCors",
                builder =>
                    builder.AllowAnyHeader()
                           .AllowAnyMethod()
                           .WithOrigins(
                                "http://localhost:5173",   // npm run dev
                                "http://localhost",        // Docker + nginx (HTTP)
                                "https://localhost",       // Docker + nginx (HTTPS)
                                "https://kazapp-trial.com" // 本番
                            )
                           .AllowCredentials());
        });

        // DTO > JSONバインド
        services.AddControllers().AddJsonOptions(options =>
        {
            // C# のプロパティ名そのまま（PascalCase）で返す。プロパティ名が完全に合致する必要あり。
            options.JsonSerializerOptions.PropertyNamingPolicy = null;
        });
    }

    public void Configure(IApplicationBuilder app, IWebHostEnvironment env, ILogger<Startup> logger)
    {
        app.UseSerilogRequestLogging();

        if (env.IsDevelopment())
        {
            app.UseDeveloperExceptionPage();
            app.UseSwagger();
            app.UseSwaggerUI();
        }
        else
        {
            //app.UseExceptionHandler("/Home/Error");
            app.UseHsts();
            //app.UseHttpsRedirection(); // Docker + nginx構成では使用しない。HTTPSはnginxで終端するため不要
        }

        app.UseRouting();
        app.UseCors("ApiCors"); // app.UseStaticFiles()より先に実行する必要がある

        app.UseStaticFiles();
        app.UseAuthentication();
        app.UseAuthorization();

        app.UseEndpoints(endpoints =>
        {
            // ブラウザのOPTIONS を必ず返す
            endpoints.MapMethods("{*path}", new[] { "OPTIONS" }, () => Results.Ok());
            // swagger
            endpoints.MapControllers();
            // mvc
            endpoints.MapControllerRoute(
                name : "default",
                pattern: "{controller=Home}/{action=Index}/{id?}");
            endpoints.MapGet("/health", () => Results.Ok("Health check: OK"));
            //endpoints.MapFallbackToFile("/index.html");
        });
        app.Use(async (context, next) =>
        {
            context.Response.Headers.Add("Content-Type", "application/json; charset=utf-8");
            await next();
        });
    }
}