package azuredevops

import (
	"fmt"
	"github.com/ellisdon-oss/terraform-provider-azuredevops/azuredevops/helper"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
)

func resourceReleaseTasks() *schema.Resource {
	return &schema.Resource{
		Create: resourceReleaseTasksCreate,
		Read:   resourceReleaseTasksRead,
		Update: resourceReleaseTasksUpdate,
		Delete: resourceReleaseTasksDelete,
		Schema: map[string]*schema.Schema{
			"task": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"job_rank": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"job_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"rank": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"before": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"after": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"stage_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"task_info": helper.WorkflowTaskSingleSchema(),
					},
				},
			},
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

func resourceReleaseTasksCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	defID := d.Get("definition_id").(int)
	projectID := d.Get("project_id").(string)

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

	userTasks := d.Get("task").([]interface{})

	for k, userTask := range userTasks {
		userTask := userTask.(map[string]interface{})

		stageName := userTask["stage_name"].(string)

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

		_, isJobNameSet := d.GetOk(fmt.Sprintf("task.%d.job_name", k))
		_, isJobRankSet := d.GetOk(fmt.Sprintf("task.%d.job_rank", k))

		if isJobNameSet && isJobRankSet {
			return fmt.Errorf("Can't set both job_name and job_rank at the same time")
		}

		_, isRankSet := d.GetOk(fmt.Sprintf("task.%d.rank", k))
		_, isAfterSet := d.GetOk(fmt.Sprintf("task.%d.after", k))
		_, isBeforeSet := d.GetOk(fmt.Sprintf("task.%d.before", k))
		amountSet := 0
		for _, v := range []bool{isRankSet, isAfterSet, isBeforeSet} {
			if v {
				amountSet++
			}
		}

		if amountSet > 1 {
			return fmt.Errorf("Can only set one of rank, after, or before")
		}

		if name, ok := d.GetOk(fmt.Sprintf("task.%d.job_name", k)); ok {
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
		} else if rank, ok := d.GetOk(fmt.Sprintf("task.%d.job_rank", k)); ok {
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

		task := userTask["task_info"].([]interface{})[0].(map[string]interface{})
		alwaysRun := task["always_run"].(bool)
		condition := task["condition"].(string)
		enabled := task["enabled"].(bool)
		inputs := convertInterfaceToStringMap(task["inputs"].(map[string]interface{}))
		taskName := task["name"].(string)
		taskID, _ := uuid.Parse(task["task_id"].(string))
		version := task["version"].(string)
		refName := task["ref_name"].(string)
		continueOnError := task["continue_on_error"].(bool)
		definitionType := task["definition_type"].(string)

		newTask := map[string]interface{}{
			"alwaysRun":       alwaysRun,
			"condition":       condition,
			"enabled":         enabled,
			"inputs":          inputs,
			"name":            taskName,
			"version":         version,
			"taskId":          taskID,
			"continueOnError": continueOnError,
			"definitionType":  definitionType,
			"refName":         refName,
		}

		if rank, ok := d.GetOk(fmt.Sprintf("task.%d.rank", k)); ok {
			if len(tasks) < rank.(int) || rank.(int) == -1 {
				tasks = append(tasks, newTask)
			} else {
				tasks = append(tasks[:rank.(int)-1], append([]interface{}{newTask}, tasks[rank.(int)-1:]...)...)
			}
		} else if after, ok := d.GetOk(fmt.Sprintf("task.%d.after", k)); ok {
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
		} else if before, ok := d.GetOk(fmt.Sprintf("task.%d.before", k)); ok {
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
	}

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: res,
	})

	if err != nil {
		return err
	}

	d.SetId(uuid.New().String())

	return resourceReleaseTasksRead(d, meta)
}

func resourceReleaseTasksUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	defID := d.Get("definition_id").(int)
	projectID := d.Get("project_id").(string)

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

	userTasks := d.Get("task").([]interface{})

	for k, userTask := range userTasks {
		userTask := userTask.(map[string]interface{})

		stageName := userTask["stage_name"].(string)

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

		_, isJobNameSet := d.GetOk(fmt.Sprintf("task.%d.job_name", k))
		_, isJobRankSet := d.GetOk(fmt.Sprintf("task.%d.job_rank", k))

		if isJobNameSet && isJobRankSet {
			return fmt.Errorf("Can't set both job_name and job_rank at the same time")
		}

		_, isRankSet := d.GetOk(fmt.Sprintf("task.%d.rank", k))
		_, isAfterSet := d.GetOk(fmt.Sprintf("task.%d.after", k))
		_, isBeforeSet := d.GetOk(fmt.Sprintf("task.%d.before", k))
		amountSet := 0
		for _, v := range []bool{isRankSet, isAfterSet, isBeforeSet} {
			if v {
				amountSet++
			}
		}

		if amountSet > 1 {
			return fmt.Errorf("Can only set one of rank, after, or before")
		}

		if name, ok := d.GetOk(fmt.Sprintf("task.%d.job_name", k)); ok {
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
		} else if rank, ok := d.GetOk(fmt.Sprintf("task.%d.job_rank", k)); ok {
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

		task := userTask["task_info"].([]interface{})[0].(map[string]interface{})
		alwaysRun := task["always_run"].(bool)
		condition := task["condition"].(string)
		enabled := task["enabled"].(bool)
		inputs := convertInterfaceToStringMap(task["inputs"].(map[string]interface{}))
		taskName := task["name"].(string)
		taskID, _ := uuid.Parse(task["task_id"].(string))
		version := task["version"].(string)
		refName := task["ref_name"].(string)
		continueOnError := task["continue_on_error"].(bool)
		definitionType := task["definition_type"].(string)

		newTask := map[string]interface{}{
			"alwaysRun":       alwaysRun,
			"condition":       condition,
			"enabled":         enabled,
			"inputs":          inputs,
			"name":            taskName,
			"version":         version,
			"taskId":          taskID,
			"continueOnError": continueOnError,
			"definitionType":  definitionType,
			"refName":  refName,
		}

		if rank, ok := d.GetOk(fmt.Sprintf("task.%d.rank", k)); ok {
			if rank.(int) == -1 || len(tasks) < rank.(int) {
				tasks[len(tasks)-1] = newTask
			} else {
				tasks[rank.(int)-1] = newTask
			}
		} else if after, ok := d.GetOk(fmt.Sprintf("task.%d.after", k)); ok {
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
		} else if before, ok := d.GetOk(fmt.Sprintf("task.%d.before", k)); ok {
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
	}

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: res,
	})

	if err != nil {
		return err
	}

	return resourceReleaseTasksRead(d, meta)
}

func resourceReleaseTasksRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	defID := d.Get("definition_id").(int)
	projectID := d.Get("project_id").(string)

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

	userTasks := d.Get("task").([]interface{})

	finalUserTasks := []interface{}{}
	for k, userTask := range userTasks {
		userTask := userTask.(map[string]interface{})

		stageName := userTask["stage_name"].(string)

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

		_, isJobNameSet := d.GetOk(fmt.Sprintf("task.%d.job_name", k))
		_, isJobRankSet := d.GetOk(fmt.Sprintf("task.%d.job_rank", k))

		if isJobNameSet && isJobRankSet {
			return fmt.Errorf("Can't set both job_name and job_rank at the same time")
		}

		_, isRankSet := d.GetOk(fmt.Sprintf("task.%d.rank", k))
		_, isAfterSet := d.GetOk(fmt.Sprintf("task.%d.after", k))
		_, isBeforeSet := d.GetOk(fmt.Sprintf("task.%d.before", k))
		amountSet := 0
		for _, v := range []bool{isRankSet, isAfterSet, isBeforeSet} {
			if v {
				amountSet++
			}
		}

		if amountSet > 1 {
			return fmt.Errorf("Can only set one of rank, after, or before")
		}

		if name, ok := d.GetOk(fmt.Sprintf("task.%d.job_name", k)); ok {
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
		} else if rank, ok := d.GetOk(fmt.Sprintf("task.%d.job_rank", k)); ok {
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

		if rank, ok := d.GetOk(fmt.Sprintf("task.%d.rank", k)); ok {
			if rank.(int) == -1 || len(tasks) < rank.(int) {
				task = tasks[len(tasks)-1]
			} else {
				task = tasks[rank.(int)-1]
			}
		} else if after, ok := d.GetOk(fmt.Sprintf("task.%d.after", k)); ok {
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
		} else if before, ok := d.GetOk(fmt.Sprintf("task.%d.before", k)); ok {
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
		userTask["task_info"] = []map[string]interface{}{
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
				"ref_name":            finalTask["refName"],
			},
		}

		finalUserTasks = append(finalUserTasks, userTask)
	}

	d.Set("task", finalUserTasks)

	return nil
}

func resourceReleaseTasksDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	defID := d.Get("definition_id").(int)
	projectID := d.Get("project_id").(string)

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
	userTasks := d.Get("task").([]interface{})

	for k, userTask := range reverseSlice(userTasks) {
		k := len(userTasks) - 1 - k
		userTask := userTask.(map[string]interface{})

		stageName := userTask["stage_name"].(string)

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

		_, isJobNameSet := d.GetOk(fmt.Sprintf("task.%d.job_name", k))
		_, isJobRankSet := d.GetOk(fmt.Sprintf("task.%d.job_rank", k))

		if isJobNameSet && isJobRankSet {
			return fmt.Errorf("Can't set both job_name and job_rank at the same time")
		}

		_, isRankSet := d.GetOk(fmt.Sprintf("task.%d.rank", k))
		_, isAfterSet := d.GetOk(fmt.Sprintf("task.%d.after", k))
		_, isBeforeSet := d.GetOk(fmt.Sprintf("task.%d.before", k))
		amountSet := 0
		for _, v := range []bool{isRankSet, isAfterSet, isBeforeSet} {
			if v {
				amountSet++
			}
		}

		if amountSet > 1 {
			return fmt.Errorf("Can only set one of rank, after, or before")
		}
		if name, ok := d.GetOk(fmt.Sprintf("task.%d.job_name", k)); ok {
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
		} else if rank, ok := d.GetOk(fmt.Sprintf("task.%d.job_rank", k)); ok {
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

		if rank, ok := d.GetOk(fmt.Sprintf("task.%d.rank", k)); ok {
			if rank.(int) == -1 || len(tasks) < rank.(int) {
				_, tasks = tasks[len(tasks)-1], tasks[:len(tasks)-1]
      } else {
        if rank.(int)-1 < len(tasks)-1 {
          copy(tasks[rank.(int)-1:], tasks[rank.(int):])
        }
        tasks[len(tasks)-1] = nil
        tasks = tasks[:len(tasks)-1]
      }
    } else if after, ok := d.GetOk(fmt.Sprintf("task.%d.after", k)); ok {
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
		} else if before, ok := d.GetOk(fmt.Sprintf("task.%d.before", k)); ok {
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
	}

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: res,
	})

	if err != nil {
		return err
	}

	return nil
}

// Ref: https://www.socketloop.com/tutorials/golang-how-to-reverse-slice-or-array-elements-order
func reverseSlice(items []interface{}) []interface{} {
	reversed := []interface{}{}

	for i := range items {
		n := items[len(items)-1-i]
		reversed = append(reversed, n)
	}

	return reversed
}
