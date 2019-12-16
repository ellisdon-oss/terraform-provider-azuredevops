# Data Source: Build Definition

Table of Contents
=================

   * [Data Source: Build Definition](#data-source-build-definition)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to find a build definition ID

## Example

```terraform
data "azuredevops_build_definition" "default" {
  project_id = "<project id>"
  name = "<build definition name>"
}
```

**NOTE:** full example can be found [here](../../examples/d/build_definition/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `project_id` | string | Required | The project ID/name |
| `name` | string | Required | The build definition name |

## Attributes

| Name | Description |
|------|-------------|
| `id` | The build definition ID | 

## Azure DevOps Reference

- [AzureDevOps Build Pipeline](https://docs.microsoft.com/en-us/azure/devops/pipelines/get-started/what-is-azure-pipelines?view=azure-devops)

