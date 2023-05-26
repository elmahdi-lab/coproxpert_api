// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Text;
using CoproXpert.Core.Encryption;

namespace CoproXpert.Test;

public class EncryptorTests
{
    private readonly byte[] _iv = Encoding.UTF8.GetBytes("fedcba9876543210");
    private readonly byte[] _key = Encoding.UTF8.GetBytes("0123456789abcdef");

    [Test]
    public void Encrypt_Decrypt_ReturnsOriginalString()
    {
        // Arrange
        var encryptor = new Encryptor(_key, _iv);
        var plainText = "Hello, world!";

        // Act
        var cipherText = encryptor.Encrypt(plainText);
        var decryptedText = encryptor.Decrypt(cipherText);

        // Assert
        Assert.That(decryptedText, Is.EqualTo(plainText));
    }
}
