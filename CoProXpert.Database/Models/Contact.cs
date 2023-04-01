// Genrate a C# model for a contact in the database


using CoProXpert.Database.Attribute;

namespace CoproXpert.Database.Models;

internal class Contact
{
    public int Id { get; set; }
    public string Name { get; set; }

    [Encrypted] public string Email { get; set; }

    public string Phone { get; set; }
    public string Address { get; set; }
    public string City { get; set; }
    public string State { get; set; }
    public string Zip { get; set; }
    public string Country { get; set; }
    public string Notes { get; set; }
    public string CreatedBy { get; set; }
    public DateTime CreatedAt { get; set; }
    public string UpdatedBy { get; set; }
    public DateTime UpdatedAt { get; set; }
}