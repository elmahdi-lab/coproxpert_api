// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Security.Cryptography;

namespace CoproXpert.Core.Security;

/// <summary>
///     Salt generator
/// </summary>
public static class KeyGenerator
{
    /// <summary>
    ///     Generate a random key
    /// </summary>
    /// <param name="length"></param>
    /// <returns></returns>
    public static byte[] GenerateBytes(int length = 32)
    {
        // Create a byte array to store the key bytes
        var keyBytes = new byte[length];
        // Create a new instance of the RandomNumberGenerator class
        using var rng = RandomNumberGenerator.Create();
        // Fill the array with a random value
        rng.GetBytes(keyBytes);
        // Return a Base64 string representation of the random value

        return keyBytes;
    }

    /// <summary>
    ///     Generate a random key
    /// </summary>
    /// <param name="length"></param>
    /// <returns></returns>
    public static string GenerateString(int length = 32)
    {
        var keyBytes = GenerateBytes(length);
        var generated = Convert.ToBase64String(keyBytes);
        return generated[..length];
    }
}
