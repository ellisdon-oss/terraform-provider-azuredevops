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
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A resource to manage service endpoints (GitHub, Kubernetes, etc)

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
}
```

**NOTE:** full example can be found [here](../../examples/r/service_endpoint/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Service endpoint name |
| `type` | string | Required | Service endpoint type (github, kubernetes, etc) |
| `project_id` | string | Required | Project ID (which project the service endpoint will be created in) |
| `allow_all_pipelines` | boolean | Optional | To allow all pipelines use this service endpoint right away |
| `data` | map | Optional | Data for the service endpoint (different for each service endpoint type) |
| `owner` | string | Required | Owner of the service endpoint type (not the user, but the owner of the type) |
| `url` | string | Required | URL for the service endpoint |
| `authorization` | [authorization](#authorization) | Required | Authorization data for the service endpoint |

## Attributes

| Name | Description |
|------|-------------|
| `id` | Service endpoint ID | 

## Extra

### Authorization

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `scheme` | string | Required | Authorization scheme |
| `parameters` | map | Required | Data for the authorization scheme |

## Azure DevOps Reference

- [Azure DevOps Service Endpoint](https://docs.microsoft.com/en-us/azure/devops/extend/develop/service-endpoints?view=azure-devops)
