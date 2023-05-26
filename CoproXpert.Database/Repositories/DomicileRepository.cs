// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database;
using CoProXpert.Database.Models.Building;

namespace CoProXpert.Database.Repositories;

public class DomicileRepository : BaseRepository<Domicile>
{
    public DomicileRepository(DataContext context) : base(context)
    {
    }
}
