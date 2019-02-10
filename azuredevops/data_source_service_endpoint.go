package azuredevops

import (
	"github.com/antihax/optional"
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
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
	name := d.Get("name").(string)
	endpointType := d.Get("type").(string)

	serviceEndpoints, _, err := config.Client.EndpointsApi.GetServiceEndpointsByNames(config.Context, config.Organization, d.Get("project_id").(string), name, config.ApiVersion, &azuredevops.GetServiceEndpointsByNamesOpts{
		IncludeDetails: optional.NewBool(true),
		Type_:          optional.NewString(endpointType),
	})

	if err != nil {
		return err
	}

	if serviceEndpoints.Count == 0 {
		return errors.New("Service Endpoint Not Found")
	}

	serviceEndpoint := serviceEndpoints.ServiceEndpoints[0]

	d.Set("owner", serviceEndpoint.Owner)
	d.Set("url", serviceEndpoint.Url)
	d.SetId(serviceEndpoint.Id)

	return nil
}
