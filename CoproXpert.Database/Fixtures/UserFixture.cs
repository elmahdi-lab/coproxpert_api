// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Core.Variables;

namespace CoproXpert.Database.Fixtures;

public class UserFixture : FixtureBase
{
    public UserFixture(DataContext dataContext) : base(dataContext)
    {
    }

    protected override void Initialize()
    {
    }

    protected override void Execute()
    {
        var position1 = new GpsPosition(42.3601, -71.0589);
        var position2 = new GpsPosition(42.3601, -71.0589);
    }
}
