module github.com/ellisdon/terraform-provider-azuredevops

require (
	github.com/ellisdon/azuredevops-go v0.0.0-20170505043639-c605e284fe17 // indirect
	github.com/hashicorp/terraform v0.12.9
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/hashicorp/tf-sdk-migrator v1.0.0 // indirect
	github.com/microsoft/azure-devops-go-api v0.0.0-20190912142452-3207b4a469d3 // indirect
	github.com/microsoft/azure-devops-go-api/azuredevops v0.0.0-20190912142452-3207b4a469d3
	github.com/pkg/errors v0.8.0
)

replace github.com/ellisdon/azuredevops-go => ../azuredevops-go
