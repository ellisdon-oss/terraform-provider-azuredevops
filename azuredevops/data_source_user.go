package azuredevops

import (
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/hashicorp/terraform/helper/validation"
	"github.com/pkg/errors"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceUserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	query := azuredevops.QueryIdentity{
		Query:           d.Get("display_name").(string),
		IdentityTypes:   []string{"user"},
		OperationScopes: []string{"ims"},
	}

	result, _, err := config.Client.IdentitiesApi.FindIdentity(config.Context, config.Organization, config.ApiVersion, query)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	if len(result.Results) == 0 {
		return errors.New("User Not Found")
	}

	d.SetId(result.Results[0].Identities[0].LocalId)
	//	d.Set("abbreviation", project.Abbreviation)
	//	d.Set("default_team_image_url", project.DefaultTeamImageUrl)
	//	d.Set("description", project.Description)
	//	d.Set("last_update_time", project.LastUpdateTime)
	//	d.Set("revision", project.Revision)
	//	d.Set("state", project.State)
	//	d.Set("url", project.Url)
	//	d.Set("visibility", project.Visibility)
	//
	//	d.SetId(project.Id)

	return nil
}
