// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;

namespace CoproXpert.Database.Models.Building;

public class Community : BaseModel
{
    [Key] public Guid Id { get; set; }

    public ICollection<Domicile>? Domiciles { get; } = new List<Domicile>();
}
