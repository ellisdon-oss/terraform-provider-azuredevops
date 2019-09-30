package azuredevops

import (
	"fmt"
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"log"
)

func resourceVariableGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVariableGroupCreate,
		Update: resourceVariableGroupUpdate,
		Delete: resourceVariableGroupDelete,
		Read:   resourceVariableGroupRead,

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
	group_id := d.Get("group_id").(int)

	group, _, err := config.Client.VariablegroupsApi.Get(config.Context, config.Organization, d.Get("project_id").(string), int32(group_id), "5.1-preview.1")

	if err != nil {
		return err
	}

	if group.Name == "" {
		return errors.New("Variable Group Not Found")
	}

	d.Set("name", group.Name)
	var variables []interface{}

	old_variables := make(map[string]string)
	for _, v := range d.Get("variable").(*schema.Set).List() {
		v := v.(map[string]interface{})
		old_variables[v["name"].(string)] = v["value"].(string)

	}

	for k, v := range group.Variables {
		if k != "" {
			value := ""
			if v.IsSecret {
				value = old_variables[k]
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

	res := schema.NewSet(schema.HashResource(testResource), variables)
	d.Set("variable", res)
	d.SetId(fmt.Sprintf("%s-%d", d.Get("project_id").(string), group.Id))

	return nil
}

func resourceVariableGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	variables := make(map[string]azuredevops.VariableValue)

	for _, v := range d.Get("variable").(*schema.Set).List() {
		v := v.(map[string]interface{})

		variables[v["name"].(string)] = azuredevops.VariableValue{
			Value:    v["value"].(string),
			IsSecret: v["is_secret"].(bool),
		}

	}

	newVariableGroup := azuredevops.VariableGroupParameters{
		Name:      d.Get("name").(string),
		Variables: variables,
	}

	log.Printf(config.Organization)
	log.Printf(d.Get("project_id").(string))
	group, _, err := config.Client.VariablegroupsApi.Add(config.Context, config.Organization, d.Get("project_id").(string), "5.1-preview.1", newVariableGroup)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.SetId(fmt.Sprintf("%s-%d", d.Get("project_id").(string), group.Id))
	d.Set("group_id", group.Id)
	return resourceVariableGroupRead(d, meta)
}

func resourceVariableGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	variables := make(map[string]azuredevops.VariableValue)

	for _, v := range d.Get("variable").(*schema.Set).List() {
		v := v.(map[string]interface{})

		variables[v["name"].(string)] = azuredevops.VariableValue{
			Value:    v["value"].(string),
			IsSecret: v["is_secret"].(bool),
		}

	}

	newVariableGroup := azuredevops.VariableGroupParameters{
		Name:      d.Get("name").(string),
		Variables: variables,
	}

	group, _, err := config.Client.VariablegroupsApi.Update(config.Context, config.Organization, d.Get("project_id").(string), int32(d.Get("group_id").(int)), "5.1-preview.1", newVariableGroup)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.SetId(fmt.Sprintf("%s-%d", d.Get("project_id").(string), group.Id))

	return resourceVariableGroupRead(d, meta)
}

func resourceVariableGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	_, err := config.Client.VariablegroupsApi.Delete(config.Context, config.Organization, d.Get("project_id").(string), int32(d.Get("group_id").(int)), "5.1-preview.1")

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	return nil
}
