# Data Source: Release Definition

Table of Contents
=================

   * [Data Source: Release Definition](#data-source-release-definition)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to find release definition(release pipeline) on AzureDevOps

## Example

```terraform
data "azuredevops_release_definition" "default" {
  project_id = "<project id>"
  name = "<Release Pipeline Name>"
}
```

**NOTE:** full example can be found [here](../../examples/d/release_definition/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| project_id | string | Required | The Project ID/Name |
| name | string | Required | The Release Definition Name |

## Attributes

| Name | Description |
|------|-------------|
| id | Release Definition ID | 

## AzureDevOps Reference

- [AzureDevOps Release Pipeline](https://docs.microsoft.com/en-us/azure/devops/pipelines/get-started/what-is-azure-pipelines?view=azure-devops)
