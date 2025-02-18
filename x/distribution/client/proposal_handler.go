package client

import (
	"https://github.com/ArjavJP/Cosmos-sdk/x/distribution/client/cli"
	"https://github.com/ArjavJP/Cosmos-sdk/x/distribution/client/rest"
	govclient "https://github.com/ArjavJP/Cosmos-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
