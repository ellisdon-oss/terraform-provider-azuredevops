package azuredevops

import (
	"fmt"
	"github.com/ellisdon/azuredevops-go"
	"github.com/ellisdon/terraform-provider-azuredevops/azuredevops/helper"
	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/hashicorp/terraform/helper/validation"
	"encoding/json"
	"github.com/pkg/errors"
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
			"environment":      helper.EnvironmentResourceSchema(),
			"release_variable": helper.ReleaseVariableSchema(),
			"artifact":         helper.ArtifactSchema(),
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
	newReleaseDefinition := azuredevops.ReleaseDefinition{
		Name:         d.Get("name").(string),
		Environments: extractEnvironments(envs),
	}

	if v, ok := d.GetOk("release_variable"); ok {
		newReleaseDefinition.Variables = extractReleaseVariables(v.(*schema.Set))
	}

	if v, ok := d.GetOk("artifact"); ok {
		newReleaseDefinition.Artifacts = extractArtifact(v.(*schema.Set))
	}

	releaseDef, _, err := config.Client.ReleaseDefinitionsApi.CreateReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, newReleaseDefinition)
	//
	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}
	//
	d.SetId(fmt.Sprint(releaseDef.Id))

	return resourceReleaseDefinitionRead(d, meta)
}

func resourceReleaseDefinitionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	envs := d.Get("environment").([]interface{})

	id, _ := strconv.ParseInt(d.Id(), 10, 32)

	newReleaseDefinition := azuredevops.ReleaseDefinition{
		Id:           int32(id),
		Name:         d.Get("name").(string),
		Environments: extractEnvironments(envs),
		Revision:     int32(d.Get("revision").(int)),
	}

	if v, ok := d.GetOk("release_variable"); ok {
		newReleaseDefinition.Variables = extractReleaseVariables(v.(*schema.Set))
	}

	if v, ok := d.GetOk("artifact"); ok {
		newReleaseDefinition.Artifacts = extractArtifact(v.(*schema.Set))
	}
	_, _, err := config.Client.ReleaseDefinitionsApi.UpdateReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, newReleaseDefinition)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	return resourceReleaseDefinitionRead(d, meta)
}

func resourceReleaseDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, _ := strconv.ParseInt(d.Id(), 10, 32)

	res, _, err := config.Client.ReleaseDefinitionsApi.GetReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), int32(id), config.ApiVersion, nil)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.Set("revision", res.Revision)
	d.Set("name", res.Name)

	return nil
}

func resourceReleaseDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, _ := strconv.ParseInt(d.Id(), 10, 32)
	_, err := config.Client.ReleaseDefinitionsApi.DeleteReleaseDefinition(config.Context, config.Organization, d.Get("project_id").(string), int32(id), config.ApiVersion, nil)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	return nil
}

