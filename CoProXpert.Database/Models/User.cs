using CoProXpert.Database.Models;

namespace CoproXpert.Database.Models;

public class User : BaseModel, IModel
{
    public int Id { get; set; }

    public string Username { get; set; }
    //
    // public Credential Credential { get; set; } = new();
    // public ICollection<Permission> Permissions { get; set; } = new List<Permission>();
}