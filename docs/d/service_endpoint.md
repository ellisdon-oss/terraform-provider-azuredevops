# Data Source: Service Endpoint

Table of Contents
=================

   * [Data Source: Service Endpoint](#data-source-service-endpoint)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to find service endpoints (Kubernetes, AWS, Azure, etc.) on AzureDevOps

## Example

```terraform
data "azuredevops_service_endpoint" "default" {
  project_id = "<project id>"
	type = "kubernetes"
  name = "<service endpoint name>"
}
```

**NOTE:** full example can be found [here](../../examples/d/service_endpoint/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `project_id` | string | Required | The Project ID/Name |
| `name` | string | Required | The service endpoint name |
| `type` | string | Required | The type of service endpoint (`github`, `kubernetes`, etc.) |

## Attributes

| Name | Description |
|------|-------------|
| `id` | Service Endpoint ID | 
| `owner` | The owner of the service endpoint | 
| `url` | The URL of the service endpoint | 

## Azure DevOps Reference

- [Azure DevOps Service Endpoint](https://docs.microsoft.com/en-us/azure/devops/extend/develop/service-endpoints?view=azure-devops)
