// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Bogus;
using CoproXpert.Database.Models.Information;

namespace CoproXpert.Database.Fixtures.Fakers;

public sealed class CityFaker : Faker<City>
{
    public CityFaker()
    {
        RuleFor(c => c.Name, f => f.Address.City());
        RuleFor(c => c.ZipCode, f => f.Address.ZipCode());
        RuleFor(c => c.Country, _ => new CountryFaker().Generate());
    }
}
