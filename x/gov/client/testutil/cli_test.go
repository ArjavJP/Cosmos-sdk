// +build norace

package testutil

import (
	"testing"
	"time"

	"https://github.com/ArjavJP/Cosmos-sdk/testutil/network"
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	"https://github.com/ArjavJP/Cosmos-sdk/x/gov/types"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	genesisState := types.DefaultGenesisState()
	genesisState.DepositParams = types.NewDepositParams(sdk.NewCoins(sdk.NewCoin(cfg.BondDenom, types.DefaultMinDepositTokens)), time.Duration(15)*time.Second)
	genesisState.VotingParams = types.NewVotingParams(time.Duration(5) * time.Second)
	bz, err := cfg.Codec.MarshalJSON(genesisState)
	require.NoError(t, err)
	cfg.GenesisState["gov"] = bz
	suite.Run(t, NewDepositTestSuite(cfg))
}
