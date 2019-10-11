resource "azuredevops_service_endpoint" "dockerhost" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "dockerhost"
  
  url = "<docker host url>" 

  authorization {
    scheme = "Certificate"
    
    parameters = {
      cacert = "<cacert>"
      cert = "<cert>"
      key = "<key>"
      certificate = "<certificate>"
    }
  }
}
