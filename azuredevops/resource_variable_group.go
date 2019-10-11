package azuredevops

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/taskagent"
	"github.com/pkg/errors"
	"strings"
)

func resourceVariableGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVariableGroupCreate,
		Update: resourceVariableGroupUpdate,
		Delete: resourceVariableGroupDelete,
		Read:   resourceVariableGroupRead,
		Importer: &schema.ResourceImporter{
			State: resourceVariableGroupImport,
		},

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"variable": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
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

func resourceVariableGroupRead(d *schema.ResourceData, meta interface{}) error {
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

	config := meta.(*Config)
	groupID := d.Get("group_id").(int)
	projectID := d.Get("project_id").(string)

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	group, err := agentClient.GetVariableGroup(config.Context, taskagent.GetVariableGroupArgs{
		Project: &projectID,
		GroupId: &groupID,
	})

	if err != nil {
		return err
	}

	if (*(*group).Name) == "" {
		return errors.New("Variable Group Not Found")
	}

	d.Set("name", *group.Name)
	var variables []interface{}

	oldVariables := make(map[string]string)
	for _, v := range d.Get("variable").(*schema.Set).List() {
		v := v.(map[string]interface{})
		oldVariables[v["name"].(string)] = v["value"].(string)

	}

	for k, v := range *group.Variables {
		if k != "" {
			value := ""
			var isSecret bool
			if v.IsSecret != nil && *v.IsSecret {
				isSecret = true
			} else {
				isSecret = false
			}

			if isSecret {
				value = oldVariables[k]
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
	d.Set("variable", res)
	d.SetId(fmt.Sprintf("%s-%d", d.Get("project_id").(string), *group.Id))

	return nil
}

func resourceVariableGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	variables := make(map[string]taskagent.VariableValue)

	for _, v := range d.Get("variable").(*schema.Set).List() {
		v := v.(map[string]interface{})

		value := v["value"].(string)
		isSecret := v["is_secret"].(bool)
		variables[v["name"].(string)] = taskagent.VariableValue{
			Value:    &value,
			IsSecret: &isSecret,
		}
	}

	name := d.Get("name").(string)
	newVariableGroup := taskagent.VariableGroupParameters{
		Name:      &name,
		Variables: &variables,
	}

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)

	group, err := agentClient.AddVariableGroup(config.Context, taskagent.AddVariableGroupArgs{
		Project: &projectID,
		Group:   &newVariableGroup,
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s-%d", d.Get("project_id").(string), &group.Id))
	d.Set("group_id", &group.Id)

	return resourceVariableGroupRead(d, meta)
}

func resourceVariableGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	variables := make(map[string]taskagent.VariableValue)

	for _, v := range d.Get("variable").(*schema.Set).List() {
		v := v.(map[string]interface{})

		value := v["value"].(string)
		isSecret := v["is_secret"].(bool)
		variables[v["name"].(string)] = taskagent.VariableValue{
			Value:    &value,
			IsSecret: &isSecret,
		}
	}

	name := d.Get("name").(string)
	newVariableGroup := taskagent.VariableGroupParameters{
		Name:      &name,
		Variables: &variables,
	}

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)
	groupID := d.Get("group_id").(int)

	group, err := agentClient.UpdateVariableGroup(config.Context, taskagent.UpdateVariableGroupArgs{
		Group:   &newVariableGroup,
		Project: &projectID,
		GroupId: &groupID,
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s-%d", d.Get("project_id").(string), *group.Id))

	return resourceVariableGroupRead(d, meta)
}

func resourceVariableGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)
	groupID := d.Get("group_id").(int)

	err = agentClient.DeleteVariableGroup(config.Context, taskagent.DeleteVariableGroupArgs{
		Project: &projectID,
		GroupId: &groupID,
	})

	if err != nil {
		return err
	}

	return nil
}

func resourceVariableGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	name := d.Id()
	if name == "" {
		return nil, fmt.Errorf("variable group id cannot be empty")
	}

	res := strings.Split(name, "/")
	if len(res) != 2 {
		return nil, fmt.Errorf("the format has to be in <project-id>/<variable-group-id>")
	}

	d.Set("project_id", res[0])
	d.Set("group_id", res[1])

	return []*schema.ResourceData{d}, nil
}
