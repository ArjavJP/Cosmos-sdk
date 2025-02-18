package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ArjavJP/Cosmos-sdk/codec"
	codectypes "github.com/ArjavJP/Cosmos-sdk/codec/types"
	"github.com/ArjavJP/Cosmos-sdk/std"
	"github.com/ArjavJP/Cosmos-sdk/testutil/testdata"
	sdk "github.com/ArjavJP/Cosmos-sdk/types"
	"github.com/ArjavJP/Cosmos-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
