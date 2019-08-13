# Data Source: Project

## Description

A Data source to find AzureDevops Project

## Example

```terraform
data "azuredevops_project" "default" {
  project_id = "Project ID"
}
```

## Arguments

- `project_id`: (Required, string) The project id (ID/Name)

## Attributes

- `id`: Project ID
- `name`: Project Name
- `abbreviation`: Project name abbreviation
- `default_team_image_url`: Avatar Image Url
- `description`: Project Description
- `last_update_time`: Last Updated Time for Project
- `revision`: The Revision number of project
- `state`: Project State(active/not active)
- `url`: Project URL
- `visibility`: Project visibility(private/public)
- `links`: Project Links
- `default_team`: Project Default Team
- `capabilities`: Project Capabilities(git enabled/scrum/etc)
