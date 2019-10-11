// Service Principal
resource "azuredevops_service_endpoint" "azurerm_serviceprincipal" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "azurerm"
  
  // Possible values: https://management.core.windows.net/, https://management.core.chinacloudapi.cn/, https://management.core.usgovcloudapi.net/, https://management.core.cloudapi.de/, Empty for Azure Stack
  url = "<azure url>"

  authorization {
    scheme = "ServicePrincipal"
    
    parameters = {
      accessToken = "<token>"
      role = "<role>"
      scope = "<scope>"
      accessTokenFetchingMethod = "<method>"
      servicePrincipalId = "<id>"
      authenticationType = "<type>"
      servicePrincipalKey = "<key>"
      servicePrincipalCertificate = "<cert>"
      tenantId = "<id>"
    }
  }

  data = {
    // Possible values: AzureCloud, AzureChinaCloud, AzureUSGovernment, AzureGermanCloud
    environment = "<env>"

    // Possible values: Subscription, ManagementGroup, MachineLearningWorkspace
    scopeLevel = "<level>"

    subscriptionId = "<id>"
    subscriptionName = "<name>"
    managementGroupId = "<group id>"
    managementGroupName = "<group name>"
    oboAuthorization = <bool>
    creationMode = "<mode>"
    azureSpnRoleAssignmentId = "<id>"
    azureSpnPermissions = "<permissions>"
    spnObjectId = "<id>"
    appObjectId = "<id>"
  }
}

// Certificates
resource "azuredevops_service_endpoint" "azurerm_managed" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "azurerm"
  
  // Possible values: https://management.core.windows.net/, https://management.core.chinacloudapi.cn/, https://management.core.usgovcloudapi.net/, https://management.core.cloudapi.de/, Empty for Azure Stack
  url = "<azure url>"

  authorization {
    scheme = "ManagedServiceIdentity"
    
    parameters = {
      tenantId = "<id>"
    }
  }

  data = {
    // Possible values: AzureCloud, AzureChinaCloud, AzureUSGovernment, AzureGermanCloud
    environment = "<env>"

    // Possible values: Subscription, ManagementGroup, MachineLearningWorkspace
    scopeLevel = "<level>"

    subscriptionId = "<id>"
    subscriptionName = "<name>"
    managementGroupId = "<group id>"
    managementGroupName = "<group name>"
    oboAuthorization = <bool>
    creationMode = "<mode>"
    azureSpnRoleAssignmentId = "<id>"
    azureSpnPermissions = "<permissions>"
    spnObjectId = "<id>"
    appObjectId = "<id>"
  }
}
