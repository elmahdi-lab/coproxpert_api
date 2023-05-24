using CoproXpert.Api.Sources;

namespace CoProXpert.Test;

[TestFixture]
public class TranslatorTests
{
    [SetUp]
    public void Setup()
    {
        _translator = new Translator("../../../../CoproXpert.Api/Sources/Resources");
    }

    private Translator _translator;

    [Test]
    public void Translate_SingleTranslation_ReturnsCorrectTranslation()
    {
        var translation = _translator.Translate("index.welcome");
        Assert.That(translation, Is.EqualTo("Welcome"));
    }

    [Test]
    public void Translate_TranslationWithPlaceholder_ReturnsTranslatedStringWithPlaceholderReplaced()
    {
        var namedArgs = new Dictionary<string, object> { { "name", "John" } };
        var translation = _translator.Translate("index.greeting", "en-CA", namedArgs);
        Assert.That(translation, Is.EqualTo("Hello, John!"));
    }

    [Test]
    public void Translate_NestedTranslation_ReturnsCorrectTranslation()
    {
        var translation = _translator.Translate("errors.backend.unauthorized", "en-CA");
        Assert.That(translation, Is.EqualTo("Unauthorized"));
    }

    [Test]
    public void Translate_MissingTranslationFile_ThrowsArgumentException()
    {
        Assert.Throws<ArgumentException>(() => _translator.Translate("index.welcome", "es-ES"));
    }

    [Test]
    public void Translate_MissingTranslationKey_ThrowsKeyNotFoundException()
    {
        Assert.Throws<KeyNotFoundException>(() => _translator.Translate("invalid.key"));
    }

    [Test]
    public void Translate_PluralTranslation_ReturnsCorrectTranslation()
    {
        var translation1 = _translator.Translate("apples", "en-CA", new Dictionary<string, object> { { "count", 0 } });
        Assert.That(translation1, Is.EqualTo("There are no apples"));

        var translation2 = _translator.Translate("apples", "en-CA", new Dictionary<string, object> { { "count", 25 } });
        Assert.That(translation2, Is.EqualTo("There are many apples"));

        var translation3 = _translator.Translate("apples", "en-CA", new Dictionary<string, object> { { "count", 1 } });
        Assert.That(translation3, Is.EqualTo("There is one apple"));

        var translation4 = _translator.Translate("apples", "en-CA", new Dictionary<string, object> { { "count", 3 } });
        Assert.That(translation4, Is.EqualTo("There are 3 apples"));
    }
}
