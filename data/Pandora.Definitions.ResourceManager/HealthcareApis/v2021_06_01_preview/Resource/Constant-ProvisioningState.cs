using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.HealthcareApis.v2021_06_01_preview.Resource;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum ProvisioningStateConstant
{
    [Description("Accepted")]
    Accepted,

    [Description("Canceled")]
    Canceled,

    [Description("Creating")]
    Creating,

    [Description("Deleting")]
    Deleting,

    [Description("Deprovisioned")]
    Deprovisioned,

    [Description("Failed")]
    Failed,

    [Description("Succeeded")]
    Succeeded,

    [Description("Updating")]
    Updating,

    [Description("Verifying")]
    Verifying,
}