func extractEnvironments(environments []interface{}) []azuredevops.ReleaseDefinitionEnvironment {
	var result []azuredevops.ReleaseDefinitionEnvironment
	rank := 1
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
		var finalConditions []azuredevops.Condition
		for _, v := range conditions {

			v := v.(map[string]interface{})

			finalConditions = append(finalConditions, azuredevops.Condition{
				ConditionType: v["condition_type"].(string),
				Name:          v["name"].(string),
				Value:         v["value"].(string),
			})
		}

		var predeployapprovals []interface{}
		var finalPreDeployApprovals []azuredevops.ReleaseDefinitionApprovalStep

		finalApprovalOptions := azuredevops.ApprovalOptions{}

		if env["pre_deploy_approval"] != nil && len(env["pre_deploy_approval"].([]interface{})) != 0 {

			predeployapprovals = env["pre_deploy_approval"].([]interface{})[0].(map[string]interface{})["approvals"].([]interface{})

			if env["pre_deploy_approval"].([]interface{})[0].(map[string]interface{})["options"] != nil {
				options := env["pre_deploy_approval"].([]interface{})[0].(map[string]interface{})["options"]
				if options.([]interface{})[0].(map[string]interface{})["execution_order"] != nil {
					finalApprovalOptions.ExecutionOrder = options.([]interface{})[0].(map[string]interface{})["execution_order"].(string)
				} else {
					finalApprovalOptions.ExecutionOrder = "beforeGates"
				}

				finalApprovalOptions.TimeoutInMinutes = int32(options.([]interface{})[0].(map[string]interface{})["timeout_in_minutes"].(int))

				if options.([]interface{})[0].(map[string]interface{})["release_creator_can_be_approver"] != nil {
					finalApprovalOptions.ReleaseCreatorCanBeApprover = options.([]interface{})[0].(map[string]interface{})["release_creator_can_be_approver"].(bool)
				} else {
					finalApprovalOptions.ReleaseCreatorCanBeApprover = false
				}
			}
		} else {
			finalPreDeployApprovals = []azuredevops.ReleaseDefinitionApprovalStep{
				azuredevops.ReleaseDefinitionApprovalStep{
					Rank:             1,
					IsAutomated:      true,
					IsNotificationOn: false,
				},
			}
		}
		approvalRank := 1
		for _, v := range predeployapprovals {

			v := v.(map[string]interface{})

			finalPreDeployApprovals = append(finalPreDeployApprovals, azuredevops.ReleaseDefinitionApprovalStep{
				Rank: int32(approvalRank),
				Approver: azuredevops.IdentityRef{
					Id: v["approver_id"].(string),
				},
				IsAutomated:      v["is_automated"].(bool),
				IsNotificationOn: v["is_notification_on"].(bool),
			})

			approvalRank++
		}

		deployPhases := env["deploy_phase"].([]interface{})
		var finalDeployPhases []azuredevops.DeployPhase
		for _, v := range deployPhases {

			v := v.(map[string]interface{})

			var finalTasks []azuredevops.WorkflowTask

			tasks := v["workflow_task"].([]interface{})
			for _, task := range tasks {
				task := task.(map[string]interface{})
				finalTasks = append(finalTasks, azuredevops.WorkflowTask{
					AlwaysRun:       task["always_run"].(bool),
					Condition:       task["condition"].(string),
					Enabled:         task["enabled"].(bool),
					Inputs:          task["inputs"].(map[string]interface{}),
					Name:            task["name"].(string),
					Version:         task["version"].(string),
					TaskId:          task["task_id"].(string),
					ContinueOnError: task["continue_on_error"].(bool),
					DefinitionType:  task["definition_type"].(string),
				})
			}

			finalDeployPhases = append(finalDeployPhases, azuredevops.DeployPhase{
				DeploymentInput: v["deployment_input"].(map[string]interface{}),
				PhaseType:       v["phase_type"].(string),
				Rank:            int32(v["rank"].(int)),
				Name:            v["name"].(string),
				WorkflowTasks:   finalTasks,
			})
		}

		result = append(result, azuredevops.ReleaseDefinitionEnvironment{
			Name:         env["name"].(string),
			Conditions:   finalConditions,
			DeployPhases: finalDeployPhases,
			Rank:         int32(rank),
			RetentionPolicy: azuredevops.EnvironmentRetentionPolicy{
				DaysToKeep:     30,
				ReleasesToKeep: 3,
				RetainBuild:    true,
			},
			PostDeployApprovals: azuredevops.ReleaseDefinitionApprovals{
				Approvals: []azuredevops.ReleaseDefinitionApprovalStep{
					azuredevops.ReleaseDefinitionApprovalStep{
						Rank:             1,
						IsAutomated:      true,
						IsNotificationOn: false,
					},
				},
			},
			Variables: extractReleaseVariables(env["variable"].(*schema.Set)),
			PreDeployApprovals: azuredevops.ReleaseDefinitionApprovals{
				ApprovalOptions: finalApprovalOptions,
				Approvals:       finalPreDeployApprovals,
			},
		})

		rank++
	}

	return result
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

func extractReleaseVariables(variables *schema.Set) map[string]azuredevops.ConfigurationVariableValue {

	finalVariables := make(map[string]azuredevops.ConfigurationVariableValue)

	for _, value := range variables.List() {
		variableMap := value.(map[string]interface{})

		finalVariables[variableMap["name"].(string)] = azuredevops.ConfigurationVariableValue{
			IsSecret: variableMap["is_secret"].(bool),
			Value:    variableMap["value"].(string),
		}
	}

	return finalVariables
}

func extractArtifact(variables *schema.Set) []azuredevops.Artifact {

	var finalArtifacts []azuredevops.Artifact
	for _, value := range variables.List() {
		artifact := value.(map[string]interface{})
		refs := make(map[string]azuredevops.ArtifactSourceReference)

		ref := artifact["definition_reference"].(string)

		json.Unmarshal([]byte(ref), &refs)

		finalArtifacts = append(finalArtifacts, azuredevops.Artifact{
			DefinitionReference: refs,
			Alias:               artifact["alias"].(string),
			SourceId:            artifact["source_id"].(string),
			Type:                artifact["type"].(string),
		})
	}

	return finalArtifacts
}
