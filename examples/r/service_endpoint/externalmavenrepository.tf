// Token
resource "azuredevops_service_endpoint" "maven_token" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externalmavenrepository"
  
  url = "<maven url>"

  authorization {
    scheme = "Token"
    

    parameters = {
      apitoken = "<maven repo token>"
    }
  }

  data = {
    RepositoryId = "<repo id>"
  }
}

// Username & Password
resource "azuredevops_service_endpoint" "maven_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externalmavenrepository"

  url = "<maven url>"

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
