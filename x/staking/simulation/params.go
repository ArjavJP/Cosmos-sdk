package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"https://github.com/ArjavJP/Cosmos-sdk/x/simulation"

	simtypes "https://github.com/ArjavJP/Cosmos-sdk/types/simulation"
	"https://github.com/ArjavJP/Cosmos-sdk/x/staking/types"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMaxValidators),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%d", GenMaxValidators(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyUnbondingTime),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%d\"", GenUnbondingTime(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyHistoricalEntries),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%d", GetHistEntries(r))
			},
		),
	}
}
