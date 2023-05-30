// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using CoproXpert.Database.Models.Building.Enums;

namespace CoproXpert.Database.Models.Building;

public class SharedFeature : BaseModel
{
    [Key] public Guid Id { get; set; }
    public FeatureType? Type { get; set; }
    public string? Description { get; set; }

    public bool IsFunctioning { get; set; }
    // TODO: handle maintenance.

    public Community? Community { get; set; }
}
