package azuredevops

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/identity"
	"github.com/pkg/errors"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"is_email": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
		},
	}
}

func dataSourceUserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	var searchFilter string
	if d.Get("is_email").(bool) {
		searchFilter = "AccountName"
	} else {
		searchFilter = "DisplayName"
	}
	filterValue := d.Get("name").(string)

	identityClient, err := identity.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	res, err := identityClient.ReadIdentities(config.Context, identity.ReadIdentitiesArgs{
		SearchFilter: &searchFilter,
		FilterValue:  &filterValue,
	})

	if err != nil {
		return err
	}

	if len(*res) == 0 {
		return errors.New("User Not Found")
	}

	for _, v := range *res {
		identityType := v.Properties.(map[string]interface{})["SchemaClassName"].(map[string]interface{})["$value"].(string)

		if identityType != "User" {
			return errors.New("Name is not a user")
		}

		d.SetId(v.Id.String())
	}
	return nil
}
