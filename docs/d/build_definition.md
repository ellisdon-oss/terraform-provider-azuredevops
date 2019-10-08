# Data Source: Build Definition

Table of Contents
=================

   * [Data Source: Build Definition](#data-source-build-definition)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to find build definition ID

## Example

```terraform
data "azuredevops_build_definition" "default" {
  project_id = "<project id>"
  name = "Some Build Definition Name"
}
```

**NOTE:** full example can be found [here](../../examples/d/build_definition/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| project_id | string | Required | The Project ID/Name |
| name | string | Required | The Build Definition Name |

## Attributes

| Name | Description |
|------|-------------|
| id | Build Definition ID | 

## AzureDevOps Reference

- [AzureDevOps Build Pipeline](https://docs.microsoft.com/en-us/azure/devops/pipelines/get-started/what-is-azure-pipelines?view=azure-devops)

