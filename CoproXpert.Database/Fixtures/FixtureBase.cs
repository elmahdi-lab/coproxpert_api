// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Database.Fixtures;

public abstract class FixtureBase : IFixture
{
    private readonly DataContext _dataContext;

    protected FixtureBase(DataContext dataContext)
    {
        _dataContext = dataContext;
    }

    protected new abstract void Initialize();

    protected new abstract void Execute();
}
