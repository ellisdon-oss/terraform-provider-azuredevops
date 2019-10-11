resource "azuredevops_service_endpoint" "jenkins" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "jenkins"
  
  url = "<jenkins url>" 

  authorization {
    scheme = "UsernamePassword"
    
    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }

  data = {
    acceptUntrustedCerts = <bool>
  }
}
