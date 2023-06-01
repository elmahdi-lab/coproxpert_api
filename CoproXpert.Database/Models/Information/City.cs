// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using CoproXpert.Core.Variables;

namespace CoproXpert.Database.Models.Information;

public class City : BaseModel
{
    [Key] public Guid Id { get; set; }
    public string? Name { get; set; }

    [NotMapped]
    public GpsPosition? Location
    {
        get => GpsPosition.Parse(LocationString, out var position) ? position : null;
        set => LocationString = value?.ToString();
    }

    private string? LocationString { get; set; }

    public string? ZipCode { get; set; } = null!;

    public Country Country { get; set; } = null!;
}
