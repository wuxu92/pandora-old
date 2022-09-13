package processors

import (
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

func modelDefinitionsMatch(t *testing.T, first map[string]resourcemanager.TerraformSchemaModelDefinition, second map[string]resourcemanager.TerraformSchemaModelDefinition) bool {
	if len(first) != len(second) {
		t.Fatalf("first had %d models but second had %d models", len(first), len(second))
		return false
	}

	for key, firstVal := range first {
		secondVal, ok := second[key]
		if !ok {
			t.Fatalf("key %q was present in first but not second", key)
			return false
		}

		if !fieldDefinitionsMatch(t, firstVal.Fields, secondVal.Fields) {
			t.Fatalf("field definitions didn't match")
			return false
		}
	}
	return true
}

func fieldDefinitionsMatch(t *testing.T, first map[string]resourcemanager.TerraformSchemaFieldDefinition, second map[string]resourcemanager.TerraformSchemaFieldDefinition) bool {
	// we can't use reflect.DeepEqual since there's pointers involved, so we'll do this the old-fashioned way
	if len(first) != len(second) {
		t.Fatalf("first had %d fields but second had %d fields", len(first), len(second))
		return false
	}

	for key, firstVal := range first {
		secondVal, ok := second[key]
		if !ok {
			t.Fatalf("key %q was present in first but not second", key)
			return false
		}

		if firstVal.Computed != secondVal.Computed {
			t.Fatalf("first computed %t != second computed %t", firstVal.Computed, secondVal.Computed)
		}
		if firstVal.ForceNew != secondVal.ForceNew {
			t.Fatalf("first forcenew %t != second forcenew %t", firstVal.ForceNew, secondVal.ForceNew)
		}
		if firstVal.HclName != secondVal.HclName {
			t.Fatalf("first hclName %q != second hclName %q", firstVal.HclName, secondVal.HclName)
		}
		if firstVal.Optional != secondVal.Optional {
			t.Fatalf("first optional %t != second optional %t", firstVal.Optional, secondVal.Optional)
		}
		if firstVal.Required != secondVal.Required {
			t.Fatalf("first required %t != second required %t", firstVal.Required, secondVal.Required)
		}
		if !reflect.DeepEqual(firstVal.Documentation, secondVal.Documentation) {
			t.Fatalf("first documentation %+v != second documentation %+v", firstVal.Documentation, secondVal.Documentation)
		}

		if !objectDefinitionsMatch(t, &firstVal.ObjectDefinition, &secondVal.ObjectDefinition) {
			t.Fatalf("object definitions didn't match")
			return false
		}

		if !mappingsMatch(t, firstVal.Mappings, secondVal.Mappings) {
			t.Fatalf("mappings didn't match")
			return false
		}

		if !validatorsMatch(t, firstVal.Validation, secondVal.Validation) {
			t.Fatalf("validation didn't match")
			return false
		}
	}

	return true
}

func objectDefinitionsMatch(t *testing.T, first *resourcemanager.TerraformSchemaFieldObjectDefinition, second *resourcemanager.TerraformSchemaFieldObjectDefinition) bool {
	if first == nil && second == nil {
		return true
	}
	if first != nil && second == nil {
		t.Fatalf("first was %+v but second was nil", *first)
		return false
	}
	if first == nil && second != nil {
		t.Fatalf("first was nil but second wasn't: %+v", *second)
		return false
	}

	if first.Type != second.Type {
		t.Fatalf("type's didn't match - first %q / second %q", string(first.Type), string(second.Type))
		return false
	}
	firstRefName := valueForNilableString(first.ReferenceName)
	secondRefName := valueForNilableString(second.ReferenceName)
	if firstRefName != secondRefName {
		t.Fatalf("reference name's didn't match - first %q / second %q", firstRefName, secondRefName)
		return false
	}
	if !objectDefinitionsMatch(t, first.NestedObject, second.NestedObject) {
		t.Fatalf("Nested ObjectDefinition's didn't match - first %+v / second %+v", *first.NestedObject, *second.NestedObject)
		return false
	}

	return true
}

func mappingsMatch(t *testing.T, first resourcemanager.TerraformSchemaFieldMappingDefinition, second resourcemanager.TerraformSchemaFieldMappingDefinition) bool {
	if valueForNilableString(first.ResourceIdSegment) != valueForNilableString(second.ResourceIdSegment) {
		t.Fatalf("ResourceIdSegment didn't match - first %q / second %q", valueForNilableString(first.ResourceIdSegment), valueForNilableString(second.ResourceIdSegment))
		return false
	}
	if valueForNilableString(first.SdkPathForCreate) != valueForNilableString(second.SdkPathForCreate) {
		t.Fatalf("SdkPathForCreate didn't match - first %q / second %q", valueForNilableString(first.SdkPathForCreate), valueForNilableString(second.SdkPathForCreate))
		return false
	}
	if valueForNilableString(first.SdkPathForRead) != valueForNilableString(second.SdkPathForRead) {
		t.Fatalf("SdkPathForRead didn't match - first %q / second %q", valueForNilableString(first.SdkPathForRead), valueForNilableString(second.SdkPathForRead))
		return false
	}
	if valueForNilableString(first.SdkPathForUpdate) != valueForNilableString(second.SdkPathForUpdate) {
		t.Fatalf("SdkPathForUpdate didn't match - first %q / second %q", valueForNilableString(first.SdkPathForUpdate), valueForNilableString(second.SdkPathForUpdate))
		return false
	}

	return true
}

func validatorsMatch(t *testing.T, first *resourcemanager.TerraformSchemaValidationDefinition, second *resourcemanager.TerraformSchemaValidationDefinition) bool {
	if first == nil && second == nil {
		return true
	}
	if first != nil && second == nil {
		t.Fatalf("first was %+v but second was nil", *first)
		return false
	}
	if first == nil && second != nil {
		t.Fatalf("first was nil but second wasn't: %+v", *second)
		return false
	}

	if first.Type != second.Type {
		t.Fatalf("type's didn't match - first %q / second %q", string(first.Type), string(second.Type))
		return false
	}
	firstValues := strings.Join(stringifyValues(first.PossibleValues.Values), ", ")
	secondValues := strings.Join(stringifyValues(second.PossibleValues.Values), ", ")
	if !reflect.DeepEqual(firstValues, secondValues) {
		t.Fatalf("possible values didn't match - first %q / second %q", firstValues, secondValues)
		return false
	}

	return true
}

func stringifyValues(input []interface{}) []string {
	output := make([]string, 0)

	if input != nil {
		for _, val := range input {
			v, ok := val.(string)
			if ok {
				output = append(output, v)
			}
		}
	}

	return output
}

func valueForNilableString(input *string) string {
	if input == nil {
		return ""
	}

	return *input
}

func stringPointer(input string) *string {
	return &input
}