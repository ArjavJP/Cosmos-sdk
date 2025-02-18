package client

import (
	"github.com/ArjavJP/Cosmos-sdk/x/distribution/client/cli"
	"github.com/ArjavJP/Cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/ArjavJP/Cosmos-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
