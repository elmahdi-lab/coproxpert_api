// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoProXpert.Database.Models;
using CoProXpert.Database.Models.Security;

namespace CoproXpert.Database.Models;

public class User : BaseModel, IModel
{
    public int Id { get; set; }
    public string Username { get; set; }

    public Credential? Credential { get; set; }
    public ICollection<Permission>? Permissions { get; set; }
}
