using KazApi.Common._Filter;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.IdentityModel.Tokens;
using System.Net;
using System.Text;

public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("KazApp is starting...");
        CreateHostBuilder(args).Build().Run();
    }

    public static IHostBuilder CreateHostBuilder(string[] args) =>
        Host.CreateDefaultBuilder(args)
            .ConfigureWebHostDefaults(webBuilder =>
            {
                webBuilder.UseStartup<Startup>();
                webBuilder.ConfigureKestrel((context, options) => {
                    //options.Listen(IPAddress.Any, 5000); // HTTP
                    if (context.HostingEnvironment.IsProduction())
                    {
                        options.Listen(IPAddress.Any, 5001, listenOptions =>
                        {
                            listenOptions.UseHttps(
                                "/etc/letsencrypt/live/kazapp-trial.com/kazapp-trial.pfx",
                                "kaz_5050");
                        });
                    }
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
                Title = "KazApi",
                Version = "v1"
            });
        });

        services.AddControllers(options =>
        {
            options.Filters.Add<ExceptionFilter>();
            options.Filters.Add<AuthActionFilter>();
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
    }

    public void Configure(IApplicationBuilder app, IWebHostEnvironment env, ILogger<Startup> logger)
    {
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
            //app.UseHttpsRedirection();
        }

        app.UseRouting();
        app.UseCors("ApiCors"); // app.UseStaticFiles()より先に実行する必要がある

        app.UseStaticFiles();
        app.UseAuthentication();
        app.UseAuthorization();

        app.UseEndpoints(endpoints =>
        {
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