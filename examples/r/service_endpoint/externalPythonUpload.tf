// Token
resource "azuredevops_service_endpoint" "python_upload_token" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externalPythonUploadFeed"
  
  url = "<python feed url>"

  authorization {
    scheme = "Token"
    
    parameters = {
      apitoken = "<token>"
    }
  }

  data = {
    EndpointName = "<name>"
  }
}

// Username & Password
resource "azuredevops_service_endpoint" "python_upload_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externalPythonUploadFeed"
  
  url = "<python feed url>"

  authorization {
    scheme = "UsernamePassword"
    
    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }

  data = {
    EndpointName = "<name>"
  }
}
