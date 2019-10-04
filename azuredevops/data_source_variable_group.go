package azuredevops

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/taskagent"
	"github.com/pkg/errors"
)

func dataSourceVariableGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVariableGroupRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": &schema.Schema{
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"name"},
			},
			"name": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"group_id"},
			},
			"variables": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_secret": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVariableGroupRead(d *schema.ResourceData, meta interface{}) error {
	testResource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_secret": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}

	config := meta.(*Config)

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)
	groupID := d.Get("group_id").(int)
	name := d.Get("name").(string)

	var group taskagent.VariableGroup

	if name == "" {
		res, err := agentClient.GetVariableGroup(config.Context, taskagent.GetVariableGroupArgs{
			Project: &projectID,
			GroupId: &groupID,
		})

		if err != nil {
			return err
		}

		if (*(*res).Name) == "" {
			return errors.New("Variable Group Not Found")
		}

		group = *res
		d.Set("name", group.Name)
	} else {
		groups, err := agentClient.GetVariableGroups(config.Context, taskagent.GetVariableGroupsArgs{
			Project:   &projectID,
			GroupName: &name,
		})

		if err != nil {
			return err
		}

		if len(*groups) == 0 {
			return errors.New("Variable Group Not Found")
		}

		group = (*groups)[0]
		d.Set("group_id", group.Id)
	}

	var variables []interface{}

	for k, v := range *group.Variables {
		variables = append(variables, map[string]interface{}{
			"is_secret": v.IsSecret,
			"value":     *v.Value,
			"name":      k,
		})
	}

	res := schema.NewSet(schema.HashResource(testResource), variables)
	d.Set("variables", res)
	d.SetId(fmt.Sprintf("%s-%d", projectID, *group.Id))

	return nil
}
