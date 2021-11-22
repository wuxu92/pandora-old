package resourceids

import (
	"testing"

	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

func TestCodeForDataSourceNoSegments(t *testing.T) {
	input := DataSourceForResourceIdGenerator{
		ProviderName:   "myprovider",
		PackageName:    "somepackage",
		Debug:          true,
		ResourceIdName: "IdWithNoSegments",
		ResourceIdDetails: resourcemanager.ResourceIdDefinition{
			CommonAlias:   nil,
			ConstantNames: []string{},
			Id:            "/",
			Segments:      []resourcemanager.ResourceIdSegment{},
		},
	}
	actual, err := input.codeForDataSource()
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	if actual != nil {
		t.Fatalf("expected no result but got one: %+v", *actual)
	}
}
