package azuredevops

import (
	"fmt"
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceTaskGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskGroupCreate,
		Update: resourceTaskGroupUpdate,
		Delete: resourceTaskGroupDelete,
		Read:   resourceTaskGroupRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				ForceNew: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_test": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"major": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"minor": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"patch": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"task": taskGroupTaskSchema(),
		},
	}
}

func taskGroupTaskSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
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
	}
}

func resourceTaskGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	groupID := d.Get("group_id").(string)

	res, _, err := config.Client.TaskgroupsApi.GetTaskGroups(config.Context, config.Organization, d.Get("project_id").(string), "", config.ApiVersion, &azuredevops.GetTaskGroupsOpts{})

	if err != nil {
		return err
	}

	var group azuredevops.TaskGroup

	for _, v := range res.TaskGroups {
		if v.Id == groupID {
			group = v
			break
		}
	}

	if group.Id == "" {
		d.SetId("")
		return nil
	}

	var tasks []interface{}

	for _, v := range group.Tasks {
		tasks = append(tasks, map[string]interface{}{
			"name":              v.DisplayName,
			"version":           v.Task.VersionSpec,
			"task_id":           v.Task.Id,
			"definition_type":   v.Task.DefinitionType,
			"inputs":            v.Inputs,
			"enabled":           v.Enabled,
			"condition":         v.Condition,
			"continue_on_error": v.ContinueOnError,
			"environment":       v.Environment,
		})

	}

	d.Set("task", tasks)
	d.Set("revision", group.Revision)
	d.SetId(fmt.Sprintf("%s-%s", d.Get("project_id").(string), group.Id))

	return nil
}

func resourceTaskGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	tasks := d.Get("task").([]interface{})

	var finalTasks []azuredevops.TaskGroupStep

	for _, v := range tasks {
		v := v.(map[string]interface{})
		temp := azuredevops.TaskGroupStep{
			DisplayName: v["name"].(string),
			Inputs:      v["inputs"].(map[string]interface{}),
			Task: azuredevops.TaskDefinitionReference{
				Id: v["task_id"].(string),
			},
		}

		if r := v["timeout_in_minutes"]; r != nil {
			temp.TimeoutInMinutes = int32(r.(int))
		}

		if r := v["always_run"]; r != nil {
			temp.AlwaysRun = r.(bool)
		}

		if r := v["enabled"]; r != nil {
			temp.Enabled = r.(bool)
		}

		if r := v["condition"]; r != nil {
			temp.Condition = r.(string)
		}

		if r := v["continue_on_error"]; r != nil {
			temp.ContinueOnError = r.(bool)
		}

		if r := v["environment"]; r != nil {
			temp.Environment = r.(map[string]interface{})
		}

		if r := v["version"]; r != nil {
			temp.Task.VersionSpec = r.(string)
		}

		if r := v["definition_type"]; r != nil {
			temp.Task.DefinitionType = r.(string)
		}

		finalTasks = append(finalTasks, temp)
	}

	taskGroup := azuredevops.TaskGroupCreateParameter{
		Name:  d.Get("name").(string),
		Tasks: finalTasks,
	}

	if v := d.Get("version"); len(v.([]interface{})) != 0 {
		v := v.([]interface{})
		t := v[0].(map[string]interface{})
		version := azuredevops.TaskVersion{
			IsTest: t["is_test"].(bool),
			Major:  int32(t["major"].(int)),
			Minor:  int32(t["minor"].(int)),
			Patch:  int32(t["patch"].(int)),
		}

		taskGroup.Version = version

	}

	group, _, err := config.Client.TaskgroupsApi.AddTaskGroup(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, taskGroup)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.SetId(fmt.Sprintf("%s-%s", d.Get("project_id").(string), group.Id))
	d.Set("group_id", group.Id)
	return resourceTaskGroupRead(d, meta)
}

func resourceTaskGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	groupID := d.Get("group_id").(string)
	tasks := d.Get("task").([]interface{})

	res, _, err := config.Client.TaskgroupsApi.GetTaskGroups(config.Context, config.Organization, d.Get("project_id").(string), "", config.ApiVersion, &azuredevops.GetTaskGroupsOpts{})

	if err != nil {
		return err
	}

	var group azuredevops.TaskGroup

	for _, v := range res.TaskGroups {
		if v.Id == groupID {
			group = v
			break
		}
	}

	if group.Id == "" {
		d.SetId("")
		return nil
	}

	var finalTasks []azuredevops.TaskGroupStep

	for _, v := range tasks {
		v := v.(map[string]interface{})
		temp := azuredevops.TaskGroupStep{
			DisplayName: v["name"].(string),
			Inputs:      v["inputs"].(map[string]interface{}),
			Task: azuredevops.TaskDefinitionReference{
				Id: v["task_id"].(string),
			},
		}

		if r := v["timeout_in_minutes"]; r != nil {
			temp.TimeoutInMinutes = int32(r.(int))
		}

		if r := v["always_run"]; r != nil {
			temp.AlwaysRun = r.(bool)
		}

		if r := v["enabled"]; r != nil {
			temp.Enabled = r.(bool)
		}

		if r := v["condition"]; r != nil {
			temp.Condition = r.(string)
		}

		if r := v["continue_on_error"]; r != nil {
			temp.ContinueOnError = r.(bool)
		}

		if r := v["environment"]; r != nil {
			temp.Environment = r.(map[string]interface{})
		}

		if r := v["version"]; r != nil {
			temp.Task.VersionSpec = r.(string)
		}

		if r := v["definition_type"]; r != nil {
			temp.Task.DefinitionType = r.(string)
		}

		finalTasks = append(finalTasks, temp)
	}

	taskGroup := azuredevops.TaskGroupUpdateParameter{
		Id:       groupID,
		Name:     d.Get("name").(string),
		Tasks:    finalTasks,
		Revision: int32(d.Get("revision").(int)),
		Version: azuredevops.TaskVersion{
			IsTest: group.Version.IsTest,
			Major:  group.Version.Major,
			Minor:  group.Version.Minor,
			Patch:  group.Version.Patch,
		},
	}

	updatedGroup, _, err := config.Client.TaskgroupsApi.UpdateTaskGroup(config.Context, config.Organization, d.Get("project_id").(string), "", config.ApiVersion, taskGroup)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.SetId(fmt.Sprintf("%s-%s", d.Get("project_id").(string), updatedGroup.Id))

	return resourceTaskGroupRead(d, meta)
}

func resourceTaskGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	_, err := config.Client.TaskgroupsApi.DeleteTaskGroup(config.Context, config.Organization, d.Get("project_id").(string), d.Get("group_id").(string), config.ApiVersion, &azuredevops.DeleteTaskGroupOpts{})

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	return nil
}
