package azuredevops

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/identity"
	"github.com/pkg/errors"
)

func dataSourceGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGroupRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	var searchFilter string
	searchFilter = "DisplayName"
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
		return errors.New("Group Not Found")
	}

	for _, v := range *res {
		identityType := v.Properties.(map[string]interface{})["SchemaClassName"].(map[string]interface{})["$value"].(string)

		if identityType != "Group" {
			return errors.New("Name is not a group")
		}

		d.SetId(v.Id.String())
	}
	return nil
}
