namespace CoproXpert.Sources.Models;

public class User
{
    public int Id { get; set; }
    public string Username { get; set; }
    
    public DateTime CreatedAt { get; set; }
    public DateTime UpdatedAt { get; set; }

    public Credential Credential { get; set; } = new Credential();
    public ICollection<Permission> Permissions { get; set; } = new List<Permission>();
}

