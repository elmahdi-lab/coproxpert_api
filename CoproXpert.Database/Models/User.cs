// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using CoproXpert.Core.Security;
using CoproXpert.Database.Models.Information;
using CoproXpert.Database.Models.Security;
using Microsoft.EntityFrameworkCore;

namespace CoproXpert.Database.Models;

[Index(nameof(Username), IsUnique = true)]
public class User : BaseModel
{
    private const int MaxFailedAttempts = 5;
    private const int LockTime = 5; // in minutes
    private const int TokenLength = 100;

    public User()
    {
        Token = new Token();
    }

    [Key] public Guid Id { get; set; }

    public ICollection<Social> Socials { get; } = new List<Social>();
    public ICollection<Claim> Claims { get; } = new List<Claim>();

    public Contact Contact { get; set; } = null!;
    public Token Token { get; set; }

    public string Username { get; set; } = null!;

    [DataType(DataType.Password)] public string HashedPassword { get; set; } = null!;

    public int FailedAttempts { get; set; }

    public string? PasswordForgetToken { get; set; }

    public DateTime? LockedUntil { get; set; }

    public bool IsLocked => LockedUntil > DateTime.UtcNow;
    public bool IsFailedAttemptsExceeded => FailedAttempts >= MaxFailedAttempts;
    public DateTime? ResetTokenExpiration { get; set; }

    public void RefreshPasswordForgetToken()
    {
        PasswordForgetToken = KeyGenerator.GenerateString(TokenLength);
        ResetTokenExpiration = DateTime.UtcNow.AddMinutes(LockTime);
    }

    public void IncrementFailedAttempts()
    {
        FailedAttempts++;
        if (IsFailedAttemptsExceeded)
        {
            LockedUntil = DateTime.UtcNow.AddMinutes(LockTime);
        }
    }
}
