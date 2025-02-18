// types/query.go

package types

type QueryGPUResourcesRequest struct {
    // Any filters you might want (e.g., specific node ID, etc.)
}

type QueryGPUResourcesResponse struct {
    Resources []GPUResource `json:"resources"`
}

type GPUResource struct {
    Owner    string `json:"owner"`
    GPUCount int    `json:"gpu_count"`
}
