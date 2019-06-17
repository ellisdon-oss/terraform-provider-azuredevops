package azuredevops

import (
	"fmt"
	//"github.com/antihax/optional"
	//"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/pkg/errors"
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
										Type:      schema.TypeString,
										Computed:  true,
										Sensitive: true,
										Optional:  true,
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

	groups, _, err := config.Client.VariablegroupsApi.GetVariableGroups(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, nil)

	if err != nil {
		return err
	}

	var resGroups []interface{}
	for _, v := range groups.VariableGroups {
		group := map[string]interface{}{
			"name":     v.Name,
			"group_id": int(v.Id),
		}

		var variables []interface{}

		for l, p := range v.Variables {
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
	d.SetId(fmt.Sprintf("%s-variable_groups", d.Get("project_id").(string)))

	return nil
}
