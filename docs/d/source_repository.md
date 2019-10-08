# Data Source: Source Repository

Table of Contents
=================

   * [Data Source: Source Repository](#data-source-source-repository)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to find source repository(git repo) on AzureDevOps

## Example

```terraform
data "azuredevops_source_repository" "default" {
  project_id = <project id>
	type = "github"
  org_name = "<Github Organization Name>"
  repo_name = "<Github Repo Name>"
  service_endpoint_id = "<Github Service Endpoint ID>"
}
```

**NOTE:** full example can be found [here](../../examples/d/source_repository/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| project_id | string | Required | The Project ID/Name |
| service_endpoint_id | string | Required | The Service Endpoint ID needed for grabbing source repository |
| type | string | Required | The type of source repository(github, etc) |
| org_name | string | Required | The organization name for the source repository |
| repo_name | string | Required | The repo name for the source repository |

## Attributes

| Name | Description |
|------|-------------|
| id | Source Repository ID | 
| url | Url for the source repository | 
| default_branch | Default branch for the source repository | 
| properties | A Map of the properties of the repo | 

## AzureDevOps Reference

- [AzureDevOps Source Repository](https://docs.microsoft.com/en-us/azure/devops/pipelines/repos/?view=azure-devops)

