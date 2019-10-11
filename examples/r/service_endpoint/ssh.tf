// Username & Password
resource "azuredevops_service_endpoint" "ssh_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "ssh"
  
  url = "<server url>"

  authorization {
    scheme = "UsernamePassword"
    
    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }

  data = {
    Host = "<host>"
    Port = "<port>"
    PrivateKey = "<privateKey>"
  }
}
