// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Bogus;
using CoproXpert.Database.Models.Information;

namespace CoproXpert.Database.Fixtures.Fakers;

public sealed class CountryFaker : Faker<Country>
{
    public CountryFaker()
    {
        RuleFor(c => c.Name, f => f.Address.Country());
        RuleFor(c => c.Code, f => f.Address.CountryCode());
    }
}
