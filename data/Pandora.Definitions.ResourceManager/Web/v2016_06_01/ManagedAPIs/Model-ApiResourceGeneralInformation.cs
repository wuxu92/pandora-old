using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Web.v2016_06_01.ManagedAPIs;


internal class ApiResourceGeneralInformationModel
{
    [JsonPropertyName("description")]
    public string? Description { get; set; }

    [JsonPropertyName("displayName")]
    public string? DisplayName { get; set; }

    [JsonPropertyName("iconUrl")]
    public string? IconUrl { get; set; }

    [JsonPropertyName("releaseTag")]
    public string? ReleaseTag { get; set; }

    [JsonPropertyName("termsOfUseUrl")]
    public string? TermsOfUseUrl { get; set; }
}
