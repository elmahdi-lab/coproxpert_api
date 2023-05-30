// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;

namespace CoproXpert.Api.Sources.Authentication;

/// <summary>
/// </summary>
public sealed class ApiKeyAuthFilter : Attribute, IAsyncAuthorizationFilter
{
    private readonly ApiKeyAuthenticator _apiKeyAuthenticator;

    /// <summary>
    /// </summary>
    /// <param name="apiKeyAuthenticator"></param>
    public ApiKeyAuthFilter(ApiKeyAuthenticator apiKeyAuthenticator)
    {
        _apiKeyAuthenticator = apiKeyAuthenticator;
    }

    /// <summary>
    /// </summary>
    /// <param name="context"></param>
    /// <returns></returns>
    public Task OnAuthorizationAsync(AuthorizationFilterContext context)
    {
        // If configuration is required you can fetch it by:
        // var configuration = context.HttpContext.RequestServices.GetRequiredService<IConfiguration>();

        if (!context.HttpContext.Request.Headers.TryGetValue(ApiKeyAuthenticator.ApiKeyHeaderName,
                out var extractedApiKey))
        {
            context.Result = new UnauthorizedObjectResult("Api Key was not provided. (Missing header)");
            return Task.CompletedTask;
        }

        // check the api key is valid
        if (!_apiKeyAuthenticator.IsKeyValid(extractedApiKey!))
        {
            context.Result = new UnauthorizedObjectResult("Unauthorized client.");
        }

        return Task.CompletedTask;
    }

    public ApiKeyAuthenticator ApiKeyAuthenticator { get; }
}
