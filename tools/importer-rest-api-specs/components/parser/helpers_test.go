package parser

import (
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

func ParseSwaggerFileForTesting(t *testing.T, file string) (*models.AzureApiDefinition, error) {
	parsed, err := load("testdata/", file, hclog.New(hclog.DefaultOptions))
	if err != nil {
		t.Fatalf("loading: %+v", err)
	}

	resourceIds, err := parsed.ParseResourceIds()
	if err != nil {
		t.Fatalf("parsing Resource Ids: %+v", err)
	}

	out, err := parsed.parse("Example", "2020-01-01", *resourceIds)
	if err != nil {
		t.Fatalf("parsing file %q: %+v", file, err)
	}

	return out, nil
}
