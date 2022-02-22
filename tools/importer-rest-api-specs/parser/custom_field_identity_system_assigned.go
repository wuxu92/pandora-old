package parser

import (
	"reflect"
	"strings"

	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

var _ customFieldMatcher = systemAssignedIdentityMatcher{}

type systemAssignedIdentityMatcher struct{}

func (systemAssignedIdentityMatcher) customFieldType() models.CustomFieldType {
	return models.CustomFieldTypeSystemAssignedIdentity
}

func (systemAssignedIdentityMatcher) isMatch(_ models.FieldDetails, definition models.ObjectDefinition, known parseResult) bool {
	if definition.Type != models.ObjectDefinitionReference {
		return false
	}

	// retrieve the model from the reference
	model, ok := known.models[*definition.ReferenceName]
	if !ok {
		return false
	}

	hasMatchingType := false
	hasPrincipalId := false
	hasTenantId := false

	for fieldName, fieldVal := range model.Fields {
		if strings.EqualFold(fieldName, "PrincipalId") {
			hasPrincipalId = true
			continue
		}

		if strings.EqualFold(fieldName, "TenantId") {
			hasTenantId = true
			continue
		}

		if strings.EqualFold(fieldName, "Type") {
			// it's either a FixedValue, in which case they should match
			if fieldVal.FixedValue != nil && strings.EqualFold(*fieldVal.FixedValue, "SystemAssigned") {
				hasMatchingType = true
				continue
			}

			// else it's a constant
			if fieldVal.ObjectDefinition == nil || fieldVal.ObjectDefinition.Type != models.ObjectDefinitionReference {
				continue
			}
			constant, ok := known.constants[*fieldVal.ObjectDefinition.ReferenceName]
			if !ok {
				continue
			}
			expected := map[string]string{
				"SystemAssigned": "SystemAssigned",
			}
			hasMatchingType = validateIdentityConstantValues(constant, expected)
			continue
		}

		// other fields
		return false
	}

	return hasMatchingType && hasPrincipalId && hasTenantId
}

func validateIdentityConstantValues(input models.ConstantDetails, expected map[string]string) bool {
	if input.FieldType != models.StringConstant {
		return false
	}

	// we can't guarantee the casing on these, so we should parse this insensitively since it'll be swapped
	// out anyway
	actual := make(map[string]string, 0)
	for k, v := range input.Values {
		// Some Identity Constants define 'None', others don't - we're not particularly concerned
		// for this comparison however, so remove them to normalize those values
		if strings.EqualFold(k, "None") {
			continue
		}

		actual[strings.ToLower(k)] = strings.ToLower(v)
	}

	normalizedExpected := make(map[string]string, 0)
	for k, v := range expected {
		normalizedExpected[strings.ToLower(k)] = strings.ToLower(v)
	}

	return reflect.DeepEqual(actual, normalizedExpected)
}
