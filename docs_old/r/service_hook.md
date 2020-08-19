# Resource: Service Hook

Table of Contents
=================

   * [Resource: Service Hook](#resource-service-hook)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
          * [Publisher](#publisher)
          * [Consumer](#consumer)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A resource to manage service hook (Slack, webhook, etc.)

## Example

```terraform
resource "azuredevops_service_hook" "slack" {
  consumer {
    id        = "slack"
    action_id = "postMessageToChannel"
    inputs = {
      url = "<slack webhook url>"
    }
  }

  publisher {
    id = "tfs"
    inputs = {
      buildStatus    = "Failed"
      definitionName = "<definition name>"
      projectId      = "<project id>"
    }
  }

  event_type = "build.complete"
}
```

**NOTE:** full example can be found [here](../../examples/r/service_hook/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `publisher` | [publisher](#publisher) | Required | Publisher (the sender) in the message chain |
| `custom_path` | string | Optional | Use for special events like the ones in vsrm.dev.azure.com |
| `consumer` | [consumer](#consumer) | Required | Consumer (the receiver) in the message chain |
| `event_type` | string | Required | Event type of the message chain |

## Attributes

| Name | Description |
|------|-------------|
| `id` | Service Hook ID | 

## Extra

### Publisher

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `id` | string | Required | The id of the publisher (TFS, etc.) |
| `inputs` | map | Required | Inputs for the publisher( settings for the publisher) |

### Consumer

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `id` | string | Required | The id of the consumer (Slack, etc.) |
| `action_id` | string | Required | The specific action wish to perform for the consumer |
| `inputs` | map | Required | Inputs for the consumer (settings for the consumer) |

## Azure DevOps Reference

- [Azure DevOps Service Hook](https://docs.microsoft.com/en-us/azure/devops/organizations/projects/create-project?view=azure-devops)
