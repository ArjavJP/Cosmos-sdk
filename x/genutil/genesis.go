package genutil

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ArjavJP/Cosmos-sdk/client"
	sdk "github.com/ArjavJP/Cosmos-sdk/types"
	"github.com/ArjavJP/Cosmos-sdk/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(
	ctx sdk.Context, stakingKeeper types.StakingKeeper,
	deliverTx deliverTxfn, genesisState types.GenesisState,
	txEncodingConfig client.TxEncodingConfig,
) (validators []abci.ValidatorUpdate, err error) {
	if len(genesisState.GenTxs) > 0 {
		validators, err = DeliverGenTxs(ctx, genesisState.GenTxs, stakingKeeper, deliverTx, txEncodingConfig)
	}
	return
}
