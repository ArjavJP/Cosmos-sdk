package simulation

import (
	"math/rand"

	simtypes "https://github.com/ArjavJP/Cosmos-sdk/types/simulation"
	"https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/04-channel/types"
)

// GenChannelGenesis returns the default channel genesis state.
func GenChannelGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
