package azuredevops

import (
	"encoding/json"
	"fmt"
	"github.com/ellisdon/terraform-provider-azuredevops/azuredevops/helper"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
	"strconv"
	"strings"
)

func resourceReleaseDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceReleaseDefinitionCreate,
		Read:   resourceReleaseDefinitionRead,
		Update: resourceReleaseDefinitionUpdate,
		Delete: resourceReleaseDefinitionDelete,
		Importer: &schema.ResourceImporter{
			State: resourceReleaseDefinitionImport,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "\\",
			},
			"environment":      helper.EnvironmentResourceSchema(),
			"release_variable": helper.ReleaseVariableSchema(),
			"release_variable_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"artifact": helper.ArtifactSchema(),
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceReleaseDefinitionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	envs := d.Get("environment").([]interface{})
	name := d.Get("name").(string)
	path := d.Get("path").(string)
	extractedEnvs := extractEnvironments(envs)
	newReleaseDefinition := release.ReleaseDefinition{
		Name:         &name,
		Path:         &path,
		Environments: &extractedEnvs,
	}

	if v, ok := d.GetOk("release_variable"); ok {
		releaseVariables := extractReleaseVariables(v.(*schema.Set))
		newReleaseDefinition.Variables = &releaseVariables
	}

	if l, ok := d.GetOk("release_variable_groups"); ok {
		var varGroups []int

		for _, v := range l.([]interface{}) {
			varGroups = append(varGroups, v.(int))
		}

		newReleaseDefinition.VariableGroups = &varGroups
	}

	if v, ok := d.GetOk("artifact"); ok {
		artifacts := extractArtifact(v.(*schema.Set))
		newReleaseDefinition.Artifacts = &artifacts
	}

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)

	releaseDef, err := releaseClient.CreateReleaseDefinition(config.Context, release.CreateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: &newReleaseDefinition,
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprint(*releaseDef.Id))

	return resourceReleaseDefinitionRead(d, meta)
}

func resourceReleaseDefinitionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	envs := d.Get("environment").([]interface{})

	id, _ := strconv.ParseInt(d.Id(), 10, 32)

	defID := int(id)
	name := d.Get("name").(string)
	path := d.Get("path").(string)
	revision := d.Get("revision").(int)
	extractedEnvs := extractEnvironments(envs)

	newReleaseDefinition := release.ReleaseDefinition{
		Id:           &defID,
		Name:         &name,
		Path:         &path,
		Environments: &extractedEnvs,
		Revision:     &revision,
	}

	if v, ok := d.GetOk("release_variable"); ok {
		releaseVariables := extractReleaseVariables(v.(*schema.Set))
		newReleaseDefinition.Variables = &releaseVariables
	}

	if l, ok := d.GetOk("release_variable_groups"); ok {
		var varGroups []int
		for _, v := range l.([]interface{}) {
			varGroups = append(varGroups, v.(int))

		}
		newReleaseDefinition.VariableGroups = &varGroups
	}

	if v, ok := d.GetOk("artifact"); ok {
		artifacts := extractArtifact(v.(*schema.Set))
		newReleaseDefinition.Artifacts = &artifacts
	}

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: &newReleaseDefinition,
	})

	if err != nil {
		return err
	}

	return resourceReleaseDefinitionRead(d, meta)
}

func resourceReleaseDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, _ := strconv.ParseInt(d.Id(), 10, 32)

	defID := int(id)
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

	d.Set("revision", *res.Revision)
	d.Set("name", *res.Name)
	d.Set("path", *res.Path)
	d.Set("release_variable_groups", *res.VariableGroups)

	if v, ok := d.GetOk("release_variable"); ok {
		d.Set("release_variable", readReleaseVariables(v.(*schema.Set), *res.Variables))
	} else {
		d.Set("release_variable", readReleaseVariables(nil, *res.Variables))
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
		result = append(result, convertEnvToMap(v, modifiedEnvs[*v.Name]))
	}

	d.Set("environment", result)
	d.Set("artifact", readArtifact(*res.Artifacts))

	return nil
}

func resourceReleaseDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, _ := strconv.ParseInt(d.Id(), 10, 32)
	defID := int(id)
	projectID := d.Get("project_id").(string)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	err = releaseClient.DeleteReleaseDefinition(config.Context, release.DeleteReleaseDefinitionArgs{
		Project:      &projectID,
		DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

	return nil
}

func extractEnvironments(environments []interface{}) []release.ReleaseDefinitionEnvironment {
	var result []release.ReleaseDefinitionEnvironment
	for _, env := range environments {
		env := env.(map[string]interface{})
		//{
		//          "deploymentInput": {
		//            "parallelExecution": {
		//              "parallelExecutionType": "none"
		//            },
		//            "skipArtifactsDownload": false,
		//            "artifactsDownloadInput": {},
		//            "queueId": 15,
		//            "demands": [],
		//            "enableAccessToken": false,
		//            "timeoutInMinutes": 0,
		//            "jobCancelTimeoutInMinutes": 1,
		//            "condition": "succeeded()",
		//            "overrideInputs": {}
		//          },
		//          "rank": 1,
		//          "phaseType": "agentBasedDeployment",
		//          "name": "Run on agent",
		//          "workflowTasks": []
		//        }
		conditions := env["condition"].([]interface{})
		var finalConditions []release.Condition
		for _, v := range conditions {

			v := v.(map[string]interface{})

			conditionType := release.ConditionType(v["condition_type"].(string))
			conditionName := v["name"].(string)
			conditionValue := v["value"].(string)
			finalConditions = append(finalConditions, release.Condition{
				ConditionType: &conditionType,
				Name:          &conditionName,
				Value:         &conditionValue,
			})
		}

		var predeployapprovals []interface{}
		var finalPreDeployApprovals []release.ReleaseDefinitionApprovalStep

		finalApprovalOptions := release.ApprovalOptions{}

		if env["pre_deploy_approval"] != nil && len(env["pre_deploy_approval"].([]interface{})) != 0 {

			predeployapprovals = env["pre_deploy_approval"].([]interface{})[0].(map[string]interface{})["approvals"].([]interface{})

			if env["pre_deploy_approval"].([]interface{})[0].(map[string]interface{})["options"] != nil {
				options := env["pre_deploy_approval"].([]interface{})[0].(map[string]interface{})["options"]
				if len(options.([]interface{})) != 0 && options.([]interface{})[0].(map[string]interface{})["execution_order"] != nil {
					executionOrder := release.ApprovalExecutionOrder(options.([]interface{})[0].(map[string]interface{})["execution_order"].(string))
					finalApprovalOptions.ExecutionOrder = &executionOrder
				} else {
					executionOrder := release.ApprovalExecutionOrder("beforeGates")
					finalApprovalOptions.ExecutionOrder = &executionOrder
				}

				if len(options.([]interface{})) != 0 && options.([]interface{})[0].(map[string]interface{})["timeout_in_minutes"] != nil {
					timeoutInMinutes := options.([]interface{})[0].(map[string]interface{})["timeout_in_minutes"].(int)
					finalApprovalOptions.TimeoutInMinutes = &timeoutInMinutes
				}

				if len(options.([]interface{})) != 0 && options.([]interface{})[0].(map[string]interface{})["release_creator_can_be_approver"] != nil {
					releaseCreatorCanBeApprover := options.([]interface{})[0].(map[string]interface{})["release_creator_can_be_approver"].(bool)
					finalApprovalOptions.ReleaseCreatorCanBeApprover = &releaseCreatorCanBeApprover
				} else {
					releaseCreatorCanBeApprover := false
					finalApprovalOptions.ReleaseCreatorCanBeApprover = &releaseCreatorCanBeApprover
				}
			}
		} else {
			rank := 1
			isAutomated := true
			isNotificationOn := false
			finalPreDeployApprovals = []release.ReleaseDefinitionApprovalStep{
				release.ReleaseDefinitionApprovalStep{
					Rank:             &rank,
					IsAutomated:      &isAutomated,
					IsNotificationOn: &isNotificationOn,
				},
			}
		}

		globalApprovalRank := 1

		for _, v := range predeployapprovals {

			v := v.(map[string]interface{})
			approverID := v["approver_id"].(string)
			isAutomated := v["is_automated"].(bool)
			isNotificationOn := v["is_notification_on"].(bool)
			approvalRank := globalApprovalRank

			finalPreDeployApprovals = append(finalPreDeployApprovals, release.ReleaseDefinitionApprovalStep{
				Rank: &approvalRank,
				Approver: &webapi.IdentityRef{
					Id: &approverID,
				},
				IsAutomated:      &isAutomated,
				IsNotificationOn: &isNotificationOn,
			})

			globalApprovalRank++
		}

		deployPhases := env["deploy_phase"].([]interface{})
		var finalDeployPhases []interface{}
		for _, v := range deployPhases {

			v := v.(map[string]interface{})

			var finalTasks []release.WorkflowTask

			tasks := v["workflow_task"].([]interface{})
			for _, task := range tasks {
				task := task.(map[string]interface{})
				alwaysRun := task["always_run"].(bool)
				condition := task["condition"].(string)
				enabled := task["enabled"].(bool)
				inputs := convertInterfaceToStringMap(task["inputs"].(map[string]interface{}))
				taskName := task["name"].(string)
				taskID, _ := uuid.Parse(task["task_id"].(string))
				version := task["version"].(string)
				continueOnError := task["continue_on_error"].(bool)
				definitionType := task["definition_type"].(string)

				finalTasks = append(finalTasks, release.WorkflowTask{
					AlwaysRun:       &alwaysRun,
					Condition:       &condition,
					Enabled:         &enabled,
					Inputs:          &inputs,
					Name:            &taskName,
					Version:         &version,
					TaskId:          &taskID,
					ContinueOnError: &continueOnError,
					DefinitionType:  &definitionType,
				})
			}

			finalDeployPhases = append(finalDeployPhases, map[string]interface{}{
				"deploymentInput": v["deployment_input"].(map[string]interface{}),
				"phaseType":       v["phase_type"].(string),
				"rank":            v["rank"].(int),
				"name":            v["name"].(string),
				"workflowTasks":   finalTasks,
			})
		}

		var varGroups []int

		if l := env["variable_groups"]; l != nil {
			for _, v := range l.([]interface{}) {
				varGroups = append(varGroups, v.(int))
			}
		}

		envName := env["name"].(string)
		rank := env["rank"].(int)
		postRank := 1

		daysToKeep := 30
		releasesToKeep := 3
		retainBuild := true

		isAutomated := true
		isNotificationOn := false
		releaseVariables := extractReleaseVariables(env["variable"].(*schema.Set))

		result = append(result, release.ReleaseDefinitionEnvironment{
			Name:           &envName,
			Conditions:     &finalConditions,
			DeployPhases:   &finalDeployPhases,
			Rank:           &rank,
			VariableGroups: &varGroups,
			RetentionPolicy: &release.EnvironmentRetentionPolicy{
				DaysToKeep:     &daysToKeep,
				ReleasesToKeep: &releasesToKeep,
				RetainBuild:    &retainBuild,
			},
			PostDeployApprovals: &release.ReleaseDefinitionApprovals{
				Approvals: &[]release.ReleaseDefinitionApprovalStep{
					release.ReleaseDefinitionApprovalStep{
						Rank:             &postRank,
						IsAutomated:      &isAutomated,
						IsNotificationOn: &isNotificationOn,
					},
				},
			},
			Variables: &releaseVariables,
			PreDeployApprovals: &release.ReleaseDefinitionApprovals{
				ApprovalOptions: &finalApprovalOptions,
				Approvals:       &finalPreDeployApprovals,
			},
		})
	}

	return result
}

//func extractDemands(demands *schema.Set) []release.Demand {
//
//	var finalDemands []release.Demand
//
//	for _, v := range demands.List() {
//		v := v.(map[string]interface{})
//		finalDemands = append(finalDemands, release.Demand{
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

func resourceReleaseDefinitionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

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

func readReleaseVariables(oldVars *schema.Set, vars map[string]release.ConfigurationVariableValue) *schema.Set {
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

	var variables []interface{}

	old_variables := make(map[string]string)
	if oldVars != nil {
		for _, v := range oldVars.List() {
			v := v.(map[string]interface{})
			old_variables[v["name"].(string)] = v["value"].(string)
		}
	}

	for k, v := range vars {
		if k != "" {
			value := ""

			var isSecret bool
			if v.IsSecret != nil && *v.IsSecret {
				isSecret = true
			} else {
				isSecret = false
			}

			if isSecret {
				if oldVars != nil {
					value = old_variables[k]
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

	res := schema.NewSet(schema.HashResource(testResource), variables)
	return res
}

func extractReleaseVariables(variables *schema.Set) map[string]release.ConfigurationVariableValue {

	finalVariables := make(map[string]release.ConfigurationVariableValue)

	for _, value := range variables.List() {
		variableMap := value.(map[string]interface{})
		isSecret := variableMap["is_secret"].(bool)
		varValue := variableMap["value"].(string)

		finalVariables[variableMap["name"].(string)] = release.ConfigurationVariableValue{
			IsSecret: &isSecret,
			Value:    &varValue,
		}
	}

	return finalVariables
}

func extractArtifact(variables *schema.Set) []release.Artifact {

	var finalArtifacts []release.Artifact
	for _, value := range variables.List() {
		artifact := value.(map[string]interface{})
		refs := make(map[string]release.ArtifactSourceReference)

		ref := artifact["definition_reference"].(string)

		json.Unmarshal([]byte(ref), &refs)
		alias := artifact["alias"].(string)
		sourceID := artifact["source_id"].(string)
		artifactType := artifact["type"].(string)

		finalArtifacts = append(finalArtifacts, release.Artifact{
			DefinitionReference: &refs,
			Alias:               &alias,
			SourceId:            &sourceID,
			Type:                &artifactType,
		})
	}

	return finalArtifacts
}

func readArtifact(artifacts []release.Artifact) *schema.Set {
	var finalArtifacts []interface{}

	testResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"definition_reference": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}

	for _, value := range artifacts {
		res, _ := json.Marshal(value.DefinitionReference)
		finalArtifacts = append(finalArtifacts, map[string]interface{}{
			"definition_reference": string(res),
			"alias":                *value.Alias,
			"source_id":            *value.SourceId,
			"type":                 *value.Type,
		})
	}

	return schema.NewSet(schema.HashResource(testResource), finalArtifacts)
}
