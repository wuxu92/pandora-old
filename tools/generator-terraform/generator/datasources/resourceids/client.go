package resourceids

import (
	"fmt"
	"path"

	"github.com/hashicorp/pandora/tools/generator-terraform/internal"
	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

type DataSourceForResourceIdGenerator struct {
	ProviderName     string
	PackageName      string
	WorkingDirectory string
	Debug            bool

	ResourceIdName    string
	ResourceIdDetails resourcemanager.ResourceIdDefinition
}

func (g DataSourceForResourceIdGenerator) Generate() error {
	// NOTE: for now everything gets generated into a single folder, fix that in time

	g.log("[DEBUG] Generating a Data Source for the Resource ID %q..", g.ResourceIdName)
	snakeCasedName := internal.SnakeCase(g.ResourceIdName)

	// if there's no segments which can be set - don't generate anything
	// an aside: this is a bug which should be caught elsewhere, but extra validation isn't a bad thing
	segment := g.fieldNameToSnakeCaseMap()
	if len(segment) == 0 {
		g.log("[DEBUG] No user configurable segments found - skipping generating data source")
		return nil
	}

	// Data Source
	dataSourceFilePath := path.Join(g.WorkingDirectory, fmt.Sprintf("data_source_%s.gen.go", snakeCasedName))
	g.log("[DEBUG] Generating Data Source (%q)..", dataSourceFilePath)
	code, err := g.codeForDataSource()
	if err != nil {
		return fmt.Errorf("generating code for data source %q: %+v", g.ResourceIdName, err)
	}
	if err := internal.WriteToPathRecreatingIfNeeded(dataSourceFilePath, *code); err != nil {
		return fmt.Errorf("writing code for data source %q: %+v", g.ResourceIdName, err)
	}
	g.log("[DEBUG] Generated Data Source.")

	// Acceptance Tests
	acceptanceTestsFilePath := path.Join(g.WorkingDirectory, fmt.Sprintf("data_source_%s_test.gen.go", snakeCasedName))
	g.log("[DEBUG] Generating Acceptance Tests for Data Source (%q)..", acceptanceTestsFilePath)
	code, err = g.codeForDataSourceTests()
	if err != nil {
		return fmt.Errorf("generating acceptance test code for data source %q: %+v", g.ResourceIdName, err)
	}
	if err := internal.WriteToPathRecreatingIfNeeded(acceptanceTestsFilePath, *code); err != nil {
		return fmt.Errorf("writing acceptance test code for data source %q: %+v", g.ResourceIdName, err)
	}
	g.log("[DEBUG] Generated Documentation for Data Source.")

	// Docs
	docsFilePath := path.Join(g.WorkingDirectory, fmt.Sprintf("%s.html.markdown", snakeCasedName))
	g.log("[DEBUG] Generating Documentation for Data Source (%q)..", acceptanceTestsFilePath)
	code, err = g.codeForDataSourceTests()
	if err != nil {
		return fmt.Errorf("generating documentation code for data source %q: %+v", g.ResourceIdName, err)
	}
	if err := internal.WriteToPathRecreatingIfNeeded(docsFilePath, *code); err != nil {
		return fmt.Errorf("writing docs code for data source %q: %+v", g.ResourceIdName, err)
	}
	g.log("[DEBUG] Generated Documentation for Data Source.")

	return nil
}
