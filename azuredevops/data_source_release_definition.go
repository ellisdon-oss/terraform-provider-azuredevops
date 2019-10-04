package azuredevops

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"github.com/pkg/errors"
)

func dataSourceReleaseDefinition() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceReleaseDefinitionRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceReleaseDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)
	name := d.Get("name").(string)

	releaseDefs, err := releaseClient.GetReleaseDefinitions(config.Context, release.GetReleaseDefinitionsArgs{
		Project:    &projectID,
		SearchText: &name,
	})

	if err != nil {
		return err
	}

	if len((*releaseDefs).Value) == 0 {
		return errors.New("Release Definition Not Found")
	}

	releaseDef := (*releaseDefs).Value[0]

	d.SetId(fmt.Sprint(*releaseDef.Id))

	return nil
}
