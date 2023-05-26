// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models;

namespace CoproXpert.Database.Models.Information;

/// <summary>
/// Represents an address with street number, street address, country, city, state and postal code.
/// </summary>
public class Address : BaseModel
{
    /// <summary>
    /// Gets or sets the street number of the address.
    /// </summary>
    public string? StreetNumber { get; set; }

    /// <summary>
    /// Gets or sets the street address of the address.
    /// </summary>
    public string? StreetAddress { get; set; }

    /// <summary>
    /// Gets or sets the country of the address.
    /// </summary>
    public Country Country { get; set; }

    /// <summary>
    /// Gets or sets the city of the address.
    /// </summary>
    public City City { get; set; }

    /// <summary>
    /// Gets or sets the state of the address.
    /// </summary>
    public string? State { get; set; }

    /// <summary>
    /// Gets or sets the postal code of the address.
    /// </summary>
    public string? PostalCode { get; set; }
}
