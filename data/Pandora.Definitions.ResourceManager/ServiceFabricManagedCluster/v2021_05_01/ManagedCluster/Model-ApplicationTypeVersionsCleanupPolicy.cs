using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;

namespace Pandora.Definitions.ResourceManager.ServiceFabricManagedCluster.v2021_05_01.ManagedCluster
{

    internal class ApplicationTypeVersionsCleanupPolicyModel
    {
        [JsonPropertyName("maxUnusedVersionsToKeep")]
        [Required]
        public int MaxUnusedVersionsToKeep { get; set; }
    }
}