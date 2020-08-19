# Data Source: Variable Groups

Table of Contents
=================

   * [Data Source: Variable Groups](#data-source-variable-groups)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
         * [Variable Group](#variable-group)
         * [Variable](#variable)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to read all variable groups from an Azure DevOps project

Note: if you want to read single variable groups from Azure DevOps, then check [variable_group](./variable_group.md)

## Example

```terraform
data "azuredevops_variable_groups" "default" {
  project_id = "<project id>"
}
```

**NOTE:** full example can be found [here](../../examples/d/variable_groups/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `project_id` | string | Required | The Project ID/Name |

## Attributes

| Name | Description |
|------|-------------|
| `id` | User ID |
| `groups` | An array of [variable groups](#variable-group) |

## Extra

### Variable Group

| Name | Type | Description |
|------|-------------|-------|
| `group_id` | integer | Variable Group ID |
| `name` | string | Variable Group Name |
| `variables` | [variable[]](#variable) | The variables in the group |

### Variable

| Name | Type | Description |
|------|-------------|-------|
| `name` | string | Variable name |
| `value` | string | Variable value |
| `is_secret` | boolean | Mark variable as secret or not |

## Azure DevOps Reference

- [Azure DevOps Variable Group](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/variable-groups?view=azure-devops&tabs=yaml)
