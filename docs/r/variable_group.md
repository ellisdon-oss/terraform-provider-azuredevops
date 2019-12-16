# Resource: Variable Group

Table of Contents
=================

   * [Resource: Variable Group](#resource-variable-group)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
          * [Variable](#variable)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A resource to manage a variable group

## Example

```terraform
resource "azuredevops_variable_group" "default" {
  project_id = "<project id>"

  name = "a-variable-group"

  variable {
    name = "NormalVar"
    value = "notsosecretvalue"
  }

  variable {
    name = "SecretVar"
    value = "supersecretvalue"
    is_secret = true
  }
}
```

**NOTE:** full example can be found [here](../../examples/r/variable_group/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `project_id` | string | Required | Project ID |
| `name` | string | Required | Variable group name |
| `variable` | [variable](#variable) | Required | Variables |

## Attributes

| Name | Description |
|------|-------------|
| `id` | Variable Group ID | 

## Extra

### Variable

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Variable name |
| `value` | string | Required | Variable value |
| `is_secret` | boolean | Optional | Mark variable as secret or not |

## Azure DevOps Reference

- [Azure DevOps Variable Group](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/variable-groups?view=azure-devops&tabs=yaml)
