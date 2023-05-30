// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Database.Models.Security;

//[Index(nameof(User), nameof(System.Type), IsUnique = true)]
public class Claim : BaseModel
{
    public Guid Id { get; set; }
    public string Type { get; set; } = null!;
    public string Value { get; set; } = null!;
}
