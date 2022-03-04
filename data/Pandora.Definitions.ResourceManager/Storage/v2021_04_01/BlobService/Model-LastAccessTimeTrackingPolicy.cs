using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;

namespace Pandora.Definitions.ResourceManager.Storage.v2021_04_01.BlobService;


internal class LastAccessTimeTrackingPolicyModel
{
    [JsonPropertyName("blobType")]
    public List<string>? BlobType { get; set; }

    [JsonPropertyName("enable")]
    [Required]
    public bool Enable { get; set; }

    [JsonPropertyName("name")]
    public NameConstant? Name { get; set; }

    [JsonPropertyName("trackingGranularityInDays")]
    public int? TrackingGranularityInDays { get; set; }
}