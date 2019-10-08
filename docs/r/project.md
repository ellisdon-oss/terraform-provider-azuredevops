# Resource: Project

Table of Contents
=================

   * [Resource: Project](#resource-project)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
          * [Capabilities](#capabilities)
          * [Version Control](#version-control)
          * [Process Template](#process-template)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Resource to manage project in AzureDevOps

## Example

```terraform
resource "azuredevops_project" "default" {
  name       = "a-azuredevops-project"
  visibility = "private"

  capabilities {

    version_control {
      source_control_type = "git"
    }

    process_template {
      template_type_id = "adcc42ab-9882-485e-a3ed-7678f01f66bc"
    }
  }
}
```

**NOTE:** full example can be found [here](../../examples/r/project/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| name | string | Required | Project Name |
| description | string | Optional | Project Description |
| visibility | string | Optional | Visibility of the project(public/private) |
| capabilities | [capabilities](#capabilities) | Optional | Capabilities of the project(source control type,etc) |

## Attributes

| Name | Description |
|------|-------------|
| id | Project ID | 

## Extra

### Capabilities

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| version_control | [version_control](#version-control) | Optional | Version control setting for project |
| process_template | [process_template](#process-template) | Optional | Process Template |

### Version Control

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| source_control_type | string | Optional | Source Control Type |

### Process Template

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| template_type_id | string | Optional | Template Type ID |

## AzureDevOps Reference

- [AzureDevOps Project](https://docs.microsoft.com/en-us/azure/devops/organizations/projects/create-project?view=azure-devops)
