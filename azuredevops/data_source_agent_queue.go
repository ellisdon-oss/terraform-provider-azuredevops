package azuredevops

import (
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/hashicorp/terraform/helper/validation"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func dataSourceAgentQueue() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAgentQueueRead,

		Schema: map[string]*schema.Schema{
			"queue_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceAgentQueueRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	result, _, err := config.Client.QueuesApi.GetAgentQueuesByNames(config.Context, config.Organization, d.Get("queue_name").(string), d.Get("project_id").(string), config.ApiVersion, &azuredevops.GetAgentQueuesByNamesOpts{})

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	if result.Count == 0 {
		return errors.New("AgentQueue Not Found")
	}

	log.Print(result.TaskAgentQueues[0])
	d.SetId(fmt.Sprint(result.TaskAgentQueues[0].Id))

	return nil
}
