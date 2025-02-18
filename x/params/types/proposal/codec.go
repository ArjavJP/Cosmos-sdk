package proposal

import (
	"https://github.com/ArjavJP/Cosmos-sdk/codec"
	"https://github.com/ArjavJP/Cosmos-sdk/codec/types"
	govtypes "https://github.com/ArjavJP/Cosmos-sdk/x/gov/types"
)

// RegisterLegacyAminoCodec registers all necessary param module types with a given LegacyAmino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&ParameterChangeProposal{}, "cosmos-sdk/ParameterChangeProposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&ParameterChangeProposal{},
	)
}
