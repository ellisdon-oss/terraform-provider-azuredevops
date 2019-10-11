// Token
resource "azuredevops_service_endpoint" "github_enterprise_token" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "githubenterprise"
  
  url = "<github enterprise url>" 

  authorization {
    scheme = "Token"
    

    parameters = {
      apitoken = "<access token>"
    }
  }

  data = {
    acceptUntrustedCerts = <boolean>
  }
}

// Username & Password
resource "azuredevops_service_endpoint" "github_enterprise_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "githubenterprise"
  
  url = "<github enterprise url>" 

  authorization {
    scheme = "UsernamePassword"
    

    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }

  data = {
    acceptUntrustedCerts = <boolean>
  }
}
