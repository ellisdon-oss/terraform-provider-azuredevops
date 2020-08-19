# Resource: Release Variables

Table of Contents
=================

   * [Resource: Release Variables](#resource-release-variables)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Extra](#extra)
          * [Variable](#variable)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A resource to manage release variables in release pipelines

## Example

```terraform
resource "azuredevops_release_variables" "default" {
  project_id = "<project id>"

  definition_id = 1

  variable {
    name = "NormalVar"
    value = "notsosecretvalue"
  }

  variable {
    stage_name = "TestStage"
    name = "SecretVar"
    value = "supersecretvalue"
    is_secret = true
  }
}
```

**NOTE:** full example can be found [here](../../examples/r/release_variables/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `project_id` | string | Required | Project ID |
| `definition_id` | integer | Required | Definition ID |
| `variable` | [variable](#variable) | Required | Variables |

## Attributes

| Name | Description |
|------|-------------|
| `id` | Release Variables UUID | 

## Extra

### Variable

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `name` | string | Required | Variable name |
| `value` | string | Required | Variable Value |
| `stage_name` | string | Optional | Specify stage to add the variable to |
| `is_secret` | boolean | Optional | Mark variable as secret or not |

## Azure DevOps Reference

- [Azure DevOps Release Variables](https://docs.microsoft.com/en-us/azure/devops/pipelines/release/variables?view=azure-devops&tabs=batch)

