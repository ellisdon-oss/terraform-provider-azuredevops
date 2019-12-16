# Data Source: Project

Table of Contents
=================

   * [Data Source: Project](#data-source-project)
      * [Description](#description)
      * [Example](#example)
      * [Arguments](#arguments)
      * [Attributes](#attributes)
      * [Azure DevOps Reference](#azure-devops-reference)

## Description

A data source to find the Azure Devops project

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
| `project_id` | string | Required | The project ID/name |

## Attributes

| Name | Description |
|------|-------------|
| `id` | The project ID | 
| `name` | The project name | 
| `abbreviation` | The project name abbreviation | 
| `default_team_image_url` | The avatar image URL | 
| `description` | The project description | 
| `last_update_time` | The last updated time for the project | 
| `revision` | The revision number of project | 
| `state` | The state of the project (active/not active) | 
| `url` | Project URL | 
| `visibility` | The project's visibility (private/public) | 
| `links` | Project links | 
| `default_team` | The default team for the project | 
| `capabilities` | Project capabilities (git enabled/scrum/etc) | 

## Azure DevOps Reference

- [Azure DevOps Project](https://docs.microsoft.com/en-us/azure/devops/organizations/projects/create-project?view=azure-devops)
