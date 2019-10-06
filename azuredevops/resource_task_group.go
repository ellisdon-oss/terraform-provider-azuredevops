package azuredevops

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/taskagent"
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
							ForceNew: true,
							Default:  false,
						},
						"major": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"minor": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"patch": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"task": taskGroupTaskSchema(),
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Deploy",
			},
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

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)
	groupID := d.Get("group_id").(string)

	taskGroups, err := agentClient.GetTaskGroups(config.Context, taskagent.GetTaskGroupsArgs{
		Project: &projectID,
	})

	if err != nil {
		return err
	}

	var group taskagent.TaskGroup

	for _, v := range *taskGroups {
		if v.Id.String() == groupID {
			group = v
			break
		}
	}

	if group.Id.String() == "" {
		d.SetId("")
		return nil
	}

	var tasks []interface{}

	for _, v := range *group.Tasks {
		tasks = append(tasks, map[string]interface{}{
			"name":              *v.DisplayName,
			"version":           *v.Task.VersionSpec,
			"task_id":           v.Task.Id.String(),
			"definition_type":   *v.Task.DefinitionType,
			"inputs":            *v.Inputs,
			"enabled":           *v.Enabled,
			"condition":         *v.Condition,
			"continue_on_error": *v.ContinueOnError,
			"environment":       *v.Environment,
		})
	}

	d.Set("name", *group.Name)
	d.Set("task", tasks)
	d.Set("category", group.Category)
	d.Set("inputs", *group.Inputs)
	d.Set("revision", *group.Revision)
	d.SetId(fmt.Sprintf("%s-%s", d.Get("project_id").(string), group.Id.String()))

	return nil
}

func resourceTaskGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	tasks := d.Get("task").([]interface{})

	var finalTasks []taskagent.TaskGroupStep

	for _, v := range tasks {
		v := v.(map[string]interface{})
		name := v["name"].(string)

		inputs := convertInterfaceToStringMap(v["inputs"].(map[string]interface{}))

		taskID, _ := uuid.Parse(v["task_id"].(string))

		temp := taskagent.TaskGroupStep{
			DisplayName: &name,
			Inputs:      &inputs,
			Task: &taskagent.TaskDefinitionReference{
				Id: &taskID,
			},
		}

		if r := v["timeout_in_minutes"]; r != nil {
			timeoutInMinutes := r.(int)
			temp.TimeoutInMinutes = &timeoutInMinutes
		}

		if r := v["always_run"]; r != nil {
			alwaysRun := r.(bool)
			temp.AlwaysRun = &alwaysRun
		}

		if r := v["enabled"]; r != nil {
			enabled := r.(bool)
			temp.Enabled = &enabled
		}

		if r := v["condition"]; r != nil {
			condition := r.(string)
			temp.Condition = &condition
		}

		if r := v["continue_on_error"]; r != nil {
			continueOnError := r.(bool)
			temp.ContinueOnError = &continueOnError
		}

		if r := v["environment"]; r != nil {
			environment := convertInterfaceToStringMap(r.(map[string]interface{}))
			temp.Environment = &environment
		}

		if r := v["version"]; r != nil {
			version := r.(string)
			temp.Task.VersionSpec = &version
		}

		if r := v["definition_type"]; r != nil {
			definitionType := r.(string)
			temp.Task.DefinitionType = &definitionType
		}

		finalTasks = append(finalTasks, temp)
	}

	name := d.Get("name").(string)
	category := d.Get("category").(string)

	taskGroup := taskagent.TaskGroupCreateParameter{
		Name:     &name,
		Tasks:    &finalTasks,
		Category: &category,
	}

	if v := d.Get("version"); len(v.([]interface{})) != 0 {
		v := v.([]interface{})
		t := v[0].(map[string]interface{})

		isTest := t["is_test"].(bool)
		major := t["major"].(int)
		minor := t["minor"].(int)
		patch := t["patch"].(int)

		version := taskagent.TaskVersion{
			IsTest: &isTest,
			Major:  &major,
			Minor:  &minor,
			Patch:  &patch,
		}

		taskGroup.Version = &version

	}

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	projectID := d.Get("project_id").(string)

	createdTaskGroup, err := agentClient.AddTaskGroup(config.Context, taskagent.AddTaskGroupArgs{
		Project:   &projectID,
		TaskGroup: &taskGroup,
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s-%s", d.Get("project_id").(string), createdTaskGroup.Id.String()))
	d.Set("group_id", createdTaskGroup.Id.String())
	return resourceTaskGroupRead(d, meta)
}

func resourceTaskGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return nil
	}

	groupID := d.Get("group_id").(string)
	tasks := d.Get("task").([]interface{})
	projectID := d.Get("project_id").(string)

	taskGroups, err := agentClient.GetTaskGroups(config.Context, taskagent.GetTaskGroupsArgs{
		Project: &projectID,
	})

	if err != nil {
		return err
	}

	var group taskagent.TaskGroup

	for _, v := range *taskGroups {
		if v.Id.String() == groupID {
			group = v
			break
		}
	}

	if group.Id.String() == "" {
		d.SetId("")
		return nil
	}

	var finalTasks []taskagent.TaskGroupStep

	for _, v := range tasks {
		v := v.(map[string]interface{})

		name := v["name"].(string)
		inputs := convertInterfaceToStringMap(v["inputs"].(map[string]interface{}))

		taskID, _ := uuid.Parse(v["task_id"].(string))

		temp := taskagent.TaskGroupStep{
			DisplayName: &name,
			Inputs:      &inputs,
			Task: &taskagent.TaskDefinitionReference{
				Id: &taskID,
			},
		}

		if r := v["timeout_in_minutes"]; r != nil {
			timeoutInMinutes := r.(int)
			temp.TimeoutInMinutes = &timeoutInMinutes
		}

		if r := v["always_run"]; r != nil {
			alwaysRun := r.(bool)
			temp.AlwaysRun = &alwaysRun
		}

		if r := v["enabled"]; r != nil {
			enabled := r.(bool)
			temp.Enabled = &enabled
		}

		if r := v["condition"]; r != nil {
			condition := r.(string)
			temp.Condition = &condition
		}

		if r := v["continue_on_error"]; r != nil {
			continueOnError := r.(bool)
			temp.ContinueOnError = &continueOnError
		}

		if r := v["environment"]; r != nil {
			environment := convertInterfaceToStringMap(r.(map[string]interface{}))
			temp.Environment = &environment
		}

		if r := v["version"]; r != nil {
			version := r.(string)
			temp.Task.VersionSpec = &version
		}

		if r := v["definition_type"]; r != nil {
			definitionType := r.(string)
			temp.Task.DefinitionType = &definitionType
		}

		finalTasks = append(finalTasks, temp)
	}

	name := d.Get("name").(string)
	category := d.Get("category").(string)
	parsedGroupID, _ := uuid.Parse(groupID)
	revision := d.Get("revision").(int)

	taskGroup := taskagent.TaskGroupUpdateParameter{
		Id:       &parsedGroupID,
		Name:     &name,
		Category: &category,
		Tasks:    &finalTasks,
		Revision: &revision,
		Version: &taskagent.TaskVersion{
			IsTest: group.Version.IsTest,
			Major:  group.Version.Major,
			Minor:  group.Version.Minor,
			Patch:  group.Version.Patch,
		},
	}

	updatedGroup, err := agentClient.UpdateTaskGroup(config.Context, taskagent.UpdateTaskGroupArgs{
		TaskGroup:   &taskGroup,
		Project:     &projectID,
		TaskGroupId: &parsedGroupID,
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s-%s", d.Get("project_id").(string), updatedGroup.Id))

	return resourceTaskGroupRead(d, meta)
}

func resourceTaskGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}
	projectID := d.Get("project_id").(string)
	groupID, _ := uuid.Parse(d.Get("group_id").(string))

	agentClient.DeleteTaskGroup(config.Context, taskagent.DeleteTaskGroupArgs{
		Project:     &projectID,
		TaskGroupId: &groupID,
	})

	if err != nil {
		return err
	}

	return nil
}
