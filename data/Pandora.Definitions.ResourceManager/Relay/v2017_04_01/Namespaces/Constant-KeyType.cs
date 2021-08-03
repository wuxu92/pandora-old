using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Relay.v2017_04_01.Namespaces
{
    [ConstantType(ConstantTypeAttribute.ConstantType.String)]
    internal enum KeyType
    {
        [Description("PrimaryKey")]
        PrimaryKey,

        [Description("SecondaryKey")]
        SecondaryKey,
    }
}