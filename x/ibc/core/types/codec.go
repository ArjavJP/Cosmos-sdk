package types

import (
	codectypes "https://github.com/ArjavJP/Cosmos-sdk/codec/types"
	clienttypes "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/02-client/types"
	connectiontypes "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/03-connection/types"
	channeltypes "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/04-channel/types"
	commitmenttypes "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/23-commitment/types"
	solomachinetypes "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/light-clients/06-solomachine/types"
	ibctmtypes "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/light-clients/07-tendermint/types"
	localhosttypes "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/light-clients/09-localhost/types"
)

// RegisterInterfaces registers x/ibc interfaces into protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	clienttypes.RegisterInterfaces(registry)
	connectiontypes.RegisterInterfaces(registry)
	channeltypes.RegisterInterfaces(registry)
	solomachinetypes.RegisterInterfaces(registry)
	ibctmtypes.RegisterInterfaces(registry)
	localhosttypes.RegisterInterfaces(registry)
	commitmenttypes.RegisterInterfaces(registry)
}
