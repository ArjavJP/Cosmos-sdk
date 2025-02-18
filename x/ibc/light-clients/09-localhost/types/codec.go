package types

import (
	codectypes "https://github.com/ArjavJP/Cosmos-sdk/codec/types"
	"https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/exported"
)

// RegisterInterfaces register the ibc interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
}
