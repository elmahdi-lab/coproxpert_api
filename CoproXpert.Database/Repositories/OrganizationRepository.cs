// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Core.Attributes;
using CoproXpert.Core.Enums;
using CoproXpert.Database.Models;

namespace CoproXpert.Database.Repositories;

[Autowire(Lifetime.Scoped)]
public class OrganizationRepository : BaseRepository<Organization>
{
    public OrganizationRepository(DataContext context) : base(context)
    {
    }
}
