package gpuresource

import (
    "github.com/ArjavJP/Cosmos-sdk/types"
    "my-chain/x/gpuresource/types"
    "fmt"
)

// Querying available GPU resources
func GetGPUResources(ctx types.Context, req *types.QueryGPUResourcesRequest) (*types.QueryGPUResourcesResponse, error) {
    // Fetch GPU resource data from the store
    store := ctx.KVStore(k.storeKey)
    res := []types.GPUResource{}

    // Example of fetching resources from storage
    iterator := store.Iterator(nil, nil)
    defer iterator.Close()
    for ; iterator.Valid(); iterator.Next() {
        var resource types.GPUResource
        k.cdc.MustUnmarshal(iterator.Value(), &resource)
        res = append(res, resource)
    }

    return &types.QueryGPUResourcesResponse{
        Resources: res,
    }, nil
}
