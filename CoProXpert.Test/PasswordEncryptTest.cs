using System.Text;
using CoproXpert.Sources.Security;

namespace CoProXpert.Test;

public class PasswordEncryptTest
{
    private string _password = "P@ssw0rd!";
    private string _salt = null!;
    
    
    [SetUp]
    public void Setup()
    {
        _salt = SaltGenerator.GenerateString(16);
    }

    [Test]
    public void EncryptPasswordTest()
    {
        var encrypted = PasswordEncrypt.Encrypt(Encoding.UTF8.GetBytes(_password), Encoding.UTF8.GetBytes(_salt));
        Assert.That(encrypted, Is.Not.Null);
    }
    
    [Test]
    public void DecryptPasswordTest()
    {
        var encrypted = PasswordEncrypt.Encrypt(Encoding.UTF8.GetBytes(_password), Encoding.UTF8.GetBytes(_salt));
        var isValid = PasswordEncrypt.PasswordCheck(Encoding.UTF8.GetBytes(_password), Encoding.UTF8.GetBytes(_salt), encrypted!);
        Assert.That(isValid, Is.True);
    }
}