package azuredevops

import (
	"context"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
)

type Config struct {
	Organization string
	Connection   *azuredevops.Connection
	Context      context.Context
}
