package azuredevops

import (
	//	"log"
	"time"

	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "private",
			},
			"capabilities": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version_control": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_control_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Default:  "Git",
									},
								},
							},
						},
						"process_template": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template_type_id": &schema.Schema{
										Type:     schema.TypeString,
										ForceNew: true,
										Optional: true,
										Default:  "6b724908-ef14-45cf-84f8-768b5384da45",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	visibility := d.Get("visibility").(string)

	newProject := azuredevops.TeamProject{
		Name:        name,
		Description: description,
		Visibility:  visibility,
		Capabilities: map[string]map[string]string{
			"processTemplate": {
				"templateTypeId": d.Get("capabilities.0.process_template.0.template_type_id").(string),
			},
			"versioncontrol": {
				"sourceControlType": d.Get("capabilities.0.version_control.0.source_control_type").(string),
			},
		},
	}

	o, _, err := config.Client.ProjectsApi.QueueCreateProject(config.Context, config.Organization, config.ApiVersion, newProject)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	for {
		time.Sleep(2 * time.Second)
		operation, _, err := config.Client.OperationsApi.GetOperation(config.Context, o.Id, config.Organization, config.ApiVersion, nil)

		if err != nil {
			return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
		}

		if operation.Status == "succeeded" {
			break
		}
	}

	time.Sleep(3 * time.Second)

	projects, _, err := config.Client.ProjectsApi.GetProjects(config.Context, config.Organization, config.ApiVersion, nil)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	for _, proj := range projects.Projects {
		if proj.Name == name {
			d.SetId(proj.Id)
			break
		}
	}

	return resourceProjectRead(d, meta)
}

func resourceProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, _, err := config.Client.ProjectsApi.GetProject(config.Context, config.Organization, d.Id(), config.ApiVersion, nil)

	if err != nil {
		return err
	}

	d.Set("description", project.Description)
	d.Set("visibility", project.Visibility)

	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	newProject := azuredevops.TeamProject{
		Description: d.Get("description").(string),
		Visibility:  d.Get("visibility").(string),
	}

	o, _, err := config.Client.ProjectsApi.UpdateProject(config.Context, config.Organization, d.Id(), config.ApiVersion, newProject)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	for {
		time.Sleep(2 * time.Second)
		operation, _, err := config.Client.OperationsApi.GetOperation(config.Context, o.Id, config.Organization, config.ApiVersion, nil)

		if err != nil {
			return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
		}

		if operation.Status == "succeeded" {
			break
		}
	}

	return resourceProjectRead(d, meta)
}

func resourceProjectDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	o, _, err := config.Client.ProjectsApi.QueueDeleteProject(config.Context, config.Organization, d.Id(), config.ApiVersion)

	if err != nil {
		return err
	}

	for {
		time.Sleep(2 * time.Second)
		operation, _, err := config.Client.OperationsApi.GetOperation(config.Context, o.Id, config.Organization, config.ApiVersion, nil)

		if err != nil {
			return err
		}

		if operation.Status == "succeeded" {
			break
		}
	}

	return nil
}
