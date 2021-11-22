package resourceids

import (
	"github.com/hashicorp/pandora/tools/generator-terraform/internal"
	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

var userConfiguredSegments = map[resourcemanager.ResourceIdSegmentType]struct{}{
	resourcemanager.ConstantSegment:       {},
	resourcemanager.ResourceGroupSegment:  {},
	resourcemanager.SubscriptionIdSegment: {},
	resourcemanager.UserSpecifiedSegment:  {},
}

func (g DataSourceForResourceIdGenerator) fieldNameToSnakeCaseMap() map[string]string {
	output := make(map[string]string, 0)

	for _, segment := range g.ResourceIdDetails.Segments {
		if _, userConfigurable := userConfiguredSegments[segment.Type]; !userConfigurable {
			continue
		}

		output[segment.Name] = internal.SnakeCase(segment.Name)
	}

	return output
}
