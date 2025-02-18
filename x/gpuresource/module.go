package gpuresource

import (
    "https://github.com/ArjavJP/Cosmos-sdk/types"
    "my-chain/x/gpuresource/keeper"
    "my-chain/x/gpuresource/handler"
    "my-chain/x/gpuresource/types"
)

// ModuleName defines the name of the GPU resource-sharing module
const ModuleName = "gpuresource"

// RegisterCodec registers the module codec
func RegisterCodec(cdc *types.Codec) {
    cdc.RegisterConcrete(&types.SubmitGPUJobRequest{}, "gpuresource/SubmitGPUJobRequest", nil)
    cdc.RegisterConcrete(&types.GPUResource{}, "gpuresource/GPUResource", nil)
}

// NewHandler returns the handler for the GPU resource-sharing module
func NewHandler(k keeper.Keeper) types.Handler {
    return func(ctx types.Context, msg types.Msg) types.Response {
        switch msg := msg.(type) {
        case types.SubmitGPUJobRequest:
            return handler.handleSubmitGPUJobRequest(ctx, msg, k)
        default:
            return types.Response{Code: types.CodeTypeUnknownRequest, Log: "unrecognized request"}
        }
    }
}

// Initialize the module
func NewKeeper(storeKey types.StoreKey, cdc types.Codec) keeper.Keeper {
    return keeper.Keeper{
        storeKey: storeKey,
        cdc:      cdc,
    }
}
