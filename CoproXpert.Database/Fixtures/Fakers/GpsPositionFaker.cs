// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Bogus;
using CoproXpert.Core.Variables;

namespace CoproXpert.Database.Fixtures.Fakers;

public sealed class GpsPositionFaker : Faker<GpsPosition>
{
    public GpsPositionFaker()
    {
        RuleFor(p => p.Latitude, f => f.Random.Double(-90, 90));
        RuleFor(p => p.Longitude, f => f.Random.Double(-180, 180));
    }
}
