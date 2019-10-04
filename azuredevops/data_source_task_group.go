package azuredevops

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/taskagent"
	"github.com/pkg/errors"
)

func dataSourceTaskGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskGroupRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceTaskGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)
	projectID := d.Get("project_id").(string)

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	res, err := agentClient.GetTaskGroups(config.Context, taskagent.GetTaskGroupsArgs{
		Project: &projectID,
	})

	if err != nil {
		return err
	}

	if len(*res) == 0 {
		return errors.New("Project have no task groups")
	}

	var group taskagent.TaskGroup

	for _, v := range *res {
		if (*v.Name) == name {
			group = v
			break
		}
	}

	if group.Id.String() == "" {
		return errors.New("Task group not found")
	}

	d.Set("group_id", group.Id.String())
	d.SetId(fmt.Sprintf("%s/%s", d.Get("project_id").(string), group.Id))

	return nil
}
