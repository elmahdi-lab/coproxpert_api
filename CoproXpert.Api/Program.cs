// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Api.Sources.Helpers;
using CoproXpert.Database;

var builder = WebApplication.CreateBuilder(args);
// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// define appsettings file name with the environment
var appSettingsFileName = $"appsettings.{builder.Environment.EnvironmentName}.json";

// Load appsettings.json
var configuration = new ConfigurationBuilder()
    .SetBasePath(Directory.GetCurrentDirectory())
    .AddJsonFile(appSettingsFileName, false, true)
    .Build();

// // Use Settings class to access appsettings.json
// var settings = new Settings();
// configuration.Bind(settings);
// builder.Services.AddSingleton<Settings>();

// builder.Services.Configure<Settings>(configuration.GetSection("Settings"));

builder.Services.AddDbContext<DataContext>();
// Create a list of services to be injected

ServiceInitializer.Init(builder.Services);

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseRequestLocalization();

app.UseHttpsRedirection();

app.UseAuthentication();
app.UseAuthorization();

app.MapControllers();
app.Run();
