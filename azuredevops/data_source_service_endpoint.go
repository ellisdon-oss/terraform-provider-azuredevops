package azuredevops

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/serviceendpoint"
	"github.com/pkg/errors"
)

func dataSourceServiceEndpoint() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServiceEndpointRead,

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceServiceEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	endpointClient, err := serviceendpoint.NewClient(config.Context, config.Connection)
	projectID := d.Get("project_id").(string)

	name := []string{d.Get("name").(string)}
	endpointType := d.Get("type").(string)
	detailInclude := true

	serviceEndpoints, err := endpointClient.GetServiceEndpointsByNames(config.Context, serviceendpoint.GetServiceEndpointsByNamesArgs{
		Project:        &projectID,
		EndpointNames:  &name,
		Type:           &endpointType,
		IncludeDetails: &detailInclude,
	})

	if err != nil {
		return err
	}

	if len(*serviceEndpoints) == 0 {
		return errors.New("Service Endpoint Not Found")
	}

	serviceEndpoint := (*serviceEndpoints)[0]

	d.Set("owner", serviceEndpoint.Owner)
	d.Set("url", serviceEndpoint.Url)
	d.SetId(serviceEndpoint.Id.String())

	return nil
}
