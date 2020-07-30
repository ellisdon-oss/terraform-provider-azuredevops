package azuredevops

import (
  "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
  "errors"
)

func dataSourceReleaseDefinitionEnvironments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceReleaseDefinitionEnvironmentsRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"environments": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
        Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceReleaseDefinitionEnvironmentsRead(d *schema.ResourceData, meta interface{}) error {
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
    Expand: &release.ReleaseDefinitionExpandsValues.Environments,
	})

	if err != nil {
		return err
	}

	if len((*releaseDefs).Value) == 0 {
		return errors.New("Release Definition Not Found")
	}

	releaseDef := (*releaseDefs).Value[0]

  environments := make([]string, 0)

  for _, env := range *releaseDef.Environments {
    environments = append(environments, *env.Name)
  }

	d.SetId(fmt.Sprintf("%d-%s", *releaseDef.Id, name))

  d.Set("environments", environments)
  
	return nil
}
