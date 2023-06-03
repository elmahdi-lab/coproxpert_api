// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Api.Sources.Models;
using CoproXpert.Core.Security;
using CoproXpert.Database.Repositories;
using Microsoft.AspNetCore.Mvc;

namespace CoproXpert.Api.Sources.Controllers;

/// <summary>
///     Controller for managing security access.
/// </summary>
[Route("[controller]", Name = "SecurityRoute")]
public class SecurityController : ControllerBase
{
    private readonly ILogger<SecurityController> _logger;
    private readonly PasswordHasher _passwordHasher;
    private readonly UserRepository _userRepository;

    /// <summary>
    ///     Initializes a new instance of the <see cref="SecurityController" /> class.
    /// </summary>
    /// <param name="userRepository"></param>
    /// <param name="passwordHasher"></param>
    /// <param name="logger"></param>
    public SecurityController(UserRepository userRepository, PasswordHasher passwordHasher,
        ILogger<SecurityController> logger)
    {
        _userRepository = userRepository;
        _passwordHasher = passwordHasher;
        _logger = logger;
    }

    /// <summary>
    ///     Generate a token for a user.
    /// </summary>
    /// <param name="login"></param>
    /// <returns></returns>
    [HttpPost("login", Name = "LoginRoute")]
    public async Task<ActionResult<string>> Login([FromBody] LoginModel login)
    {
        var user = await _userRepository.GetByUserName(login.Username)!.ConfigureAwait(false);
        if (user is null)
        {
            return NotFound();
        }

        if (user.IsLocked)
        {
            _logger.LogWarning("User {UserId} is locked until {LockedUntil}", user.Id, user.LockedUntil);
            return Unauthorized();
        }

        var passwordVerificationResult = _passwordHasher.Verify(user.HashedPassword, login.Password);
        if (passwordVerificationResult != true)
        {
            user.IncrementFailedAttempts();
            return Unauthorized();
        }

        user.Token.RefreshToken();
        _userRepository.Update(user);

        return Ok(user.Token.Value);
    }

    /// <summary>
    ///     Reset the password of a user.
    /// </summary>
    /// <param name="token"></param>
    /// <param name="password"></param>
    /// <returns></returns>
    [HttpPost("password-reset")]
    public async Task<ActionResult> PasswordReset(string token, string password)
    {
        var user = await _userRepository.GetByForgotPasswordToken(token)!.ConfigureAwait(false);
        if (user is null)
        {
            return NotFound();
        }

        if (user.Token.IsExpired())
        {
            _logger.LogWarning("User {UserId} forgot password token is expired", user.Id);
            return Unauthorized();
        }

        user.HashedPassword = _passwordHasher.Hash(password);
        var isUpdated = _userRepository.Update(user);
        if (isUpdated)
        {
            user.RefreshPasswordForgetToken();
        }

        return Ok();
    }
}
