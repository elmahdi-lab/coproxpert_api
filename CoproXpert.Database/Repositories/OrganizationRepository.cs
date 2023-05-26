// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database;
using CoproXpert.Database.Models;

namespace CoProXpert.Database.Repositories;

public class OrganizationRepository : BaseRepository<Organization>
{
    public OrganizationRepository(DataContext context) : base(context)
    {
    }
}
