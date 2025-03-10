package simulation

import (
	"math/rand"

	simtypes "github.com/ArjavJP/Cosmos-sdk/types/simulation"
	"github.com/ArjavJP/Cosmos-sdk/x/ibc/core/02-client/types"
)

// GenClientGenesis returns the default client genesis state.
func GenClientGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
