// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Core.Enums;

namespace CoproXpert.Core.Attributes;

[AttributeUsage(AttributeTargets.Class)]
public class AutowireAttribute : Attribute
{
    public AutowireAttribute(Lifetime lifetime)
    {
        Lifetime = lifetime;
    }

    public Lifetime Lifetime { get; }
}
