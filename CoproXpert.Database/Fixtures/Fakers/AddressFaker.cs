// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Bogus;
using CoproXpert.Database.Models.Information;

namespace CoproXpert.Database.Fixtures.Fakers;

public sealed class AddressFaker : Faker<Address>
{
    public AddressFaker()
    {
        RuleFor(a => a.StreetNumber, f => f.Address.BuildingNumber());
        RuleFor(a => a.StreetAddress, f => f.Address.StreetAddress());
        RuleFor(a => a.Country, f => new CountryFaker().Generate());
        RuleFor(a => a.City, f => new CityFaker().Generate());
        RuleFor(a => a.State, f => f.Address.State());
        RuleFor(a => a.PostalCode, f => f.Address.ZipCode());
    }
}
