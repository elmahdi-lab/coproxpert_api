namespace CoproXpert.Database.Models;

public abstract class BaseModel
{
    protected BaseModel()
    {
        CreatedAt = DateTime.UtcNow;
    }

    public DateTime CreatedAt { get; set; }
    public DateTime? UpdatedAt { get; set; }

    public void SetUpdatedAt()
    {
        UpdatedAt = DateTime.UtcNow;
    }
}