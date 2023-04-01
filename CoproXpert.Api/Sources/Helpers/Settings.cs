namespace CoproXpert.Api.Sources.Helpers;

public class Database
{
    public string ConnectionString { get; set; }
}

public class Settings
{
    public Database Database;
}