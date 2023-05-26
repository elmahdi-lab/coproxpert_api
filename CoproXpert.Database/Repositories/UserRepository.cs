// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database;
using CoproXpert.Database.Models;
using Microsoft.EntityFrameworkCore;

namespace CoProXpert.Database.Repositories;

public class UserRepository : BaseRepository<User>
{
    public UserRepository(DataContext context) : base(context)
    {
    }

    public Task<User?>? GetByUserName(string userName)
    {
        // Find User by UserName
        return Context.Users?.FirstOrDefaultAsync(u => u.Username == userName);
    }
}
