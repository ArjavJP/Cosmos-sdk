// Package tx 's xauthclient.go file is copy-pasted from
// https://github.com/ArjavJP/Cosmos-sdk/blob/v0.41.3/x/auth/client/query.go
// It is duplicated as to not introduce any breaking change in 0.41.4, see PR:
// https://github.com/ArjavJP/Cosmos-sdk/pull/8732#discussion_r584746947
package tx

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/ArjavJP/Cosmos-sdk/client"
	codectypes "github.com/ArjavJP/Cosmos-sdk/codec/types"
	sdk "github.com/ArjavJP/Cosmos-sdk/types"
)

// QueryTxsByEvents performs a search for transactions for a given set of events
// via the Tendermint RPC. An event takes the form of:
// "{eventAttribute}.{attributeKey} = '{attributeValue}'". Each event is
// concatenated with an 'AND' operand. It returns a slice of Info object
// containing txs and metadata. An error is returned if the query fails.
// If an empty string is provided it will order txs by asc
func queryTxsByEvents(goCtx context.Context, clientCtx client.Context, events []string, page, limit int, orderBy string) (*sdk.SearchTxsResult, error) {
	if len(events) == 0 {
		return nil, errors.New("must declare at least one event to search")
	}

	if page <= 0 {
		return nil, errors.New("page must greater than 0")
	}

	if limit <= 0 {
		return nil, errors.New("limit must greater than 0")
	}

	// XXX: implement ANY
	query := strings.Join(events, " AND ")

	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	// TODO: this may not always need to be proven
	// https://github.com/ArjavJP/Cosmos-sdk/issues/6807
	resTxs, err := node.TxSearch(goCtx, query, true, &page, &limit, orderBy)
	if err != nil {
		return nil, err
	}

	resBlocks, err := getBlocksForTxResults(goCtx, clientCtx, resTxs.Txs)
	if err != nil {
		return nil, err
	}

	txs, err := formatTxResults(clientCtx.TxConfig, resTxs.Txs, resBlocks)
	if err != nil {
		return nil, err
	}

	result := sdk.NewSearchTxsResult(uint64(resTxs.TotalCount), uint64(len(txs)), uint64(page), uint64(limit), txs)

	return result, nil
}

// QueryTx queries for a single transaction by a hash string in hex format. An
// error is returned if the transaction does not exist or cannot be queried.
func queryTx(goCtx context.Context, clientCtx client.Context, hashHexStr string) (*sdk.TxResponse, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return nil, err
	}

	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	//TODO: this may not always need to be proven
	// https://github.com/ArjavJP/Cosmos-sdk/issues/6807
	resTx, err := node.Tx(goCtx, hash, true)
	if err != nil {
		return nil, err
	}

	resBlocks, err := getBlocksForTxResults(goCtx, clientCtx, []*ctypes.ResultTx{resTx})
	if err != nil {
		return nil, err
	}

	out, err := mkTxResult(clientCtx.TxConfig, resTx, resBlocks[resTx.Height])
	if err != nil {
		return out, err
	}

	return out, nil
}

// formatTxResults parses the indexed txs into a slice of TxResponse objects.
func formatTxResults(txConfig client.TxConfig, resTxs []*ctypes.ResultTx, resBlocks map[int64]*ctypes.ResultBlock) ([]*sdk.TxResponse, error) {
	var err error
	out := make([]*sdk.TxResponse, len(resTxs))
	for i := range resTxs {
		out[i], err = mkTxResult(txConfig, resTxs[i], resBlocks[resTxs[i].Height])
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

func getBlocksForTxResults(goCtx context.Context, clientCtx client.Context, resTxs []*ctypes.ResultTx) (map[int64]*ctypes.ResultBlock, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	resBlocks := make(map[int64]*ctypes.ResultBlock)

	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := node.Block(goCtx, &resTx.Height)
			if err != nil {
				return nil, err
			}

			resBlocks[resTx.Height] = resBlock
		}
	}

	return resBlocks, nil
}

func mkTxResult(txConfig client.TxConfig, resTx *ctypes.ResultTx, resBlock *ctypes.ResultBlock) (*sdk.TxResponse, error) {
	txb, err := txConfig.TxDecoder()(resTx.Tx)
	if err != nil {
		return nil, err
	}
	p, ok := txb.(intoAny)
	if !ok {
		return nil, fmt.Errorf("expecting a type implementing intoAny, got: %T", txb)
	}
	any := p.AsAny()
	return sdk.NewResponseResultTx(resTx, any, resBlock.Block.Time.Format(time.RFC3339)), nil
}

// Deprecated: this interface is used only internally for scenario we are
// deprecating (StdTxConfig support)
type intoAny interface {
	AsAny() *codectypes.Any
}
