package azuredevops

import (
	//"log"

	"github.com/hashicorp/terraform/helper/schema"
	//	"github.com/pkg/errors"
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

	project, _, err := config.Client.ProjectsApi.GetProject(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, nil)

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

	d.SetId(project.Id)

	return nil
}
