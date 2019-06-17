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
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
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

	project, _, err := config.Client.ProjectsApi.GetProject(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, nil)

	newBuildDefinition := azuredevops.BuildDefinition{
		Name: d.Get("name").(string),
		Project: &azuredevops.TeamProjectReference{
			Id: project.Id,
		},
		Queue: &azuredevops.AgentPoolQueue{
			Name: agentName,
		},
		Repository: &azuredevops.BuildRepository{
			DefaultBranch:      d.Get("repository.0.default_branch").(string),
			Clean:              d.Get("repository.0.clean").(string),
			Name:               d.Get("repository.0.name").(string),
			Url:                d.Get("repository.0.url").(string),
			Type:               d.Get("repository.0.type").(string),
			CheckoutSubmodules: d.Get("repository.0.checkout_submodules").(bool),
			Properties:         properties,
		},
		Process: &azuredevops.BuildProcess{
			Type:         processType,
			YamlFileName: yamlFileName,
		},
	}

	if v, ok := d.GetOk("build_variable"); ok {
		newBuildDefinition.Variables = extractBuildVariables(v.(*schema.Set))
	}

	//	if v, ok := d.GetOk("demand"); ok {
	//		newBuildDefinition.Demands = extractDemands(v.(*schema.Set))
	//	}

	if v, ok := d.GetOk("triggers"); ok {
		newBuildDefinition.Triggers = extractBuildTriggers(v.([]interface{}))
	}

	res, _, err := config.Client.DefinitionsApi.CreateDefinition(config.Context, config.Organization, project.Id, config.ApiVersion, newBuildDefinition, nil)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.SetId(fmt.Sprint(res.Id))

	return resourceBuildDefinitionRead(d, meta)
}

func resourceBuildDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, _ := strconv.ParseInt(d.Id(), 10, 32)

	project, _, err := config.Client.ProjectsApi.GetProject(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, nil)

	definition, _, err := config.Client.DefinitionsApi.GetDefinition(config.Context, config.Organization, project.Id, int32(id), config.ApiVersion, nil)

	if err != nil {
		return err
	}
	//
	d.Set("revision", int(definition.Revision))
	d.Set("name", definition.Name)

	var finalList []interface{}
	buildVariables := definition.Variables

	oldBuildVariables := d.Get("build_variable").(*schema.Set).List()

	for key, value := range buildVariables {
		var mapValue string
		if value.IsSecret {
			for _, v := range oldBuildVariables {
				if v.(map[string]interface{})["name"].(string) == key {
					mapValue = v.(map[string]interface{})["value"].(string)
				}
			}
		} else {
			mapValue = value.Value
		}

		temp := map[string]interface{}{
			"is_secret": value.IsSecret,
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

	project, _, err := config.Client.ProjectsApi.GetProject(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, nil)

	id, _ := strconv.ParseInt(d.Id(), 10, 32)
	revision := int32(d.Get("revision").(int))
	newBuildDefinition := azuredevops.BuildDefinition{
		Id:       int32(id),
		Revision: revision,
		Name:     d.Get("name").(string),
		Project: &azuredevops.TeamProjectReference{
			Id: project.Id,
		},
		Queue: &azuredevops.AgentPoolQueue{
			Name: agentName,
		},
		Repository: &azuredevops.BuildRepository{
			DefaultBranch:      d.Get("repository.0.default_branch").(string),
			Clean:              d.Get("repository.0.clean").(string),
			Name:               d.Get("repository.0.name").(string),
			Url:                d.Get("repository.0.url").(string),
			Type:               d.Get("repository.0.type").(string),
			CheckoutSubmodules: d.Get("repository.0.checkout_submodules").(bool),
			Properties:         properties,
		},
		Process: &azuredevops.BuildProcess{
			Type:         processType,
			YamlFileName: yamlFileName,
		},
	}

	if v, ok := d.GetOk("build_variable"); ok {
		newBuildDefinition.Variables = extractBuildVariables(v.(*schema.Set))
	}

	//	if v, ok := d.GetOk("demand"); ok {
	//		newBuildDefinition.Demands = extractDemands(v.(*schema.Set))
	//	}

	if v, ok := d.GetOk("triggers"); ok {
		newBuildDefinition.Triggers = extractBuildTriggers(v.([]interface{}))
	}

	res, _, err := config.Client.DefinitionsApi.UpdateDefinition(config.Context, config.Organization, project.Id, int32(id), config.ApiVersion, newBuildDefinition, nil)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.SetId(fmt.Sprint(res.Id))

	return resourceBuildDefinitionRead(d, meta)
}

func resourceBuildDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, _ := strconv.ParseInt(d.Id(), 10, 32)
	_, err := config.Client.DefinitionsApi.DeleteDefinition(config.Context, config.Organization, d.Get("project_id").(string), int32(id), config.ApiVersion)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
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

func extractBuildTriggers(triggers []interface{}) []azuredevops.BuildTrigger {

	var finalTriggers []azuredevops.BuildTrigger
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

			trigger := azuredevops.BuildTrigger{
				TriggerType:   "pullRequest",
				BranchFilters: convertInterfaceSliceToStringSlice(v.(map[string]interface{})["branch_filters"].([]interface{})),
			}

			v = v.(map[string]interface{})

			if val := v.(map[string]interface{})["auto_cancel"]; val != nil {
				trigger.BatchChanges = val.(bool)
			}

			if val := v.(map[string]interface{})["is_comment_required_for_pull_request"]; val != nil {
				trigger.IsCommentRequiredForPullRequest = val.(bool)
			}

			if val := v.(map[string]interface{})["forks"]; len(val.([]interface{})) != 0 {
				val = val.([]interface{})[0]
				trigger.Forks = &azuredevops.Forks{
					AllowSecrets: val.(map[string]interface{})["allow_secrets"].(bool),
					Enabled:      val.(map[string]interface{})["enabled"].(bool),
				}
			}

			if val := v.(map[string]interface{})["path_filters"]; val != nil {
				trigger.PathFilters = convertInterfaceSliceToStringSlice(val.([]interface{}))
			}

			if val := v.(map[string]interface{})["settings_source_type"]; val != nil {
				trigger.SettingsSourceType = int32(val.(int))
			}

			finalTriggers = append(finalTriggers, trigger)

		case "build_completion":
		case "schedule":
			if len(v.([]interface{})) == 0 {
				continue
			}

			var schedules []azuredevops.Schedule

			for _, val := range v.([]interface{}) {

				schedule := azuredevops.Schedule{
					BranchFilters: convertInterfaceSliceToStringSlice(val.(map[string]interface{})["branch_filters"].([]interface{})),
				}

				if temp := val.(map[string]interface{})["days_to_build"]; len(temp.([]interface{})) != 0 {
					days := convertInterfaceSliceToStringSlice(temp.([]interface{}))
					schedule.DaysToBuild = helper.CalcScheduleDays(days)
				}

				if temp := val.(map[string]interface{})["schedule_job_id"]; temp != nil {
					schedule.ScheduleJobId = temp.(string)
				}

				schedules = append(schedules, schedule)
			}

			if len(schedules) == 0 {
				continue
			}

			trigger := azuredevops.BuildTrigger{
				TriggerType: "schedule",
				Schedules:   schedules,
			}

			finalTriggers = append(finalTriggers, trigger)

		case "continuous_integration":
			if len(v.([]interface{})) == 0 {
				continue
			}
			v = v.([]interface{})[0]

			trigger := azuredevops.BuildTrigger{
				TriggerType:   "continuousIntegration",
				BranchFilters: convertInterfaceSliceToStringSlice(v.(map[string]interface{})["branch_filters"].([]interface{})),
			}

			v = v.(map[string]interface{})

			if val := v.(map[string]interface{})["batch_changes"]; val != nil {
				trigger.BatchChanges = val.(bool)
			}

			if val := v.(map[string]interface{})["max_concurrent_builds_per_branch"]; val != nil {
				trigger.MaxConcurrentBuildsPerBranch = int32(val.(int))
			}

			if val := v.(map[string]interface{})["path_filters"]; val != nil {
				trigger.PathFilters = convertInterfaceSliceToStringSlice(val.([]interface{}))
			}

			if val := v.(map[string]interface{})["polling_interval"]; val != nil {
				trigger.PollingInterval = int32(val.(int))
			}

			if val := v.(map[string]interface{})["polling_job_id"]; val != nil {
				trigger.PollingJobId = val.(string)
			}

			if val := v.(map[string]interface{})["settings_source_type"]; val != nil {
				trigger.SettingsSourceType = int32(val.(int))
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

func extractBuildVariables(variables *schema.Set) map[string]azuredevops.BuildDefinitionVariable {

	finalVariables := make(map[string]azuredevops.BuildDefinitionVariable)

	for _, value := range variables.List() {
		variableMap := value.(map[string]interface{})

		finalVariables[variableMap["name"].(string)] = azuredevops.BuildDefinitionVariable{
			IsSecret: variableMap["is_secret"].(bool),
			Value:    variableMap["value"].(string),
		}
	}

	return finalVariables
}
