# Resource: Service Endpoint

Table of Contents
=================

   * [Resource: Service Endpoint](#resource-service-endpoint)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
          * [Authorization](#authorization)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Resource to manage service hook(slack, webhook, etc)

## Example

```terraform
resource "azuredevops_service_endpoint" "github" {
  name       = "github-example"
  owner      = "Library"
  project_id = "<project id>"
  type       = "github"
  url        = "http://github.com"

  allow_all_pipelines = true

  authorization {
    scheme = "PersonalAccessToken"
    parameters = {
      accessToken = "<github-token>"
    }
  }

  data = {}
}
```

**NOTE:** full example can be found [here](../../examples/r/service_endpoint/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| name | string | Required | Service Endpoint Name |
| type | string | Required | Service Endpoint Type(github, kubernetes, etc) |
| project_id | string | Required | Project ID(which project the service endpoint will be create in) |
| allow_all_pipelines | boolean | Optional | To allow all pipelines use this service endpoint right away |
| data | map | Required | Data for the service endpoint(different for each service endpoint type) |
| owner | string | Required | Owner of the Service Endpoint Type(not the user, but the type owner) |
| url | string | Required | URL for the Service Endpoint |
| authorization | [authorization](#authorization) | Required | Authorization data for the Service Endpoint |

## Attributes

| Name | Description |
|------|-------------|
| id | Service Endpoint ID | 

## Extra

### Authorization

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| scheme | string | Required | scheme type for the authorization |
| parameters | map | Required | Data for the authorization scheme |

## AzureDevOps Reference

- [AzureDevOps Service Endpoint](https://docs.microsoft.com/en-us/azure/devops/extend/develop/service-endpoints?view=azure-devops)
