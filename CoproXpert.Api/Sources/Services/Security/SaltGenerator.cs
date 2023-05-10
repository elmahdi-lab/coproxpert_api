// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Security.Cryptography;

namespace CoproXpert.Api.Sources.Services.Security;

/// <summary>
///   Salt generator
/// </summary>
public static class SaltGenerator
{
    /// <summary>
    ///  Generate a random salt
    /// </summary>
    /// <param name="length"></param>
    /// <returns></returns>
    public static byte[] GenerateBytes(int length = 32)
    {
        // Create a byte array to store the salt bytes
        var saltBytes = new byte[length];
        // Create a new instance of the RandomNumberGenerator class
        var rng = RandomNumberGenerator.Create();
        // Fill the array with a random value
        rng.GetBytes(saltBytes);
        // Return a Base64 string representation of the random value
        return saltBytes;
    }

    /// <summary>
    /// Generate a random salt
    /// </summary>
    /// <param name="length"></param>
    /// <returns></returns>
    public static string GenerateString(int length = 32)
    {
        var saltBytes = GenerateBytes(length);
        var generated = Convert.ToBase64String(saltBytes);
        return generated[..length];
    }
}
