# Data Source: Service Endpoint

## Description

A Data source to find service endpoint(kube, aws, azure, etc) on AzureDevOps

## Example

```terraform
data "azuredevops_service_endpoint" "default" {
  project_id = <project id>
	type = "kubernetes"
  name = "<Service Endpoint Name>"
}
```

## Arguments

- `project_id`: (Required, string) The Project ID (ID/Name)
- `name`: (Required, string) The Service Endpoint Name
- `type`: (Required, string) The type of service endpoint(github, kubernetes, etc)

## Attributes

- `id`: Service Endpoint ID
- `owner`: The owner of the service endpoint
- `url`: The url of the service endpoint
