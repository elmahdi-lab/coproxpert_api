// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Core.Variables;

public readonly struct GpsPosition
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
}
