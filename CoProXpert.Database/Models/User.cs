// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using CoProXpert.Database.Models.Security;

namespace CoproXpert.Database.Models;

public class User : BaseModel
{
    [Key] public Guid Id { get; set; }

    public ICollection<Social> Socials { get; set; } = null!;
    public ICollection<Permission> Permissions { get; set; } = null!;

    public string Username { get; set; } = null!;

    public string Password { get; set; } = null!;

    public int FailedAttempts { get; set; } = 0;

    public string? ResetToken { get; set; } = null!;

    public DateTime? LockedUntil { get; set; } = null;

    public bool IsLocked => LockedUntil > DateTime.Now;

    public DateTime? ResetTokenExpiration { get; set; } = null;

    public void GenerateResetToken()
    {
        ResetToken = Guid.NewGuid().ToString().Replace("-", "");
        ResetTokenExpiration = DateTime.Now.AddMinutes(30);
    }
}
