using System.Collections.Generic;
using Pandora.Definitions.Interfaces;

namespace Pandora.Definitions.ResourceManager.EventHub.v2018_01_01_preview.NetworkRuleSets
{
    internal class Definition : ApiDefinition
    {
        public string ApiVersion => "2018-01-01-preview";
        public string Name => "NetworkRuleSets";
        public IEnumerable<ApiOperation> Operations => new List<ApiOperation>
        {
            new NamespacesCreateOrUpdateNetworkRuleSet(),
            new NamespacesGetNetworkRuleSet(),
        };
    }
}