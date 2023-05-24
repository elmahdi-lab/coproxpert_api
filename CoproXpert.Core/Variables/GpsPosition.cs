// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Core.Variables;

public readonly struct GpsPosition : IEquatable<GpsPosition>
{
    private double Latitude { get; }
    private double Longitude { get; }

    private GpsPosition(double latitude, double longitude)
    {
        Latitude = latitude;
        Longitude = longitude;
    }

    public override string ToString()
    {
        return $"{Latitude},{Longitude}";
    }

    public static GpsPosition Parse(string value)
    {
        var split = value.Split(',');
        return new GpsPosition(double.Parse(split[0]), double.Parse(split[1]));
    }

    public override bool Equals(object obj)
    {
        throw new NotImplementedException();
    }

    public override int GetHashCode()
    {
        throw new NotImplementedException();
    }

    public static bool operator ==(GpsPosition left, GpsPosition right)
    {
        return left.Equals(right);
    }

    public static bool operator !=(GpsPosition left, GpsPosition right)
    {
        return !(left == right);
    }

    public bool Equals(GpsPosition other)
    {
        throw new NotImplementedException();
    }
}
