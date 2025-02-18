package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"https://github.com/ArjavJP/Cosmos-sdk/codec"
	codectypes "https://github.com/ArjavJP/Cosmos-sdk/codec/types"
	"https://github.com/ArjavJP/Cosmos-sdk/std"
	"https://github.com/ArjavJP/Cosmos-sdk/testutil/testdata"
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	"https://github.com/ArjavJP/Cosmos-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
