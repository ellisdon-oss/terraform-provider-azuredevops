# Data Source: Variable Groups

Table of Contents
=================

   * [Data Source: Variable Groups](#data-source-variable-groups)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
         * [Variable](#variable-group)
         * [Variable](#variable)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to read all variable groups from a AzureDevOps project

Note: if you want to read single the variable groups from AzureDevOps, then check [variable_group](./variable_group.md)

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
| project_id | string | Required | The Project ID/Name |

## Attributes

| Name | Description |
|------|-------------|
| id | User ID |
| groups | Array of [Variable Group](#variable-group) |

## Extra

### Variable Group

| Name | Type | Description |
|------|-------------|-------|
| group_id | integer | Variable Group ID |
| name | string | Variable Group Name |
| variables | array of [variable](#variable) | The variables in the group |

### Variable

| Name | Type | Description |
|------|-------------|-------|
| name | string | Variable name |
| value | string | Variable Value |
| is_secret | boolean | Status of Value being secret, bool |

## AzureDevOps Reference

- [AzureDevOps Variable Group](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/variable-groups?view=azure-devops&tabs=yaml)
