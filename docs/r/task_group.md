# Resource: Task Group

Table of Contents
=================

   * [Resource: Task Group](#resource-task-group)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
          * [Version](#version)
          * [Task](#task)
          * [Input](#input)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A resource to manage task group

## Example

```terraform
resource "azuredevops_task_group" "helm" {
  name       = "Helm Tasks"
  project_id = "<project id>"

  task {
    task_id = "068d5909-43e6-48c5-9e01-7c8a94816220"
    name    = "Install Helm 2.10.0"

    inputs = {
      kubectlVersion         = "1.8.9"
      checkLatestHelmVersion = "false"
      installKubeCtl         = "true"
      checkLatestKubeCtl     = "true"
      helmVersion            = "2.10.0"
    }
  }
}
```

**NOTE:** full example can be found [here](../../examples/r/task_group/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `project_id` | string | Required | Project ID |
| `name` | string | Required | Task group name |
| `runs_on` | list of string | Optional | "Runs On" parameter(default: `"Agent, DeploymentGroup"`) |
| `category` | string | Optional | Task group category(default to Deploy) |
| `version` | [version](#version) | Required | Task group version |
| `task` | [task](#task) | Required | Tasks |
| `input` | collection of [input](#input) | Optional | Input parameter |

## Attributes

| Name | Description |
|------|-------------|
| `group_id` | Task group ID | 
| `revision` | Task group revision | 

## Extra

### Version

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `is_test` | bool | Optional | Mark task group as test version |
| `major` | integer | Required | Major version for task group |
| `minor` | integer | Required | Minor version for task group |
| `patch` | integer | Required | Patch version for task group |

### Task

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Name of the task |
| `definition_type` | string | Optional | Definition type (default: `task`), there is also `metaTask` for adding a task group into a task group |
| `version` | string | Optional | Version of the task |
| `task_id` | string | Required | UUID of the task (recommend using [workflow_task](../d/workflow_task.md) to get the ID)  |
| `enabled` | boolean | Optional | Enable/disable the task |
| `always_run` | boolean | Optional | Enable/disable "Always Run" option in the task |
| `continue_on_error` | boolean | Optional | Enable/disable continue on error option in the task |
| `condition` | string | Optional | [Condition](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/expressions?view=azure-devops#job-status-functions) of the task |
| `environment` | map | Optional | Key/value map of environment variables for task |
| `inputs` | map | Required | Key/value map of settings for task |

### Input

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Name of the parameter |
| `label` | string | Optional | Label of the parameter (leave empty to not display a label) |
| `default` | string | Optional | Default value of the parameter |
| `required` | boolean | Optional | Is the parameter a required value |
| `help_text` | string | Optional | Parameter description |
| `type` | string | Required | Parameter type |

## Azure DevOps Reference

- [Azure DevOps Task Group](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/task-groups?view=azure-devops)
