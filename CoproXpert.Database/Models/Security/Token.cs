// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Core.Security;
using Microsoft.EntityFrameworkCore;

namespace CoproXpert.Database.Models.Security;

// TODO: Confirm that this is work, we have unique values for token so each user would be able to have only one token
//      but the case where we have a duplicate token is not handled.

[Index(nameof(Value), IsUnique = true)]
public class Token : BaseModel
{
    private const int TokenLength = 256;
    public const int ExpirationTime = 30; // in minutes

    public Token()
    {
        ExpirationDate = DateTime.UtcNow.AddMinutes(ExpirationTime);
    }

    public Guid Id { get; set; }

    public string Value { get; set; } = null!;
    public DateTime ExpirationDate { get; set; }

    public bool IsExpired()
    {
        return ExpirationDate < DateTime.UtcNow;
    }

    public void RefreshToken()
    {
        Value = KeyGenerator.GenerateString(TokenLength);
        ExtendExpirationDate();
    }

    public void ExtendExpirationDate()
    {
        ExpirationDate = DateTime.UtcNow.AddMinutes(ExpirationTime);
    }
}
