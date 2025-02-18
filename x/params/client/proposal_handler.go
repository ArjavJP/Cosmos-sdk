package client

import (
	govclient "https://github.com/ArjavJP/Cosmos-sdk/x/gov/client"
	"https://github.com/ArjavJP/Cosmos-sdk/x/params/client/cli"
	"https://github.com/ArjavJP/Cosmos-sdk/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd, rest.ProposalRESTHandler)
