// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Diagnostics;
using System.Globalization;
using CoproXpert.Core.Variables;
using CoproXpert.Database.Models.Information;

namespace CoproXpert.Database.Fixtures;

public class CountryFixture : FixtureBase
{
    private readonly DataContext _dataContext;

    public CountryFixture(DataContext dataContext)
    {
        _dataContext = dataContext;
    }

    public override Task Initialize()
    {
        return Task.CompletedTask;
    }

    public override Task Execute()
    {
        const string Countries = "/home/elmahdi/Downloads/countryInfo.txt"; // Replace with the actual file path
        const string cities = "/home/elmahdi/Downloads/allCountries.txt"; // Replace with the actual file path
        const bool LoadedCountries = true;
        if (LoadedCountries)
        {
            using var fileStream = new StreamReader(Countries);
            while (fileStream.ReadLine() is { } line)
            {
                var columns = line.Split('\t');
                if (columns.Length < 19)
                {
                    continue;
                }

                var countryCode = columns[0];
                var countryName = columns[4];
                var currency = columns[10];
                var extension = columns[12];
                var country = new Country
                {
                    Id = new Guid(),
                    Code = countryCode,
                    Name = countryName,
                    Currency = currency,
                    Extention = extension
                };
                _dataContext.Countries!.Add(country);

                Debug.WriteLine($"Country: {country.Name}, Code: {country.Code}");
                _dataContext.SaveChanges();
            }
        }

        using var countryStream = new StreamReader(cities);
        while (countryStream.ReadLine() is { } line)
        {
            var columns = line.Split('\t');
            string[] validFeatureCodes = { "PPLC", "PPLA", "PPLA2", "PPLA3", "PPLA4", "PPLA5" };

            if (columns.Length < 19 || !validFeatureCodes.Contains(columns[7]))
            {
                continue;
            }

            var country = _dataContext.Countries!.FirstOrDefault(c => c.Code == columns[8]);
            if (country is null)
            {
                continue;
            }

            country.Timezone ??= columns[17];

            var cityName = columns[1];
            var existingCity =
                _dataContext.Cities!.FirstOrDefault(c => c.Name == cityName && c.Country.Id == country.Id);
            if (existingCity is not null)
            {
                continue;
            }

            var city = new City
            {
                Id = new Guid(),
                Name = cityName,
                Country = country,
                Location = new GpsPosition(double.Parse(columns[4], CultureInfo.InvariantCulture),
                    double.Parse(columns[5], CultureInfo.InvariantCulture))
            };

            _dataContext.Cities!.Add(city);
            Debug.WriteLine("City: " + city.Name + ", Country: " + city.Country.Name + ", Location: " +
                            city.Location);
            _dataContext.SaveChanges();
        }

        return Task.CompletedTask;
    }
}
