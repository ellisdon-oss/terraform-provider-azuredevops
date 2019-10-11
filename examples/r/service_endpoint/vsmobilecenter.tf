resource "azuredevops_service_endpoint" "vsmobilecenter" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "vsmobilecenter"
  
  // Default: https://api.appcenter.ms/v0.1
  url = "<url>" 

  authorization {
    scheme = "Token"
    
    parameters = {
      apitoken = "<token>"
    }
  }
}
