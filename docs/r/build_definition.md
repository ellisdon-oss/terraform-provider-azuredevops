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
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Resource to manage project in AzureDevOps

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
| name | string | Required | Build Definition Name |
| project_id | string | Required | Project ID |
| demand | [demand](#demand) | Optional | Demand for agent pool |
| process | [process](#process) | Optional | Process settings for the definition |
| triggers | [triggers](#triggers) | Optional | Trigger configuration for the definition |
| queue | [queue](#queue) | Optional | Queue setting  |
| repository | [repository](#repository) | Required | Repository setting  |
| build_variable | [build_variable](#build-variable) | Optional | Build Variable for the definition |

## Attributes

| Name | Description |
|------|-------------|
| id | Build Definition ID | 
| revision | Build Definition Revision | 

## Extra

### Demand

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| name | string | Required | Demand name |
| value | string | Required | Demand value |

### Process

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| type | string | Optional | Process type |
| yaml_file_name | string | Optional | Name for the pipeline yaml(default to `./azure-pipelines.yml`) |

### Triggers

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| pull_request | [pull_request](#pull-request) | Optional | Pull Request Trigger |
| continuous_integration | [continuous_integration](#continuous-integration) | Optional | Continuous Integration Trigger |

### Pull Request

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| auto_cancel | boolean | Optional | Auto cancel setting |
| is_comment_required_for_pull_request | boolean | Optional | Enable/Disable comment requirement |
| settings_source_type | integer | Optional | 1(override from the yaml file), 2(not override from yaml file) |
| path_filters | list of strings | Optional | Path filter for branch |
| branch_filters | list of strings | Optional | Branch filter for branch |
| forks | [forks](#forks) | Optional | Branch filter for branch |

### Continuous Integration

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| branch_filters | list of string | Required | Branch filters for the trigger |
| path_filters | list of string | Optional | Path filters for the trigger |
| batch_changes | boolean | Optional | Batch changes while a build is in progress|
| max_concurrent_builds_per_branch | integer | Optional | Max parallel job per branch |
| polling_interval | integer | Optional | Polling interval setting |
| polling_job_id | string | Optional | Polling job id |
| settings_source_type | integer | Optional | 1(override from the yaml file), 2(not override from yaml file) |

### Queue

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| agent_name | string | Required | Agent for the build(defalut to `Hosted`) |

### Repository

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| name | string | Required | name of the repo |
| url | string | Required | url of the repo |
| clean | string | Optional | enable cleaning of the repo(defualt to `"true"`) |
| type | string | Optional | Type of the repo(default to `Github`) |
| properties | map | Required | Repo settings(can be acquire with [source_repository](../d/source_repository.md) |
| default_branch | string | Optional | set default branch for repo |
| checkout_submodules | boolean | Optional | enable to checkout submodules on build |

### Build Variable

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| name | string | Required | Variable name |
| value | string | Required | Variable Value |
| is_secret | boolean | Optional | Mark variable as secret or not |

### Forks

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| allow_secrets | boolean | Optional | Allow secret vars into forked repo build |
| enabled | boolean | Optional | Enabled forked repo to be build by this definition |

## AzureDevOps Reference

- [AzureDevOps Build Pipeline](https://docs.microsoft.com/en-us/azure/devops/pipelines/get-started/what-is-azure-pipelines?view=azure-devops)
