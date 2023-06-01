// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Core.Security;
using CoproXpert.Database.Fixtures.Fakers;

namespace CoproXpert.Database.Fixtures;

public class UserFixture : FixtureBase
{
    private readonly DataContext _dataContext;

    public UserFixture(DataContext dataContext)
    {
        _dataContext = dataContext;
    }

    public override Task Initialize()
    {
        return Task.CompletedTask;
    }

    public override Task Execute()
    {
        _dataContext.Database.EnsureCreated();

        using var transaction = _dataContext.Database.BeginTransaction();
        try
        {
            var user = new UserFaker().Generate();
            _dataContext.Users?.Add(user);
            _dataContext.SaveChanges();
            transaction.Commit();
        }
        catch
        {
            transaction.Rollback();
            throw;
        }

        return Task.CompletedTask;
    }
}
