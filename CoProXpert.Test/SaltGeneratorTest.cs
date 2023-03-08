using CoproXpert.Sources.Security;

namespace CoProXpert.Test;

public class SaltGeneratorTest
{
    private string _salt = null!;
    
    [SetUp]
    public void Setup()
    {
        _salt = SaltGenerator.GenerateString(16);
    }

    [Test]
    public void AssertLength()
    {
        Assert.That(_salt, Has.Length.EqualTo(16));
        Assert.That(_salt, Is.Not.Null);
    }
    
    [Test]
    public void AssertType()
    {
        Assert.That(_salt.GetType(), Is.EqualTo(typeof(string)));
    }
}