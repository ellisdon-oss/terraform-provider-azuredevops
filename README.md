# Terraform Provider for AzureDevOps

## Requirements

-    [Terraform](https://www.terraform.io/downloads.html) 0.12.0+
-    [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)
-    [Microsoft's AzureDevOps Go SDK](https://github.com/microsoft/azure-devops-go-api) 

## Building The Provider

```sh
go get github.com/ellisdon-oss/terraform-provider-azuredevops
```

## Install The Provider

1. Build the provider or grab the provider binary from the [release page](https://github.com/EllisDon-Aegean/terraform-provider-azuredevops/releases)
2. Extra and copy the provider binary to the terraform global folder(Mac/Linux `~/.terraform.d/plugins` or Windows `%APPDATA%\terraform.d\plugins`)

## [Docs](./docs)

## Contributing

1. Fork the repo
2. Make the changes 
3. Preferably before PR, the commits should be rebase and squash to 1 commit
4. create a new PR (new features should be in a new features/<name> branch and patches should be in patches/<name> branch)

Note: please add yourself into CONTRIBUTORS.md before submitting PR

## Todo-List

- [ ] Add Docker Image for provider
- [ ] Full Documentation
