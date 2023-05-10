// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using CoproXpert.Database.Models;

namespace CoProXpert.Database.Models.Building;

public class Community : BaseModel, IModel
{
    [Key] public Guid Id { get; set; }

    public ICollection<Domicile> Domiciles { get; set; }
}
