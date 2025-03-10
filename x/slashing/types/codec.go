package types

import (
	"github.com/ArjavJP/Cosmos-sdk/codec"
	"github.com/ArjavJP/Cosmos-sdk/codec/types"
	cryptocodec "github.com/ArjavJP/Cosmos-sdk/crypto/codec"
	sdk "github.com/ArjavJP/Cosmos-sdk/types"
	"github.com/ArjavJP/Cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers concrete types on LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUnjail{}, "cosmos-sdk/MsgUnjail", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnjail{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/slashing module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
	// is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/slashing and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
