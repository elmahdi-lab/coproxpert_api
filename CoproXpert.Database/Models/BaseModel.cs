// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Database.Models;

/// <summary>
/// Base class for all domain models.
/// </summary>
public abstract class BaseModel
{
    /// <summary>
    /// Initializes a new instance of the <see cref="BaseModel"/> class.
    /// </summary>
    protected BaseModel()
    {
        CreatedAt = DateTime.UtcNow;
    }

    /// <summary>
    /// Gets or sets the date and time when the model was created.
    /// </summary>
    public DateTime CreatedAt { get; set; }

    /// <summary>
    /// Gets or sets the date and time when the model was last updated.
    /// </summary>
    public DateTime? UpdatedAt { get; set; }

    /// <summary>
    /// Sets the value of <see cref="UpdatedAt"/> to the current UTC date and time.
    /// </summary>
    public void SetUpdatedAt()
    {
        UpdatedAt = DateTime.UtcNow;
    }
}
