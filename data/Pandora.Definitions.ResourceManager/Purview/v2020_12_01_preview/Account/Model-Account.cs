using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;

namespace Pandora.Definitions.ResourceManager.Purview.v2020_12_01_preview.Account
{

    internal class AccountModel
    {
        [JsonPropertyName("id")]
        public string? Id { get; set; }

        [JsonPropertyName("identity")]
        public CustomTypes.SystemAssignedIdentity? Identity { get; set; }

        [JsonPropertyName("location")]
        public CustomTypes.Location? Location { get; set; }

        [JsonPropertyName("name")]
        public string? Name { get; set; }

        [JsonPropertyName("properties")]
        public AccountPropertiesModel? Properties { get; set; }

        [JsonPropertyName("sku")]
        public AccountSkuModel? Sku { get; set; }

        [JsonPropertyName("tags")]
        public CustomTypes.Tags? Tags { get; set; }

        [JsonPropertyName("type")]
        public string? Type { get; set; }
    }
}