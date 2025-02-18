package store

import (
	dbm "github.com/tendermint/tm-db"

	"https://github.com/ArjavJP/Cosmos-sdk/store/cache"
	"https://github.com/ArjavJP/Cosmos-sdk/store/rootmulti"
	"https://github.com/ArjavJP/Cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
