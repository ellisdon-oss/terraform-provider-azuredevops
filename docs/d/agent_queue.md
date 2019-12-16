# Data Source: Agent Queue

Table of Contents
=================

   * [Data Source: Agent Queue](#data-source-agent-queue)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to find the agent queue ID when the agent queue is different for each project

## Example

```terraform
data "azuredevops_agent_queue" "default" {
  project_id = "<project id>"
  queue_name = "<queue name>"
}
```

**NOTE:** full example can be found [here](../../examples/d/agent_queue/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `project_id` | string | Required | The project name/ID |
| `queue_name` | string | Required | The agent queue name |

## Attributes

| Name | Description |
|------|-------------|
| `id` | The agent queue ID | 

## Azure DevOps Reference

- [Azure DevOps Agent](https://docs.microsoft.com/en-us/azure/devops/pipelines/agents/agents?view=azure-devops)

