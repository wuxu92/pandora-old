package dataapigenerator

import (
	"fmt"

	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

func (s Generator) generateServiceDefinitions(apiVersions []models.AzureApiDefinition) error {
	s.logger.Debug(fmt.Sprintf("Processing Service %q..", s.serviceName))

	excludeList := []string{
		// TODO: presumably we can remove this once https://github.com/hashicorp/pandora/issues/403
		// is resolved
		"ServiceDefinition-GenerationSetting.cs",
	}
	if err := recreateDirectoryExcludingFiles(s.workingDirectoryForService, excludeList, s.logger); err != nil {
		return fmt.Errorf("recreating %q: %+v", s.workingDirectoryForService, err)
	}

	// finally let's output the new Service Definition
	if err := s.generateServiceDefinition(apiVersions); err != nil {
		return fmt.Errorf("generating Service Definition for Namespace %q: %+v", s.namespaceForService, err)
	}

	return nil
}
