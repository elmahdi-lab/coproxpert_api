// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;

namespace CoproXpert.Database.Models.Building;

public class Domicile : BaseModel
{
    [Key] public Guid Id { get; set; }

    public Community Community { get; set; } = null!;
}
