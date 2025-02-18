package keeper_test

import (
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"https://github.com/ArjavJP/Cosmos-sdk/simapp"
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	authtypes "https://github.com/ArjavJP/Cosmos-sdk/x/auth/types"
)

// returns context and app with params set on account keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	return app, ctx
}
