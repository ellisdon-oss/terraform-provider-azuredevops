package azuredevops

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
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
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
							Optional:  true,
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
				Type:      schema.TypeString,
				Computed:  true,
				Optional:  true,
				Sensitive: true,
			},
		},
	}

	config := meta.(*Config)
	group_id := d.Get("group_id").(int)
	name := d.Get("name").(string)

	var group azuredevops.VariableGroup

	if name == "" {
		res, _, err := config.Client.VariablegroupsApi.Get(config.Context, config.Organization, d.Get("project_id").(string), int32(group_id), config.ApiVersion)

		if err != nil {
			return err
		}

		if res.Name == "" {
			return errors.New("Variable Group Not Found")
		}

		group = res
		d.Set("name", group.Name)
	} else {
		groups, _, err := config.Client.VariablegroupsApi.GetVariableGroups(config.Context, config.Organization, d.Get("project_id").(string), config.ApiVersion, &azuredevops.GetVariableGroupsOpts{
			GroupName: optional.NewString(name),
		})

		if err != nil {
			return err
		}

		if groups.Count == 0 {
			return errors.New("Variable Group Not Found")
		}

		group = groups.VariableGroups[0]
		d.Set("group_id", group.Id)
	}

	var variables []interface{}

	for k, v := range group.Variables {
		variables = append(variables, map[string]interface{}{
			"is_secret": v.IsSecret,
			"value":     v.Value,
			"name":      k,
		})
	}

	res := schema.NewSet(schema.HashResource(testResource), variables)
	d.Set("variables", res)
	d.SetId(fmt.Sprintf("%s-%d", d.Get("project_id").(string), group.Id))

	return nil
}
