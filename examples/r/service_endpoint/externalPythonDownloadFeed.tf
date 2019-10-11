// Token
resource "azuredevops_service_endpoint" "python_download_token" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externalPythonDownloadFeed"
  
  url = "<python feed url>"

  authorization {
    scheme = "Token"
    

    parameters = {
      apitoken = "<token>"
    }
  }
}

// Username & Password
resource "azuredevops_service_endpoint" "python_download_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externalPythonDownloadFeed"
  
  url = "<python feed url>"

  authorization {
    scheme = "UsernamePassword"
    

    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }
}
