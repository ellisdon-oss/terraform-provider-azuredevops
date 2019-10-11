resource "azuredevops_service_endpoint" "azureservicebus" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "AzureServiceBus"
  
  url = "<endpoint url>" 

  authorization {
    scheme = "None"

    parameters = {
      serviceBusConnectionString = "<connection string>"
    }
  }

  data = {
    serviceBusQueueName = "<name>"
  }
}
