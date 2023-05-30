// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using Microsoft.EntityFrameworkCore;

namespace CoproXpert.Database.Models.Information;

[Index(nameof(Code), IsUnique = true)]
public class Country : BaseModel
{
    [Key] public Guid Id { get; set; }

    public string? Name { get; set; }

    public string? Code { get; set; }
}
