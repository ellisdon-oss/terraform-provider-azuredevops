package azuredevops

import (
	"fmt"
	"github.com/ellisdon-oss/terraform-provider-azuredevops/azuredevops/helper"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"strconv"
	"strings"
)

func resourceReleaseTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceReleaseTaskCreate,
		Read:   resourceReleaseTaskRead,
		Update: resourceReleaseTaskUpdate,
		Delete: resourceReleaseTaskDelete,
		Importer: &schema.ResourceImporter{
			State: resourceReleaseTaskImport,
		},
		Schema: map[string]*schema.Schema{
			"job_rank": &schema.Schema{
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"job_name"},
			},
			"job_name": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"job_rank"},
			},
			"rank": &schema.Schema{
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"after", "before"},
			},
			"before": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"rank", "after"},
			},
			"after": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"rank", "before"},
			},
			"stage_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"task_info": helper.WorkflowTaskSingleSchema(),
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"definition_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceReleaseTaskCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	defID := d.Get("definition_id").(int)
	projectID := d.Get("project_id").(string)
	stageName := d.Get("stage_name").(string)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	res, err := releaseClient.GetReleaseDefinition(config.Context, release.GetReleaseDefinitionArgs{
		Project:      &projectID,
		DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

	var resultEnv release.ReleaseDefinitionEnvironment
	var stageIndex int
	var jobIndex int

	for k, env := range *res.Environments {
		if *env.Name == stageName {
			resultEnv = env
			stageIndex = k
			break
		}
	}

	if resultEnv.Name == nil {
		return fmt.Errorf("No stage with the name %s", stageName)
	}

	var job map[string]interface{}

	if name, ok := d.GetOk("job_name"); ok {
		for k, v := range *resultEnv.DeployPhases {
			v := v.(map[string]interface{})
			if v["name"].(string) == name {
				job = v
				jobIndex = k
				break
			}
		}

		if job == nil {
			return fmt.Errorf("No job found with %s", name)
		}
	} else if rank, ok := d.GetOk("job_rank"); ok {
		for k, v := range *resultEnv.DeployPhases {
			v := v.(map[string]interface{})

			if int(v["rank"].(float64)) == rank {
				job = v
				jobIndex = k
				break
			}
		}

		if job == nil {
			return fmt.Errorf("No job found with %d", rank)
		}
	} else {
		return fmt.Errorf("Either job_name or job_rank has to be set")
	}

	tasks := job["workflowTasks"].([]interface{})

	task := d.Get("task_info").([]interface{})[0].(map[string]interface{})
	alwaysRun := task["always_run"].(bool)
	condition := task["condition"].(string)
	enabled := task["enabled"].(bool)
	inputs := convertInterfaceToStringMap(task["inputs"].(map[string]interface{}))
	taskName := task["name"].(string)
	taskID, _ := uuid.Parse(task["task_id"].(string))
	version := task["version"].(string)
	continueOnError := task["continue_on_error"].(bool)
	definitionType := task["definition_type"].(string)

	newTask := release.WorkflowTask{
		AlwaysRun:       &alwaysRun,
		Condition:       &condition,
		Enabled:         &enabled,
		Inputs:          &inputs,
		Name:            &taskName,
		Version:         &version,
		TaskId:          &taskID,
		ContinueOnError: &continueOnError,
		DefinitionType:  &definitionType,
	}

	if rank, ok := d.GetOk("rank"); ok {
		if len(tasks) < rank.(int) {
			tasks = append(tasks, newTask)
		} else {
			tasks = append(tasks[:rank.(int)-1], append([]interface{}{newTask}, tasks[rank.(int)-1:]...)...)
		}
	} else if after, ok := d.GetOk("after"); ok {
		found := false
		for k, v := range tasks {
			v := v.(map[string]interface{})
			if v["name"].(string) == after.(string) {
				if k+1 == len(tasks) {
					tasks = append(tasks, newTask)
					found = true
					break
				} else {
					tasks = append(tasks[:k+1], append([]interface{}{newTask}, tasks[k+1:]...)...)
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("No target task of %s not found", after.(string))
		}
	} else if before, ok := d.GetOk("before"); ok {
		found := false
		for k, v := range tasks {
			v := v.(map[string]interface{})
			if v["name"].(string) == before.(string) {
				tasks = append(tasks[:k], append([]interface{}{newTask}, tasks[k:]...)...)
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("No target task of %s not found", before.(string))
		}
	} else {
		return fmt.Errorf("Either after or before or rank has to be set")
	}

	(*resultEnv.DeployPhases)[jobIndex].(map[string]interface{})["workflowTasks"] = tasks

	(*res.Environments)[stageIndex] = resultEnv

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: res,
	})

	if err != nil {
		return err
	}

	d.SetId(uuid.New().String())

	return resourceReleaseTaskRead(d, meta)
}

func resourceReleaseTaskUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	defID := d.Get("definition_id").(int)
	projectID := d.Get("project_id").(string)
	stageName := d.Get("stage_name").(string)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	res, err := releaseClient.GetReleaseDefinition(config.Context, release.GetReleaseDefinitionArgs{
		Project:      &projectID,
		DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

	var resultEnv release.ReleaseDefinitionEnvironment
	var stageIndex int
	var jobIndex int

	for k, env := range *res.Environments {
		if *env.Name == stageName {
			resultEnv = env
			stageIndex = k
			break
		}
	}

	if resultEnv.Name == nil {
		return fmt.Errorf("No stage with the name %s", stageName)
	}

	var job map[string]interface{}

	if name, ok := d.GetOk("job_name"); ok {
		for k, v := range *resultEnv.DeployPhases {
			v := v.(map[string]interface{})
			if v["name"].(string) == name {
				job = v
				jobIndex = k
				break
			}
		}

		if job == nil {
			return fmt.Errorf("No job found with %s", name)
		}
	} else if rank, ok := d.GetOk("job_rank"); ok {
		for k, v := range *resultEnv.DeployPhases {
			v := v.(map[string]interface{})

			if int(v["rank"].(float64)) == rank {
				job = v
				jobIndex = k
				break
			}
		}

		if job == nil {
			return fmt.Errorf("No job found with %d", rank)
		}
	} else {
		return fmt.Errorf("Either job_name or job_rank has to be set")
	}

	tasks := job["workflowTasks"].([]interface{})

	task := d.Get("task_info").([]interface{})[0].(map[string]interface{})
	alwaysRun := task["always_run"].(bool)
	condition := task["condition"].(string)
	enabled := task["enabled"].(bool)
	inputs := convertInterfaceToStringMap(task["inputs"].(map[string]interface{}))
	taskName := task["name"].(string)
	taskID, _ := uuid.Parse(task["task_id"].(string))
	version := task["version"].(string)
	continueOnError := task["continue_on_error"].(bool)
	definitionType := task["definition_type"].(string)

	newTask := release.WorkflowTask{
		AlwaysRun:       &alwaysRun,
		Condition:       &condition,
		Enabled:         &enabled,
		Inputs:          &inputs,
		Name:            &taskName,
		Version:         &version,
		TaskId:          &taskID,
		ContinueOnError: &continueOnError,
		DefinitionType:  &definitionType,
	}

	if rank, ok := d.GetOk("rank"); ok {
		if len(tasks) < rank.(int) {
			tasks[len(tasks)-1] = newTask
		} else {
			tasks[rank.(int)-1] = newTask
		}
	} else if after, ok := d.GetOk("after"); ok {
		found := false
		for k, v := range tasks {
			v := v.(map[string]interface{})
			if v["name"].(string) == after.(string) {
				if k+1 == len(tasks) {
					tasks[len(tasks)-1] = newTask
					found = true
					break
				} else {
					tasks[k+1] = newTask
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("No target task of %s not found", after.(string))
		}
	} else if before, ok := d.GetOk("before"); ok {
		found := false
		for k, v := range tasks {
			v := v.(map[string]interface{})
			if v["name"].(string) == before.(string) {
				tasks[k-1] = newTask
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("No target task of %s not found", before.(string))
		}
	} else {
		return fmt.Errorf("Either after or before or rank has to be set")
	}

	(*resultEnv.DeployPhases)[jobIndex].(map[string]interface{})["workflowTasks"] = tasks

	(*res.Environments)[stageIndex] = resultEnv

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: res,
	})

	if err != nil {
		return err
	}

	return resourceReleaseTaskRead(d, meta)
}

func resourceReleaseTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	defID := d.Get("definition_id").(int)
	projectID := d.Get("project_id").(string)
	stageName := d.Get("stage_name").(string)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	res, err := releaseClient.GetReleaseDefinition(config.Context, release.GetReleaseDefinitionArgs{
		Project:      &projectID,
		DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

	var resultEnv release.ReleaseDefinitionEnvironment

	for _, env := range *res.Environments {
		if *env.Name == stageName {
			resultEnv = env
			break
		}
	}

	if resultEnv.Name == nil {
		return fmt.Errorf("No stage with the name %s", stageName)
	}

	var job map[string]interface{}

	if name, ok := d.GetOk("job_name"); ok {
		for _, v := range *resultEnv.DeployPhases {
			v := v.(map[string]interface{})
			if v["name"].(string) == name {
				job = v
				break
			}
		}

		if job == nil {
			return fmt.Errorf("No job found with %s", name)
		}
	} else if rank, ok := d.GetOk("job_rank"); ok {
		for _, v := range *resultEnv.DeployPhases {
			v := v.(map[string]interface{})

			if int(v["rank"].(float64)) == rank {
				job = v
				break
			}
		}

		if job == nil {
			return fmt.Errorf("No job found with %d", rank)
		}
	} else {
		return fmt.Errorf("Either job_name or job_rank has to be set")
	}

	tasks := job["workflowTasks"].([]interface{})

	var task interface{}

	if rank, ok := d.GetOk("rank"); ok {
		if len(tasks) < rank.(int) {
			task = tasks[len(tasks)-1]
		} else {
			task = tasks[rank.(int)-1]
		}
	} else if after, ok := d.GetOk("after"); ok {
		found := false
		for k, v := range tasks {
			v := v.(map[string]interface{})
			if v["name"].(string) == after.(string) {
				if k+1 == len(tasks) {
					task = tasks[len(tasks)-1]
					found = true
					break
				} else {
					task = tasks[k+1]
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("No target task of %s not found", after.(string))
		}
	} else if before, ok := d.GetOk("before"); ok {
		found := false
		for k, v := range tasks {
			v := v.(map[string]interface{})
			if v["name"].(string) == before.(string) {
				task = tasks[k-1]
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("No target task of %s not found", before.(string))
		}
	} else {
		return fmt.Errorf("Either after or before or rank has to be set")
	}

	if err != nil {
		return err
	}

	finalTask := task.(map[string]interface{})

	d.Set("task_info", []map[string]interface{}{
		map[string]interface{}{
			"name":              finalTask["name"],
			"definition_type":   finalTask["definitionType"],
			"version":           finalTask["version"],
			"task_id":           finalTask["taskId"],
			"enabled":           finalTask["enabled"],
			"always_run":        finalTask["alwaysRun"],
			"continue_on_error": finalTask["continueOnError"],
			"condition":         finalTask["condition"],
			"inputs":            finalTask["inputs"],
		},
	})

	return nil
}

func resourceReleaseTaskDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	defID := d.Get("definition_id").(int)
	projectID := d.Get("project_id").(string)
	stageName := d.Get("stage_name").(string)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	res, err := releaseClient.GetReleaseDefinition(config.Context, release.GetReleaseDefinitionArgs{
		Project:      &projectID,
		DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

	var resultEnv release.ReleaseDefinitionEnvironment
	var stageIndex int
	var jobIndex int

	for k, env := range *res.Environments {
		if *env.Name == stageName {
			resultEnv = env
			stageIndex = k
			break
		}
	}

	if resultEnv.Name == nil {
		return fmt.Errorf("No stage with the name %s", stageName)
	}

	var job map[string]interface{}

	if name, ok := d.GetOk("job_name"); ok {
		for k, v := range *resultEnv.DeployPhases {
			v := v.(map[string]interface{})
			if v["name"].(string) == name {
				job = v
				jobIndex = k
				break
			}
		}

		if job == nil {
			return fmt.Errorf("No job found with %s", name)
		}
	} else if rank, ok := d.GetOk("job_rank"); ok {
		for k, v := range *resultEnv.DeployPhases {
			v := v.(map[string]interface{})

			if int(v["rank"].(float64)) == rank {
				job = v
				jobIndex = k
				break
			}
		}

		if job == nil {
			return fmt.Errorf("No job found with %d", rank)
		}
	} else {
		return fmt.Errorf("Either job_name or job_rank has to be set")
	}

	tasks := job["workflowTasks"].([]interface{})

	if rank, ok := d.GetOk("rank"); ok {
		if len(tasks) < rank.(int) {
			_, tasks = tasks[len(tasks)-1], tasks[:len(tasks)-1]
		} else {
			if rank.(int)-1 < len(tasks)-1 {
				copy(tasks[rank.(int)-1:], tasks[rank.(int):])
			}
			tasks[len(tasks)-1] = nil
			tasks = tasks[:len(tasks)-1]
		}
	} else if after, ok := d.GetOk("after"); ok {
		found := false
		for k, v := range tasks {
			v := v.(map[string]interface{})
			if v["name"].(string) == after.(string) {
				if k+1 == len(tasks) {
					_, tasks = tasks[len(tasks)-1], tasks[:len(tasks)-1]
					found = true
					break
				} else {
					if k+1 < len(tasks)-1 {
						copy(tasks[k+1:], tasks[k+2:])
					}
					tasks[len(tasks)-1] = nil
					tasks = tasks[:len(tasks)-1]
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("No target task of %s not found", after.(string))
		}
	} else if before, ok := d.GetOk("before"); ok {
		found := false
		for k, v := range tasks {
			v := v.(map[string]interface{})
			if v["name"].(string) == before.(string) {
				if k < len(tasks)-1 {
					copy(tasks[k-1:], tasks[k:])
				}
				tasks[len(tasks)-1] = nil
				tasks = tasks[:len(tasks)-1]
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("No target task of %s not found", before.(string))
		}
	} else {
		return fmt.Errorf("Either after or before or rank has to be set")
	}

	(*resultEnv.DeployPhases)[jobIndex].(map[string]interface{})["workflowTasks"] = tasks

	(*res.Environments)[stageIndex] = resultEnv

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: res,
	})

	if err != nil {
		return err
	}

	return nil
}

func resourceReleaseTaskImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	name := d.Id()
	if name == "" {
		return nil, fmt.Errorf("release task info cannot be empty")
	}

	res := strings.Split(name, "/")
	if len(res) != 5 {
		return nil, fmt.Errorf("the format has to be in <project-id>/<release-definition-id>/<stage-name>/<job-rank>/<task-rank>")
	}

	defID, err := strconv.Atoi(res[1])

	if err != nil {
		return nil, err
	}

	jobRank, err := strconv.Atoi(res[3])

	if err != nil {
		return nil, err
	}

	rank, err := strconv.Atoi(res[4])

	if err != nil {
		return nil, err
	}

	d.Set("project_id", res[0])
	d.Set("definition_id", defID)
	d.Set("stage_name", res[2])
	d.Set("job_rank", jobRank)
	d.Set("rank", rank)

	d.SetId(uuid.New().String())

	return []*schema.ResourceData{d}, nil
}
