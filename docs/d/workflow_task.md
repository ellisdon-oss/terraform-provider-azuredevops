# Data Source: Workflow Task

Table of Contents
=================

   * [Data Source: Workflow Task](#data-source-workflow-task)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to find task, useful for finding id for task to be use in pipelines

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
| name | string | Required | Task Name |

## Attributes

| Name | Description |
|------|-------------|
| id | Task ID | 

## AzureDevOps Reference

- [AzureDevOps Task](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/tasks?view=azure-devops&tabs=yaml)

