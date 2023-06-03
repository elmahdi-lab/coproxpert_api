// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Security.Cryptography;
using CoproXpert.Core.Attributes;
using CoproXpert.Core.Enums;

namespace CoproXpert.Core.Security;

[Autowire(Lifetime.Singleton)]
public class PasswordHasher : IPasswordHasher
{
    private const int SaltSize = 128 / 8;
    private const int KeySize = 256 / 8;
    private const int HashIteration = 1000;
    private const string Delimiter = ";";
    private static readonly HashAlgorithmName s_hashAlgorithm = HashAlgorithmName.SHA256;

    public bool Verify(string passwordHash, string inputPassword)
    {
        if (passwordHash is null)
        {
            throw new ArgumentNullException(nameof(passwordHash));
        }

        var elements = passwordHash.Split(Delimiter);
        var salt = Convert.FromBase64String(elements[0]);
        var hash = Convert.FromBase64String(elements[1]);

        var hashInput = Rfc2898DeriveBytes.Pbkdf2(inputPassword, salt, HashIteration, s_hashAlgorithm, KeySize);

        return CryptographicOperations.FixedTimeEquals(hash, hashInput);
    }

    public string Hash(string password)
    {
        var salt = RandomNumberGenerator.GetBytes(SaltSize);
        var hash = Rfc2898DeriveBytes.Pbkdf2(password, salt, HashIteration, s_hashAlgorithm, KeySize);
        return string.Join(Delimiter, Convert.ToBase64String(salt), Convert.ToBase64String(hash));
    }
}

public interface IPasswordHasher
{
    bool Verify(string passwordHash, string inputPassword);
    string Hash(string password);
}
