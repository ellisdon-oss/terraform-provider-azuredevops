package azuredevops

import (
  "fmt"
  "strconv"
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
  "github.com/microsoft/azure-devops-go-api/azuredevops/release"
)

func dataSourceReleaseStageVariables() *schema.Resource {
  return &schema.Resource{
    Read: dataSourceReleaseStageVariablesRead,

    Schema: map[string]*schema.Schema{
      "project_id": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "definition_id": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
        ForceNew: true,
      },
      "stage_name": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "variables": &schema.Schema{
        Type:     schema.TypeSet,
        Computed: true,
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

func dataSourceReleaseStageVariablesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	releaseClient, err := release.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	projectID := d.Get("project_id").(string)
	definition_id, _ := strconv.ParseInt(d.Get("definition_id").(string), 10, 32)

	defID := int(definition_id)

	releaseDef, err := releaseClient.GetReleaseDefinition(config.Context, release.GetReleaseDefinitionArgs{
		Project:    &projectID,
    DefinitionId: &defID,
	})

	if err != nil {
		return err
	}

	var variables []interface{}

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

  for _, env := range *releaseDef.Environments {
    if *env.Name == d.Get("stage_name").(string) {
      for k, v := range *env.Variables {
        if k != "" {
          value := ""
          isSecret := false

          if v.IsSecret != nil && !(*v.IsSecret) {
            isSecret = true
          } else if v.Value == nil {
            value = ""
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
    }
  }

  d.Set("variables", schema.NewSet(schema.HashResource(testResource), variables))
  
	d.SetId(fmt.Sprintf("%d-%s",*releaseDef.Id, d.Get("stage_name").(string)))

	return nil
}
