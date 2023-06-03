// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Core.Attributes;
using CoproXpert.Core.Enums;
using CoproXpert.Database.Models.Security;
using Microsoft.EntityFrameworkCore;

namespace CoproXpert.Database.Repositories;

[Autowire(Lifetime.Scoped)]
public class TokenRepository : BaseRepository<Token>
{
    public TokenRepository(DataContext context) : base(context)
    {
    }

    public Task<Token?>? GetByToken(string token)
    {
        // Find Token by TokenName
        return Context.Tokens?.FirstOrDefaultAsync(t => t.Value == token);
    }
    //
    // public Task<Token?>? GetByUser(User user)
    // {
    //     //return Context.Tokens?.FirstOrDefaultAsync(t => t.User == user);
    // }
}
