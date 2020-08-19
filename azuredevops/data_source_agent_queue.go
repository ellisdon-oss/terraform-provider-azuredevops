package azuredevops

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/taskagent"
	"github.com/pkg/errors"
)

func dataSourceAgentQueue() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAgentQueueRead,

		Schema: map[string]*schema.Schema{
			"queue_name": &schema.Schema{
				Type:     schema.TypeString,
        Description: "The agent queue name",
				Required: true,
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
        Description: "The project name/ID",
				Required: true,
			},
		},
	}
}

func dataSourceAgentQueueRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	queueName := []string{d.Get("queue_name").(string)}
	projectID := d.Get("project_id").(string)

	agentQueue, err := agentClient.GetAgentQueuesByNames(config.Context, taskagent.GetAgentQueuesByNamesArgs{
		QueueNames: &queueName,
		Project:    &projectID,
	})

	if err != nil {
		return err
	}

	if len(*agentQueue) == 0 {
		return errors.New("AgentQueue Not Found")
	}

	d.SetId(fmt.Sprint(*(*agentQueue)[0].Id))

	return nil
}
