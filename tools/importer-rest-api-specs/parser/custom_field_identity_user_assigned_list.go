package parser

import (
	"strings"

	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

var _ customFieldMatcher = userAssignedIdentityListMatcher{}

type userAssignedIdentityListMatcher struct{}

func (userAssignedIdentityListMatcher) customFieldType() models.CustomFieldType {
	return models.CustomFieldTypeUserAssignedIdentityList
}

func (userAssignedIdentityListMatcher) isMatch(_ models.FieldDetails, definition models.ObjectDefinition, known parseResult) bool {
	if definition.Type != models.ObjectDefinitionReference {
		return false
	}

	// retrieve the model from the reference
	model, ok := known.models[*definition.ReferenceName]
	if !ok {
		return false
	}

	hasUserAssignedIdentities := false

	for fieldName, fieldVal := range model.Fields {
		if strings.EqualFold(fieldName, "UserAssignedIdentities") {
			// this should be a List of Strings
			if fieldVal.ObjectDefinition == nil || fieldVal.ObjectDefinition.Type != models.ObjectDefinitionList {
				continue
			}
			if fieldVal.ObjectDefinition.NestedItem == nil || fieldVal.ObjectDefinition.NestedItem.Type != models.ObjectDefinitionString {
				continue
			}

			hasUserAssignedIdentities = true
			continue
		}

		// Type is an optional check due to some badly defined Swaggers
		// https://github.com/Azure/azure-rest-api-specs/blob/c803720c6bcfcb0fcf4c97f3463ec33a18f9e55c/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabricManagedClusters/stable/2021-05-01/nodetype.json#L763
		// as such we're only concerned if it's defined and doesn't match
		if strings.EqualFold(fieldName, "Type") {
			// it's either a FixedValue, in which case the values should match else we fail
			if fieldVal.FixedValue != nil && !strings.EqualFold(*fieldVal.FixedValue, "UserAssigned") {
				return false
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
				"UserAssigned": "UserAssigned",
			}
			if !validateIdentityConstantValues(constant, expected) {
				return false
			}
			continue
		}

		// other fields
		return false
	}

	return hasUserAssignedIdentities
}
