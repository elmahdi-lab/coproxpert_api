// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Core.Security;

namespace CoproXpert.Test;

[TestFixture]
public class KeyGeneratorTests
{
    [Test]
    public void GenerateBytes_ShouldGenerateRandomKeyWithSpecifiedLength()
    {
        // Arrange
        const int Length = 32;

        // Act
        var keyBytes = KeyGenerator.GenerateBytes();

        // Assert
        Assert.That(keyBytes, Has.Length.EqualTo(Length));
        Assert.That(keyBytes.All(_ => true), Is.True);
    }

    [Test]
    public void GenerateString_ShouldGenerateRandomStringWithSpecifiedLength()
    {
        // Arrange
        const int Length = 32;

        // Act
        var keyString = KeyGenerator.GenerateString();

        // Assert
        Assert.That(keyString, Has.Length.EqualTo(Length));
        Assert.That(keyString.All(c => char.IsLetterOrDigit(c) || c == '+' || c == '/' || c == '='), Is.True);
    }

    [Test]
    public void GenerateString_WithLength256_ShouldGenerateRandomStringWithSpecifiedLength()
    {
        // Arrange
        const int Length = 256;

        // Act
        var keyString = KeyGenerator.GenerateString(Length);

        // Assert
        Assert.That(keyString, Has.Length.EqualTo(Length));
        Assert.That(keyString.All(c => char.IsLetterOrDigit(c) || c == '+' || c == '/' || c == '='), Is.True);
    }
}
