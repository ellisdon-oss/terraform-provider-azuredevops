package azuredevops

import (
	"fmt"
  "strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
)

func dataSourceReleaseTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceReleaseTasksRead,

    Schema: map[string]*schema.Schema{
      "project_id": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
			"definition_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
      "job_rank": &schema.Schema{
        Type:     schema.TypeInt,
        Required: true,
      },
      "stage_name": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "tasks": &schema.Schema{
        Type:     schema.TypeList,
        Computed: true,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
              Type:     schema.TypeString,
              Required: true,
            },
            "definition_type": &schema.Schema{
              Type:     schema.TypeString,
              Optional: true,
              Default:  "task",
            },
            "version": &schema.Schema{
              Type:     schema.TypeString,
              Optional: true,
              Default:  "0.*",
            },
            "task_id": &schema.Schema{
              Type:     schema.TypeString,
              Required: true,
            },
            "enabled": &schema.Schema{
              Type:     schema.TypeBool,
              Optional: true,
              Default:  true,
            },
            "ref_name": &schema.Schema{
              Type:     schema.TypeString,
              Optional: true,
            },
            "always_run": &schema.Schema{
              Type:     schema.TypeBool,
              Optional: true,
              Default:  false,
            },
            "continue_on_error": &schema.Schema{
              Type:     schema.TypeBool,
              Optional: true,
              Default:  false,
            },
            "condition": &schema.Schema{
              Type:     schema.TypeString,
              Optional: true,
              Default:  "succeeded()",
            },
            "environment": &schema.Schema{
              Type:     schema.TypeMap,
              Optional: true,
              Elem: &schema.Schema{
                Type: schema.TypeString,
              },
            },
            "inputs": &schema.Schema{
              Type:     schema.TypeMap,
              Required: true,
              Elem: &schema.Schema{
                Type: schema.TypeString,
              },
            },
          },
        },
      },
    },
	}
}

func dataSourceReleaseTasksRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)
	definition_id, _ := strconv.ParseInt(d.Get("definition_id").(string), 10, 32)

	defID := int(definition_id)

	releaseDef, err := releaseClient.GetReleaseDefinition(config.Context, release.GetReleaseDefinitionArgs{
		Project:    &projectID,
    DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

  finalTasks := make([]map[string]interface{}, 0)

  for _, env := range *releaseDef.Environments {
    if *env.Name == d.Get("stage_name").(string) {
      tasks := (*env.DeployPhases)[d.Get("job_rank").(int)].(map[string]interface{})["workflowTasks"].([]interface{})

      for _, task := range tasks {
        task := task.(map[string]interface{})
        finalTasks = append(finalTasks, map[string]interface{}{
          "name":              task["name"],
          "definition_type":   task["definitionType"],
          "version":           task["version"],
          "task_id":           task["taskId"],
          "enabled":           task["enabled"],
          "always_run":        task["alwaysRun"],
          "continue_on_error": task["continueOnError"],
          "condition":         task["condition"],
          "environment":       task["environment"],
          "ref_name":          task["refName"],
          "inputs":            task["inputs"],
        })
      }
    }
  }

  d.Set("tasks", finalTasks)
  
	d.SetId(fmt.Sprintf("%s-%s-%d",*releaseDef.Id, d.Get("stage_name").(string), d.Get("job_rank").(int)))

	return nil
}
