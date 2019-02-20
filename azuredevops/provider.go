package azuredevops

import (
	"os"

	"context"
	//	"github.com/ellisdon/azuredevops-go/core"
	//	"github.com/ellisdon/azuredevops-go/operations"
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
)

// Provider returns a schema.Provider for Example.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("AZUREDEVOPS_USERNAME"),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("AZUREDEVOPS_PASSWORD"),
			},
			"organization": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("AZUREDEVOPS_ORG"),
			},
			"api_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "5.1-preview",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"azuredevops_project":            resourceProject(),
			"azuredevops_build_definition":   resourceBuildDefinition(),
			"azuredevops_release_definition": resourceReleaseDefinition(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"azuredevops_project":           dataSourceProject(),
			"azuredevops_service_endpoint":  dataSourceServiceEndpoint(),
			"azuredevops_source_repository": dataSourceSourceRepository(),
			"azuredevops_workflow_task":     dataSourceWorkflowTask(),
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
	auth := context.WithValue(context.Background(), azuredevops.ContextBasicAuth, azuredevops.BasicAuth{
		UserName: d.Get("username").(string),
		Password: d.Get("password").(string),
	})

	cfg := azuredevops.NewConfiguration()
	client := azuredevops.NewAPIClient(cfg)
	config := Config{
		Client:       client,
		Organization: d.Get("organization").(string),
		Context:      auth,
		ApiVersion:   d.Get("api_version").(string),
	}

	log.Printf("[INFO] AzureDevOps Client configured for use")

	return &config, nil
}
