package simulation

import (
	"math/rand"

	simappparams "https://github.com/ArjavJP/Cosmos-sdk/simapp/params"
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	simtypes "https://github.com/ArjavJP/Cosmos-sdk/types/simulation"
	"https://github.com/ArjavJP/Cosmos-sdk/x/gov/types"
	"https://github.com/ArjavJP/Cosmos-sdk/x/simulation"
)

// OpWeightSubmitTextProposal app params key for text proposal
const OpWeightSubmitTextProposal = "op_weight_submit_text_proposal"

// ProposalContents defines the module weighted proposals' contents
func ProposalContents() []simtypes.WeightedProposalContent {
	return []simtypes.WeightedProposalContent{
		simulation.NewWeightedProposalContent(
			OpWeightMsgDeposit,
			simappparams.DefaultWeightTextProposal,
			SimulateTextProposalContent,
		),
	}
}

// SimulateTextProposalContent returns a random text proposal content.
func SimulateTextProposalContent(r *rand.Rand, _ sdk.Context, _ []simtypes.Account) simtypes.Content {
	return types.NewTextProposal(
		simtypes.RandStringOfLength(r, 140),
		simtypes.RandStringOfLength(r, 5000),
	)
}
