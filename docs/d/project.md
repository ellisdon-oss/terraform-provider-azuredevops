# Data Source: Project

Table of Contents
=================

   * [Data Source: Project](#data-source-project)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [AzureDevOps Reference](#azuredevops-reference)

## Description

A Data source to find AzureDevops Project

## Example

```terraform
data "azuredevops_project" "default" {
  project_id = "<project id>"
}
```

**NOTE:** full example can be found [here](../../examples/d/project/main.tf)

## Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| project_id | string | Required | The Project ID/Name |

## Attributes

| Name | Description |
|------|-------------|
| id | Project ID | 
| name | Project Name | 
| abbreviation | Project name abbreviation | 
| default_team_image_url | Avatar Image Url | 
| description | Project Description | 
| last_update_time | Last Updated Time for Project | 
| revision | The Revision number of project | 
| state | Project State(active/not active) | 
| url | Project URL | 
| visibility | Project visibility(private/public) | 
| links | Project Links | 
| default_team | Project Default Team | 
| capabilities | Project Capabilities(git enabled/scrum/etc) | 

## AzureDevOps Reference

- [AzureDevOps Project](https://docs.microsoft.com/en-us/azure/devops/organizations/projects/create-project?view=azure-devops)
