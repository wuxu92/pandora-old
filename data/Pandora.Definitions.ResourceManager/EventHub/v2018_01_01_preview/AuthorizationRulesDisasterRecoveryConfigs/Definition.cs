using System.Collections.Generic;
using Pandora.Definitions.Interfaces;

namespace Pandora.Definitions.ResourceManager.EventHub.v2018_01_01_preview.AuthorizationRulesDisasterRecoveryConfigs
{
    internal class Definition : ApiDefinition
    {
        public string ApiVersion => "2018-01-01-preview";
        public string Name => "AuthorizationRulesDisasterRecoveryConfigs";
        public IEnumerable<ApiOperation> Operations => new List<ApiOperation>
        {
            new DisasterRecoveryConfigsGetAuthorizationRule(),
            new DisasterRecoveryConfigsListAuthorizationRules(),
            new DisasterRecoveryConfigsListKeys(),
        };
    }
}