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

**NOTE:** Guide to get token are [here](../examples/)

### Arguments

| Name | Type | Required/Optional | Description |
|------|------|-------------------|-------------|
| token | string | Required | AzureDevOps token |
| organization_url | string | Required | Organization URL |

## Data Sources

| Name | Description | Docs | Examples |
|------|-------------|------|----------|
| azuredevops_agent_queue | Agent Queue | [Link](./d/agent_queue.md) | [Link](../examples/d/agent_queue/README.md) |
| azuredevops_build_definition | Build Definition(Build Pipeline) | [Link](./d/build_definition.md) | [Link](../examples/d/build_definition/README.md) |
| azuredevops_group | Group | [Link](./d/group.md) | [Link](../examples/d/group/README.md) |
| azuredevops_user | User | [Link](./d/user.md) | [Link](../examples/d/user/README.md) |
| azuredevops_project | Project | [Link](./d/project.md) | [Link](../examples/d/project/README.md) |
| azuredevops_release_definition | Release Definition(Release Pipeline) | [Link](./d/release_definition.md) | [Link](../examples/d/release_definition/README.md) |
| azuredevops_service_endpoint | Service Endpoint | [Link](./d/service_endpoint.md) | [Link](../examples/d/service_endpoint/README.md) |
| azuredevops_source_repository | Source Repository | [Link](./d/source_repository.md) | [Link](../examples/d/source_repository/README.md) |
| azuredevops_task_group | Task Group | [Link](./d/task_group.md) | [Link](../examples/d/task_group/README.md) |
| azuredevops_variable_group | Variable Group(Single Variable Group) | [Link](./d/variable_group.md) | [Link](../examples/d/variable_group/README.md) |
| azuredevops_variable_groups | Variable Groups(Multiple Variable Group) | [Link](./d/variable_groups.md) | [Link](../examples/d/variable_groups/README.md) |
| azuredevops_workflow_task | Workflow Task(for Pipelines) | [Link](./d/workflow_task.md) | [Link](../examples/d/workflow_task/README.md) |

## Resources

| Name | Description | Docs | Examples |
|------|-------------|------|----------|
| azuredevops_build_definition | Manage Build Definition | [Link](./r/build_definition.md) | [Link](../examples/r/build_definition/README.md) |
| azuredevops_project | Manage Project | [Link](./r/project.md) | [Link](../examples/r/project/README.md) |
| azuredevops_release_definition | Manage Full Release Definition | [Link](./r/release_definition.md) | [Link](../examples/r/release_definition/README.md) |
| azuredevops_release_environment | Manage Partial Release Definition | [Link](./r/release_environment.md) | [Link](../examples/r/release_environment/README.md) |
| azuredevops_release_task | Manage Single Task in Release Pipeline | [Link](./r/release_task.md) | [Link](../examples/r/release_task/README.md) |
| azuredevops_release_tasks | Manage Group of Tasks in Release Pipeline | [Link](./r/release_tasks.md) | [Link](../examples/r/release_tasks/README.md) |
| azuredevops_release_variables | Manage Variables in Release Pipeline | [Link](./r/release_variables.md) | [Link](../examples/r/release_variables/README.md) |
| azuredevops_service_endpoint | Manage Service Endpoint(Github, Kubernetes, etc) | [Link](./r/service_endpoint.md) | [Link](../examples/r/service_endpoint/README.md) |
| azuredevops_service_hook | Manage Service Hook(Slack, etc) | [Link](./r/service_hook.md) | [Link](../examples/r/service_hook/README.md) |
| azuredevops_task_group | Manage Task Group | [Link](./r/task_group.md) | [Link](../examples/r/task_group/README.md) |
| azuredevops_variable_group | Manage Variable Group | [Link](./r/variable_group.md) | [Link](../examples/r/variable_group/README.md) |
| azuredevops_extension | Manage Extension | [Link](./r/extension.md) | [Link](../examples/r/extension/README.md) |
