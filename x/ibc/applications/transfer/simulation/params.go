package simulation

import (
	"math/rand"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/ArjavJP/Cosmos-sdk/x/simulation"

	simtypes "github.com/ArjavJP/Cosmos-sdk/types/simulation"
	"github.com/ArjavJP/Cosmos-sdk/x/ibc/applications/transfer/types"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeySendEnabled),
			func(r *rand.Rand) string {
				sendEnabled := RadomEnabled(r)
				return string(types.ModuleCdc.MustMarshalJSON(&gogotypes.BoolValue{Value: sendEnabled}))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyReceiveEnabled),
			func(r *rand.Rand) string {
				receiveEnabled := RadomEnabled(r)
				return string(types.ModuleCdc.MustMarshalJSON(&gogotypes.BoolValue{Value: receiveEnabled}))
			},
		),
	}
}
