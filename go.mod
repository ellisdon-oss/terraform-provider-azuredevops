module github.com/ellisdon/terraform-provider-azuredevops

require (
	github.com/antihax/optional v0.0.0-20180407024304-ca021399b1a6
	github.com/ellisdon/azuredevops-go v0.0.0-20170505043639-c605e284fe17
	github.com/hashicorp/terraform v0.12.6
	github.com/pkg/errors v0.0.0-20170505043639-c605e284fe17
)

replace github.com/ellisdon/azuredevops-go => ../azuredevops-go
