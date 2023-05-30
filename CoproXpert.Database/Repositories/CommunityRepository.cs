// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models.Building;

namespace CoproXpert.Database.Repositories;

public class CommunityRepository : BaseRepository<Community>
{
    public CommunityRepository(DataContext context) : base(context)
    {
    }
}
