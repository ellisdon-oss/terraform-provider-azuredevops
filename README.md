# Terraform Provider for AzureDevOps

Table of Contents
=================

   * [Terraform Provider for AzureDevOps](#terraform-provider-for-azuredevops)
      * [About EllisDon-OSS](#about-ellisdon-oss)
      * [Requirements](#requirements)
      * [Building The Provider](#building-the-provider)
         * [Requirements](#requirements-1)
         * [Build](#build)
      * [Install The Provider](#install-the-provider)
      * [Examples](#examples)
         * [Creating a Project](#creating-a-project)
         * [Creating Service Endpoint (GitHub)](#creating-service-endpoint-github)
         * [Creating Service Endpoint (Kubernetes)](#creating-service-endpoint-kubernetes)
      * [<a href="./docs">Docs</a>](#docs)
      * [Contributing](#contributing)
      * [Todo-List](#todo-list)
      * [License](#license)

## About EllisDon-OSS

Placeholder

## Requirements

-    [Terraform](https://www.terraform.io/downloads.html) 0.12.0+
-    [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)
-    [Microsoft's AzureDevOps Go SDK](https://github.com/microsoft/azure-devops-go-api) 

## Building The Provider

### Requirements

- You will need Go with GO111MODULE enabled

### Build

First build by running the following command
```sh
go get github.com/ellisdon-oss/terraform-provider-azuredevops
```
then copy the `terraform-provider-azuredevops` binary to your terraform plugin folder or your terraform project folder(plugin folder can be either `$HOME/.terraform.d/plugins/` or `%APPDATA%\terraform.d\plugins`)

## Install The Provider

1. Build the provider or grab the provider binary from the [release page](https://github.com/EllisDon-Aegean/terraform-provider-azuredevops/releases)
2. Extra and copy the provider binary to the terraform global folder(Mac/Linux `~/.terraform.d/plugins` or Windows `%APPDATA%\terraform.d\plugins`)

## Examples

**NOTE:** All examples have their full version in [examples](./examples) folder

### Creating a Project

```terraform
resource "azuredevops_project" "default" {
  name       = "a-azuredevops-project"
  visibility = "private"
  capabilities {
    version_control {
      source_control_type = "git"
    }
    process_template {
      template_type_id = "adcc42ab-9882-485e-a3ed-7678f01f66bc"
    }
  }
}
```

ref: [Project](https://docs.microsoft.com/en-us/azure/devops/organizations/projects/create-project?view=azure-devops)

### Creating Service Endpoint (GitHub)

```terraform
...

resource "azuredevops_service_endpoint" "github" {
  name       = "github-example"
  owner      = "Library"
  project_id = azuredevops_project.default.id
  type       = "github"
  url        = "http://github.com"

  # To enable all Pipeline to use this service endpoint
  allow_all_pipelines = true

  authorization {
    scheme = "PersonalAccessToken"
    parameters = {
      accessToken = "<github-token>"
    }
  }

  data = {}
}
```

ref: [Service Endpoint](https://docs.microsoft.com/en-us/azure/devops/extend/develop/service-endpoints?view=azure-devops)

### Creating Service Endpoint (Kubernetes)

```terraform
...

resource "azuredevops_service_endpoint" "kubernetes" {
  name       = "kubernetes-example"
  owner      = "Library"
  project_id = azuredevops_project.default.id
  type       = "kubernetes"
  url        = "http://<kube-cluster-url>"

  authorization {
    parameters = {
      apiToken                  = "<kube-token>"
      serviceAccountCertificate = "<kube-cert>"
    }

    scheme = "Token"
  }

  data = {
    acceptUntrustedCerts = "true"
    authorizationType    = "ServiceAccount"
  }
}
```

ref: [Service Endpoint](https://docs.microsoft.com/en-us/azure/devops/extend/develop/service-endpoints?view=azure-devops)

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

## License

[Mozilla Public License v2.0](./LICENSE.md)
