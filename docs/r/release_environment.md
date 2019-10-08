# Resource: Release Environment

Table of Contents
=================

   * [Resource: Release Environment](#resource-release-environment)
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
          * [Variable](#variable)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Resource to partial release pipeline(a stage) in AzureDevOps

## Example

```terraform
resource "azuredevops_release_environment" "default" {
  project_id    = "<project id>"
  definition_id = "<definition id>"

  environment {
    name = "Stage 1"
    rank = 1

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

    deploy_phase {
      deployment_input = {
        queueId = "<queue id>"
      }

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
}
```

**NOTE:** full example can be found [here](../../examples/r/release_environment/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| definition_id | string | Required | Release Definition ID |
| project_id | string | Required | Project ID |
| environment | [environment](#environment) | Required | Single Release Stage |

## Attributes

## Extra

### Environment

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| name | string | Required | name of the stage |
| rank | integer | Required | the order of the stage(have to be incremental) |
| pre_deploy_approval | [pre_deploy_approval](#pre-deploy-approval) | Optional | Pre-deploy approval setting for the stage |
| variable | [variable](#variable) | Optional | Variable for the stage(same parameters as release variables) |
| variable_groups | list of integers | Optional | IDs of variable groups that will get associate into the stage |
| condition | [condition](#condition) | Optional | Condition for the stage |
| deploy_phase | [deploy_phase](#deploy-phase) | Required | Deploy Phase for stage(where the actual tasks reside) |

### Pre Deploy Approval

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| options | [options](#options) | Optional | options for the approval |
| approvals | [approvals](#approvals) | Optional | settings for the actual approval |

### Options

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| execution_order | string | Optional | the order of the executions(default to `beforeGates`) |
| timeout_in_minutes | integer | Required | set how long for timeout(in minutes) for approval |
| release_creator_can_be_approver | boolean | Optional | toggle for allowing release creator to be approver |

### Approvals 

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| approver_id | string | Optional | UUID of the approver(can use [user](../d/user.md) or [group](../d/group.md) to get the id) |
| is_automated | boolean | Optional | toggle for approval |
| is_notification_on | boolean | Optional | enable notification |

### Condition

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| condition_type | string | Required | condition type for the stage |
| name | string | Required | condition name |
| value | string | Optional | condition value |

### Deploy Phase

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| workflow_task | [workflow_task](#workflow-task) | Required | Task definition |
| name | string | Required | Name of the Phase |
| rank | integer | Required | Rank(order) of the Phase |
| deployment_input | map | Required | Input for the deployment(like setting agent queue id) |
| phase_type | string | Required | Phase Type(like if is agent based) |

### Workflow Task

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| name | string | Required | Name of the task |
| definition_type | string | Optional | Definition Type(default to `task`), there is also `metaTask` for adding a task group into a task group |
| version | string | Optional | Version of the task |
| task_id | string | Required | UUID of the task(recommended using [workflow_task](../d/workflow_task.md) to get the id)  |
| enabled | boolean | Optional | Enable/Disable the task |
| always_run | boolean | Optional | Enable/Disable Always Run option in the task |
| continue_on_error | boolean | Optional | Enable/Disable continue on error option in the task |
| condition | string | Optional | [Condition](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/expressions?view=azure-devops#job-status-functions) of the task |
| inputs | map | Required | Key/Value Map of settings for task |

### Variable

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| name | string | Required | Variable name |
| value | string | Required | Variable Value |
| is_secret | boolean | Optional | Mark variable as secret or not |

## AzureDevOps Reference

- [AzureDevOps Release Pipeline](https://docs.microsoft.com/en-us/azure/devops/pipelines/get-started/what-is-azure-pipelines?view=azure-devops)
