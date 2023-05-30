// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using CoproXpert.Database.Models.Security.Type;

namespace CoproXpert.Database.Models.Security;

public class Social : BaseModel
{
    [Key] public Guid Id { get; set; }
    public SocialProvider Provider { get; set; }
    public string SocialId { get; set; } = null!;
    public string Token { get; set; } = null!;
}
