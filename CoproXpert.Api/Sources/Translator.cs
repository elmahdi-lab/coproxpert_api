using System.Diagnostics;
using System.Globalization;
using Newtonsoft.Json.Linq;

namespace CoproXpert.Api.Sources;

/// <summary>
/// </summary>
public class Translator
{
    private readonly string _jsonFilesPath = "./Resources";
    private CultureInfo _defaultLanguage = new("en-CA");
    private Dictionary<string, JObject>? _translations;

    /// <summary>
    /// </summary>
    /// <param name="jsonFilesPath"></param>
    /// <param name="defaultLanguage"></param>
    public Translator(string? jsonFilesPath = null, CultureInfo? defaultLanguage = null)
    {
        if (jsonFilesPath != null)
        {
            _jsonFilesPath = jsonFilesPath;
        }

        if (defaultLanguage != null)
        {
            _defaultLanguage = defaultLanguage;
        }

        LoadTranslations();
    }

    private void LoadTranslations()
    {
        _translations = new Dictionary<string, JObject>();

        foreach (var cultureFolder in Directory.GetDirectories(_jsonFilesPath))
        {
            var cultureCode = Path.GetFileName(cultureFolder);

            var filePath = Path.Combine(cultureFolder, "index.json");
            if (!File.Exists(filePath))
            {
                continue;
            }

            var json = File.ReadAllText(filePath);
            var translationObject = JObject.Parse(json);

            _translations[cultureCode] = translationObject;
        }
    }

    /// <summary>
    /// </summary>
    /// <param name="key"></param>
    /// <param name="cultureCode"></param>
    /// <param name="namedArgs"></param>
    /// <returns></returns>
    /// <exception cref="ArgumentException"></exception>
    public string Translate(string key, string? cultureCode = null, Dictionary<string, object>? namedArgs = null)
    {
        cultureCode ??= _defaultLanguage.Name;

        if (cultureCode != _defaultLanguage.Name)
        {
            _defaultLanguage = new CultureInfo(cultureCode);
            LoadTranslations();
        }

        Debug.Assert(_translations != null, nameof(_translations) + " != null");
        if (!_translations.ContainsKey(cultureCode))
        {
            throw new ArgumentException($"Translation file for culture '{cultureCode}' not found.");
        }

        var translationObject = _translations[cultureCode];
        var translation = GetTranslation(key, translationObject);

        if (namedArgs is { Count: > 0 })
        {
            translation = ReplaceNamedPlaceholders(translation, namedArgs);
        }

        return translation;
    }

    private static string ReplaceNamedPlaceholders(string translation, Dictionary<string, object> namedArgs)
    {
        foreach (var namedArg in namedArgs)
        {
            var placeholder = $"%{namedArg.Key}%";
            translation = translation.Replace(placeholder, namedArg.Value.ToString());
        }

        return translation;
    }

    private static string GetTranslation(string key, JObject translationObject)
    {
        var keys = key.Split('.');
        var currentObject = translationObject;

        foreach (var k in keys)
        {
            if (!currentObject.TryGetValue(k, out var value))
            {
                continue;
            }

            if (value.Type != JTokenType.Object)
            {
                if (value.Type == JTokenType.String)
                {
                    return value.ToString();
                }
            }
            else
            {
                currentObject = (JObject)value;
            }
        }

        throw new KeyNotFoundException($"Translation key '{key}' not found.");
    }
}
