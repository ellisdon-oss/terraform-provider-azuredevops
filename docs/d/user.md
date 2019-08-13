# Data Source: User

## Description

A Data source to find AzureDevOps user (useful when doing approval step in release pipeline)

## Example

```terraform
data "azuredevops_user" "default" {
  display_name = "AzureDevops User"
}
```

## Arguments

- `display_name`: (Required, string) The name on AzureDevOps UI for the user

## Attributes

- `id`: User ID
