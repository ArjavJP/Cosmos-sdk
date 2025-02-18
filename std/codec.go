package std

import (
	"https://github.com/ArjavJP/Cosmos-sdk/codec"
	"https://github.com/ArjavJP/Cosmos-sdk/codec/types"
	cryptocodec "https://github.com/ArjavJP/Cosmos-sdk/crypto/codec"
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	txtypes "https://github.com/ArjavJP/Cosmos-sdk/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
