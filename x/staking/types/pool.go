package types

import (
	sdk "github.com/ArjavJP/Cosmos-sdk/types"
)

// names used as root for pool module accounts:
//
// - NotBondedPool -> "not_bonded_tokens_pool"
//
// - BondedPool -> "bonded_tokens_pool"
const (
	NotBondedPoolName = "not_bonded_tokens_pool"
	BondedPoolName    = "bonded_tokens_pool"
)

// NewPool creates a new Pool instance used for queries
func NewPool(notBonded, bonded sdk.Int) Pool {
	return Pool{
		NotBondedTokens: notBonded,
		BondedTokens:    bonded,
	}
}
