// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Api.Sources.Services.Encryption;
using CoproXpert.Api.Sources.Services.Security;

namespace CoproXpert.Test;

public class EncryptServiceTest
{
    private const string Text = "Hello World!";
    private string _encrypted = null!;

    [SetUp]
    public void Setup()
    {
        _encrypted = Encryptor.Encrypt(Text);
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
        // Padding is invalid Error

        var decrypted = Encryptor.Decrypt(_encrypted);
        Assert.That(decrypted, Is.EqualTo(Text));
    }
}
