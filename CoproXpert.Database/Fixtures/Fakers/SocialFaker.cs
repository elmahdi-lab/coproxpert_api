// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Bogus;
using CoproXpert.Database.Models.Security;
using CoproXpert.Database.Models.Security.Type;

namespace CoproXpert.Database.Fixtures.Fakers;

public sealed class SocialFaker : Faker<Social>
{
    public SocialFaker()
    {
        RuleFor(s => s.Provider, f => f.PickRandom<SocialProvider>());
        RuleFor(s => s.SocialId, f => f.Random.Guid().ToString());
        RuleFor(s => s.Token, f => f.Random.AlphaNumeric(10));
    }
}
