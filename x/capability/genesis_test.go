package capability_test

import (
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/ArjavJP/Cosmos-sdk/simapp"
	sdk "github.com/ArjavJP/Cosmos-sdk/types"
	banktypes "github.com/ArjavJP/Cosmos-sdk/x/bank/types"
	"github.com/ArjavJP/Cosmos-sdk/x/capability"
	"github.com/ArjavJP/Cosmos-sdk/x/capability/keeper"
	"github.com/ArjavJP/Cosmos-sdk/x/capability/types"
	stakingtypes "github.com/ArjavJP/Cosmos-sdk/x/staking/types"
)

func (suite *CapabilityTestSuite) TestGenesis() {
	sk1 := suite.keeper.ScopeToModule(banktypes.ModuleName)
	sk2 := suite.keeper.ScopeToModule(stakingtypes.ModuleName)

	cap1, err := sk1.NewCapability(suite.ctx, "transfer")
	suite.Require().NoError(err)
	suite.Require().NotNil(cap1)

	err = sk2.ClaimCapability(suite.ctx, cap1, "transfer")
	suite.Require().NoError(err)

	cap2, err := sk2.NewCapability(suite.ctx, "ica")
	suite.Require().NoError(err)
	suite.Require().NotNil(cap2)

	genState := capability.ExportGenesis(suite.ctx, *suite.keeper)

	// create new app that does not share persistent or in-memory state
	// and initialize app from exported genesis state above.
	db := dbm.NewMemDB()
	encCdc := simapp.MakeTestEncodingConfig()
	newApp := simapp.NewSimApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, simapp.DefaultNodeHome, 5, encCdc, simapp.EmptyAppOptions{})

	newKeeper := keeper.NewKeeper(suite.cdc, newApp.GetKey(types.StoreKey), newApp.GetMemKey(types.MemStoreKey))
	newSk1 := newKeeper.ScopeToModule(banktypes.ModuleName)
	newSk2 := newKeeper.ScopeToModule(stakingtypes.ModuleName)
	deliverCtx, _ := newApp.BaseApp.NewUncachedContext(false, tmproto.Header{}).WithBlockGasMeter(sdk.NewInfiniteGasMeter()).CacheContext()

	capability.InitGenesis(deliverCtx, *newKeeper, *genState)

	// check that all previous capabilities exist in new app after InitGenesis
	sk1Cap1, ok := newSk1.GetCapability(deliverCtx, "transfer")
	suite.Require().True(ok, "could not get first capability after genesis on first ScopedKeeper")
	suite.Require().Equal(*cap1, *sk1Cap1, "capability values not equal on first ScopedKeeper")

	sk2Cap1, ok := newSk2.GetCapability(deliverCtx, "transfer")
	suite.Require().True(ok, "could not get first capability after genesis on first ScopedKeeper")
	suite.Require().Equal(*cap1, *sk2Cap1, "capability values not equal on first ScopedKeeper")

	sk2Cap2, ok := newSk2.GetCapability(deliverCtx, "ica")
	suite.Require().True(ok, "could not get second capability after genesis on second ScopedKeeper")
	suite.Require().Equal(*cap2, *sk2Cap2, "capability values not equal on second ScopedKeeper")
}
