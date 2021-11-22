package resourceids

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"

	"github.com/hashicorp/pandora/tools/generator-terraform/internal"
)

// TODO: switch to using gotemplate

type dataSourceGenerator struct {
	data                  DataSourceForResourceIdGenerator
	fieldNamesToSnakeCase map[string]string
}

func (g dataSourceGenerator) structDefinition() string {
	conflictsWithKeys := make([]string, 0)
	schemaLines := make([]string, 0)
	structLines := []string{
		"AzureResourceId string `tfschema:\"azure_resource_id\"`",
	}
	for fieldName, snakeCasedFieldName := range g.fieldNamesToSnakeCase {
		conflictsWithKeys = append(conflictsWithKeys, fmt.Sprintf("%q", snakeCasedFieldName))
		// TODO: this also needs RequiredWith
		// TODO: if it's a constant we need have the values validated here too
		schemaLines = append(schemaLines, fmt.Sprintf(`
		%q: {
			Type: pluginsdk.TypeString,
			Optional: true,
			Computed: true,
			ConflictsWith: []string{ "azure_resource_id" },
		},`, snakeCasedFieldName))
		structLines = append(structLines, fmt.Sprintf("%s string `tfschema:%q`", strings.Title(fieldName), snakeCasedFieldName))
	}
	// TODO: validating the resource type
	schemaLines = append(schemaLines, fmt.Sprintf(`
		"azure_resource_id": {
			Type: 			pluginsdk.TypeString,
			Optional: 		true,
			Computed: 		true,
			ConflictsWith:  []string{ %s },
		},
	`, strings.Join(conflictsWithKeys, ",")))

	resourceType := internal.SnakeCase(fmt.Sprintf("%s%s", g.data.ProviderName, g.data.ResourceIdName))

	return fmt.Sprintf(`
var _ sdk.DataSource = %[1]sDataSource{}

type %[1]sDataSource struct {}

type %[1]sDataSourceModel struct {
	%[2]s
}

func (ds %[1]sDataSource) Arguments() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		%[4]s
	}
}

func (ds %[1]sDataSource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func (ds %[1]sDataSource) ModelObject() interface{} {
	return %[1]sDataSourceModel{}
}

func (ds %[1]sDataSource) ResourceType() string {
	return %[3]q
}
`, g.data.ResourceIdName, strings.Join(structLines, "\n"), resourceType, strings.Join(schemaLines, "\n"))
}

func (g dataSourceGenerator) readMethod() (*string, error) {
	resourceIdWithoutSuffix := strings.TrimSuffix(g.data.ResourceIdName, "Id")
	arguments := make([]string, 0)
	assignments := make([]string, 0)

	for _, segment := range g.data.ResourceIdDetails.Segments {
		if _, userConfigurable := userConfiguredSegments[segment.Type]; !userConfigurable {
			continue
		}

		if segment.Type == resourcemanager.ConstantSegment {
			if segment.ConstantReference == nil {
				return nil, fmt.Errorf("constant segment %q is missing a constant reference", segment.Name)
			}

			arguments = append(arguments, fmt.Sprintf("%s(model.%s)", *segment.ConstantReference, strings.Title(segment.Name)))
		} else {
			arguments = append(arguments, fmt.Sprintf("model.%s", strings.Title(segment.Name)))
		}

		assignments = append(assignments, fmt.Sprintf("model.%[1]s = id.%[1]s", strings.Title(segment.Name)))
	}

	out := fmt.Sprintf(`
func (ds %[1]sDataSource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var model %[1]sDataSourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding model: %%+v", err)
			}

			var id somesdk.%[1]s
			if model.AzureResourceId != "" {
				metadata.Logger.Infof("[DEBUG] Parsing the Resource ID..")
				resourceId, err := somesdk.Parse%[2]sID(model.AzureResourceId)
				if err != nil {
					return fmt.Errorf("parsing 'azure_resource_id' %%q: %%+v", model.AzureResourceId, err)
				}

				id = *resourceId
			} else {
				metadata.Logger.Infof("[DEBUG] Building the Resource ID..")
				id = somesdk.New%[2]sID(%[3]s)
			}

			model.AzureResourceId = id.String()
			model.ClusterName = id.ClusterName

			return metadata.Encode(&model)
		},
		Timeout: 5 * time.Minute,
	}
}
`, g.data.ResourceIdName, resourceIdWithoutSuffix, strings.Join(arguments, ", "), strings.Join(assignments, "\n"))
	return &out, nil
}

func (g DataSourceForResourceIdGenerator) codeForDataSource() (*string, error) {
	if len(g.ResourceIdDetails.Segments) == 0 {
		return nil, fmt.Errorf("resource id has no segments")
	}

	fieldNamesToSnakeCase := g.fieldNameToSnakeCaseMap()
	gen := dataSourceGenerator{
		data:                  g,
		fieldNamesToSnakeCase: fieldNamesToSnakeCase,
	}

	structCode := gen.structDefinition()
	readMethod, err := gen.readMethod()
	if err != nil {
		return nil, fmt.Errorf("generating read method: %+v", err)
	}

	out := fmt.Sprintf(`package %[1]s

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

%[3]s
%[4]s
`, g.PackageName, g.ResourceIdName, structCode, *readMethod)
	log.Print(out)
	return &out, nil
}
