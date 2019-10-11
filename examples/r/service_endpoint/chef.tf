resource "azuredevops_service_endpoint" "chef" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "chef"
  
  url = "<chef url>" 

  authorization {
    scheme = "UsernamePassword"
    

    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }
}
