resource "azuredevops_service_endpoint" "sonarqube" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "sonarqube"
  
  url = "<url>" 

  authorization {
    scheme = "UsernamePassword"
    
    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }
}
