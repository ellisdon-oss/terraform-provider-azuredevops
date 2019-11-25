# Data Source: Workflow Task

Table of Contents
=================

   * [Data Source: Workflow Task](#data-source-workflow-task)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to find a task, useful for finding the ID for a task to be used in pipelines

## Example

```terraform
data "azuredevops_workflow_task" "default" {
  name = "HelmDeploy"
}
```

**NOTE:** full example can be found [here](../../examples/d/workflow_task/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Task Name |
| `extension_id` | string | Optional | Extension ID(only for reference purposes) |
| `wait` | integer | Optional | Wait time for reading, useful when reading extension tasks |

## Attributes

| Name | Description |
|------|-------------|
| `id` | Task ID | 

## Azure DevOps Reference

- [Azure DevOps Task](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/tasks?view=azure-devops&tabs=yaml)

