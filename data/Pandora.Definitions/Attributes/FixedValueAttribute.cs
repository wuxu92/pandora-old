using System;

namespace Pandora.Definitions.Attributes;

public class FixedValueAttribute : Attribute
{
    public string FixedValue { get; }

    public FixedValueAttribute(string value)
    {
        FixedValue = value;
    }
}