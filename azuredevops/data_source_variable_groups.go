package azuredevops

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/taskagent"
)

func dataSourceVariableGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVariableGroupsRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"variables": &schema.Schema{
							Type:     schema.TypeList,
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
				},
			},
		},
	}
}

func dataSourceVariableGroupsRead(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	agentClient, err := taskagent.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)

	groups, err := agentClient.GetVariableGroups(config.Context, taskagent.GetVariableGroupsArgs{
		Project: &projectID,
	})

	if err != nil {
		return err
	}

	var resGroups []interface{}
	for _, v := range *groups {
		group := map[string]interface{}{
			"name":     v.Name,
			"group_id": int(*v.Id),
		}

		var variables []interface{}

		for l, p := range *v.Variables {
			variables = append(variables, map[string]interface{}{
				"is_secret": p.IsSecret,
				"value":     p.Value,
				"name":      l,
			})
		}

		group["variables"] = variables

		resGroups = append(resGroups, group)
	}

	d.Set("groups", resGroups)
	d.SetId(fmt.Sprintf("%s-variable_groups", projectID))

	return nil
}
