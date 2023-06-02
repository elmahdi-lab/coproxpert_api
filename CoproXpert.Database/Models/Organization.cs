// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models.Building;
using CoproXpert.Database.Models.Information;

namespace CoproXpert.Database.Models;

public class Organization : BaseModel
{
    public Guid Id { get; set; }
    public string Name { get; set; } = null!;
    public string? Description { get; set; }
    public string? Logo { get; set; }
    public string? Website { get; set; }
    public Contact? Contact { get; set; }

    public ICollection<Community>? Communities { get; } = new List<Community>();
}
