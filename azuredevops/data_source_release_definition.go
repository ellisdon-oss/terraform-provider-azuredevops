package azuredevops

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
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

	projId := d.Get("project_id").(string)
	name := d.Get("name").(string)
	releaseDefs, _, err := config.Client.ReleaseDefinitionsApi.GetReleaseDefinitions(config.Context, config.Organization, projId, config.ApiVersion, &azuredevops.GetReleaseDefinitionsOpts{
		SearchText: optional.NewString(name),
	})

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	if releaseDefs.Count == 0 {
		return errors.New("Release Definition Not Found")
	}

	releaseDef := releaseDefs.ReleaseDefinitions[0]

	d.SetId(fmt.Sprint(releaseDef.Id))

	return nil
}
