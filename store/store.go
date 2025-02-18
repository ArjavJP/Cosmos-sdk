package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/ArjavJP/Cosmos-sdk/store/cache"
	"github.com/ArjavJP/Cosmos-sdk/store/rootmulti"
	"github.com/ArjavJP/Cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
