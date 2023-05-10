// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using CoProXpert.Database.Attribute;
using CoproXpert.Database.Models;

namespace CoProXpert.Database.Models.Security;

public class Permission : BaseModel
{
    [Key] public Guid Id { get; set; }
    public User User { get; set; }
    public Type Entity { get; set; }
    public Guid EntityId { get; set; }
    public string Role { get; set; }
    public DateTime? IsExpiringAt { get; set; }
}
