package keeper

import (
    sdk "github.com/ArjavJP/Cosmos-sdk/types"
    "my-chain/x/gpuresource/types"
    "fmt"
)

// Keeper struct for GPU resource module
type Keeper struct {
    storeKey sdk.StoreKey
    cdc      sdk.Codec
}

// Aggregate GPU resources from stakers
func (k Keeper) AggregateGPUResources(ctx sdk.Context, gpuRequired int) ([]types.GPUResource, error) {
    var aggregatedResources []types.GPUResource
    var currentGPUCount int

    // Iterate through the stakers and collect GPU resources
    iter := ctx.KVStore(k.storeKey).Iterator(nil, nil)
    for ; iter.Valid(); iter.Next() {
        var resource types.GPUResource
        k.cdc.MustUnmarshal(iter.Value(), &resource)

        if currentGPUCount >= gpuRequired {
            break
        }

        aggregatedResources = append(aggregatedResources, resource)
        currentGPUCount += resource.GPUCount
    }

    // Ensure there are enough GPU resources
    if currentGPUCount < gpuRequired {
        return nil, fmt.Errorf("not enough GPU resources available")
    }

    return aggregatedResources, nil
}
