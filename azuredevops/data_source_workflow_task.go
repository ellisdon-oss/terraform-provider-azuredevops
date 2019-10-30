package azuredevops

import (
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
)

func dataSourceWorkflowTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceWorkflowTaskRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"wait": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func dataSourceWorkflowTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	wait := d.Get("wait").(int)

	time.Sleep(time.Duration(wait) * time.Second)

	tasks, err := getAllTasks(config)

	if err != nil {
		return err
	}

	id := getTaskID(d.Get("name").(string), (*tasks))

	if id == "" {
		return errors.New("Task Not Found")
	}

	d.SetId(id)

	return nil
}

func getTaskID(name string, tasks []distributedWorkflowTask) string {
	for _, task := range tasks {

		if (*task.Name) == name {
			return task.ID.String()
		}
	}

	return ""
}

func getAllTasks(config *Config) (*[]distributedWorkflowTask, error) {

	fakeId, _ := uuid.Parse("efc2f575-36ef-48e9-b672-0c6fb4a48ac5")
	generalClient, _ := config.Connection.GetClientByResourceAreaId(config.Context, fakeId)

	fullUrl := strings.TrimRight(config.Connection.BaseUrl, "/") + "/" + strings.TrimLeft("_apis/distributedtask/tasks", "/")

	req, err := generalClient.CreateRequestMessage(config.Context, http.MethodGet, fullUrl, "", nil, "application/json", "application/json", nil)

	if err != nil {
		return nil, err
	}

	resp, err := generalClient.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var responseValue []distributedWorkflowTask
	err = generalClient.UnmarshalCollectionBody(resp, &responseValue)

	return &responseValue, err
}

type distributedWorkflowTask struct {
	ID   *uuid.UUID `json:"id,omitempty"`
	Name *string    `json:"name,omitempty"`
}
