using Pandora.Definitions.Attributes;
using Pandora.Definitions.Interfaces;
using Pandora.Definitions.Operations;
using System.Collections.Generic;
using System.Net;

namespace Pandora.Definitions.ResourceManager.EventHub.v2021_01_01_preview.AuthorizationRulesDisasterRecoveryConfigs
{
	internal class DisasterRecoveryConfigsListKeys : PostOperation
	{
		public override IEnumerable<HttpStatusCode> ExpectedStatusCodes()
		{
			return new List<HttpStatusCode>
			{
				HttpStatusCode.OK,
			};
		}

		public override object? RequestObject()
		{
			return null;
		}

		public override ResourceID? ResourceId()
		{
			return new AuthorizationRuleId();
		}

		public override object? ResponseObject()
		{
			return new AccessKeys();
		}

		public override string? UriSuffix()
		{
			return "/listKeys";
		}
	}
}