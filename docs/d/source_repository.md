# Data Source: Source Repository

## Description

A Data source to find source repository(git repo) on AzureDevOps

## Example

```terraform
data "azuredevops_source_repository" "default" {
  project_id = <project id>
	type = "github"
  org_name = "<Github Organization Name>"
  repo_name = "<Github Repo Name>"
  service_endpoint_id = "<Github Service Endpoint ID>"
}
```

## Arguments

- `project_id`: (Required, string) The Project ID (ID/Name)
- `service_endpoint_id`: (Required, string) The Service Endpoint ID needed for grabbing source repository
- `type`: (Required, string) The type of source repository(github, etc)
- `org_name`: (Required, string) The organization name for the source repository
- `repo_name`: (Required, string) The repo name for the source repository

## Attributes

- `id`: Service Endpoint ID
- `properties`: A Map of the properties of the repo
- `url`: Url for the source repository
- `default_branch`: Default branch for the source repository
