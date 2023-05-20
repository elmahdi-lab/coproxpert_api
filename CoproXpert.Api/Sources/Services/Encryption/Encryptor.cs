// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Security.Cryptography;
using System.Text;

namespace CoproXpert.Api.Sources.Services.Encryption;

/// <summary>
///     Provides methods for encrypting and decrypting strings using AES encryption.
/// </summary>
public static class Encryptor
{
    /// <summary>
    ///     Encrypts a plain text string using the given key and initialization vector.
    /// </summary>
    /// <param name="plainText">The plain text string to encrypt.</param>
    /// <returns>The encrypted string as a base64-encoded string.</returns>
    public static string Encrypt(string plainText)
    {
        using var aes = Aes.Create();
        aes.KeySize = 256;
        aes.GenerateKey();
        aes.GenerateIV();
        var encryptor = aes.CreateEncryptor();
        using var memoryStream = new MemoryStream();
        using var cryptoStream = new CryptoStream(memoryStream, encryptor, CryptoStreamMode.Write);
        using var streamWriter = new StreamWriter(cryptoStream);
        streamWriter.Write(plainText);
        cryptoStream.FlushFinalBlock();
        var cipheredBytes = memoryStream.ToArray();
        var result = new byte[aes.IV.Length + cipheredBytes.Length];
        Array.Copy(aes.IV, result, aes.IV.Length);
        Array.Copy(cipheredBytes, 0, result, aes.IV.Length, cipheredBytes.Length);
        return Convert.ToBase64String(result);
    }

    /// <summary>
    ///     Decrypts a base64-encoded cipher text string using the given key and initialization vector.
    /// </summary>
    /// <param name="cipherText">The cipher text string to decrypt.</param>
    /// <returns>The decrypted plain text string.</returns>
    public static string Decrypt(string cipherText)
    {
        using var aes = Aes.Create();
        aes.KeySize = 256;
        var encryptedBytes = Convert.FromBase64String(cipherText);
        var iv = new byte[aes.IV.Length];
        Array.Copy(encryptedBytes, iv, aes.IV.Length);
        aes.IV = iv;
        var decrypt = aes.CreateDecryptor();
        using var memoryStream = new MemoryStream();
        using var cryptoStream = new CryptoStream(memoryStream, decrypt, CryptoStreamMode.Write);
        cryptoStream.Write(encryptedBytes, aes.IV.Length, encryptedBytes.Length - aes.IV.Length);
        cryptoStream.FlushFinalBlock();
        var plainBytes = memoryStream.ToArray();
        return Encoding.UTF8.GetString(plainBytes);
    }
}
