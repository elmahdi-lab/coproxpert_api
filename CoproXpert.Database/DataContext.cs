// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Reflection;
using CoproXpert.Database.EntityConfiguration;
using CoproXpert.Database.Models;
using CoProXpert.Database.Models.Building;
using CoProXpert.Database.Models.Information;
using CoProXpert.Database.Models.Security.Type;
using Microsoft.EntityFrameworkCore;

namespace CoproXpert.Database;

public class DataContext : DbContext
{
    public DbSet<User>? Users { get; set; }
    public DbSet<Organization>? Organizations { get; set; }

    public DbSet<Contact>? Contacts { get; set; }
    public DbSet<Address>? Addresses { get; set; }
    public DbSet<City>? Cities { get; set; }
    public DbSet<Country>? Countries { get; set; }

    public DbSet<Community>? Communities { get; set; }
    public DbSet<Domicile>? Domiciles { get; set; }
    public DbSet<SharedFeature>? SharedFeatures { get; set; }
    public DbSet<SharedSpace>? SharedSpaces { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder options)
    {
        const string Connection =
            "Host=localhost;Port=5532;Database=cx_db;Username=cx_user;Password=Password";
        if (Connection.GetType() != typeof(string))
        {
            throw new Exception("Connection string is not a string.");
        }

        options.UseNpgsql(Connection);
    }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.HasPostgresEnum<SocialProvider>();
        CollectionToString.ApplyCustomConfigurations(modelBuilder);
        modelBuilder.ApplyConfigurationsFromAssembly(Assembly.GetExecutingAssembly()); // This may not be necessary
        base.OnModelCreating(modelBuilder);
    }
}
