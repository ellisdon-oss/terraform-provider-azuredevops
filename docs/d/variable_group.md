# Data Source: Variable Group

Table of Contents
=================

   * [Data Source: Variable Group](#data-source-variable-group)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
         * [Variable](#variable)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A Data source to read a single variable group from Azure DevOps

Note: if you want to read all the variable groups from Azure DevOps, then check [variable_groups](./variable_groups.md)

## Example

```terraform
data "azuredevops_variable_group" "default" {
  project_id = "<project id>"
  name = "<variable group name>"
}
```

**NOTE:** full example can be found [here](../../examples/d/variable_group/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `project_id` | string | Required | The project ID/name |
| `name` | string | Optional | The variable group name (mutually exclusive with `group_id`) |
| `group_id` | string | Optional | The variable group ID (mutually exclusive with `name`) |

## Attributes

| Name | Description |
|------|-------------|
| `id` | User ID |
| `group_id` | Group ID |
| `name` | Group name |
| `variables` | Array of [Variables](#variable) |

## Extra

### Variable

| Name | Type | Description |
|------|-------------|-------|
| `name` | string | Variable name |
| `value` | string | Variable value |
| `is_secret` | boolean | Mark variable as secret or not |

## Azure DevOps Reference

- [Azure DevOps Variable Group](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/variable-groups?view=azure-devops&tabs=yaml)
