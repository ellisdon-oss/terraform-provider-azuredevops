package azuredevops

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

func dataSourceProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProjectRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"abbreviation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_team_image_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_update_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"links": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			"capabilities": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			"default_team": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	projectID := d.Get("project_id").(string)
	coreClient, err := core.NewClient(config.Context, config.Connection)

	project, err := coreClient.GetProject(config.Context, core.GetProjectArgs{
		ProjectId: &projectID,
	})

	if err != nil {
		return err
	}

	d.Set("name", project.Name)
	d.Set("abbreviation", project.Abbreviation)
	d.Set("default_team_image_url", project.DefaultTeamImageUrl)
	d.Set("description", project.Description)
	d.Set("last_update_time", project.LastUpdateTime)
	d.Set("revision", project.Revision)
	d.Set("state", project.State)
	d.Set("url", project.Url)
	d.Set("visibility", project.Visibility)

	d.SetId(project.Id.String())

	return nil
}
