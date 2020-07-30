package azuredevops

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
)

func dataSourceReleaseDefinitions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceReleaseDefinitionsRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"definitions": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
        Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceReleaseDefinitionsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)

	releaseDefs, err := releaseClient.GetReleaseDefinitions(config.Context, release.GetReleaseDefinitionsArgs{
		Project:    &projectID,
	})

	if err != nil {
		return err
	}

  definitions := make([]string, 0)

	d.SetId(projectID)

  for _, definition := range (*releaseDefs).Value {
    definitions = append(definitions, *definition.Name)
  }

  d.Set("definitions", definitions)
  
	return nil
}
