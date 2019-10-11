// Token
resource "azuredevops_service_endpoint" "externaltfs_token" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externaltfs"

  url = "<tfs url>"

  authorization {
    scheme = "Token"
    

    parameters = {
      apitoken = "<tfs token>"
    }
  }
}

// Username & Password
resource "azuredevops_service_endpoint" "externaltfs_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externaltfs"

  url = "<tfs url>"

  authorization {
    scheme = "UsernamePassword"
    

    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }
}
