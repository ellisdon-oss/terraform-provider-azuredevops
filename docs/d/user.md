# Data Source: User

Table of Contents
=================

   * [Data Source: User](#data-source-user)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to find AzureDevOps user (useful when doing approval step in release pipeline)

## Example

```terraform
data "azuredevops_user" "default" {
  name = "AzureDevops User"
}
```

**NOTE:** full example can be found [here](../../examples/d/user/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| name | string | Required | The name on AzureDevOps UI for the user |
| is_email | boolean | Optional | Switch search between Display Name and Email |

## Attributes

| Name | Description |
|------|-------------|
| id | User ID | 

## AzureDevOps Reference

- [AzureDevOps User](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/manage-users-table-view?view=azure-devops&tabs=browser)

