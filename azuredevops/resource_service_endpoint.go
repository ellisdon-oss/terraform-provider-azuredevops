package azuredevops

import (
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceServiceEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceEndpointCreate,
		Update: resourceServiceEndpointUpdate,
		Delete: resourceServiceEndpointDelete,
		Read:   resourceServiceEndpointRead,

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
				ForceNew: true,
				Required: true,
			},
			"allow_all_pipelines": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"data": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"authorization": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scheme": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"parameters": &schema.Schema{
							Type:     schema.TypeMap,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceServiceEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serviceEndpoint, _, err := config.Client.EndpointsApi.GetServiceEndpointDetails(config.Context, config.Organization, d.Get("project_id").(string), d.Id(), config.ApiVersion)

	if err != nil {
		return err
	}

	if serviceEndpoint.Name == "" {
		d.SetId("")
		return nil
	}

	d.Set("name", serviceEndpoint.Name)
	d.Set("url", serviceEndpoint.Url)
	d.Set("owner", serviceEndpoint.Owner)
	d.Set("type", serviceEndpoint.Type)
	d.Set("data", serviceEndpoint.Data)
	d.SetId(serviceEndpoint.Id)

	return nil
}

func resourceServiceEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	newServiceEndpoint := azuredevops.ServiceEndpoint{
		Name:  d.Get("name").(string),
		Owner: d.Get("owner").(string),
		Url:   d.Get("url").(string),
		Type:  d.Get("type").(string),
		Data:  d.Get("data").(map[string]interface{}),
		Authorization: azuredevops.EndpointAuthorization{
			Parameters: d.Get("authorization.0.parameters").(map[string]interface{}),
			Scheme:     d.Get("authorization.0.scheme").(string),
		},
	}

	serviceEndpoint, _, err := config.Client.EndpointsApi.CreateServiceEndpoint(config.Context, config.Organization, d.Get("project_id").(string), "5.1-preview.2", newServiceEndpoint)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.SetId(serviceEndpoint.Id)

	_, err2 := config.Client.PipelinePermissionsApi.SetPipelinePermissions(config.Context, config.Organization, d.Get("project_id").(string), d.Id(), "5.1-preview.1", azuredevops.PipelinePermissionRequest{
		AllPipelines: azuredevops.PipelinePermissionRequestAllPipelines{
			Authorized: d.Get("allow_all_pipelines").(bool),
		},
	})

	if err2 != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}
	return resourceServiceEndpointRead(d, meta)
}

func resourceServiceEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	newServiceEndpoint := azuredevops.ServiceEndpoint{
		Name:  d.Get("name").(string),
		Owner: d.Get("owner").(string),
		Url:   d.Get("url").(string),
		Type:  d.Get("type").(string),
		Data:  d.Get("data").(map[string]interface{}),
		Authorization: azuredevops.EndpointAuthorization{
			Parameters: d.Get("authorization.0.parameters").(map[string]interface{}),
			Scheme:     d.Get("authorization.0.scheme").(string),
		},
	}

	serviceEndpoint, _, err := config.Client.EndpointsApi.UpdateServiceEndpoint(config.Context, config.Organization, d.Get("project_id").(string), d.Id(), "5.1-preview.2", newServiceEndpoint, nil)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.SetId(serviceEndpoint.Id)

	_, err2 := config.Client.PipelinePermissionsApi.SetPipelinePermissions(config.Context, config.Organization, d.Get("project_id").(string), d.Id(), "5.1-preview.1", azuredevops.PipelinePermissionRequest{
		AllPipelines: azuredevops.PipelinePermissionRequestAllPipelines{
			Authorized: d.Get("allow_all_pipelines").(bool),
		},
	})

	if err2 != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	return resourceServiceEndpointRead(d, meta)
}

func resourceServiceEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	_, err := config.Client.EndpointsApi.DeleteServiceEndpoint(config.Context, config.Organization, d.Get("project_id").(string), d.Id(), "5.1-preview.2", nil)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	return nil
}
