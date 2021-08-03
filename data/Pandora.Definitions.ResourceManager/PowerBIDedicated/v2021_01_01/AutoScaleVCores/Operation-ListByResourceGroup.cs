using Pandora.Definitions.Attributes;
using Pandora.Definitions.Interfaces;
using Pandora.Definitions.Operations;
using System.Collections.Generic;
using System.Net;

namespace Pandora.Definitions.ResourceManager.PowerBIDedicated.v2021_01_01.AutoScaleVCores
{
    internal class ListByResourceGroupOperation : Operations.GetOperation
    {
        public override ResourceID? ResourceId()
        {
            return new ResourceGroupId();
        }

        public override object? ResponseObject()
        {
            return new AutoScaleVCoreListResult();
        }

        public override string? UriSuffix()
        {
            return "/providers/Microsoft.PowerBIDedicated/autoScaleVCores";
        }


    }
}