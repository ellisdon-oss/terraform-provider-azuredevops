// None(for push only)
resource "azuredevops_service_endpoint" "nuget_none" {
  
  project_id = "<project id>"

  name = "<endpoint name>"
  owner = "Library"
  type = "externalnugetfeed"
  
  url = "<nuget url>"

  authorization {
    scheme = "None"
    

    parameters = {
      nugetkey = "<api key>"
    }
  }
}

// Token
resource "azuredevops_service_endpoint" "nuget_token" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externalnugetfeed"
  
  url = "<nuget url>"

  authorization {
    scheme = "Token"
    

    parameters = {
      apitoken = "<api token>"
    }
  }
}

// Username & Password
resource "azuredevops_service_endpoint" "nuget_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "externalnugetfeed"
  
  url = "<nuget url>"

  authorization {
    scheme = "UsernamePassword"
    

    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }
}
