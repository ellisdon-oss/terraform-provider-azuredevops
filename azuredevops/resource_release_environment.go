package azuredevops

import (
	"fmt"
	"github.com/ellisdon-oss/terraform-provider-azuredevops/azuredevops/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"strconv"
	"strings"
)

func resourceReleaseEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceReleaseEnvironmentCreate,
		Read:   resourceReleaseEnvironmentRead,
		Update: resourceReleaseEnvironmentUpdate,
		Delete: resourceReleaseEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			State: resourceReleaseEnvironmentImport,
		},
		Schema: map[string]*schema.Schema{
			"definition_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment": helper.EnvironmentSingleSchema(),
		},
	}
}

func resourceReleaseEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	definition_id, _ := strconv.ParseInt(d.Get("definition_id").(string), 10, 32)
	defID := int(definition_id)
	projectID := d.Get("project_id").(string)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	releaseDef, err := releaseClient.GetReleaseDefinition(config.Context, release.GetReleaseDefinitionArgs{
		Project:      &projectID,
		DefinitionId: &defID,
	})

	//
	if err != nil {
		d.SetId("")
		return err
	}

	envs := d.Get("environment").([]interface{})
	parsedEnvs := extractEnvironments(envs)
	combinedEnvs := append(*releaseDef.Environments, parsedEnvs...)
	releaseDef.Environments = &combinedEnvs

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		ReleaseDefinition: releaseDef,
		Project:           &projectID,
	})

	if err != nil {
		d.SetId("")
		return err
	}

	resultId := fmt.Sprintf("%s-%d", d.Get("project_id").(string), definition_id)

	tempEnvNames := "["
	for _, v := range parsedEnvs {
		tempEnvNames = fmt.Sprintf("%s %s", tempEnvNames, v.Name)
	}
	tempEnvNames = fmt.Sprintf("%s]", tempEnvNames)

	d.SetId(fmt.Sprintf("%s-%s", resultId, tempEnvNames))

	return resourceReleaseEnvironmentRead(d, meta)
}

func resourceReleaseEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	definition_id, _ := strconv.ParseInt(d.Get("definition_id").(string), 10, 32)

	defID := int(definition_id)
	projectID := d.Get("project_id").(string)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	releaseDef, err := releaseClient.GetReleaseDefinition(config.Context, release.GetReleaseDefinitionArgs{
		Project:      &projectID,
		DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

	envs := d.Get("environment").([]interface{})
	parsedEnvs := extractEnvironments(envs)
	modifiedEnvs := map[string]release.ReleaseDefinitionEnvironment{}

	for _, v := range parsedEnvs {
		modifiedEnvs[*v.Name] = v
	}
	tempEnvs := releaseDef.Environments

	names := []string{}

	for _, v := range parsedEnvs {
		names = append(names, *v.Name)
	}

	for k, v := range *releaseDef.Environments {
		if stringInSlice(*v.Name, names) {
			(*tempEnvs)[k] = modifiedEnvs[*v.Name]
		}
	}

	releaseDef.Environments = tempEnvs

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: releaseDef,
	})

	if err != nil {
		return err
	}

	return resourceReleaseEnvironmentRead(d, meta)
}

func resourceReleaseEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	definition_id, _ := strconv.ParseInt(d.Get("definition_id").(string), 10, 32)

	defID := int(definition_id)
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

	envs := d.Get("environment").([]interface{})
	parsedEnvs := extractEnvironments(envs)
	modifiedEnvs := map[string]release.ReleaseDefinitionEnvironment{}

	for _, v := range parsedEnvs {
		modifiedEnvs[*v.Name] = v
	}
	names := []string{}

	for _, v := range parsedEnvs {
		names = append(names, *v.Name)
	}

	result := []interface{}{}

	for _, v := range *res.Environments {
		if stringInSlice(*v.Name, names) {
			result = append(result, convertEnvToMap(v, modifiedEnvs[*v.Name]))
		}
	}

	d.Set("environment", result)
	return nil
}

func resourceReleaseEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	definition_id, _ := strconv.ParseInt(d.Get("definition_id").(string), 10, 32)

	defID := int(definition_id)
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

	envs := d.Get("environment").([]interface{})
	parsedEnvs := extractEnvironments(envs)
	names := []string{}

	for _, v := range parsedEnvs {
		names = append(names, *v.Name)
	}
	tempEnvs := res.Environments

	for k, v := range *res.Environments {
		if stringInSlice(*v.Name, names) {
			(*tempEnvs) = append((*tempEnvs)[:k], (*tempEnvs)[k+1:]...)
		}
	}

	res.Environments = tempEnvs

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: res,
	})

	if err != nil {
		return err
	}

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//func extractDemands(demands *schema.Set) []azuredevops.Demand {
//
//	var finalDemands []azuredevops.Demand
//
//	for _, v := range demands.List() {
//		v := v.(map[string]interface{})
//		finalDemands = append(finalDemands, azuredevops.Demand{
//			Name:  v["name"].(string),
//			Value: v["value"].(string),
//		})
//	}
//
//	return finalDemands
//}

//func convertInterfaceSliceToStringSlice(t []interface{}) []string {
//	s := make([]string, len(t))
//	for i, v := range t {
//		s[i] = fmt.Sprint(v)
//	}
//}

func resourceReleaseEnvironmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	name := d.Id()
	if name == "" {
		return nil, fmt.Errorf("release definition id cannot be empty")
	}

	res := strings.Split(name, "/")
	if len(res) != 2 {
		return nil, fmt.Errorf("the format has to be in <project-id>/<release-definition-id>")
	}

	d.Set("project_id", res[0])
	d.SetId(res[1])

	return []*schema.ResourceData{d}, nil
}

func convertEnvToMap(env release.ReleaseDefinitionEnvironment, oldEnv release.ReleaseDefinitionEnvironment) map[string]interface{} {
	result := make(map[string]interface{})
	testResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_secret": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}

	deployPhases := []interface{}{}
	preDeployApprovals := []interface{}{}
	conditions := []interface{}{}

	var variables []interface{}

	for k, v := range *env.DeployPhases {
		workflowTasks := []interface{}{}
		val := v.(map[string]interface{})["workflowTasks"].([]interface{})

		for k2, l := range val {
			inputs := make(map[string]interface{})
			if oldEnv.Name != nil && oldEnv.DeployPhases != nil && *oldEnv.Name != "" {
				phase := (*oldEnv.DeployPhases)[k]
				if k2 < len(phase.(map[string]interface{})["workflowTasks"].([]release.WorkflowTask)) && phase.(map[string]interface{})["workflowTasks"].([]release.WorkflowTask)[k2].Name == l.(map[string]interface{})["name"] {
					oldInputs := phase.(map[string]interface{})["workflowTasks"].([]release.WorkflowTask)[k2].Inputs
					for m := range *oldInputs {
						inputs[m] = l.(map[string]interface{})["inputs"].(map[string]interface{})[m]
					}
				} else {
					inputs = l.(map[string]interface{})["inputs"].(map[string]interface{})
				}
			} else {
				for k, v := range l.(map[string]interface{})["inputs"].(map[string]interface{}) {
					inputs[k] = v
				}
			}

			mapL := l.(map[string]interface{})

			workflowTasks = append(workflowTasks, map[string]interface{}{
				"name":              mapL["name"],
				"definition_type":   mapL["definitionType"],
				"version":           mapL["version"],
				"task_id":           mapL["taskId"],
				"enabled":           mapL["enabled"],
				"always_run":        mapL["alwaysRun"],
				"continue_on_error": mapL["continueOnError"],
				"condition":         mapL["condition"],
				"environment":       mapL["environment"],
				"inputs":            inputs,
			})
		}
		deploymentInput := make(map[string]interface{})

		if v.(map[string]interface{})["deploymentInput"].(map[string]interface{})["queueId"] != nil {
			deploymentInput["queueId"] = strconv.Itoa(int(v.(map[string]interface{})["deploymentInput"].(map[string]interface{})["queueId"].(float64)))
		}

		//log.Panic(deploymentInput)
		deployPhases = append(deployPhases, map[string]interface{}{
			"phase_type":       v.(map[string]interface{})["phaseType"],
			"name":             v.(map[string]interface{})["name"],
			"rank":             v.(map[string]interface{})["rank"],
			"deployment_input": deploymentInput,
			"workflow_task":    workflowTasks,
		})
	}

	approvals := []interface{}{}

	for _, v := range *env.PreDeployApprovals.Approvals {
		var approverID string
		if v.Approver == nil {
			approverID = ""
		} else {
			approverID = *(*v.Approver).Id
		}
		approvals = append(approvals, map[string]interface{}{
			"approver_id":        approverID,
			"is_automated":       *v.IsAutomated,
			"is_notification_on": *v.IsNotificationOn,
		})
	}

	options := []interface{}{
		map[string]interface{}{
			"execution_order":                 *env.PreDeployApprovals.ApprovalOptions.ExecutionOrder,
			"timeout_in_minutes":              *env.PreDeployApprovals.ApprovalOptions.TimeoutInMinutes,
			"release_creator_can_be_approver": *env.PreDeployApprovals.ApprovalOptions.ReleaseCreatorCanBeApprover,
		},
	}

	preDeployApprovals = append(preDeployApprovals, map[string]interface{}{
		"options":   options,
		"approvals": approvals,
	})

	for _, v := range *env.Conditions {
		conditions = append(conditions, map[string]interface{}{
			"condition_type": *v.ConditionType,
			"name":           *v.Name,
			"value":          *v.Value,
		})
	}

	var oldVariables map[string]release.ConfigurationVariableValue

	if oldEnv.Name != nil && *oldEnv.Name != "" {
		oldVariables = *oldEnv.Variables
	}

	for k, v := range *env.Variables {
		if k != "" {
			value := ""

			var isSecret bool
			if v.IsSecret != nil && *v.IsSecret {
				isSecret = true
			} else {
				isSecret = false
			}

			if isSecret {
				if *oldEnv.Name != "" {
					value = *oldVariables[k].Value
				}
			} else {
				value = *v.Value
			}

			variables = append(variables, map[string]interface{}{
				"is_secret": isSecret,
				"value":     value,
				"name":      k,
			})
		}
	}

	result = map[string]interface{}{
		"name":                *env.Name,
		"rank":                *env.Rank,
		"deploy_phase":        deployPhases,
		"pre_deploy_approval": preDeployApprovals,
		"condition":           conditions,
		"variable":            schema.NewSet(schema.HashResource(testResource), variables),
		"variable_groups":     *env.VariableGroups,
	}

	return result
}
