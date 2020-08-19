# Resource: Extension

Table of Contents
=================

   * [Resource: Extension](#resource-extension)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A resource to manage an extension

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
| `publisher` | string | Required | Publisher name |
| `name` | string | Required | Extension name |
| `version` | string | Optional | Version of extension to install (This has no effect on hosted Azure Devops since it is automatically updated to the latest version) | 
| `state` | string | Optional | `"none"` (default) to enable, and `"disabled"` to disable the extension |

## Attributes

| Name | Description |
|------|-------------|
| `id` | Extension ID | 

## Azure DevOps Reference

- [Azure DevOps Extension](https://docs.microsoft.com/en-us/azure/devops/marketplace/install-extension?view=azure-devops)
