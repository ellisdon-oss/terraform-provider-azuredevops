# Data Source: Group

Table of Contents
=================

   * [Data Source: Group](#data-source-group)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to find AzureDevOps group (useful when doing approval step in release pipeline)

## Example

```terraform
data "azuredevops_group" "default" {
  display_name = "AzureDevops Group"
}
```

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| display_name | string | Required | The name on AzureDevOps UI for the group |

## Attributes

| Name | Description |
|------|-------------|
| id | Group ID | 

## AzureDevOps Reference

- [AzureDevOps Group](https://docs.microsoft.com/en-us/azure/devops/organizations/security/permissions?view=azure-devops&tabs=preview-page)
