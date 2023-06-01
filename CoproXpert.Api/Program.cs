// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Api.Sources.Authentication;
using CoproXpert.Api.Sources.Helpers;
using CoproXpert.Database;
using CoproXpert.Database.Fixtures;

var builder = WebApplication.CreateBuilder(args);
// Create a list of services to be injected
ServiceInitializer.Init(builder.Services);

builder.Services.AddAuthentication();
builder.Services.AddAuthorization();

builder.Services.AddScoped<ApiKeyAuthenticator>();
builder.Services.AddScoped<ApiKeyAuthFilter>();

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

builder.Services.AddDbContext<DataContext>();
if (builder.Environment.IsDevelopment())
{
    builder.Services.AddTransient<UserFixture>();
    builder.Services.AddScoped<FixtureLoader>();
}

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();

    using var scope = app.Services.CreateScope();
    var fixtureLoader = scope.ServiceProvider.GetRequiredService<FixtureLoader>();
    fixtureLoader.ExecuteAllFixtures();
}

app.UseRequestLocalization();

app.UseHttpsRedirection();

app.UseAuthentication();
app.UseAuthorization();

app.MapControllers();
app.Run();
