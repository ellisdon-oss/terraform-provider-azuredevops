// Username & Password
resource "azuredevops_service_endpoint" "azureclassic_basic" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "azure"
  
  // Possible values: https://management.core.windows.net/, https://management.core.chinacloudapi.cn/, https://management.core.usgovcloudapi.net/, https://management.core.cloudapi.de/
  url = "<azure url>"

  authorization {
    scheme = "UsernamePassword"
    
    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }

  data = {
    // Possible values: AzureCloud, AzureChinaCloud, AzureUSGovernment, AzureGermanCloud
    environment = "<env>"
    subscriptionId = "<id>"
    subscriptionName = "<name>"
  }
}

// Certificates
resource "azuredevops_service_endpoint" "azureclassic_cert" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "azure"
  
  // Possible values: https://management.core.windows.net/, https://management.core.chinacloudapi.cn/, https://management.core.usgovcloudapi.net/, https://management.core.cloudapi.de/
  url = "<azure url>"

  authorization {
    scheme = "Certificate"
    
    parameters = {
      certificate = "<cert>"
    }
  }

  data = {
    // Possible values: AzureCloud, AzureChinaCloud, AzureUSGovernment, AzureGermanCloud
    environment = "<env>"
    subscriptionId = "<id>"
    subscriptionName = "<name>"
  }
}
