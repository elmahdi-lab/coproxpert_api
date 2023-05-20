// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoProXpert.Database.Models.Building;
using CoProXpert.Database.Models.Communication;

namespace CoproXpert.Database.Models;

public class Organization : BaseModel
{
    public int Id { get; set; }
    public string Name { get; set; } = null!;
    public string? Description { get; set; }
    public string? Website { get; set; }
    public Contact Contact { get; set; } = null!;
    public ICollection<Community>? Communities { get; set; }
}
