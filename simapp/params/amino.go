// +build test_amino

package params

import (
	"github.com/ArjavJP/Cosmos-sdk/codec"
	"github.com/ArjavJP/Cosmos-sdk/codec/types"
	"github.com/ArjavJP/Cosmos-sdk/x/auth/legacy/legacytx"
)

// MakeTestEncodingConfig creates an EncodingConfig for an amino based test configuration.
// This function should be used only internally (in the SDK).
// App user should'nt create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
func MakeTestEncodingConfig() EncodingConfig {
	cdc := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewAminoCodec(cdc)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          legacytx.StdTxConfig{Cdc: cdc},
		Amino:             cdc,
	}
}
