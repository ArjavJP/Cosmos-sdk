package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"https://github.com/ArjavJP/Cosmos-sdk/codec"
	cryptoAmino "https://github.com/ArjavJP/Cosmos-sdk/crypto/codec"
	"https://github.com/ArjavJP/Cosmos-sdk/testutil/testdata"
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	"https://github.com/ArjavJP/Cosmos-sdk/x/auth/legacy/legacytx"
	"https://github.com/ArjavJP/Cosmos-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
