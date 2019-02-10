package azuredevops

import (
	"github.com/antihax/optional"
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func dataSourceSourceRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSourceRepositoryRead,

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"org_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"repo_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_endpoint_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"properties": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_branch": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSourceRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := optional.NewString(d.Get("org_name").(string) + "/" + d.Get("repo_name").(string))
	projId := d.Get("project_id").(string)
	sourceType := d.Get("type").(string)
	serviceId := d.Get("service_endpoint_id").(string)

	repos, _, err := config.Client.SourceProvidersApi.ListRepositories(config.Context, config.Organization, projId, sourceType, config.ApiVersion, &azuredevops.ListRepositoriesOpts{
		ServiceEndpointId: optional.NewInterface(serviceId),
		Repository:        name,
	})

	if err != nil {
		return err
	}

	if repos.PageLength == 0 {
		return errors.New("Repo Not Found")
	}

	repo := repos.Repositories[0]

	d.Set("properties", repo.Properties)
	d.Set("url", repo.Url)
	d.Set("default_branch", repo.DefaultBranch)

	d.SetId(repo.Id)

	return nil
}
