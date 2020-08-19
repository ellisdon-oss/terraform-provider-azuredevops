# Data Source: Group

Table of Contents
=================

   * [Data Source: Group](#data-source-group)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to find the Azure DevOps group (useful when doing an approval step in a release pipeline)

## Example

```terraform
data "azuredevops_group" "default" {
  display_name = "<group display name>"
}
```

**NOTE:** full example can be found [here](../../examples/d/group/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `display_name` | string | Required | The name in the Azure DevOps UI for the group |

## Attributes

| Name | Description |
|------|-------------|
| `id` | Group ID | 

## Azure DevOps Reference

- [Azure DevOps Group](https://docs.microsoft.com/en-us/azure/devops/organizations/security/permissions?view=azure-devops&tabs=preview-page)
