package schema

import (
	"log"

	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

func (b Builder) identifyFieldsWithinPropertiesBlock(input operationPayloads) (*map[string]FieldDefinition, error) {
	allFields := make(map[string]struct{}, 0)
	for _, model := range input.createReadUpdatePayloadsProperties(b.models) {
		for k, v := range model.Fields {
			if fieldShouldBeIgnored(k, v, b.constants) {
				continue
			}

			allFields[k] = struct{}{}
		}
	}

	// find the model for the `properties` field within the read response
	readPropertiesModel := input.getPropertiesModelWithinModel(input.readPayload, b.models)
	createPropertiesModel := input.getPropertiesModelWithinModel(input.createPayload, b.models)
	var updatePropertiesModel *resourcemanager.ModelDetails
	if input.updatePayload != nil {
		updatePropertiesModel = input.getPropertiesModelWithinModel(*input.updatePayload, b.models)
	}

	out := make(map[string]FieldDefinition, 0)
	if readPropertiesModel != nil {
		for k := range allFields {
			var readField *resourcemanager.FieldDetails
			hasRead := false
			if readPropertiesModel != nil {
				readField, hasRead = getField(*readPropertiesModel, k)
			}

			var createField *resourcemanager.FieldDetails
			hasCreate := false
			if createPropertiesModel != nil {
				createField, hasCreate = getField(*createPropertiesModel, k)
			}

			var updateField *resourcemanager.FieldDetails
			hasUpdate := false
			if updatePropertiesModel != nil {
				updateField, hasUpdate = getField(*updatePropertiesModel, k)
			}

			// based on this information
			isComputed := false
			isForceNew := false
			isRequired := false
			isOptional := false
			isWriteOnly := false

			if !hasCreate && !hasUpdate && hasRead {
				isComputed = true
			}
			if hasCreate || hasUpdate {
				if !hasRead {
					isWriteOnly = true
					isForceNew = hasUpdate && !updateField.ForceNew
				} else if hasCreate {
					isRequired = createField.Required
					isOptional = createField.Optional
					isForceNew = hasUpdate
				} else if hasUpdate {
					isRequired = updateField.Required
					isOptional = updateField.Optional
					isForceNew = updateField.ForceNew
				}
			}

			typedModelName := ""

			if hasRead {
				typedModelName = b.determineNameForSchemaField(*readPropertiesModel, k)
			} else if hasCreate {
				typedModelName = b.determineNameForSchemaField(*createPropertiesModel, k)
			} else if hasUpdate {
				typedModelName = b.determineNameForSchemaField(*updatePropertiesModel, k)
			}

			schemaFieldName := convertToSnakeCase(typedModelName)
			log.Printf("[DEBUG] Properties Field %q would be output as %q / %q", k, typedModelName, schemaFieldName)

			definition := FieldDefinition{
				Required:  isRequired,
				ForceNew:  isForceNew,
				Optional:  isOptional,
				Computed:  isComputed,
				WriteOnly: isWriteOnly,
				// TODO: also need to add the mappings & any validation
			}

			if hasRead {
				definition.Definition = readField.ObjectDefinition
			} else if hasCreate {
				definition.Definition = createField.ObjectDefinition
			} else if hasUpdate {
				definition.Definition = updateField.ObjectDefinition
			}

			out[schemaFieldName] = definition
		}
	}

	return &out, nil
}
