package ante_test

import (
	"https://github.com/ArjavJP/Cosmos-sdk/codec"
	"https://github.com/ArjavJP/Cosmos-sdk/codec/types"
	"https://github.com/ArjavJP/Cosmos-sdk/testutil/testdata"
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	"https://github.com/ArjavJP/Cosmos-sdk/x/auth/ante"
	"https://github.com/ArjavJP/Cosmos-sdk/x/auth/tx"
)

type setFeeGranter interface {
	SetFeeGranter(feeGranter sdk.AccAddress)
}

func (suite *AnteTestSuite) TestRejectFeeGranter() {
	suite.SetupTest(true) // setup
	txConfig := tx.NewTxConfig(codec.NewProtoCodec(types.NewInterfaceRegistry()), tx.DefaultSignModes)
	txBuilder := txConfig.NewTxBuilder()
	d := ante.NewRejectFeeGranterDecorator()
	antehandler := sdk.ChainAnteDecorators(d)

	_, err := antehandler(suite.ctx, txBuilder.GetTx(), false)
	suite.Require().NoError(err)

	setGranterTx := txBuilder.(setFeeGranter)
	_, _, addr := testdata.KeyTestPubAddr()
	setGranterTx.SetFeeGranter(addr)

	_, err = antehandler(suite.ctx, txBuilder.GetTx(), false)
	suite.Require().Error(err)
}
