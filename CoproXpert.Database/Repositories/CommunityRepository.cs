// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database;
using CoProXpert.Database.Models.Building;

namespace CoProXpert.Database.Repositories;

public class CommunityRepository : BaseRepository<Community>
{
    public CommunityRepository(DataContext context) : base(context)
    {
    }
}
