// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Bogus;
using CoproXpert.Database.Models.Information;

namespace CoproXpert.Database.Fixtures.Fakers;

public sealed class ContactFaker : Faker<Contact>
{
    public ContactFaker()
    {
        RuleFor(c => c.Name, f => f.Person.FullName);
        RuleFor(c => c.Email, f => f.Person.Email);
        RuleFor(c => c.Phone, f => f.Person.Phone);
        RuleFor(c => c.Location, f => new GpsPositionFaker().Generate());
        RuleFor(c => c.Notes, f => f.Lorem.Sentence());
        RuleFor(c => c.Address, f => new AddressFaker().Generate());
    }
}
