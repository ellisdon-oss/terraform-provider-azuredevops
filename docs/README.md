# Documentation Listing

Table of Contents
=================

   * [Documentation Listing](#documentation-listing)
      * [Provider](#provider)
        * [Example](#example)
        * [Arguments](#arguments)
      * [Data Sources](#data-sources)
      * [Resources](#resources)

## Provider

### Examples

```terraform
provider "azuredevops" {
  token = "<azuredevops token>"
  organization_url = "https://dev.azure.com/<organization name>"
}
```

**NOTE:** Guide to get token are [here](../../examples/)

### Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| token | string | Required | AzureDevOps token |
| organization_url | string | Required | Organization URL |

## Data Sources

| Name | Description | Docs | Examples |
|------|-------------|------|----------|
| azuredevops_agent_queue | Agent Queue | [Link](./d/agent_queue.md) | [Link](../examples/d/agent_queue/main.tf) |
| azuredevops_build_definition | Build Definition(Build Pipeline) | [Link](./d/build_definition.md) | [Link](../examples/d/build_definition/main.tf) |
| azuredevops_group | Group | [Link](./d/group.md) | [Link](../examples/d/group/main.tf) |
| azuredevops_user | User | [Link](./d/user.md) | [Link](../examples/d/user/main.tf) |
| azuredevops_project | Project | [Link](./d/project.md) | [Link](../examples/d/project/main.tf) |
| azuredevops_release_definition | Release Definition(Release Pipeline) | [Link](./d/release_definition.md) | [Link](../examples/d/release_definition/main.tf) |
| azuredevops_service_endpoint | Service Endpoint | [Link](./d/service_endpoint.md) | [Link](../examples/d/service_endpoint/main.tf) |
| azuredevops_source_repository | Source Repository | [Link](./d/source_repository.md) | [Link](../examples/d/source_repository/main.tf) |
| azuredevops_task_group | Task Group | [Link](./d/task_group.md) | [Link](../examples/d/task_group/main.tf) |
| azuredevops_variable_group | Variable Group(Single Variable Group) | [Link](./d/variable_group.md) | [Link](../examples/d/variable_group/main.tf) |
| azuredevops_variable_groups | Variable Groups(Multiple Variable Group) | [Link](./d/variable_groups.md) | [Link](../examples/d/variable_groups/main.tf) |
| azuredevops_workflow_task | Workflow Task(for Pipelines) | [Link](./d/workflow_task.md) | [Link](../examples/d/workflow_task/main.tf) |

## Resources

| Name | Description | Docs | Examples |
|------|-------------|------|----------|
| azuredevops_build_definition | Manage Build Definition | [Link](./r/build_definition.md) | [Link](../examples/r/build_definition/main.tf) |
| azuredevops_project | Manage Project | [Link](./r/project.md) | [Link](../examples/r/project/main.tf) |
| azuredevops_release_definition | Manage Full Release Definition | [Link](./r/release_definition.md) | [Link](../examples/r/release_definition/main.tf) |
| azuredevops_release_environment | Manage Partial Release Definition | [Link](./r/release_environment.md) | [Link](../examples/r/release_environment/main.tf) |
| azuredevops_service_endpoint | Manage Service Endpoint(Github, Kubernetes, etc) | [Link](./r/service_endpoint.md) | [Link](../examples/r/service_endpoint/main.tf) |
| azuredevops_service_hook | Manage Service Hook(Slack, etc) | [Link](./r/service_hook.md) | [Link](../examples/r/service_hook/main.tf) |
| azuredevops_task_group | Manage Task Group | [Link](./r/task_group.md) | [Link](../examples/r/task_group/main.tf) |
| azuredevops_variable_group | Manage Variable Group | [Link](./r/variable_group.md) | [Link](../examples/r/variable_group/main.tf) |
