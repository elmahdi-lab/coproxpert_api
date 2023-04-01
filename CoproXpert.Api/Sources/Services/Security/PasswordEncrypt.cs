using System.Security.Cryptography;

namespace CoproXpert.Api.Sources.Security;

public static class PasswordEncrypt
{
    private const int Iterations = 10000;
    private const int BytesNumber = 64;
    private static readonly HashAlgorithmName HashingAlgorithm = HashAlgorithmName.SHA512;

    public static string? Encrypt(byte[] password, byte[] salt)
    {
        var hash = new Rfc2898DeriveBytes(password, salt, Iterations, HashingAlgorithm);
        return Convert.ToBase64String(hash.GetBytes(BytesNumber));
    }

    public static bool PasswordCheck(byte[] password, byte[] salt, string storedHash)
    {
        var hash = new Rfc2898DeriveBytes(password, salt, Iterations, HashingAlgorithm);
        var enteredHash = hash.GetBytes(BytesNumber);
        var decodedHash = Convert.FromBase64String(storedHash);
        return enteredHash.SequenceEqual(decodedHash);
    }
}