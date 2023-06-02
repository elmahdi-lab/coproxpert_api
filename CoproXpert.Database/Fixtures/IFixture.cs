// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Database.Fixtures;

internal interface IFixture
{
    Task Initialize();

    Task<Task> Execute();
}
