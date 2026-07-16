using CSLib.Logging;
using CSLib.Middleware;
using CSLib.Notify;
using Serilog;

var builder = WebApplication.CreateBuilder(args);

Log.Logger = LoggerConfig.Create("logs/public-api-.log");
builder.Host.UseSerilog();

// Add services to the container.

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.AddHttpClient<INotify, DiscordNotify>(); // notify setting

builder.Services.AddCors(options =>
{
    options.AddPolicy("PublicApi", policy =>
    {
        policy
            .AllowAnyOrigin()
            .AllowAnyHeader()
            .AllowAnyMethod();
    });
});

var app = builder.Build();

app.UseSerilogRequestLogging();

app.UseMiddleware<ExceptionMiddleware>();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseCors("PublicApi");

app.UseAuthorization();

app.MapControllers();

app.Run();
