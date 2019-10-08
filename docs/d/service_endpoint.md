# Data Source: Service Endpoint

Table of Contents
=================

   * [Data Source: Release Definition](#data-source-release-definition)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to find service endpoint(kube, aws, azure, etc) on AzureDevOps

## Example

```terraform
data "azuredevops_service_endpoint" "default" {
  project_id = "<project id>"
	type = "kubernetes"
  name = "<Service Endpoint Name>"
}
```

**NOTE:** full example can be found [here](../../examples/d/service_endpoint/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| project_id | string | Required | The Project ID/Name |
| name | string | Required | The Service Endpoint Name |
| type | string | Required | The type of service endpoint(github, kubernetes, etc) |

## Attributes

| Name | Description |
|------|-------------|
| id | Service Endpoint ID | 
| owner | The owner of the service endpoint | 
| url | The url of the service endpoint | 

## AzureDevOps Reference

- [AzureDevOps Service Endpoint](https://docs.microsoft.com/en-us/azure/devops/extend/develop/service-endpoints?view=azure-devops)
