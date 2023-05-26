// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Security.Cryptography;

namespace CoproXpert.Core.Encryption;

public class Encryptor
{
    private readonly byte[] _iv;
    private readonly byte[] _key;

    public Encryptor(byte[] key, byte[] iv)
    {
        _key = key;
        _iv = iv;
    }

    public string Encrypt(string plainText)
    {
        byte[] encrypted;

        using (var aes = Aes.Create())
        {
            aes.Key = _key;
            aes.IV = _iv;

            var encryptor = aes.CreateEncryptor(aes.Key, aes.IV);

            using (var ms = new MemoryStream())
            {
                using (var cs = new CryptoStream(ms, encryptor, CryptoStreamMode.Write))
                {
                    using (var sw = new StreamWriter(cs))
                    {
                        sw.Write(plainText);
                    }

                    encrypted = ms.ToArray();
                }
            }
        }

        return Convert.ToBase64String(encrypted);
    }

    public string Decrypt(string cipherText)
    {
        var cipherBytes = Convert.FromBase64String(cipherText);

        using var aes = Aes.Create();
        aes.Key = _key;
        aes.IV = _iv;

        var decryptor = aes.CreateDecryptor(aes.Key, aes.IV);

        using var ms = new MemoryStream(cipherBytes);
        using var cs = new CryptoStream(ms, decryptor, CryptoStreamMode.Read);
        using var sr = new StreamReader(cs);
        return sr.ReadToEnd();
    }
}
