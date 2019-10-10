resource "azuredevops_service_endpoint" "fossa" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "fossa"

  url = "<fossa url>"

  authorization {
    scheme = "Token"

    parameters = {
      apitoken = "<fossa token>"
    }
  }
}
