// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Bogus;
using CoproXpert.Core.Security;
using CoproXpert.Database.Models;
using CoproXpert.Database.Models.Security;

namespace CoproXpert.Database.Fixtures.Fakers;

public sealed class UserFaker : Faker<User>
{
    // password
    private const string PasswordHash = "dv020Q6dJqWj9eoZO91TXg==;imN+jV4CBbqc9mB1pJ6v8M/wXirXLB5I13todQpmDSk=";

    public UserFaker()
    {

        RuleFor(u => u.Socials, f => new List<Social>());
        RuleFor(u => u.Claims, f => new List<Claim>());
        RuleFor(u => u.Contact, f => new ContactFaker().Generate());
        RuleFor(u => u.Token, f => new TokenFaker().Generate());
        RuleFor(u => u.Username, f => f.Person.Email);
        RuleFor(u => u.HashedPassword, f => PasswordHash);
        RuleFor(u => u.FailedAttempts, f => 0);
        RuleFor(u => u.PasswordForgetToken, f => f.Random.Hash());
        RuleFor(u => u.LockedUntil, f => null);
        RuleFor(u => u.ResetTokenExpiration, f => DateTime.UtcNow.AddMinutes(Token.ExpirationTime));
        FinishWith((f, u) =>
        {
            u.Id = Guid.NewGuid();
            u.RefreshPasswordForgetToken();
        });
    }
}
