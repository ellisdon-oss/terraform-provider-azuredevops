# Data Source: Variable Group

Table of Contents
=================

   * [Data Source: Variable Group](#data-source-variable-group)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to read single variable group from AzureDevOps

Note: if you want to read all the variable groups from AzureDevOps, then check [variable_groups](./variable_groups.md)

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
| project_id | string | Required | The Project ID/Name |
| name | string | Optional | The Variable Group Name(conflict with group_id) |
| group_id | string | Optional | The Variable Group ID(conflict with name) |

## Attributes

| Name | Description | Notes |
|------|-------------|-------|
| id | User ID ||
| group_id | Group ID ||
| name | Group Name ||
| variables | Variables of the Variable Group | check [Variable](#variable) |

## Extra

### Variable

| Name | Type | Description |
|------|-------------|-------|
| name | string | Variable name |
| value | string | Variable Value |
| is_secret | boolean | Status of Value being secret, bool |

## AzureDevOps Reference

- [AzureDevOps Variable Group](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/variable-groups?view=azure-devops&tabs=yaml)

