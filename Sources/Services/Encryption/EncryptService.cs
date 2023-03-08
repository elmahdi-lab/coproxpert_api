using System.Security.Cryptography;

namespace CoproXpert.Sources.Services.Encryption;

public class EncryptService
{
    private static Aes AesInit(byte[] key)
    {
        using var aes = Aes.Create();
        aes.Padding = PaddingMode.None;
        aes.Key = key;
        return aes;
    }

    public static string Encrypt(string plainText, byte[] key, byte[] iv)
    {
        byte[] cipheredText;
        using (var aes = Aes.Create())
        {
            var encryptor = aes.CreateEncryptor(key, iv);
            using (var memoryStream = new MemoryStream())
            {
                using (var cryptoStream = new CryptoStream(memoryStream, encryptor, CryptoStreamMode.Write))
                {
                    using (var streamWriter = new StreamWriter(cryptoStream))
                    {
                        streamWriter.Write(plainText);
                    }

                    cipheredText = memoryStream.ToArray();
                }
            }
        }

        return Convert.ToBase64String(cipheredText);
    }


    public static string Decrypt(string cipherText, byte[] key, byte[] iv)
    {
        using var aes = Aes.Create();
        var decryptor = aes.CreateDecryptor(key, iv);
        using var memoryStream = new MemoryStream(Convert.FromBase64String(cipherText));
        using var cryptoStream = new CryptoStream(memoryStream, decryptor, CryptoStreamMode.Read);
        using var streamReader = new StreamReader(cryptoStream);
        var simpleText = streamReader.ReadToEnd();

        return simpleText;
    }
}