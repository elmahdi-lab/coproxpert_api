// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Reflection;
using CoproXpert.Database.Models;
using CoProXpert.Database.Attribute;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;

namespace CoproXpert.Database;

public class DataContext : DbContext
{
    public DbSet<User>? Users { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder options)
    {
        const string connection = "Host=localhost;Database=copro_xpert_db;Username=postgres;Password=postgres";
        if (connection.GetType() != typeof(string))
        {
            throw new Exception("Connection string is not a string.");
        }

        options.UseNpgsql(connection);
    }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
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
