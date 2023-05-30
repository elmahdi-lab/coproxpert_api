// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Database.Fixtures;

public class FixtureLoader
{
    public FixtureLoader(DataContext dataContext)
    {
        Fixtures = new List<IFixture>();
        Fixtures.Add(new UserFixture(dataContext));
    }

    private ICollection<IFixture> Fixtures { get; }

    public void ExecuteAllFixtures()
    {
        foreach (var fixture in Fixtures)
        {
            fixture.Initialize();
            fixture.Execute();
        }
    }
}
