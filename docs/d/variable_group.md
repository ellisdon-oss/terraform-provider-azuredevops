# Data Source: Variable Group

## Description

A Data source to read single variable group from AzureDevOps

Note: if you want to read all the variable groups from AzureDevOps, then check [variable_groups](./variable_groups.md)

## Example

```terraform
data "azuredevops_agent_queue" "default" {
  project_id = <project id>
  queue_name = "VS2017 Hosted"
}
```

## Arguments

- `project_id`: (Required, string) The Project ID (ID/Name)
- `queue_name`: (Required, string) The Queue Name

## Attributes

- `id`: Queue ID
