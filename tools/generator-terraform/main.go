package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/pandora/tools/generator-terraform/generator/datasources/resourceids"

	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

func main() {
	log.Printf("[DEBUG] Started")

	input := runInput{
		endpoint: "http://localhost:5000",
		debug:    true,

		dataSourcesForResourceIds: true,
	}

	if err := run(context.TODO(), input); err != nil {
		log.Fatalf("running: %+v", err)
		return
	}

}

type runInput struct {
	endpoint string
	debug    bool

	// dataSourcesForResourceIds specifies whether data sources should be generated for each Resource ID
	dataSourcesForResourceIds bool
}

func (i runInput) log(format string, v ...interface{}) {
	if !i.debug {
		return
	}

	log.Printf(format, v...)
}

func (i runInput) generateDataSourcesForResourceIds(service string, version string, resource string, ids map[string]resourcemanager.ResourceIdDefinition) error {
	if !i.dataSourcesForResourceIds {
		i.log("[DEBUG] Skipping generating Data Sources for these Resource ID's since this is disabled")
		return nil
	}

	for name, resourceId := range ids {
		generator := resourceids.DataSourceForResourceIdGenerator{
			ProviderName:      "myprovider",
			PackageName:       "somepackage",
			WorkingDirectory:  "./generated",
			Debug:             i.debug,
			ResourceIdName:    name,
			ResourceIdDetails: resourceId,
		}
		if err := generator.Generate(); err != nil {
			return fmt.Errorf("generating Data Source for Resource ID %q (Service %q / Version %q / Resource %q): %+v", name, service, version, resource, err)
		}
	}

	return nil
}

func run(context context.Context, input runInput) error {
	client := resourcemanager.NewClient(input.endpoint)

	input.log("[DEBUG] ")
	services, err := client.Services().Get()
	if err != nil {
		return fmt.Errorf("loading services: %+v", err)
	}

	for serviceName, service := range *services {
		log.Printf("[DEBUG] Procesing Service %q..", serviceName)

		serviceDetails, err := client.ServiceDetails().Get(service)
		if err != nil {
			return fmt.Errorf("loading details for service %q: %+v", serviceName, err)
		}

		for versionNumber, versionSummary := range serviceDetails.Versions {
			if !versionSummary.Generate {
				input.log("[DEBUG] Skipping Service %q / Version %q since it's marked as do not generate", serviceName, versionNumber)
				continue
			}

			versionDetails, err := client.ServiceVersion().Get(versionSummary)
			if err != nil {
				return fmt.Errorf("loading details for version %q of service %q: %+v", versionNumber, serviceName, err)
			}

			for resourceName, resourceSummary := range versionDetails.Resources {
				schema, err := client.ApiSchema().Get(resourceSummary)
				if err != nil {
					return fmt.Errorf("loading schema for resource %q / version %q / service %q: %+v", resourceName, versionNumber, serviceName, err)
				}

				input.log("[DEBUG] Generating Data Sources for Resource ID's  (Service %q / Version %q / Resource %q)..", serviceName, versionNumber, resourceName)
				if err := input.generateDataSourcesForResourceIds(serviceName, versionNumber, resourceName, schema.ResourceIds); err != nil {
					return fmt.Errorf("generating Data Sources for Resource ID's (Service %q / Version %q / Resource %q): %v", serviceName, versionNumber, resourceName, err)
				}
				input.log("[DEBUG] Generated Data Sources for Resource ID's  (Service %q / Version %q / Resource %q)", serviceName, versionNumber, resourceName)
			}
		}
	}

	return nil
}
