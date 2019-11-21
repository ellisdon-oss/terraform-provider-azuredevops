module github.com/ellisdon-oss/terraform-provider-azuredevops

require (
	github.com/apparentlymart/go-dump v0.0.0-20190214190832-042adf3cf4a0 // indirect
	github.com/aws/aws-sdk-go v1.22.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/microsoft/azure-devops-go-api/azuredevops v0.0.0-20190912142452-3207b4a469d3

	github.com/pkg/errors v0.8.0
	github.com/vmihailenco/msgpack v4.0.1+incompatible // indirect
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa // indirect
)

replace github.com/microsoft/azure-devops-go-api/azuredevops => github.com/ellisdon-oss/azure-devops-go-api/azuredevops v0.0.0-20191120143450-7c0fc65db71c

go 1.13
