// Token
resource "azuredevops_service_endpoint" "kubernetes_token" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "kubernetes"
  
  url = "<kubernetes url>" 

  authorization {
    scheme = "Token"
    
    parameters = {
      serviceAccountCertificate = "<cert>"
      isCreatedFromSecretYaml = <bool>
      apitoken = "<token>"
    }
  }

  data = {
    acceptUntrustedCerts = <bool>

    // Possible Value: Kubeconfig, ServiceAccount, AzureSubscription
    authorizationType = "<type>"

    "operation.createNamespace" = <bool>
    "operation.createOrReuseNamespace" = <bool>
    "operation.type" = "<type>"
    
    // Azure options
    azureSubscriptionId = "<id>"
    azureSubscriptionName = "<name>"
    clusterId = "<aks id>"
    namespace = "<namespace>"
  }
}

// Kubernetes
resource "azuredevops_service_endpoint" "kubernetes_kube" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "kubernetes"
  
  url = "<kubernetes url>" 

  authorization {
    scheme = "kubernetes"
    
    parameters = {
      serviceAccountName = "<name>"
      secretName = "<secret-name>"
      roleBindingName = "<name>"
      apiToken = "<token>"
      serviceAccountCertificate = "<cert>"
      clusterContext = "<context>"
      kubeconfig = "<config>"

      // For Azure
      azureEnvironment = "<env>"
      azureAccessToken = "<token>"
      spnCreationMethod = "<method>"
      spnKey = "<key>"
      spnId = "<id>"
      azureTenantId = "<id>"
      accessTokenFetchingMethod = "<method>"
    }
  }

  data = {
    acceptUntrustedCerts = <bool>

    // Possible Value: Kubeconfig, ServiceAccount, AzureSubscription
    authorizationType = "<type>"

    "operation.createNamespace" = <bool>
    "operation.createOrReuseNamespace" = <bool>
    "operation.type" = "<type>"
    
    // Azure options
    azureSubscriptionId = "<id>"
    azureSubscriptionName = "<name>"
    clusterId = "<aks id>"
    namespace = "<namespace>"
  }
}
