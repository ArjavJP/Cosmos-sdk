package simulation

import (
	"math/rand"

	simtypes "https://github.com/ArjavJP/Cosmos-sdk/types/simulation"
	"https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/03-connection/types"
)

// GenConnectionGenesis returns the default connection genesis state.
func GenConnectionGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
