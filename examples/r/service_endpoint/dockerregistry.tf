// Username & Password
resource "azuredevops_service_endpoint" "dockerregistry_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "dockerregistry"
  
  url = "<registry url>" 

  authorization {
    scheme = "UsernamePassword"
    
    parameters = {
      registry = "<registry>"
      username = "<username>"
      password = "<password>"
      email = "<email>"
    }
  }

  data = {
    // Possible Type: ACR, DockerHub, Others
    registrytype = "<type>"

    spnObjectId = "<id>"
    appObjectId = "<id>"
    azureSpnRoleAssignmentId = "<id>"
    azureSpnPermissions = "<permissions>"


    // ACR options
    subscriptionId = "<id>"
    subscriptionName = "<name>"
    registryId = "<id>"

  }
}

// ServicePrincipal
resource "azuredevops_service_endpoint" "dockerregistry_serviceprincipal" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "github"
  
  url = "<github url>" 

  authorization {
    scheme = "ServicePrincipal"

    parameters = {
      accessToken = "<access token>"
      scope = "<scope>"
      loginServer = "<server>"
      role = "<role>"
      accessTokenFetchingMethod = "<method>"
      servicePrincipalId = "<id>"
      authenticationType = "<type>"
      servicePrincipalKey = "<key>"
      servicePrincipalCertificate = "<cert>"
      tenantId = "<guid>"
    }
  }

  data = {
    // Possible Type: ACR, DockerHub, Others
    registrytype = "<type>"

    spnObjectId = "<id>"
    appObjectId = "<id>"
    azureSpnRoleAssignmentId = "<id>"
    azureSpnPermissions = "<permissions>"


    // ACR options
    subscriptionId = "<id>"
    subscriptionName = "<name>"
    registryId = "<id>"

  }
}
