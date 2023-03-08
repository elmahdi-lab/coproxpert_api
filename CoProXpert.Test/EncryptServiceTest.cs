using CoproXpert.Sources.Security;
using CoproXpert.Sources.Services.Encryption;

namespace CoProXpert.Test;

public class EncryptServiceTest
{
    private const string Text = "Hello World!";
    private string _encrypted = null!;
    private byte[] _iv = null!;
    private byte[] _key = null!;

    [SetUp]
    public void Setup()
    {
        _key = SaltGenerator.GenerateBytes();
        _iv = SaltGenerator.GenerateBytes(16);
        _encrypted = EncryptService.Encrypt(Text, _key, _iv);
    }

    [Test]
    public void EncryptTest()
    {
        Assert.That(_encrypted, Is.Not.EqualTo(Text));
        Assert.That(_encrypted.GetType(), Is.EqualTo(typeof(string)));
    }

    [Test]
    public void DecryptTest()
    {
        var decrypted = EncryptService.Decrypt(_encrypted, _key, _iv);
        Assert.That(decrypted, Is.EqualTo(Text));
    }
}