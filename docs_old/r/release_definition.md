# Resource: Release Definition

Table of Contents
=================

   * [Resource: Release Definition](#resource-release-definition)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
          * [Environment](#environment)
          * [Pre Deploy Approval](#pre-deploy-approval)
          * [Options](#options)
          * [Approvals](#approvals)
          * [Condition](#condition)
          * [Deploy Phase](#deploy-phase)
          * [Workflow Task](#workflow-task)
          * [Artifact](#artifact)
          * [Trigger](#trigger)
          * [Release Variable](#release-variable)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A resource for a full release pipeline in Azure DevOps

## Example

```terraform
resource "azuredevops_release_definition" "default" {
  name       = "a-release-pipeline"
  project_id = "<project id>"

  environment {
    name = "Stage 1"
    rank = 1

    variable {
      name      = "aVariable"
      value     = "aValue"
    }

    condition {
      condition_type = "artifact"
      name           = "drop"
      value = jsonencode({
        sourceBranch                = "develop"
        tags                        = []
        useBuildDefinitionBranch    = false
        createReleaseOnBuildTagging = false
      })
    }

    condition {
      condition_type = "event"
      name           = "ReleaseStarted"
      value          = ""
    }

    pre_deploy_approval {
      approvals {
        approver_id = "<user/group id>"
        is_automated = false
      }
    }

    deploy_phase {
      deployment_input = jsonencode({
        queueId = "<queue id>"
      })


      rank       = 1
      phase_type = "agentBasedDeployment"
      name       = "Run on Agent Test"

      workflow_task {
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
  }

  release_variable {
    name      = "aReleaseVariable"
    value     = "aValue"
  }

  release_variable {
    name      = "aSecretReleaseVariable"
    value     = "aSecretValue"
    is_secret = true
  }

  artifact {
    alias     = "drop"
    source_id = "<project id>:<build definition id>"
    type      = "Build"

    definition_reference = "<definition reference>"
  }
}
```

**NOTE:** full example can be found [here](../../examples/r/release_definition/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Release definition name |
| `project_id` | string | Required | Project ID |
| `path` | string | Optional | Folder path to store release definition into (default: `\\`) |
| `environment` | [environment](#environment)(can be multiple) | Optional | Release stages |
| `release_variable` | [release_variable](#release-variable) | Optional | Release variables |
| `release_variable_groups` | list of integers | Optional | IDs of variable groups to be associated with the release pipeline |
| `artifact` | [artifact](#artifact)(can be multiple) | Optional | Artifacts for the release pipeline |
| `trigger` | [trigger](#artifact)(can be multiple) | Optional | Triggers for the release pipeline |

## Attributes

| Name | Description |
|------|-------------|
| `id` | Release Pipeline ID | 

## Extra

### Environment

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | name of the stage |
| `rank` | integer | Required | the order of the stage(have to be incremental) |
| `pre_deploy_approval` | [pre_deploy_approval](#pre-deploy-approval) | Optional | Pre-deploy approval setting for the stage |
| `variable` | [variable](#release-variable) | Optional | Variable for the stage (same parameters as release variables) |
| `variable_groups` | list of integers | Optional | IDs of variable groups to be associated with the stage |
| `condition` | [condition](#condition) | Optional | Condition for the stage |
| `deploy_phase` | [deploy_phase](#deploy-phase) | Required | Deploy phase for stage (where the actual tasks reside) |

### Pre Deploy Approval

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `options` | [options](#options) | Optional | Options for the approval |
| `approvals` | [approvals](#approvals) | Optional | Settings for the actual approval |

### Options

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `execution_order` | string | Optional | The order of the executions(default to `beforeGates`) |
| `timeout_in_minutes` | integer | Required | Set how long until timeout (in minutes) for approval |
| `release_creator_can_be_approver` | boolean | Optional | Toggle for allowing the release creator to be approver |
| `required_approver_count` | int | Optional | Number of approver required |

### Approvals 

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `approver_id` | string | Optional | UUID of the approver (can use [user](../d/user.md) or [group](../d/group.md) to get the ID) |
| `rank` | int | Optional | Default to 1, is use for doing sequence approval |
| `is_automated` | boolean | Optional | Toggle for approval |
| `is_notification_on` | boolean | Optional | Enable notification |


### Condition

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `condition_type` | string | Required | Condition type for the stage |
| `name` | string | Required | Condition name |
| `value` | string | Optional | Condition value |

### Deploy Phase

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `workflow_task` | [workflow_task](#workflow-task) | Required | Task definition |
| `name` | string | Required | Name of the phase |
| `rank` | integer | Required | Rank (order) of the phase |
| `deployment_input` | json string | Required | Input for the deployment(e.g. setting agent queue id)(json string, use jsonencode() function) |
| `phase_type` | string | Required | Phase type (e.g. is agent based) |


### Workflow Task

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Name of the task |
| `definition_type` | string | Optional | Definition type (default: `task`), there is also `metaTask` for adding a task group into a task group |
| `version` | string | Optional | Version of the task |
| `ref_name` | string | Optional | Output reference name |
| `task_id` | string | Required | UUID of the task (recommended using [workflow_task](../d/workflow_task.md) to get the ID)  |
| `enabled` | boolean | Optional | Enable/disable the task |
| `always_run` | boolean | Optional | Enable/disable "Always Run" option in the task |
| `continue_on_error` | boolean | Optional | Enable/disable continue on error option in the task |
| `condition` | string | Optional | [Condition](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/expressions?view=azure-devops#job-status-functions) of the task |
| `inputs` | map | Required | Key/value Map of settings for task |

### Artifact

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `alias` | string | Required | Alias for the artifact(example `drop`) |
| `source_id` | string | Required | Source setting for artifact |
| `type` | string | Required | Artifact type |
| `definition_reference` | string | Required | JSON string of the extra definition reference (example can be found [here](../../examples/r/release_definition/main.tf) |

### Trigger

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `alias` | string | Required | Alias for target artifact(e.g. `drop`) |
| `branch_filters` | list of string | Required | Branch filters for repository |
| `trigger_type` | string | Required | Trigger type (e.g. `sourceRepo`) |

### Release Variable

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Variable name |
| `value` | string | Required | Variable Value |
| `is_secret` | boolean | Optional | Mark variable as secret or not |

## Azure DevOps Reference

- [Azure DevOps Release Pipeline](https://docs.microsoft.com/en-us/azure/devops/pipelines/get-started/what-is-azure-pipelines?view=azure-devops)
