// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Bogus;
using CoproXpert.Database.Models;
using CoproXpert.Database.Models.Security;

namespace CoproXpert.Database.Fixtures.Fakers;

public sealed class UserFaker : Faker<User>
{
    public UserFaker()
    {
        RuleFor(u => u.Socials, f => new List<Social>());
        RuleFor(u => u.Claims, f => new List<Claim>());
        RuleFor(u => u.Contact, f => new ContactFaker().Generate());
        RuleFor(u => u.Token, f => new Token());
        RuleFor(u => u.Username, f => f.Person.UserName);
        RuleFor(u => u.HashedPassword, f => "Password");
        RuleFor(u => u.FailedAttempts, f => f.Random.Int(0, 5));
        RuleFor(u => u.PasswordForgetToken, f => f.Random.Word());
        RuleFor(u => u.LockedUntil, f => f.Random.Bool() ? f.Date.Past() : null);
        RuleFor(u => u.ResetTokenExpiration, f => f.Random.Bool() ? f.Date.Future() : null);

        FinishWith((f, u) =>
        {
            u.Id = Guid.NewGuid();
            u.Token = new Token();
            u.RefreshPasswordForgetToken();
        });
    }
}
