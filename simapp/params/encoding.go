package params

import (
	"https://github.com/ArjavJP/Cosmos-sdk/client"
	"https://github.com/ArjavJP/Cosmos-sdk/codec"
	"https://github.com/ArjavJP/Cosmos-sdk/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Marshaler
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}
