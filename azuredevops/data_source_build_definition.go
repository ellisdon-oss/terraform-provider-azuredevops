package azuredevops

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
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

	projId := d.Get("project_id").(string)
	name := d.Get("name").(string)
	definitions, _, err := config.Client.DefinitionsApi.GetDefinitions(config.Context, config.Organization, projId, config.ApiVersion, &azuredevops.GetDefinitionsOpts{
		Name: optional.NewString(name),
	})

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	if definitions.Count == 0 {
		return errors.New("Build Definition Not Found")
	}

	definition := definitions.Value[0]

	d.SetId(fmt.Sprint(definition.Id))

	return nil
}
