package client

import (
	govclient "github.com/ArjavJP/Cosmos-sdk/x/gov/client"
	"github.com/ArjavJP/Cosmos-sdk/x/upgrade/client/cli"
	"github.com/ArjavJP/Cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal, rest.ProposalCancelRESTHandler)
