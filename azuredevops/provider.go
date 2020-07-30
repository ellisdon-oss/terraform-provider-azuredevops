package azuredevops

import (
	"os"

	"context"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"log"
)

// Provider returns a schema.Provider for Example.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("AZUREDEVOPS_TOKEN"),
			},
			"organization_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("AZUREDEVOPS_ORG_URL"),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"azuredevops_project":             resourceProject(),
			"azuredevops_build_definition":    resourceBuildDefinition(),
			"azuredevops_release_definition":  resourceReleaseDefinition(),
			"azuredevops_release_environment": resourceReleaseEnvironment(),
			"azuredevops_release_task":        resourceReleaseTask(),
			"azuredevops_release_tasks":       resourceReleaseTasks(),
			"azuredevops_release_variables":   resourceReleaseVariables(),
			"azuredevops_service_endpoint":    resourceServiceEndpoint(),
			"azuredevops_service_hook":        resourceServiceHook(),
			"azuredevops_variable_group":      resourceVariableGroup(),
			"azuredevops_task_group":          resourceTaskGroup(),
			"azuredevops_extension":           resourceExtension(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"azuredevops_project":            dataSourceProject(),
			"azuredevops_service_endpoint":   dataSourceServiceEndpoint(),
			"azuredevops_source_repository":  dataSourceSourceRepository(),
			"azuredevops_workflow_task":      dataSourceWorkflowTask(),
			"azuredevops_group":              dataSourceGroup(),
			"azuredevops_user":               dataSourceUser(),
			"azuredevops_build_definition":   dataSourceBuildDefinition(),
			"azuredevops_release_definition": dataSourceReleaseDefinition(),
			"azuredevops_release_definitions": dataSourceReleaseDefinitions(),
			"azuredevops_release_definition_environments": dataSourceReleaseDefinitionEnvironments(),
			"azuredevops_release_tasks":      dataSourceReleaseTasks(),
			"azuredevops_release_stage_variables": dataSourceReleaseStageVariables(),
			"azuredevops_release_environment": dataSourceReleaseEnvironment(),
			"azuredevops_agent_queue":        dataSourceAgentQueue(),
			"azuredevops_task_group":         dataSourceTaskGroup(),
			"azuredevops_variable_group":     dataSourceVariableGroup(),
			"azuredevops_variable_groups":    dataSourceVariableGroups(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func envDefaultFunc(k string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		if v := os.Getenv(k); v != "" {
			return v, nil
		}

		return nil, nil
	}
}

func envDefaultFuncAllowMissing(k string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		v := os.Getenv(k)
		return v, nil
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	connection := azuredevops.NewPatConnection(d.Get("organization_url").(string), d.Get("token").(string))

	ctx := context.WithValue(context.Background(), "Organization", d.Get("organization_url").(string))
	config := Config{
		Context:      ctx,
		Connection:   connection,
		Organization: d.Get("organization_url").(string),
	}

	log.Printf("[INFO] AzureDevOps Client configured for use")

	return &config, nil
}
