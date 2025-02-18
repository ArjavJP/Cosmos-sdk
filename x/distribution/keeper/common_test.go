package keeper_test

import (
	"https://github.com/ArjavJP/Cosmos-sdk/simapp"
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	authtypes "https://github.com/ArjavJP/Cosmos-sdk/x/auth/types"
	"https://github.com/ArjavJP/Cosmos-sdk/x/distribution/types"
)

var (
	PKS = simapp.CreateTestPubKeys(5)

	valConsPk1 = PKS[0]
	valConsPk2 = PKS[1]
	valConsPk3 = PKS[2]

	valConsAddr1 = sdk.ConsAddress(valConsPk1.Address())
	valConsAddr2 = sdk.ConsAddress(valConsPk2.Address())

	distrAcc = authtypes.NewEmptyModuleAccount(types.ModuleName)
)
