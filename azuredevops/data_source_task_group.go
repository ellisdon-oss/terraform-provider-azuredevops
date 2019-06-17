package azuredevops

import (
	"fmt"
	//	"github.com/antihax/optional"
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func dataSourceTaskGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskGroupRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			//			"tasks": &schema.Schema{
			//				Type:     schema.TypeList,
			//				Computed: true,
			//				Elem:     &schema.Schema{Type: schema.TypeMap},
			//			},
		},
	}
}

func dataSourceTaskGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	res, _, err := config.Client.TaskgroupsApi.GetTaskGroups(config.Context, config.Organization, d.Get("project_id").(string), "", config.ApiVersion, &azuredevops.GetTaskGroupsOpts{})

	if err != nil {
		return err
	}

	if res.Count == 0 {
		return errors.New("Project have no task groups")
	}

	var group azuredevops.TaskGroup

	for _, v := range res.TaskGroups {
		if v.Name == name {
			group = v
			break
		}
	}

	if group.Id == "" {
		return errors.New("Task group not found")
	}

	//  d.Set("tasks", []interface{}{group.Tasks[0].Inputs})
	//	var variables []interface{}
	//
	//	for k, v := range group.Tasks {
	//		variables = append(variables, map[string]interface{}{
	//			"is_secret": v.IsSecret,
	//			"value":     v.Value,
	//			"name":      k,
	//		})
	//	}
	//
	//	res := schema.NewSet(schema.HashResource(testResource), variables)
	//	d.Set("variables", res)
	d.Set("group_id", group.Id)
	d.SetId(fmt.Sprintf("%s-%s", d.Get("project_id").(string), group.Id))

	return nil
}
