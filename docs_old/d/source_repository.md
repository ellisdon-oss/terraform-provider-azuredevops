# Data Source: Source Repository

Table of Contents
=================

   * [Data Source: Source Repository](#data-source-source-repository)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to find source repository (git repository) on AzureDevOps

## Example

```terraform
data "azuredevops_source_repository" "default" {
  project_id = "<project id>"
	type = "github"
  org_name = "<GitHub organization name>"
  repo_name = "<GitHub repo name>"
  service_endpoint_id = "<GitHub service endpoint id>"
}
```

**NOTE:** full example can be found [here](../../examples/d/source_repository/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| `project_id` | string | Required | The project ID/name |
| `service_endpoint_id` | string | Required | The service endpoint ID needed for grabbing the source repository |
| `type` | string | Required | The type of source repository (github, etc.) |
| `org_name` | string | Required | The organization name for the source repository |
| `repo_name` | string | Required | The name for the source repository |

## Attributes

| Name | Description |
|------|-------------|
| `id` | The source repository ID | 
| `url` | The URL for the source repository | 
| `default_branch` | The default branch for the source repository | 
| `properties` | A map of the properties of the repository | 

## Azure DevOps Reference

- [Azure DevOps Source Repository](https://docs.microsoft.com/en-us/azure/devops/pipelines/repos/?view=azure-devops)

