// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using CoproXpert.Database.Models;
using CoproXpert.Database.Models.Building.Enum;

namespace CoproXpert.Database.Models.Building;

public class SharedSpace : BaseModel
{
    [Key] public Guid Id { get; set; }
    public SpaceType? Type { get; set; }
    public Community? Community { get; set; }
    public string? Description { get; set; }

    public bool IsFunctioning { get; set; }

    // TODO: handle maintenance
    public bool IsAccessible { get; set; }
    public bool IsPublic { get; set; }
    public bool IsOutdoor { get; set; }
    public bool IsIndoor { get; set; }
    public bool HasElectricity { get; set; }
    public bool HasWater { get; set; }
    public bool HasGas { get; set; }
    public bool HasInternet { get; set; }
    public bool HasTv { get; set; }
    public bool HasPhone { get; set; }
    public bool HasHeating { get; set; }
    public bool HasAirConditioning { get; set; }
    public bool HasParking { get; set; }
    public bool HasElevator { get; set; }
    public bool HasDisabledAccess { get; set; }
    public bool HasSecurity { get; set; }
    public bool HasSurveillance { get; set; }
    public bool HasAlarm { get; set; }
    public bool HasFireAlarm { get; set; }
    public bool HasFireExtinguisher { get; set; }
    public bool HasFirstAidKit { get; set; }
}
