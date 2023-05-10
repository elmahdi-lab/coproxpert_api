// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models;
using Microsoft.EntityFrameworkCore;

namespace CoproXpert.Database.Services;

public class UserService : BaseService<User>
{
    public UserService(DataContext context) : base(context)
    {
    }

    public Task<User?>? GetByUserName(string userName)
    {
        // Find User by UserName
        return Context.Users?.FirstOrDefaultAsync(u => u.Username == userName);
    }
}
