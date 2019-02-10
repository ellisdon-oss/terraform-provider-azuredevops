package azuredevops

import (
	"context"
	//"github.com/ellisdon/azuredevops-go/core"
	//"github.com/ellisdon/azuredevops-go/operations"
	"github.com/ellisdon/azuredevops-go"
)

type Config struct {
	Organization string
	Client       *azuredevops.APIClient
	Context      context.Context
	ApiVersion   string
}
