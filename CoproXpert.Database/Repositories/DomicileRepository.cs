// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database;
using CoproXpert.Database.Models.Building;

namespace CoproXpert.Database.Repositories;

public class DomicileRepository : BaseRepository<Domicile>
{
    public DomicileRepository(DataContext context) : base(context)
    {
    }
}
