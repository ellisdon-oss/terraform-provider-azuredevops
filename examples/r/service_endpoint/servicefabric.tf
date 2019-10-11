// Certificate
resource "azuredevops_service_endpoint" "servicefabric_cert" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "servicefabric"
  
  
  url = "<cluster url>" 

  authorization {
    scheme = "Certificate"

    parameters = {
      servercertthumbprint = "<thumbprint>"
      certificate = "<certificate>"
      certificatepassword = "<password>"
    }
  }
}

// Username & Password
resource "azuredevops_service_endpoint" "servicefabric_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "servicefabric"
  
  
  url = "<cluster url>" 

  authorization {
    scheme = "UsernamePassword"

    parameters = {
      servercertthumbprint = "<thumbprint>"
      username = "<username>"
      password = "<password>"
    }
  }
}

// None
resource "azuredevops_service_endpoint" "servicefabric_none" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "servicefabric"
  
  url = "<cluster url>" 

  authorization {
    scheme = "None"

    parameters = {
      Unsecured = <bool>
      ClusterSpn = "<spn>"
    }
  }
}

