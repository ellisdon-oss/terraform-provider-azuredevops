# Data Source: Task Group

Table of Contents
=================

   * [Data Source: Task Group](#data-source-task-group)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to a find task group on AzureDevOps

## Example

```terraform
data "azuredevops_task_group" "default" {
  project_id = "<project id>"
  name = "<task group name>"
}
```

**NOTE:** full example can be found [here](../../examples/d/task_group/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
|` project_id` | string | Required | The project ID/name |
| `name` | string | Required | The task group name |

## Attributes

| Name | Description |
|------|-------------|
| `group_id` | Task group ID | 

## Azure DevOps Reference

- [Azure DevOps Task Group](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/task-groups?view=azure-devops)
