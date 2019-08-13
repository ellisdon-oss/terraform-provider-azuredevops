# Data Source: Group

## Description

A Data source to find AzureDevOps group (useful when doing approval step in release pipeline)

## Example

```terraform
data "azuredevops_group" "default" {
  display_name = "AzureDevops Group"
}
```

## Arguments

- `display_name`: (Required, string) The name on AzureDevOps UI for the group

## Attributes

- `id`: Group ID
