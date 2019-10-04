package azuredevops

import (
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
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

	buildClient, err := build.NewClient(config.Context, config.Connection)

	name := d.Get("org_name").(string) + "/" + d.Get("repo_name").(string)
	projectID := d.Get("project_id").(string)
	sourceType := d.Get("type").(string)
	serviceEndpoint := d.Get("service_endpoint_id").(string)

	serviceEndpointID, _ := uuid.Parse(serviceEndpoint)
	repos, err := buildClient.ListRepositories(config.Context, build.ListRepositoriesArgs{
		ServiceEndpointId: &serviceEndpointID,
		Project:           &projectID,
		ProviderName:      &sourceType,
		Repository:        &name,
	})

	if err != nil {
		return err
	}

	if (*(*repos).PageLength) == 0 {
		return errors.New("Repo Not Found")
	}

	repo := (*(*repos).Repositories)[0]

	d.Set("properties", repo.Properties)
	d.Set("url", repo.Url)
	d.Set("default_branch", repo.DefaultBranch)

	d.SetId(*repo.Id)

	return nil
}
