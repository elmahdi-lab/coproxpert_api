// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models;
using CoproXpert.Database.Repositories;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;

namespace CoproXpert.Api.Sources.Controllers;

/// <summary>
/// </summary>
[ApiController]
[Route("[controller]", Name = "SecurityRoute")]
public class SecurityController : ControllerBase
{
    private readonly ILogger<SecurityController> _logger;
    private readonly IPasswordHasher<User> _passwordHasher;
    private readonly UserRepository _userRepository;

    /// <summary>
    /// </summary>
    /// <param name="userRepository"></param>
    /// <param name="passwordHasher"></param>
    /// <param name="logger"></param>
    public SecurityController(UserRepository userRepository, IPasswordHasher<User> passwordHasher,
        ILogger<SecurityController> logger)
    {
        _userRepository = userRepository;
        _passwordHasher = passwordHasher;
        _logger = logger;
    }

    /// <summary>
    /// </summary>
    /// <param name="username"></param>
    /// <param name="password"></param>
    /// <returns></returns>
    [HttpPost]
    public async Task<ActionResult> Login(string username, string password)
    {
        var user = await _userRepository.GetByUserName(username).ConfigureAwait(false)!;
        if (user is null)
        {
            return NotFound();
        }

        if (user.IsLocked)
        {
            _logger.LogWarning("User {UserId} is locked until {LockedUntil}", user.Id, user.LockedUntil);
            return Unauthorized();
        }

        var passwordVerificationResult = _passwordHasher.VerifyHashedPassword(user, user.HashedPassword, password);
        if (passwordVerificationResult != PasswordVerificationResult.Success)
        {
            user.IncrementFailedAttempts();
            return Unauthorized();
        }

        user.Token.RefreshToken();
        return Ok(user.Token.Value);
    }

    /// <summary>
    /// </summary>
    /// <param name="token"></param>
    /// <param name="password"></param>
    /// <returns></returns>
    [HttpPost("password-reset/{token}")]
    public async Task<ActionResult> PasswordReset(string token, string password)
    {
        var user = await _userRepository.GetByForgotPasswordToken(token).ConfigureAwait(false)!;
        if (user is null)
        {
            return NotFound();
        }

        if (user.Token.IsExpired())
        {
            _logger.LogWarning("User {UserId} forgot password token is expired", user.Id);
            return Unauthorized();
        }

        user.HashedPassword = _passwordHasher.HashPassword(user, password);
        var isUpdated = _userRepository.Update(user);
        if (isUpdated)
        {
            user.RefreshPasswordForgetToken();
        }

        return Ok();
    }
}
