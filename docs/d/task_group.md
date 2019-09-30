# Data Source: Task Group

## Description

A Data source to find service endpoint(kube, aws, azure, etc) on AzureDevOps

## Example

```terraform
data "azuredevops_task_group" "default" {
  project_id = <project id>
	type = "kubernetes"
  name = "<Task Group Name>"
}
```

## Arguments

- `project_id`: (Required, string) The Project ID (ID/Name)
- `name`: (Required, string) The Task Group Name

## Attributes

- `group_id`: Task Group ID
