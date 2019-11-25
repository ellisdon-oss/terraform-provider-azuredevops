# Data Source: Release Definition

Table of Contents
=================

   * [Data Source: Release Definition](#data-source-release-definition)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to find the release definition (release pipeline) on Azure DevOps

## Example

```terraform
data "azuredevops_release_definition" "default" {
  project_id = "<project id>"
  name = "<release pipeline name>"
}
```

**NOTE:** full example can be found [here](../../examples/d/release_definition/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `project_id` | string | Required | The Project ID/Name |
| `name` | string | Required | The release defintion name |

## Attributes

| Name | Description |
|------|-------------|
| `id` | The release definition ID | 

## Azure DevOps Reference

- [Azure DevOps Release Pipeline](https://docs.microsoft.com/en-us/azure/devops/pipelines/get-started/what-is-azure-pipelines?view=azure-devops)
