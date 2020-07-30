package azuredevops

import (
  "fmt"
  "strconv"
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/ellisdon-oss/terraform-provider-azuredevops/azuredevops/helper"
  "github.com/microsoft/azure-devops-go-api/azuredevops/release"
)

func dataSourceReleaseEnvironment() *schema.Resource {
  return &schema.Resource{
    Read: dataSourceReleaseEnvironmentRead,

    Schema: map[string]*schema.Schema{
      "project_id": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "definition_id": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
        ForceNew: true,
      },
      "stage_name": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "environment": &schema.Schema{
        Type:     schema.TypeList,
        MinItems: 1,
        Computed: true,
        Elem: &schema.Resource{
          Schema: helper.EnvironmentSchema(),
        },
      },
    },
	}
}

func dataSourceReleaseEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)
	definition_id, _ := strconv.ParseInt(d.Get("definition_id").(string), 10, 32)

	defID := int(definition_id)

	releaseDef, err := releaseClient.GetReleaseDefinition(config.Context, release.GetReleaseDefinitionArgs{
		Project:    &projectID,
    DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

	result := []interface{}{}

  for _, env := range *releaseDef.Environments {
    if *env.Name == d.Get("stage_name").(string) {
			result = append(result, convertEnvToMapDirect(env))
		}
	}

	d.Set("environment", result)
  
	d.SetId(fmt.Sprintf("%d-%s",*releaseDef.Id, d.Get("stage_name").(string)))

	return nil
}
