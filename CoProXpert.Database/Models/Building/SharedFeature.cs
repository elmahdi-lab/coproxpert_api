// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using CoproXpert.Database.Models;
using CoProXpert.Database.Models.Building.Enum;

namespace CoProXpert.Database.Models.Building;

public class SharedFeature : BaseModel
{
    [Key] public Guid Id { get; set; }
    public FeatureType? Type { get; set; }
    public string? Description { get; set; }

    public bool IsFunctioning { get; set; }
    // TODO: handle maintenance.

    public Community? Community { get; set; }
}
