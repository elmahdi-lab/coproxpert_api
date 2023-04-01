namespace CoproXpert.Database.Models;

public class Credential : BaseModel
{
    public int Id { get; set; }
    public string Password { get; set; }
    public string Salt { get; set; }
}