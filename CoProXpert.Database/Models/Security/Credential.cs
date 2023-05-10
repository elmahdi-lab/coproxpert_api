// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using CoproXpert.Database.Models;

namespace CoProXpert.Database.Models.Security;

public class Credential : BaseModel
{
    [Key] public Guid Id { get; set; }
    public string Password { get; set; } = null!;

    public int FailedAttempts { get; set; } = 0;

    public DateTime LockedUntil { get; set; }

    public bool IsLocked => LockedUntil > DateTime.Now;
    public string ResetToken { get; set; } = null!;


    [DataType(DataType.DateTime)] public DateTime ResetTokenExpiration { get; set; }


    public User User { get; set; } = null!;

    public void GenerateResetToken()
    {
        ResetToken = Guid.NewGuid().ToString().Replace("-", "");
        // TODO: minutes value should be a config, and if not configured then we can set a 5 min const.
        ResetTokenExpiration = DateTime.Now.AddMinutes(30);
    }
}
