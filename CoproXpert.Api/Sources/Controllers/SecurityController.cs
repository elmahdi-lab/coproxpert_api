// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models;
using CoproXpert.Database.Repositories;
using Microsoft.AspNetCore.Mvc;
using CoproXpert.Core.Security;
namespace CoproXpert.Api.Sources.Controllers;

/// <summary>
/// </summary>
[ApiController]
[Route("[controller]", Name = "SecurityRoute")]
public class SecurityController : ControllerBase
{
    private readonly ILogger<SecurityController> _logger;
    private readonly IPasswordHasher _passwordHasher;
    private readonly UserRepository _userRepository;

    /// <summary>
    /// </summary>
    /// <param name="userRepository"></param>
    /// <param name="passwordHasher"></param>
    /// <param name="logger"></param>
    public SecurityController(UserRepository userRepository, IPasswordHasher passwordHasher,
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
        var user = await _userRepository.GetByUserName(username)!.ConfigureAwait(false)!;
        if (user is null)
        {
            return NotFound();
        }

        if (user.IsLocked)
        {
            _logger.LogWarning("User {UserId} is locked until {LockedUntil}", user.Id, user.LockedUntil);
            return Unauthorized();
        }

        var passwordVerificationResult = _passwordHasher.Verify(user.HashedPassword, password);
        if (passwordVerificationResult != true)
        {
            user.IncrementFailedAttempts();
            return Unauthorized();
        }
        // TODO: refresh does not save to db, so the value returned is not correct.
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

        user.HashedPassword = _passwordHasher.Hash(password);
        var isUpdated = _userRepository.Update(user);
        if (isUpdated)
        {
            user.RefreshPasswordForgetToken();
        }

        return Ok();
    }


    /// <summary>
    ///
    /// </summary>
    /// <param name="password"></param>
    /// <returns></returns>
    [HttpGet("hash-password")]
    public Task<ActionResult> HashPassword(string password)
    {
        return Task.FromResult<ActionResult>(Ok(_passwordHasher.Hash(password)));
    }
}
