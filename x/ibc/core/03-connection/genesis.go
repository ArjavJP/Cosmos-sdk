package connection

import (
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	"https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/03-connection/keeper"
	"https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/03-connection/types"
)

// InitGenesis initializes the ibc connection submodule's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs types.GenesisState) {
	for _, connection := range gs.Connections {
		conn := types.NewConnectionEnd(connection.State, connection.ClientId, connection.Counterparty, connection.Versions, connection.DelayPeriod)
		k.SetConnection(ctx, connection.Id, conn)
	}
	for _, connPaths := range gs.ClientConnectionPaths {
		k.SetClientConnectionPaths(ctx, connPaths.ClientId, connPaths.Paths)
	}
	k.SetNextConnectionSequence(ctx, gs.NextConnectionSequence)
}

// ExportGenesis returns the ibc connection submodule's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Connections:            k.GetAllConnections(ctx),
		ClientConnectionPaths:  k.GetAllClientConnectionPaths(ctx),
		NextConnectionSequence: k.GetNextConnectionSequence(ctx),
	}
}
