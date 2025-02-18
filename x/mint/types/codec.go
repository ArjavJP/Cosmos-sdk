package types

import (
	"https://github.com/ArjavJP/Cosmos-sdk/codec"
	cryptocodec "https://github.com/ArjavJP/Cosmos-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
