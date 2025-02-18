package client

import (
	govclient "https://github.com/ArjavJP/Cosmos-sdk/x/gov/client"
	"https://github.com/ArjavJP/Cosmos-sdk/x/upgrade/client/cli"
	"https://github.com/ArjavJP/Cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal, rest.ProposalCancelRESTHandler)
