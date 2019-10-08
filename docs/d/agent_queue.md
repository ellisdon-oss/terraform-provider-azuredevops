# Data Source: Agent Queue

Table of Contents
=================

   * [Data Source: Agent Queue](#data-source-agent-queue)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to find agent queue ID due to agent queue being different for each project

## Example

```terraform
data "azuredevops_agent_queue" "default" {
  project_id = "<project id>"
  queue_name = "VS2017 Hosted"
}
```

**NOTE:** full example can be found [here](../../examples/d/agent_queue/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| project_id | string | Required | The Project ID/Name |
| queue_name | string | Required | The Queue Name |

## Attributes

| Name | Description |
|------|-------------|
| id | Queue ID | 

## AzureDevOps Reference

- [AzureDevOps Agent](https://docs.microsoft.com/en-us/azure/devops/pipelines/agents/agents?view=azure-devops)

