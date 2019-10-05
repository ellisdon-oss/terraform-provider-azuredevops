package azuredevops

import (
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/operations"
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

	coreClient, err := core.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	visibility := core.ProjectVisibility(d.Get("visibility").(string))

	capabilities := map[string]map[string]string{
		"processTemplate": {
			"templateTypeId": d.Get("capabilities.0.process_template.0.template_type_id").(string),
		},
		"versioncontrol": {
			"sourceControlType": d.Get("capabilities.0.version_control.0.source_control_type").(string),
		},
	}

	newProject := core.TeamProject{
		Name:         &name,
		Description:  &description,
		Visibility:   &visibility,
		Capabilities: &capabilities,
	}

	o, err := coreClient.QueueCreateProject(config.Context, core.QueueCreateProjectArgs{
		ProjectToCreate: &newProject,
	})

	if err != nil {
		return err
	}

	operationClient := operations.NewClient(config.Context, config.Connection)

	for {
		time.Sleep(2 * time.Second)

		operation, err := operationClient.GetOperation(config.Context, operations.GetOperationArgs{
			OperationId: o.Id,
		})

		if err != nil {
			return err
		}

		if (*(*operation).Status) == "succeeded" {
			break
		}
	}

	time.Sleep(3 * time.Second)

	projects, err := coreClient.GetProjects(config.Context, core.GetProjectsArgs{})

	if err != nil {
		return err
	}

	for _, proj := range (*projects).Value {
		if (*proj.Name) == name {
			d.SetId(proj.Id.String())
			break
		}
	}

	return resourceProjectRead(d, meta)
}

func resourceProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	coreClient, err := core.NewClient(config.Context, config.Connection)
	projectID := d.Id()

	project, err := coreClient.GetProject(config.Context, core.GetProjectArgs{
		ProjectId: &projectID,
	})

	if err != nil {
		return err
	}

	d.Set("description", project.Description)
	d.Set("visibility", project.Visibility)

	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	coreClient, err := core.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	description := d.Get("description").(string)
	visibility := core.ProjectVisibility(d.Get("visibility").(string))
	projectID, _ := uuid.Parse(d.Id())
	newProject := core.TeamProject{
		Description: &description,
		Visibility:  &visibility,
	}

	o, err := coreClient.UpdateProject(config.Context, core.UpdateProjectArgs{
		ProjectId:     &projectID,
		ProjectUpdate: &newProject,
	})

	if err != nil {
		return err
	}

	operationClient := operations.NewClient(config.Context, config.Connection)

	for {
		time.Sleep(2 * time.Second)

		operation, err := operationClient.GetOperation(config.Context, operations.GetOperationArgs{
			OperationId: o.Id,
		})

		if err != nil {
			return err
		}

		if (*(*operation).Status) == "succeeded" {
			break
		}
	}

	return resourceProjectRead(d, meta)
}

func resourceProjectDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	coreClient, err := core.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID, _ := uuid.Parse(d.Id())

	o, err := coreClient.QueueDeleteProject(config.Context, core.QueueDeleteProjectArgs{
		ProjectId: &projectID,
	})

	if err != nil {
		return err
	}

	operationClient := operations.NewClient(config.Context, config.Connection)

	for {
		time.Sleep(2 * time.Second)

		operation, err := operationClient.GetOperation(config.Context, operations.GetOperationArgs{
			OperationId: o.Id,
		})

		if err != nil {
			return err
		}

		if (*(*operation).Status) == "succeeded" {
			break
		}
	}

	return nil
}
