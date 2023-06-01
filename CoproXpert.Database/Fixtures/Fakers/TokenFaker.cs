// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Bogus;
using CoproXpert.Database.Models.Security;

namespace CoproXpert.Database.Fixtures.Fakers;

public sealed class TokenFaker : Faker<Token>
{
    public TokenFaker()
    {
        RuleFor(t => t.Id, f => f.Random.Guid());
        RuleFor(t => t.Value, f => f.Random.AlphaNumeric(256));
        RuleFor(t => t.ExpirationDate, f => DateTime.UtcNow.AddMinutes(Token.ExpirationTime));
    }
}
