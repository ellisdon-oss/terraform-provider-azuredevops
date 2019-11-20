module github.com/ellisdon-oss/terraform-provider-azuredevops

require (
	github.com/antihax/optional v0.0.0-20180407024304-ca021399b1a6
	github.com/google/uuid v1.1.1
	github.com/hashicorp/terraform v0.12.9
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/hashicorp/tf-sdk-migrator v1.0.0 // indirect
	github.com/microsoft/azure-devops-go-api v0.0.0-20190912142452-3207b4a469d3
	github.com/microsoft/azure-devops-go-api/azuredevops v0.0.0-20190912142452-3207b4a469d3
	github.com/pkg/errors v0.8.0
)

replace github.com/microsoft/azure-devops-go-api/azuredevops => github.com/ellisdon-oss/azure-devops-go-api/azuredevops v0.0.0-20191120143450-7c0fc65db71c

replace github.com/microsoft/azure-devops-go-api => github.com/ellisdon-oss/azure-devops-go-api v0.0.0-20191120143450-7c0fc65db71c

go 1.13
