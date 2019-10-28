# Resource: Extension

Table of Contents
=================

   * [Resource: Extension](#resource-extension)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Resource to manage extension

## Example

```terraform
resource "azuredevops_extension" "default" {
  publisher = "petergroenewegen"
  name = "PeterGroenewegen-Xpirit-Vsts-Release-Terraform"
}
```

**NOTE:** full example can be found [here](../../examples/r/extension/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| publisher | string | Required | Publisher Name |
| name | string | Required | Extension Name |
| version | string | Optional | Version of Extension to install(no affect on Hosted AzureDevOps since it auto update to latest version) |
| state | string | Optional | "none"(default) for enable, and "disabled" to disable the extension |

## Attributes

| Name | Description |
|------|-------------|
| id | Extension ID | 

## AzureDevOps Reference

- [AzureDevOps Extension](https://docs.microsoft.com/en-us/azure/devops/marketplace/install-extension?view=azure-devops)
