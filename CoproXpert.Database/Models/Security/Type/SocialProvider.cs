// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.ComponentModel.DataAnnotations;
using NpgsqlTypes;

namespace CoproXpert.Database.Models.Security.Type;

public enum SocialProvider
{
    [Display(Name = "Facebook")]
    [PgName("provider_type_facebook")]
    Facebook,

    [Display(Name = "Twitter")]
    [PgName("provider_type_twitter")]
    Twitter,

    [Display(Name = "Instagram")]
    [PgName("provider_type_instagram")]
    Instagram
}
