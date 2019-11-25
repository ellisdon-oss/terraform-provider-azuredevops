# Terraform Provider for Azure DevOps

Table of Contents
=================

   * [Terraform Provider for Azure DevOps](#terraform-provider-for-azure-devops)
      * [About EllisDon-OSS](#about-ellisdon-oss)
      * [Features](#features)
      * [Requirements](#requirements)
      * [Building The Provider](#building-the-provider)
         * [Requirements](#requirements-1)
         * [Build](#build)
      * [Installing The Provider](#install-the-provider)
      * [Examples](#examples)
         * [Creating a Project](#creating-a-project)
         * [Creating Service Endpoint (Kubernetes)](#creating-service-endpoint-kubernetes)
         * [Creating Service Endpoint (GitHub)](#creating-service-endpoint-github)
      * [<a href="./guides">Guides</a>](#guides)
      * [<a href="./docs">Docs</a>](#docs)
      * [Contributing](#contributing)
      * [Todo-List](#todo-list)
      * [License](#license)

## Features
<details>
<summary>Features List</summary>

- Query Users and Groups
- Manage the lifecycle of Variable Group
- Manage the lifecycle of Task Group
- Manage the lifecycle of Service Endpoint of any type
- Manage the lifecycle of Service hook of any type
- Manage the lifecycle of Release Pipeline
- Manage the partial section of Release Pipeline
- Manage the lifecycle of Build Pipeline for both YAML-based and direct build tasks
- Manage the lifecycle of Project
- Inject Single Tasks, or Group of Tasks into Release Pipeline
- All Resources support import
- Full Documentation

</details>

## Requirements

-    [Terraform](https://www.terraform.io/downloads.html) 0.12.0+
-    [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)
-    [Microsoft's Azure DevOps Go SDK](https://github.com/microsoft/azure-devops-go-api) 

## Building The Provider

### Requirements

- You will need Go with Go modules enabled (`GO111MODULE=on`)

### Build

First build by running the following command
```sh
go get github.com/ellisdon-oss/terraform-provider-azuredevops
```
then copy the `terraform-provider-azuredevops` binary to your terraform plugin folder or your terraform project folder (plugin folder can be either `$HOME/.terraform.d/plugins/` or `%APPDATA%\terraform.d\plugins`)

## Install The Provider

1. Build the provider or grab the provider binary from the [release page](https://github.com/EllisDon-Aegean/terraform-provider-azuredevops/releases)
2. Extract the provider binary to the terraform global folder (Mac/Linux `~/.terraform.d/plugins` or Windows `%APPDATA%\terraform.d\plugins`)

## Examples

**NOTE:** Full examples for the following can be found in [examples](./examples)

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

### Creating Service Endpoint (Kubernetes)

<details>
<summary>Code Example</summary>

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

</details>


### Creating Service Endpoint (GitHub)
<details>
<summary>Code Example</summary>

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

</details>

**More examples can be found [here](./examples)**

## [Guides](./guides)

## [Docs](./docs)

## Contributing

1. Fork the repo
2. Make the changes
3. Please add yourself into CONTRIBUTORS.md before creating a PR
4. Please rebase and squash the commits down to 1 (preferably before creating a PR)
5. Create a new PR, new features should be in a new `features/<name>` branch and patches should be in a `patches/<name>` branch 

## Todo-List

- [ ] Add Docker Image for provider
- [ ] Full Documentation

## License

[Mozilla Public License v2.0](./LICENSE.md)
