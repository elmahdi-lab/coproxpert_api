// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Api.Sources.Authentication;
using CoproXpert.Api.Sources.Helpers;
using CoproXpert.Database;
using CoproXpert.Database.Fixtures;
using Microsoft.OpenApi.Models;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddSingleton<ServiceInitializer>();
// Create a list of services to be injected
var serviceInitializer = new ServiceInitializer();
serviceInitializer.Init(builder.Services);

builder.Services.AddAuthentication();
builder.Services.AddAuthorization();


builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddDbContext<DataContext>();

if (builder.Environment.IsDevelopment())
{
    builder.Services.AddSwaggerGen(c =>
    {
        // Add the security definition for the custom API key
        c.AddSecurityDefinition("ApiKey",
            new OpenApiSecurityScheme
            {
                Description = "API Key",
                Name = ApiKeyAuthenticator.ApiKeyHeaderName,
                In = ParameterLocation.Header,
                Type = SecuritySchemeType.ApiKey
            });

        c.OperationFilter<SecurityRequirementsOperationFilter>();
    });
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
