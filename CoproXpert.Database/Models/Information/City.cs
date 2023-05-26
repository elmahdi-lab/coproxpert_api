// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using CoproXpert.Database.Models;

namespace CoproXpert.Database.Models.Information;

public class City : BaseModel
{
    [Key] public Guid Id { get; set; }
}
