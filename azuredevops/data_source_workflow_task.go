package azuredevops

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"github.com/pkg/errors"
)

func dataSourceWorkflowTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceWorkflowTaskRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceWorkflowTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	tasks, err := releaseClient.GetAllTasks(config.Context, config.Connection.BaseUrl, release.GetAllTasksArgs{})

	if err != nil {
		return err
	}

	id := getTaskID(d.Get("name").(string), (*tasks))

	if id == "" {
		return errors.New("Task Not Found")
	}

	d.SetId(id)
	//	d.Set("abbreviation", project.Abbreviation)
	//	d.Set("default_team_image_url", project.DefaultTeamImageUrl)
	//	d.Set("description", project.Description)
	//	d.Set("last_update_time", project.LastUpdateTime)
	//	d.Set("revision", project.Revision)
	//	d.Set("state", project.State)
	//	d.Set("url", project.Url)
	//	d.Set("visibility", project.Visibility)
	//
	//	d.SetId(project.Id)

	return nil
}

func getTaskID(name string, tasks []release.DistributedWorkflowTask) string {
	for _, task := range tasks {

		if (*task.Name) == name {
			return task.ID.String()
		}
	}

	return ""
}
