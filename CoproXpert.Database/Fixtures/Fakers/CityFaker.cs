// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Bogus;
using CoproXpert.Database.Models.Information;

namespace CoproXpert.Database.Fixtures.Fakers;

public sealed class CityFaker : Faker<City>
{
    public CityFaker()
    {
        RuleFor(c => c.Name, f => f.Address.City());
        RuleFor(c => c.Code, f => f.Random.AlphaNumeric(3).ToUpper());
        RuleFor(c => c.Country, f => new CountryFaker().Generate());
    }
}
