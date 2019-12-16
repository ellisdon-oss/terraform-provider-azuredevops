# Resource: Build Definition

Table of Contents
=================

   * [Resource: Build Definition](#resource-build-definition)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
          * [Demand](#demand)
          * [Process](#process)
          * [Triggers](#triggers)
          * [Pull Request](#pull-request)
          * [Continuous Integration](#continuous-integration)
          * [Queue](#queue)
          * [Repository](#repository)
          * [Build Variable](#build-variable)
          * [Forks](#forks)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A resource to manage a build definition in Azure DevOps

## Example

```terraform
resource "azuredevops_build_definition" "default" {
  name       = "a-build-pipeline"
  project_id = "<project id>"

  repository {
    name = "<source repository id>"
    url  = "<git repo url>"

    properties = {
      <source repository properties>
    }

    default_branch      = "develop"
    checkout_submodules = true
  }

  build_variable {
    name      = "testBuildVariable"
    value     = "agoodvalue"
  }

  build_variable {
    name      = "secretBuildVariable"
    value     = "asecretvalue"
    is_secret = true
  }

  triggers {
    continuous_integration {
      branch_filters = [
        "+develop",
        "+master",
      ]

      settings_source_type = 1
    }

    pull_request {
      branch_filters = [
        "+develop",
        "+master"
      ]

      settings_source_type = 1
    }
  }
}
```

**NOTE:** full example can be found [here](../../examples/r/build_definition/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Build definition nName |
| `project_id` | string | Required | Project ID |
| `demand` | [demand](#demand) | Optional | Demand for agent pool |
| `process` | [process](#process) | Optional | Process settings for the definition |
| `triggers` | [triggers](#triggers) | Optional | Trigger configuration for the definition |
| `queue` | [queue](#queue) | Optional | Queue setting  |
| `repository` | [repository](#repository) | Required | Repository setting  |
| `build_variable` | [build_variable](#build-variable) | Optional | Build variable for the definition |

## Attributes

| Name | Description |
|------|-------------|
| `id` | Build definition ID | 
| `revision` | Build definition revision | 

## Extra

### Demand

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Demand name |
| `value` | string | Required | Demand value |

### Process

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `type` | string | Optional | Process type |
| `yaml_file_name` | string | Optional | Name for the pipeline YAML file (default: `./azure-pipelines.yml`) |

### Triggers

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `pull_request` | [pull_request](#pull-request) | Optional | Pull request trigger |
| `continuous_integration` | [continuous_integration](#continuous-integration) | Optional | Continuous integration trigger |

### Pull Request

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `auto_cancel` | boolean | Optional | Auto cancel setting |
| `is_comment_required_for_pull_request` | boolean | Optional | Enable/Disable comment requirement |
| `settings_source_type` | integer | Optional | `1` (override from the YAML file), `2` (don't override from YAML file) |
| `path_filters` | list of strings | Optional | Path filter for branch |
| `branch_filters` | list of strings | Optional | Branch filter for branch |
| `forks` | [forks](#forks) | Optional | Branch filter for branch |

### Continuous Integration

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `branch_filters` | list of string | Required | Branch filters for the trigger |
| `path_filters` | list of string | Optional | Path filters for the trigger |
| `batch_changes` | boolean | Optional | Batch changes while a build is in progress|
| `max_concurrent_builds_per_branch` | integer | Optional | Max parallel jobs per branch |
| `polling_interval` | integer | Optional | Polling interval setting |
| `polling_job_id` | string | Optional | Polling job id |
| `settings_source_type` | integer | Optional | `1` (override from the YAML file), `2` (don't override from YAML file) |

### Queue

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `agent_name` | string | Required | Agent for the build (defalut: `Hosted`) |

### Repository

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | The name of the repository |
| `url` | string | Required | The URL of the repository |
| `clean` | string | Optional | Enable cleaning of the repository (defaullt: `"true"`) |
| `type` | string | Optional | Type of the repo (default: `Github`) |
| `properties` | map | Required | Repository settings (can be acquired with [source_repository](../d/source_repository.md) |
| `default_branch` | string | Optional | Set the default branch for repository |
| `checkout_submodules` | boolean | Optional | Enable checking out submodules on build |

### Build Variable

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Variable name |
| `value` | string | Required | Variable Value |
| `is_secret` | boolean | Optional | Mark variable as secret or not |

### Forks

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `allow_secrets` | boolean | Optional | Allow secret vars into forked repository build |
| `enabled` | boolean | Optional | Enable forked repository to be built by this definition |

## Azure DevOps Reference

- [Azure DevOps Build Pipeline](https://docs.microsoft.com/en-us/azure/devops/pipelines/get-started/what-is-azure-pipelines?view=azure-devops)
