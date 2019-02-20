package azuredevops

import (
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"

	//"github.com/hashicorp/terraform/helper/validation"
	"github.com/pkg/errors"
	//	"github.com/pkg/errors"
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

	tasks, _, err := config.Client.TasksApi.GetAllTasks(config.Context, config.Organization, config.ApiVersion)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	id := getTaskID(d.Get("name").(string), tasks["value"].([]interface{}))

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

func getTaskID(name string, tasks []interface{}) string {
	for _, task := range tasks {
		task := task.(map[string]interface{})

		if task["name"].(string) == name {
			return task["id"].(string)
		}
	}

	return ""
}
