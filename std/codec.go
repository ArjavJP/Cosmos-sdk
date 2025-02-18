package std

import (
	"github.com/ArjavJP/Cosmos-sdk/codec"
	"github.com/ArjavJP/Cosmos-sdk/codec/types"
	cryptocodec "github.com/ArjavJP/Cosmos-sdk/crypto/codec"
	sdk "github.com/ArjavJP/Cosmos-sdk/types"
	txtypes "github.com/ArjavJP/Cosmos-sdk/types/tx"
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
