// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Core.Attributes;
using CoproXpert.Core.Enums;
using CoproXpert.Database.Models.Building;

namespace CoproXpert.Database.Repositories;

[Autowire(Lifetime.Scoped)]
public class CommunityRepository : BaseRepository<Community>
{
    public CommunityRepository(DataContext context) : base(context)
    {
    }
}
