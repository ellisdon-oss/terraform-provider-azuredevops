# Data Source: Build Definition

## Description

A Data source to find build definition ID

## Example

```terraform
data "azuredevops_build_definition" "default" {
  project_id = <project id>
  name = "Some Build Definition Name>
}
```

## Arguments

- `project_id`: (Required, string) The Project ID (ID/Name)
- `name`: (Required, string) The Build Definition Name

## Attributes

- `id`: Build Definition ID
