// Username and Password 
resource "azuredevops_service_endpoint" "bitbucket" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "bitbucket"

  url = "<bitbucket url>"

  authorization {
    scheme = "UsernamePassword"
    

    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }

  data = {
    displayName = "<display name>"
    avatarUrl = "<avatar url>"
  }
}

// OAuth2
resource "azuredevops_service_endpoint" "bitbucket_oauth2" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "bitbucket"

  url = "<bitbucket url>"

  authorization {
    scheme = "OAuth2"
    
    parameters = {
      accessToken = "<guid>" // This can cause issue
    }
  }

  data = {
    displayName = "<display name>"
    avatarUrl = "<avatar url>"
  }
}
