// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Repositories;

namespace CoproXpert.Api.Sources.Authentication;

/// <summary>
/// </summary>
public class ApiKeyAuthenticator
{
    /// <summary>
    /// </summary>
    public const string ApiKeyHeaderName = "X-API-KEY";

    private readonly TokenRepository _tokenRepository;

    /// <summary>
    /// </summary>
    /// <param name="tokenRepository"></param>
    public ApiKeyAuthenticator(TokenRepository tokenRepository)
    {
        _tokenRepository = tokenRepository;
    }

    /// <summary>
    /// </summary>
    /// <param name="apiKey"></param>
    /// <returns></returns>
    public bool IsKeyValid(string apiKey)
    {
        var token = _tokenRepository.GetByToken(apiKey);
        return token != null;
    }
}
