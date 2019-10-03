package azuredevops

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
	"github.com/pkg/errors"
)

func dataSourceBuildDefinition() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBuildDefinitionRead,

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

func dataSourceBuildDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	projectID := d.Get("project_id").(string)
	name := d.Get("name").(string)

	buildClient, err := build.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	definitions, err := buildClient.GetDefinitions(config.Context, build.GetDefinitionsArgs{
		Name:    &name,
		Project: &projectID,
	})

	if err != nil {
		return err
	}

	if len(definitions.Value) == 0 {
		return errors.New("Build Definition Not Found")
	}

	definition := definitions.Value[0]

	d.SetId(fmt.Sprint(*definition.Id))

	return nil
}
