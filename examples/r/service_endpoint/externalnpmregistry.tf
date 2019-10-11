// Token
resource "azuredevops_service_endpoint" "npm_token" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externalnpmregistry"
  
  url = "<npm url>"

  authorization {
    scheme = "Token"
    

    parameters = {
      apitoken = "<npm repo token>"
    }
  }

  data = {
    RepositoryId = "<repo id>"
  }
}

// Username & Password
resource "azuredevops_service_endpoint" "npm_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externalnpmregistry"

  url = "<npm url>"

  authorization {
    scheme = "UsernamePassword"
    

    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }

  data = {
    RepositoryId = "<repo id>"
  }
}
