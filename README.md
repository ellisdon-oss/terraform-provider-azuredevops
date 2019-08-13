# Terraform Provider for AzureDevOps

## Requirements

-    [Terraform](https://www.terraform.io/downloads.html) 0.12.0+
-    [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)
-    [AzureDevOps-Go](https://github.com/ellisdon-oss/azuredevops-go) 

## Building The Provider

```sh
go get github.com/ellisdon-oss/terraform-provider-azuredevops
```

## Currently supported resource/data source


### Data Sources

- azuredevops_agent_queue
- azuredevops_build_definition
- azuredevops_group
- azuredevops_user
- azuredevops_project
- azuredevops_release_definition
- azuredevops_service_endpoint
- azuredevops_source_repository
- azuredevops_task_group
- azuredevops_variable_group
- azuredevops_variable_groups
- azuredevops_workflow_task

### Resources

- azuredevops_build_definition
- azuredevops_project
- azuredevops_release_definition
- azuredevops_release_environment
- azuredevops_service_endpoint
- azuredevops_service_hook
- azuredevops_task_group
- azuredevops_variable_group

## Todo-List

- [ ] Add Docker Image for provider
- [ ] Full Documentation
