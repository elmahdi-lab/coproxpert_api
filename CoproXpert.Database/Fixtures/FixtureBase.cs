// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Database.Fixtures;

public abstract class FixtureBase : IFixture
{
    public virtual Task Initialize()
    {
        throw new NotImplementedException();
    }

    public virtual Task<Task> Execute()
    {
        throw new NotImplementedException();
    }
}
