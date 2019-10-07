package azuredevops

import (
	"fmt"
	"github.com/ellisdon/terraform-provider-azuredevops/azuredevops/helper"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"

	"strconv"
	"strings"
)

func resourceBuildDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceBuildDefinitionCreate,
		Read:   resourceBuildDefinitionRead,
		Update: resourceBuildDefinitionUpdate,
		Delete: resourceBuildDefinitionDelete,
		Importer: &schema.ResourceImporter{
			State: resourceBuildDefinitionImport,
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
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"demand": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"process": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Default:  2,
						},
						"yaml_file_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "./azure-pipelines.yml",
						},
					},
				},
			},
			"triggers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//						"schedule": &schema.Schema{
						//							Type:     schema.TypeList,
						//							Optional: true,
						//							Elem: &schema.Resource{
						//								Schema: map[string]*schema.Schema{
						//									"branch_filters": &schema.Schema{
						//										Type:     schema.TypeList,
						//										Required: true,
						//										Elem: &schema.Schema{
						//											Type: schema.TypeString,
						//										},
						//									},
						//									"days_to_build": &schema.Schema{
						//										Type:     schema.TypeList,
						//										Required: true,
						//										Elem: &schema.Schema{
						//											Type: schema.TypeString,
						//											ValidateFunc: validation.StringInSlice([]string{
						//												"none",
						//												"monday",
						//												"tuesday",
						//												"wednesday",
						//												"thursday",
						//												"friday",
						//												"saturday",
						//												"sunday",
						//												"all",
						//											}, true),
						//										},
						//									},
						//									"schedule_job_id": &schema.Schema{
						//										Type:     schema.TypeString,
						//										Optional: true,
						//									},
						//									"schedule_only_with_changes": &schema.Schema{
						//										Type:     schema.TypeBool,
						//										Optional: true,
						//									},
						//									"start_hours": &schema.Schema{
						//										Type:     schema.TypeInt,
						//										Optional: true,
						//									},
						//									"start_minutes": &schema.Schema{
						//										Type:     schema.TypeInt,
						//										Optional: true,
						//									},
						//									"time_zone_id": &schema.Schema{
						//										Type:     schema.TypeString,
						//										Optional: true,
						//									},
						//								},
						//							},
						//						},
						"pull_request": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_cancel": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
									"is_comment_required_for_pull_request": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
									"settings_source_type": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"path_filters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"branch_filters": &schema.Schema{
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"forks": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allow_secrets": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"enabled": &schema.Schema{
													Type:     schema.TypeBool,
													Required: true,
												},
											},
										},
									},
								},
							},
						},
						"continuous_integration": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"branch_filters": &schema.Schema{
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"batch_changes": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"max_concurrent_builds_per_branch": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"path_filters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"polling_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"polling_job_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"settings_source_type": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"queue": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Hosted",
						},
					},
				},
			},
			"repository": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_branch": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "master",
						},
						"checkout_submodules": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"clean": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "true",
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Github",
						},
						"properties": &schema.Schema{
							Type:     schema.TypeMap,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"build_variable": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
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
				},
			},
		},
	}
}

func resourceBuildDefinitionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	orgProperties := d.Get("repository.0.properties").(map[string]interface{})

	properties := make(map[string]string)

	var agentName string

	if v, ok := d.GetOk("queue.0.agent_name"); ok {
		agentName = v.(string)
	} else {
		agentName = "Hosted"
	}

	var processType int32
	var yamlFileName string

	if v, ok := d.GetOk("process"); ok {
		process := v.([]interface{})[0].(map[string]interface{})
		processType = int32(process["type"].(int))
		yamlFileName = process["yaml_file_name"].(string)
	} else {
		processType = 2
		yamlFileName = "./azure-pipelines.yml"
	}

	for key, value := range orgProperties {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		properties[strKey] = strValue
	}

	name := d.Get("name").(string)
	projectID, _ := uuid.Parse(d.Get("project_id").(string))

	defaultBranch := d.Get("repository.0.default_branch").(string)
	clean := d.Get("repository.0.clean").(string)
	repoName := d.Get("repository.0.name").(string)
	url := d.Get("repository.0.url").(string)
	repoType := d.Get("repository.0.type").(string)
	checkoutSubmodules := d.Get("repository.0.checkout_submodules").(bool)

	newBuildDefinition := build.BuildDefinition{
		Name: &name,
		Project: &core.TeamProjectReference{
			Id: &projectID,
		},
		Queue: &build.AgentPoolQueue{
			Name: &agentName,
		},
		Repository: &build.BuildRepository{
			DefaultBranch:      &defaultBranch,
			Clean:              &clean,
			Name:               &repoName,
			Url:                &url,
			Type:               &repoType,
			CheckoutSubmodules: &checkoutSubmodules,
			Properties:         &properties,
		},
		Process: map[string]interface{}{
			"type":         &processType,
			"yamlFileName": &yamlFileName,
		},
	}

	if v, ok := d.GetOk("build_variable"); ok {
		buildVariables := extractBuildVariables(v.(*schema.Set))
		newBuildDefinition.Variables = &buildVariables
	}

	//	if v, ok := d.GetOk("demand"); ok {
	//		newBuildDefinition.Demands = extractDemands(v.(*schema.Set))
	//	}

	if v, ok := d.GetOk("triggers"); ok {
		triggers := extractBuildTriggers(v.([]interface{}))
		newBuildDefinition.Triggers = &triggers
	}

	buildClient, err := build.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectIDString := projectID.String()

	res, err := buildClient.CreateDefinition(config.Context, build.CreateDefinitionArgs{
		Project:    &projectIDString,
		Definition: &newBuildDefinition,
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprint(*res.Id))

	return resourceBuildDefinitionRead(d, meta)
}

func resourceBuildDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, _ := strconv.ParseInt(d.Id(), 10, 32)

	buildClient, err := build.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)
	defID := int(id)

	definition, err := buildClient.GetDefinition(config.Context, build.GetDefinitionArgs{
		Project:      &projectID,
		DefinitionId: &defID,
	})

	if err != nil {
		return err
	}
	//
	d.Set("revision", int(*definition.Revision))
	d.Set("name", *definition.Name)

	var finalList []interface{}
	buildVariables := *definition.Variables

	oldBuildVariables := d.Get("build_variable").(*schema.Set).List()

	for key, value := range buildVariables {
		var mapValue string
		var isSecret bool
		if value.IsSecret != nil && *value.IsSecret {
			isSecret = true
		} else {
			isSecret = false
		}

		if isSecret {
			for _, v := range oldBuildVariables {
				if v.(map[string]interface{})["name"].(string) == key {
					mapValue = v.(map[string]interface{})["value"].(string)
				}
			}
		} else {
			mapValue = *value.Value
		}

		temp := map[string]interface{}{
			"is_secret": isSecret,
			"value":     mapValue,
			"name":      key,
		}

		finalList = append(finalList, temp)
	}

	final := schema.NewSet(schema.HashResource(resourceBuildDefinition().Schema["build_variable"].Elem.(*schema.Resource)), finalList)

	d.Set("build_variable", final)

	//var finalScheduleList []interface{}
	//var scheduleTriggers []azuredevops.Schedule
	//triggers := definition.Triggers
	//for _, v := range triggers {
	//	if v.TriggerType == "schedule" {
	//		scheduleTriggers = v.Schedules
	//	}
	//}

	//for _, value := range scheduleTriggers {
	//	temp := map[string]interface{}{
	//		"schedule_job_id":            value.ScheduleJobId,
	//		"schedule_only_with_changes": value.ScheduleOnlyWithChanges,
	//		"branch_filters":             value.BranchFilters,
	//		//"days_to_build":              value.DaysToBuild,
	//	}

	//	finalScheduleList = append(finalScheduleList, temp)
	//}

	//d.Set("triggers.0.schedule", finalScheduleList)

	return nil
}

func resourceBuildDefinitionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	orgProperties := d.Get("repository.0.properties").(map[string]interface{})
	properties := make(map[string]string)

	var agentName string

	if v, ok := d.GetOk("queue.0.agent_name"); ok {
		agentName = v.(string)
	} else {
		agentName = "Hosted"
	}

	var processType int32
	var yamlFileName string

	if v, ok := d.GetOk("process"); ok {
		process := v.([]interface{})[0].(map[string]interface{})
		processType = int32(process["type"].(int))
		yamlFileName = process["yaml_file_name"].(string)
	} else {
		processType = 2
		yamlFileName = "./azure-pipelines.yml"
	}

	for key, value := range orgProperties {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		properties[strKey] = strValue
	}

	id, _ := strconv.ParseInt(d.Id(), 10, 32)
	revision := d.Get("revision").(int)
	name := d.Get("name").(string)
	projectID, _ := uuid.Parse(d.Get("project_id").(string))
	defID := int(id)

	defaultBranch := d.Get("repository.0.default_branch").(string)
	clean := d.Get("repository.0.clean").(string)
	repoName := d.Get("repository.0.name").(string)
	url := d.Get("repository.0.url").(string)
	repoType := d.Get("repository.0.type").(string)
	checkoutSubmodules := d.Get("repository.0.checkout_submodules").(bool)

	newBuildDefinition := build.BuildDefinition{
		Id:       &defID,
		Revision: &revision,
		Name:     &name,
		Project: &core.TeamProjectReference{
			Id: &projectID,
		},
		Queue: &build.AgentPoolQueue{
			Name: &agentName,
		},
		Repository: &build.BuildRepository{
			DefaultBranch:      &defaultBranch,
			Clean:              &clean,
			Name:               &repoName,
			Url:                &url,
			Type:               &repoType,
			CheckoutSubmodules: &checkoutSubmodules,
			Properties:         &properties,
		},
		Process: map[string]interface{}{
			"type":         &processType,
			"yamlFileName": &yamlFileName,
		},
	}

	if v, ok := d.GetOk("build_variable"); ok {
		buildVariables := extractBuildVariables(v.(*schema.Set))
		newBuildDefinition.Variables = &buildVariables
	}

	//	if v, ok := d.GetOk("demand"); ok {
	//		newBuildDefinition.Demands = extractDemands(v.(*schema.Set))
	//	}

	if v, ok := d.GetOk("triggers"); ok {
		triggers := extractBuildTriggers(v.([]interface{}))
		newBuildDefinition.Triggers = &triggers
	}

	buildClient, err := build.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectIDString := projectID.String()
	res, err := buildClient.UpdateDefinition(config.Context, build.UpdateDefinitionArgs{
		Project:      &projectIDString,
		Definition:   &newBuildDefinition,
		DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprint(*res.Id))

	return resourceBuildDefinitionRead(d, meta)
}

func resourceBuildDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, _ := strconv.ParseInt(d.Id(), 10, 32)

	buildClient, err := build.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	defID := int(id)
	projectID := d.Get("project_id").(string)

	err = buildClient.DeleteDefinition(config.Context, build.DeleteDefinitionArgs{
		Project:      &projectID,
		DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

	return nil
}

func resourceBuildDefinitionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	name := d.Id()
	if name == "" {
		return nil, fmt.Errorf("build definition id cannot be empty")
	}

	res := strings.Split(name, "/")
	if len(res) != 2 {
		return nil, fmt.Errorf("the format has to be in <project-id>/<build-definition-id>")
	}

	d.Set("project_id", res[0])
	d.SetId(res[1])

	return []*schema.ResourceData{d}, nil
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

func extractBuildTriggers(triggers []interface{}) []interface{} {

	var finalTriggers []interface{}
	if triggers[0] == nil {
		return finalTriggers
	}
	newTriggers := triggers[0].(map[string]interface{})

	for k, v := range newTriggers {
		switch k {
		case "pull_request":
			if len(v.([]interface{})) == 0 {
				continue
			}
			v = v.([]interface{})[0]

			trigger := map[string]interface{}{
				"triggerType":   "pullRequest",
				"branchFilters": convertInterfaceSliceToStringSlice(v.(map[string]interface{})["branch_filters"].([]interface{})),
			}

			v = v.(map[string]interface{})

			if val := v.(map[string]interface{})["auto_cancel"]; val != nil {
				trigger["batchChanges"] = val.(bool)
			}

			if val := v.(map[string]interface{})["is_comment_required_for_pull_request"]; val != nil {
				trigger["isCommentRequiredForPullRequest"] = val.(bool)
			}

			if val := v.(map[string]interface{})["forks"]; len(val.([]interface{})) != 0 {
				val = val.([]interface{})[0]
				trigger["forks"] = map[string]interface{}{
					"allowSecrets": val.(map[string]interface{})["allow_secrets"].(bool),
					"enabled":      val.(map[string]interface{})["enabled"].(bool),
				}
			}

			if val := v.(map[string]interface{})["path_filters"]; val != nil {
				trigger["pathFilters"] = convertInterfaceSliceToStringSlice(val.([]interface{}))
			}

			if val := v.(map[string]interface{})["settings_source_type"]; val != nil {
				trigger["settingsSourceType"] = int32(val.(int))
			}

			finalTriggers = append(finalTriggers, trigger)

		case "build_completion":
		case "schedule":
			if len(v.([]interface{})) == 0 {
				continue
			}

			var schedules []map[string]interface{}

			for _, val := range v.([]interface{}) {

				schedule := map[string]interface{}{
					"branchFilters": convertInterfaceSliceToStringSlice(val.(map[string]interface{})["branch_filters"].([]interface{})),
				}

				if temp := val.(map[string]interface{})["days_to_build"]; len(temp.([]interface{})) != 0 {
					days := convertInterfaceSliceToStringSlice(temp.([]interface{}))
					schedule["daysToBuild"] = helper.CalcScheduleDays(days)
				}

				if temp := val.(map[string]interface{})["schedule_job_id"]; temp != nil {
					schedule["scheduleJobId"] = temp.(string)
				}

				schedules = append(schedules, schedule)
			}

			if len(schedules) == 0 {
				continue
			}

			trigger := map[string]interface{}{
				"triggerType": "schedule",
				"schedules":   schedules,
			}

			finalTriggers = append(finalTriggers, trigger)

		case "continuous_integration":
			if len(v.([]interface{})) == 0 {
				continue
			}
			v = v.([]interface{})[0]

			trigger := map[string]interface{}{
				"triggerType":   "continuousIntegration",
				"branchFilters": convertInterfaceSliceToStringSlice(v.(map[string]interface{})["branch_filters"].([]interface{})),
			}

			v = v.(map[string]interface{})

			if val := v.(map[string]interface{})["batch_changes"]; val != nil {
				trigger["batchChanges"] = val.(bool)
			}

			if val := v.(map[string]interface{})["max_concurrent_builds_per_branch"]; val != nil {
				trigger["maxConcurrentBuildsPerBranch"] = int32(val.(int))
			}

			if val := v.(map[string]interface{})["path_filters"]; val != nil {
				trigger["pathFilters"] = convertInterfaceSliceToStringSlice(val.([]interface{}))
			}

			if val := v.(map[string]interface{})["polling_interval"]; val != nil {
				trigger["pollingInterval"] = int32(val.(int))
			}

			if val := v.(map[string]interface{})["polling_job_id"]; val != nil {
				trigger["pollingJobId"] = val.(string)
			}

			if val := v.(map[string]interface{})["settings_source_type"]; val != nil {
				trigger["settingsSourceType"] = int32(val.(int))
			}

			finalTriggers = append(finalTriggers, trigger)
		}

	}

	return finalTriggers
}

func convertInterfaceSliceToStringSlice(t []interface{}) []string {
	s := make([]string, len(t))
	for i, v := range t {
		s[i] = fmt.Sprint(v)
	}
	return s
}

func extractBuildVariables(variables *schema.Set) map[string]build.BuildDefinitionVariable {

	finalVariables := make(map[string]build.BuildDefinitionVariable)

	for _, value := range variables.List() {
		variableMap := value.(map[string]interface{})
		varValue := variableMap["value"].(string)
		isSecret := variableMap["is_secret"].(bool)
		finalVariables[variableMap["name"].(string)] = build.BuildDefinitionVariable{
			IsSecret: &isSecret,
			Value:    &varValue,
		}
	}

	return finalVariables
}
