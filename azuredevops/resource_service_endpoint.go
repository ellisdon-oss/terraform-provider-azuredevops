package azuredevops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/serviceendpoint"
	"net/http"
	"strings"
)

func resourceServiceEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceEndpointCreate,
		Update: resourceServiceEndpointUpdate,
		Delete: resourceServiceEndpointDelete,
		Read:   resourceServiceEndpointRead,
		Importer: &schema.ResourceImporter{
			State: resourceServiceEndpointImport,
		},

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
				Optional: true,
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

	endpointClient, err := serviceendpoint.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	endpointID, _ := uuid.Parse(d.Id())
	projectID := d.Get("project_id").(string)

	serviceEndpoint, err := endpointClient.GetServiceEndpointDetails(config.Context, serviceendpoint.GetServiceEndpointDetailsArgs{
		Project:    &projectID,
		EndpointId: &endpointID,
	})

	if err != nil {
		return err
	}

	if serviceEndpoint == nil {
		d.SetId("")
		return nil
	}

	if serviceEndpoint.Name == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", *serviceEndpoint.Name)
	d.Set("url", *serviceEndpoint.Url)
	d.Set("owner", *serviceEndpoint.Owner)
	d.Set("type", *serviceEndpoint.Type)
	d.Set("data", *serviceEndpoint.Data)
	d.SetId(serviceEndpoint.Id.String())

	return nil
}

func resourceServiceEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Get("name").(string)
	owner := d.Get("owner").(string)
	url := d.Get("url").(string)
	endpointType := d.Get("type").(string)
	projectID := d.Get("project_id").(string)

	data := make(map[string]string)

	if v := d.Get("data"); v != nil {
		data = convertInterfaceToStringMap(d.Get("data").(map[string]interface{}))
	}

	parameters := convertInterfaceToStringMap(d.Get("authorization.0.parameters").(map[string]interface{}))
	scheme := d.Get("authorization.0.scheme").(string)

	newServiceEndpoint := serviceendpoint.ServiceEndpoint{
		Name:  &name,
		Owner: &owner,
		Url:   &url,
		Type:  &endpointType,
		Data:  &data,
		Authorization: &serviceendpoint.EndpointAuthorization{
			Parameters: &parameters,
			Scheme:     &scheme,
		},
	}

	endpointClient, err := serviceendpoint.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	serviceEndpoint, err := endpointClient.CreateServiceEndpoint(config.Context, serviceendpoint.CreateServiceEndpointArgs{
		Endpoint: &newServiceEndpoint,
		Project:  &projectID,
	})

	if err != nil {
		return err
	}

	d.SetId(serviceEndpoint.Id.String())
	apiURL := projectID + "/_apis/pipelines/pipelinePermissions/endpoint/" + serviceEndpoint.Id.String()

	fullURL := strings.TrimRight(config.Connection.BaseUrl, "/") + "/" + strings.TrimLeft(apiURL, "/")

	body, marshalErr := json.Marshal(map[string]interface{}{
		"allPipelines": map[string]bool{
			"authorized": d.Get("allow_all_pipelines").(bool),
		},
	})

	if marshalErr != nil {
		return marshalErr
	}

	req, err := endpointClient.Client.CreateRequestMessage(config.Context, http.MethodPatch, fullURL, "5.1-preview.1", bytes.NewReader(body), "application/json", "application/json", nil)

	if err != nil {
		return err
	}

	_, err = endpointClient.Client.SendRequest(req)
	if err != nil {
		return err
	}

	return resourceServiceEndpointRead(d, meta)
}

func resourceServiceEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Get("name").(string)
	owner := d.Get("owner").(string)
	url := d.Get("url").(string)
	endpointType := d.Get("type").(string)
	projectID := d.Get("project_id").(string)
	endpointID, _ := uuid.Parse(d.Id())
	data := convertInterfaceToStringMap(d.Get("data").(map[string]interface{}))

	parameters := convertInterfaceToStringMap(d.Get("authorization.0.parameters").(map[string]interface{}))
	scheme := d.Get("authorization.0.scheme").(string)

	newServiceEndpoint := serviceendpoint.ServiceEndpoint{
		Name:  &name,
		Owner: &owner,
		Url:   &url,
		Type:  &endpointType,
		Data:  &data,
		Authorization: &serviceendpoint.EndpointAuthorization{
			Parameters: &parameters,
			Scheme:     &scheme,
		},
	}

	endpointClient, err := serviceendpoint.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	serviceEndpoint, err := endpointClient.UpdateServiceEndpoint(config.Context, serviceendpoint.UpdateServiceEndpointArgs{
		Endpoint:   &newServiceEndpoint,
		EndpointId: &endpointID,
		Project:    &projectID,
	})

	if err != nil {
		return nil
	}

	d.SetId(serviceEndpoint.Id.String())

	apiURL := projectID + "/_apis/pipelines/pipelinePermissions/endpoint/" + serviceEndpoint.Id.String()

	fullURL := strings.TrimRight(config.Connection.BaseUrl, "/") + "/" + strings.TrimLeft(apiURL, "/")

	body, marshalErr := json.Marshal(map[string]interface{}{
		"allPipelines": map[string]bool{
			"authorized": d.Get("allow_all_pipelines").(bool),
		},
	})

	if marshalErr != nil {
		return marshalErr
	}

	req, err := endpointClient.Client.CreateRequestMessage(config.Context, http.MethodPatch, fullURL, "5.1-preview.1", bytes.NewReader(body), "application/json", "application/json", nil)

	if err != nil {
		return err
	}

	_, err = endpointClient.Client.SendRequest(req)
	if err != nil {
		return err
	}

	return resourceServiceEndpointRead(d, meta)
}

func resourceServiceEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	endpointClient, err := serviceendpoint.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)
	endpointID, _ := uuid.Parse(d.Id())

	err = endpointClient.DeleteServiceEndpoint(config.Context, serviceendpoint.DeleteServiceEndpointArgs{
		Project:    &projectID,
		EndpointId: &endpointID,
	})

	if err != nil {
		return err
	}

	return nil
}

func resourceServiceEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	name := d.Id()
	if name == "" {
		return nil, fmt.Errorf("service endpoint id cannot be empty")
	}

	res := strings.Split(name, "/")
	if len(res) != 2 {
		return nil, fmt.Errorf("the format has to be in <project-id>/<service-endpoint-id>")
	}

	d.Set("project_id", res[0])
	d.SetId(res[1])

	return []*schema.ResourceData{d}, nil
}
