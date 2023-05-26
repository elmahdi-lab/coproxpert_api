// Copyright (c) COPRO XPERT - IT HUMANS All Rights Reserved.


namespace CoproXpert.Core.Variables;

/// <summary>
///     Represents a GPS position with latitude and longitude coordinates.
/// </summary>
public readonly struct GpsPosition : IEquatable<GpsPosition>
{
    /// <summary>
    ///     Gets the latitude coordinate of the GPS position.
    /// </summary>
    private double Latitude { get; }

    /// <summary>
    ///     Gets the longitude coordinate of the GPS position.
    /// </summary>
    private double Longitude { get; }

    /// <summary>
    ///     Initializes a new instance of the <see cref="GpsPosition" /> struct.
    /// </summary>
    /// <param name="latitude">The latitude coordinate.</param>
    /// <param name="longitude">The longitude coordinate.</param>
    private GpsPosition(double latitude, double longitude)
    {
        Latitude = latitude;
        Longitude = longitude;
    }

    /// <summary>
    ///     Converts the GPS position to a string representation in the format "latitude,longitude".
    /// </summary>
    /// <returns>A string representation of the GPS position.</returns>
    public override string ToString()
    {
        return $"{Latitude},{Longitude}";
    }

    /// <summary>
    ///     Parses a string representation of a GPS position into a <see cref="GpsPosition" /> object.
    /// </summary>
    /// <param name="value">The string representation of the GPS position in the format "latitude,longitude".</param>
    /// <returns>A <see cref="GpsPosition" /> object representing the parsed GPS position.</returns>
    public static GpsPosition Parse(string value)
    {
        var split = value.Split(',');
        return new GpsPosition(double.Parse(split[0]), double.Parse(split[1]));
    }

    /// <summary>
    ///     Determines whether the current <see cref="GpsPosition" /> object is equal to another <see cref="GpsPosition" />
    ///     object.
    /// </summary>
    /// <param name="obj">The object to compare with the current object.</param>
    /// <returns><c>true</c> if the objects are equal; otherwise, <c>false</c>.</returns>
    public override bool Equals(object? obj)
    {
        if (obj is GpsPosition other)
        {
            return Equals(other);
        }

        return false;
    }

    /// <summary>
    ///     Returns a hash code for the current <see cref="GpsPosition" /> object.
    /// </summary>
    /// <returns>A hash code for the current object.</returns>
    public override int GetHashCode()
    {
        return HashCode.Combine(Latitude, Longitude);
    }

    /// <summary>
    ///     Determines whether the current <see cref="GpsPosition" /> is equal to another <see cref="GpsPosition" /> object.
    /// </summary>
    /// <param name="other">The <see cref="GpsPosition" /> to compare with the current <see cref="GpsPosition" />.</param>
    /// <returns>
    ///     <c>true</c> if the specified <see cref="GpsPosition" /> is equal to the current <see cref="GpsPosition" />;
    ///     otherwise, <c>false</c>.
    /// </returns>
    public bool Equals(GpsPosition other)
    {
        const double Tolerance = 0.000001; // Define a tolerance value for comparing floating-point numbers
        return Math.Abs(Latitude - other.Latitude) < Tolerance &&
               Math.Abs(Longitude - other.Longitude) < Tolerance;
    }

    /// <summary>
    ///     Checks if two <see cref="GpsPosition" /> objects are equal.
    /// </summary>
    /// <param name="left">The first <see cref="GpsPosition" /> object to compare.</param>
    /// <param name="right">The second <see cref="GpsPosition" /> object to compare.</param>
    /// <returns><c>true</c> if the objects are equal; otherwise, <c>false</c>.</returns>
    public static bool operator ==(GpsPosition left, GpsPosition right)
    {
        return left.Equals(right);
    }

    /// <summary>
    ///     Checks if two <see cref="GpsPosition" /> objects are not equal.
    /// </summary>
    /// <param name="left">The first <see cref="GpsPosition" /> object to compare.</param>
    /// <param name="right">The second <see cref="GpsPosition" /> object to compare.</param>
    /// <returns><c>true</c> if the objects are not equal; otherwise, <c>false</c>.</returns>
    public static bool operator !=(GpsPosition left, GpsPosition right)
    {
        return !(left == right);
    }
}
