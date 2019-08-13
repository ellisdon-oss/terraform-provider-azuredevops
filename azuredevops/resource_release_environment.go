package azuredevops

import (
	"fmt"
	"github.com/ellisdon/azuredevops-go"
	"github.com/ellisdon/terraform-provider-azuredevops/azuredevops/helper"
	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/hashicorp/terraform/helper/validation"

	"github.com/pkg/errors"
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

	releaseDef, _, err := config.Client.ReleaseDefinitionsApi.GetReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), int32(definition_id), config.ApiVersion, nil)

	//
	if err != nil {
		d.SetId("")
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	envs := d.Get("environment").([]interface{})
	parsedEnvs := extractEnvironments(envs)
	releaseDef.Environments = append(releaseDef.Environments, parsedEnvs...)

	_, _, err2 := config.Client.ReleaseDefinitionsApi.UpdateReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, releaseDef)
	//

	if err2 != nil {
		d.SetId("")
		return errors.New(string(err2.(azuredevops.GenericOpenAPIError).Body()))
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

	releaseDef, _, err := config.Client.ReleaseDefinitionsApi.GetReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), int32(definition_id), config.ApiVersion, nil)
	//
	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	envs := d.Get("environment").([]interface{})
	parsedEnvs := extractEnvironments(envs)
	modifiedEnvs := map[string]azuredevops.ReleaseDefinitionEnvironment{}

	for _, v := range parsedEnvs {
		modifiedEnvs[v.Name] = v
	}
	tempEnvs := releaseDef.Environments

	names := []string{}

	for _, v := range parsedEnvs {
		names = append(names, v.Name)
	}

	for k, v := range releaseDef.Environments {
		if stringInSlice(v.Name, names) {
			tempEnvs[k] = modifiedEnvs[v.Name]
		}
	}

	releaseDef.Environments = tempEnvs

	_, _, err2 := config.Client.ReleaseDefinitionsApi.UpdateReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, releaseDef)
	//

	if err2 != nil {
		return errors.New(string(err2.(azuredevops.GenericOpenAPIError).Body()))
	}

	return resourceReleaseEnvironmentRead(d, meta)
}

func resourceReleaseEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	definition_id, _ := strconv.ParseInt(d.Get("definition_id").(string), 10, 32)

	res, _, err := config.Client.ReleaseDefinitionsApi.GetReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), int32(definition_id), config.ApiVersion, nil)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	envs := d.Get("environment").([]interface{})
	parsedEnvs := extractEnvironments(envs)
	modifiedEnvs := map[string]azuredevops.ReleaseDefinitionEnvironment{}

	for _, v := range parsedEnvs {
		modifiedEnvs[v.Name] = v
	}
	names := []string{}

	for _, v := range parsedEnvs {
		names = append(names, v.Name)
	}

	result := []interface{}{}

	for _, v := range res.Environments {
		if stringInSlice(v.Name, names) {
			result = append(result, convertEnvToMap(v, modifiedEnvs[v.Name]))
		}
	}

	d.Set("environment", result)
	return nil
}

func resourceReleaseEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	definition_id, _ := strconv.ParseInt(d.Get("definition_id").(string), 10, 32)

	res, _, err := config.Client.ReleaseDefinitionsApi.GetReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), int32(definition_id), config.ApiVersion, nil)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	envs := d.Get("environment").([]interface{})
	parsedEnvs := extractEnvironments(envs)
	names := []string{}

	for _, v := range parsedEnvs {
		names = append(names, v.Name)
	}
	tempEnvs := res.Environments

	for k, v := range res.Environments {
		if stringInSlice(v.Name, names) {
			tempEnvs = append(tempEnvs[:k], tempEnvs[k+1:]...)
		}
	}

	res.Environments = tempEnvs

	_, _, err2 := config.Client.ReleaseDefinitionsApi.UpdateReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, res)
	//

	if err2 != nil {
		return errors.New(string(err2.(azuredevops.GenericOpenAPIError).Body()))
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

func convertEnvToMap(env azuredevops.ReleaseDefinitionEnvironment, oldEnv azuredevops.ReleaseDefinitionEnvironment) map[string]interface{} {
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
				Type:      schema.TypeString,
				Sensitive: true,
				Required:  true,
			},
		},
	}

	deploy_phases := []interface{}{}
	preDeployApprovals := []interface{}{}
	conditions := []interface{}{}
	var variables []interface{}

	for k, v := range env.DeployPhases {
		workflowTasks := []interface{}{}

		for k2, l := range v.WorkflowTasks {
			inputs := make(map[string]interface{})
			for k, _ := range oldEnv.DeployPhases[k].WorkflowTasks[k2].Inputs {
				inputs[k] = l.Inputs[k]
			}

			workflowTasks = append(workflowTasks, map[string]interface{}{
				"name":              l.Name,
				"definition_type":   l.DefinitionType,
				"version":           l.Version,
				"task_id":           l.TaskId,
				"enabled":           l.Enabled,
				"always_run":        l.AlwaysRun,
				"continue_on_error": l.ContinueOnError,
				"condition":         l.Condition,
				"inputs":            inputs,
			})
		}
		deploymentInput := make(map[string]interface{})

		deploymentInput["queueId"] = strconv.Itoa(int(v.DeploymentInput["queueId"].(float64)))

		//log.Panic(deploymentInput)
		deploy_phases = append(deploy_phases, map[string]interface{}{
			"phase_type":       v.PhaseType,
			"name":             v.Name,
			"rank":             v.Rank,
			"deployment_input": deploymentInput,
			"workflow_task":    workflowTasks,
		})
	}

	approvals := []interface{}{}
	for _, v := range env.PreDeployApprovals.Approvals {
		approvals = append(approvals, map[string]interface{}{
			"approver_id":        v.Approver.Id,
			"is_automated":       v.IsAutomated,
			"is_notification_on": v.IsNotificationOn,
		})
	}

	options := []interface{}{
		map[string]interface{}{
			"execution_order":                 env.PreDeployApprovals.ApprovalOptions.ExecutionOrder,
			"timeout_in_minutes":              env.PreDeployApprovals.ApprovalOptions.TimeoutInMinutes,
			"release_creator_can_be_approver": env.PreDeployApprovals.ApprovalOptions.ReleaseCreatorCanBeApprover,
		},
	}

	preDeployApprovals = append(preDeployApprovals, map[string]interface{}{
		"options":   options,
		"approvals": approvals,
	})

	for _, v := range env.Conditions {
		conditions = append(conditions, map[string]interface{}{
			"condition_type": v.ConditionType,
			"name":           v.Name,
			"value":          v.Value,
		})
	}

	old_variables := oldEnv.Variables
	for k, v := range env.Variables {
		if k != "" {
			value := ""
			if v.IsSecret {
				value = old_variables[k].Value
			} else {
				value = v.Value
			}
			variables = append(variables, map[string]interface{}{
				"is_secret": v.IsSecret,
				"value":     value,
				"name":      k,
			})
		}
	}

	result = map[string]interface{}{
		"name":                env.Name,
		"rank":                env.Rank,
		"deploy_phase":        deploy_phases,
		"pre_deploy_approval": preDeployApprovals,
		"condition":           conditions,
		"variable":            schema.NewSet(schema.HashResource(testResource), variables),
	}

	return result
}
