# Data Source: User

Table of Contents
=================

   * [Data Source: User](#data-source-user)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to find an Azure DevOps user (useful when doing an approval step in a release pipeline)

## Example

```terraform
data "azuredevops_user" "default" {
  name = "<azure devops user>"
}
```

**NOTE:** full example can be found [here](../../examples/d/user/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | The name in the Azure DevOps UI for the user |
| `is_email` | boolean | Optional | Whether or not to search by email or user name |

## Attributes

| Name | Description |
|------|-------------|
| `id` | User ID | 

## Azure DevOps Reference

- [Azure DevOps User](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/manage-users-table-view?view=azure-devops&tabs=browser)

