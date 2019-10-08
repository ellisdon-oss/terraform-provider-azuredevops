# Data Source: Task Group

Table of Contents
=================

   * [Data Source: Task Group](#data-source-task-group)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to find task group on AzureDevOps

## Example

```terraform
data "azuredevops_task_group" "default" {
  project_id = "<project id>"
  name = "<Task Group Name>"
}
```

**NOTE:** full example can be found [here](../../examples/d/task_group/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| project_id | string | Required | The Project ID/Name |
| name | string | Required | The Task Group Name |

## Attributes

| Name | Description |
|------|-------------|
| group_id | Task Group ID | 

## AzureDevOps Reference

- [AzureDevOps Task Group](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/task-groups?view=azure-devops)
