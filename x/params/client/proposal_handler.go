package client

import (
	govclient "github.com/ArjavJP/Cosmos-sdk/x/gov/client"
	"github.com/ArjavJP/Cosmos-sdk/x/params/client/cli"
	"github.com/ArjavJP/Cosmos-sdk/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd, rest.ProposalRESTHandler)
