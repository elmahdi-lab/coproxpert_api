// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Database.Models.Information;

/// <summary>
///     Represents a contact.
/// </summary>
public class Contact : BaseModel
{
    /// <summary>
    ///     Gets or sets the ID of the contact.
    /// </summary>
    public int Id { get; set; }

    /// <summary>
    ///     Gets or sets the name of the contact.
    /// </summary>
    public string? Name { get; set; }

    /// <summary>
    ///     Gets or sets the email address of the contact.
    /// </summary>
    public string? Email { get; set; }

    /// <summary>
    ///     Gets or sets the phone number of the contact.
    /// </summary>
    public string? Phone { get; set; }

    /// <summary>
    ///     Gets or sets any notes about the contact.
    /// </summary>
    public string? Notes { get; set; }

    public Address Address { get; set; } = null!;
}
