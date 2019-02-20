package azuredevops

import (
	"fmt"
	"github.com/ellisdon/azuredevops-go"
	//"github.com/ellisdon/terraform-provider-azuredevops/azuredevops/helper"
	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/hashicorp/terraform/helper/validation"
	"github.com/pkg/errors"
	"strconv"
)

func resourceReleaseDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceReleaseDefinitionCreate,
		Read:   resourceReleaseDefinitionRead,
		Update: resourceReleaseDefinitionUpdate,
		Delete: resourceReleaseDefinitionDelete,
		//	Importer: &schema.ResourceImporter{
		//		State: resourceReleaseDefinitionImport,
		//	},

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
			"environment": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
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

		result = append(result, azuredevops.ReleaseDefinitionEnvironment{
			Name: env["name"].(string),
			DeployPhases: []azuredevops.DeployPhase{
				azuredevops.DeployPhase{
					DeploymentInput: map[string]interface{}{
						"queueId": 168,
					},
					PhaseType: "agentBasedDeployment",
					Rank:      1,
					Name:      "Test",
				},
			},
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
			PreDeployApprovals: azuredevops.ReleaseDefinitionApprovals{
				Approvals: []azuredevops.ReleaseDefinitionApprovalStep{
					azuredevops.ReleaseDefinitionApprovalStep{
						Rank:             1,
						IsAutomated:      true,
						IsNotificationOn: false,
					},
				},
			},
		})
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
