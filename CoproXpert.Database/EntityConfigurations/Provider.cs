// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models.Security;
using CoproXpert.Database.Models.Security.Type;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace CoproXpert.Database.EntityConfiguration;

public class Provider : IEntityTypeConfiguration<Social>
{
    public void Configure(EntityTypeBuilder<Social>? builder)
    {
        builder?.Property(e => e.Provider)
            .HasConversion(
                v => v.ToString(),
                v => (SocialProvider)Enum.Parse(typeof(SocialProvider), v))
            .HasColumnName("provider")
            .HasColumnType("social_provider");
    }
}
