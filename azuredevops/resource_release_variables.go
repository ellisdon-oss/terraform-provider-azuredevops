package azuredevops

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
)

func resourceReleaseVariables() *schema.Resource {
	return &schema.Resource{
		Create: resourceReleaseVariablesCreate,
		Read:   resourceReleaseVariablesRead,
		Update: resourceReleaseVariablesUpdate,
		Delete: resourceReleaseVariablesDelete,
		Schema: map[string]*schema.Schema{
			"variable": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stage_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
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

func resourceReleaseVariablesCreate(d *schema.ResourceData, meta interface{}) error {
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

	userVariables := d.Get("variable").(*schema.Set)

	for _, userVariable := range userVariables.List() {
		userVariable := userVariable.(map[string]interface{})

		stageName := userVariable["stage_name"].(string)
		value := userVariable["value"].(string)
		name := userVariable["name"].(string)
		isSecret := userVariable["is_secret"].(bool)

		if stageName == "" {
			vars := *res.Variables
			found := false
			for k, _ := range vars {
				if k == name {
					vars[k] = release.ConfigurationVariableValue{
						IsSecret: &isSecret,
						Value:    &value,
					}

					found = true
				}
			}

			if !found {
				vars[name] = release.ConfigurationVariableValue{
					IsSecret: &isSecret,
					Value:    &value,
				}
			}
		} else {
			var resultEnv release.ReleaseDefinitionEnvironment
			var stageIndex int

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

			vars := *resultEnv.Variables
			found := false
			for k, _ := range vars {
				if k == name {
					vars[k] = release.ConfigurationVariableValue{
						IsSecret: &isSecret,
						Value:    &value,
					}

					found = true
				}
			}

			if !found {
				vars[name] = release.ConfigurationVariableValue{
					IsSecret: &isSecret,
					Value:    &value,
				}
			}

			resultEnv.Variables = &vars

			(*res.Environments)[stageIndex] = resultEnv
		}
	}

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: res,
	})

	if err != nil {
		return err
	}

	d.SetId(uuid.New().String())

	return resourceReleaseVariablesRead(d, meta)
}

func resourceReleaseVariablesUpdate(d *schema.ResourceData, meta interface{}) error {
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

	userVariables := d.Get("variable").(*schema.Set)

	for _, userVariable := range userVariables.List() {
		userVariable := userVariable.(map[string]interface{})

		stageName := userVariable["stage_name"].(string)
		value := userVariable["value"].(string)
		name := userVariable["name"].(string)
		isSecret := userVariable["is_secret"].(bool)

		if stageName == "" {
			vars := *res.Variables
			found := false
			for k, _ := range vars {
				if k == name {
					vars[k] = release.ConfigurationVariableValue{
						IsSecret: &isSecret,
						Value:    &value,
					}

					found = true
				}
			}

			if !found {
				vars[name] = release.ConfigurationVariableValue{
					IsSecret: &isSecret,
					Value:    &value,
				}
			}
		} else {
			var resultEnv release.ReleaseDefinitionEnvironment
			var stageIndex int

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

			vars := *resultEnv.Variables
			found := false
			for k, _ := range vars {
				if k == name {
					vars[k] = release.ConfigurationVariableValue{
						IsSecret: &isSecret,
						Value:    &value,
					}

					found = true
				}
			}

			if !found {
				vars[name] = release.ConfigurationVariableValue{
					IsSecret: &isSecret,
					Value:    &value,
				}
			}

			resultEnv.Variables = &vars

			(*res.Environments)[stageIndex] = resultEnv
		}
	}

	_, err = releaseClient.UpdateReleaseDefinition(config.Context, release.UpdateReleaseDefinitionArgs{
		Project:           &projectID,
		ReleaseDefinition: res,
	})

	if err != nil {
		return err
	}

	return resourceReleaseVariablesRead(d, meta)
}

func resourceReleaseVariablesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

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
			"stage_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}

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

	userVariables := d.Get("variable").(*schema.Set)

	finalUserVariables := []interface{}{}
	for _, userVariable := range userVariables.List() {
		userVariable := userVariable.(map[string]interface{})

		stageName := userVariable["stage_name"].(string)
		name := userVariable["name"].(string)
		value := userVariable["value"].(string)

		if stageName == "" {
			v := (*res.Variables)[name]
			innerValue := ""
			if &v != nil {
				var innerIsSecret bool
				if v.IsSecret != nil && *v.IsSecret {
					innerIsSecret = true
				} else {
					innerIsSecret = false
				}

				if innerIsSecret {
					innerValue = value
				} else {
					innerValue = *v.Value
				}

				finalUserVariables = append(finalUserVariables, map[string]interface{}{
					"stage_name": stageName,
					"is_secret":  innerIsSecret,
					"value":      innerValue,
					"name":       name,
				})
			}
		} else {
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

			v := (*resultEnv.Variables)[name]
			innerValue := ""
			if &v != nil {
				var innerIsSecret bool
				if v.IsSecret != nil && *v.IsSecret {
					innerIsSecret = true
				} else {
					innerIsSecret = false
				}

				if innerIsSecret {
					innerValue = value
				} else {
					innerValue = *v.Value
				}

				finalUserVariables = append(finalUserVariables, map[string]interface{}{
					"stage_name": stageName,
					"is_secret":  innerIsSecret,
					"value":      innerValue,
					"name":       name,
				})
			}
		}
		//var resultEnv release.ReleaseDefinitionEnvironment

		//for _, env := range *res.Environments {
		//	if *env.Name == stageName {
		//		resultEnv = env
		//		break
		//	}
		//}

		//if resultEnv.Name == nil {
		//	return fmt.Errorf("No stage with the name %s", stageName)
		//}

		//finalUserVariables = append(finalUserVariables, userVariable)
	}

	resSet := schema.NewSet(schema.HashResource(testResource), finalUserVariables)

	d.Set("variable", resSet)

	return nil
}

func resourceReleaseVariablesDelete(d *schema.ResourceData, meta interface{}) error {
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
	userVariables := d.Get("variable").(*schema.Set)

	for _, userVariable := range reverseSlice(userVariables.List()) {
		userVariable := userVariable.(map[string]interface{})

		stageName := userVariable["stage_name"].(string)
		name := userVariable["name"].(string)

		if stageName == "" {
			vars := *res.Variables
			delete(vars, name)

			res.Variables = &vars
		} else {
			var resultEnv release.ReleaseDefinitionEnvironment
			var stageIndex int

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

			vars := *resultEnv.Variables
			delete(vars, name)

			resultEnv.Variables = &vars
			(*res.Environments)[stageIndex] = resultEnv
		}
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
