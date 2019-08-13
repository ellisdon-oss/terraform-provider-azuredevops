# Data Source: Agent Queue

## Description

A Data source to find agent queue ID due to agent queue being different for each project

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
