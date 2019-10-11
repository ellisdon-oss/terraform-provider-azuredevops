// OAuth2 - Not usable due to auth flow
//resource "azuredevops_service_endpoint" "github_oauth2" {
//  project_id = "<project id>"
//
//  name = "<endpoint name>"
//
//  owner = "Library"
//  type = "github"
//  
//
//  url = "<github url>"
//
//  authorization {
//    scheme = "OAuth"
//    
//
//    parameters = {
//      accessToken = "<guid>" // unstable, do not use
//    }
//  }
//}

// Installation Token(Github App) - Not usable due to auth flow
//resource "azuredevops_service_endpoint" "github_installation_token" {
//  project_id = "<project id>"
//
//  name = "<endpoint name>"
//
//  owner = "Library"
//  type = "github"
//  
//
//  url = "<github url>"
//
//  authorization {
//    scheme = "InstallationToken"
//    
//    parameters = {
//      token = "<token>"
//    }
//  }
//}

// Token
resource "azuredevops_service_endpoint" "github_token" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "github"
  
  url = "<github url>" 

  authorization {
    scheme = "PersonalAccessToken"
    

    parameters = {
      accessToken = "<access token>"
    }
  }
}
