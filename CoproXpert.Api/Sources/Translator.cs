// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Diagnostics;
using System.Globalization;
using System.Text.RegularExpressions;
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

        // check if translation has pluralization
        if (namedArgs is { Count: > 0 } && translation.Contains('|'))
        {
            translation = ReplacePluralization(translation, namedArgs);
        }

        if (namedArgs is not null)
        {
            translation = ReplaceNamedPlaceholders(translation, namedArgs);
        }

        return translation;
    }

    private static string ReplacePluralization(string translation, Dictionary<string, object> namedArgs)
    {
        var rules = translation.Split('|');
        var count = int.Parse(namedArgs["count"].ToString() ?? string.Empty);

        var noItemsRule = rules.FirstOrDefault(r => r.Contains("{0}"));
        rules = rules.Where(r => !r.Contains("{0}")).ToArray();

        var oneItemRule = rules.FirstOrDefault(r => r.Contains("{1}"));
        rules = rules.Where(r => !r.Contains("{1}")).ToArray();

        var manyItemsRuleRegex = new Regex(@"^\{(\d+)\}");
        var manyItemsRule = rules.FirstOrDefault(r =>
        {
            var match = manyItemsRuleRegex.Match(r);
            if (!match.Success)
            {
                return false;
            }

            var ruleCount = int.Parse(match.Groups[1].Value);
            return count >= ruleCount;
        });

        noItemsRule = noItemsRule?.Replace("{0}", "");
        oneItemRule = oneItemRule?.Replace("{1}", "");

        manyItemsRule = manyItemsRule?.Replace(manyItemsRuleRegex.Match(manyItemsRule).Value, "");

        return count switch
        {
            0 when noItemsRule != null => ReplaceNamedPlaceholders(noItemsRule, namedArgs),
            1 when oneItemRule != null => ReplaceNamedPlaceholders(oneItemRule, namedArgs),
            > 1 when manyItemsRule != null => ReplaceNamedPlaceholders(manyItemsRule, namedArgs),
            _ => translation
        };
    }

    private static string ReplaceNamedPlaceholders(string translation, Dictionary<string, object> namedArgs)
    {
        foreach (var namedArg in namedArgs)
        {
            var placeholder = $"%{namedArg.Key}%";
            translation =
                translation.Replace(placeholder, namedArg.Value.ToString(), StringComparison.InvariantCulture);
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
