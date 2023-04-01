using CoproXpert.Database.Models;
using Microsoft.EntityFrameworkCore;

namespace CoproXpert.Database;

public class DataContext : DbContext
{
    public DbSet<User>? Users { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder options)
    {
        const string connection = "Host=localhost;Database=copro_xpert_db;Username=postgres;Password=postgres";
        if (connection.GetType() != typeof(string)) throw new Exception("Connection string is not a string.");
        options.UseNpgsql(connection);
    }
}