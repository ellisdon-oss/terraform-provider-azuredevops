resource "azuredevops_service_endpoint" "artifactory" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "artifactoryService"

  url = "<artifactory url>"

  authorization {
    scheme = "UsernamePassword"
    

    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }
}
