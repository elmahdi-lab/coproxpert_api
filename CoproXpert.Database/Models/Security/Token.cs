// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Core.Security;
using Microsoft.EntityFrameworkCore;

namespace CoproXpert.Database.Models.Security;

// TODO: find a way to map user to token to add unique constraint user + value.

[Index(nameof(Value), IsUnique = true)]
public class Token : BaseModel
{
    private const int TokenLength = 256;
    public const int ExpirationTime = 30; // in minutes
    private const int MaxDuplicateValueAttempts = 3;
    private int _duplicateValueAttempts;

    [Obsolete("Has an obsolete attribute because it is only used for testing purposes.")]
    public Token()
    {
        CreateUniqueValue();
        ExpirationDate = DateTime.UtcNow.AddMinutes(ExpirationTime);
    }

    public Guid Id { get; set; }

    public string Value { get; set; } = null!;
    public DateTime ExpirationDate { get; set; }

    public bool IsExpired()
    {
        return ExpirationDate < DateTime.UtcNow;
    }

    [Obsolete("This method is only used for testing purposes.")]
    public void CreateUniqueValue()
    {
        while (_duplicateValueAttempts < MaxDuplicateValueAttempts)
        {
            Value = KeyGenerator.GenerateString(TokenLength);

            try
            {
                using var context = new DataContext();
                context.Tokens?.Add(this);
                context.SaveChanges();
                break;
            }
            catch (DbUpdateException)
            {
                _duplicateValueAttempts++;
                CreateUniqueValue();
            }
        }
    }

    public void RefreshToken()
    {
        CreateUniqueValue();
        ExpirationDate = DateTime.UtcNow.AddMinutes(ExpirationTime);
    }

    public void ExtendExpirationDate()
    {
        ExpirationDate = DateTime.UtcNow.AddMinutes(ExpirationTime);
    }
}
