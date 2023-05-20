// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Reflection;
using CoProXpert.Database.Attribute;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;

namespace CoproXpert.Database.EntityConfiguration;

public class CollectionToString
{
    public static void ApplyCustomConfigurations(ModelBuilder modelBuilder)
    {
        foreach (var entityType in modelBuilder.Model.GetEntityTypes())
        {
            foreach (var property in entityType.GetProperties())
            {
                var propertyInfo = property.PropertyInfo;
                if (propertyInfo == null || propertyInfo.GetCustomAttribute<DbCollection>() == null)
                {
                    continue;
                }

                if (propertyInfo.PropertyType.GenericTypeArguments[0] == typeof(string))
                {
                    property.SetValueConverter(
                        new ValueConverter<ICollection<string>, string>(
                            topics => string.Join(',', topics),
                            topics => topics.Split(',', StringSplitOptions.RemoveEmptyEntries).ToList()
                        )
                    );
                }
                // Add more cases for other data types if needed
            }
        }
    }
}
