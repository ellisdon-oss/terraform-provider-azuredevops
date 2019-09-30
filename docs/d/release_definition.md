# Data Source: Release Definition

## Description

A Data source to find release definition(release pipeline) on AzureDevOps

## Example

```terraform
data "azuredevops_release_definition" "default" {
  project_id = <project id>
  name = "<Release Pipeline Name>"
}
```

## Arguments

- `project_id`: (Required, string) The Project ID (ID/Name)
- `name`: (Required, string) The Release Pipeline Name

## Attributes

- `id`: Latest Release Definition ID
