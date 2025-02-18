package types

import (
    sdk "github.com/ArjavJP/Cosmos-sdk/types"
)

// GPUResource represents a single GPU resource stake
type GPUResource struct {
    Owner      sdk.AccAddress `json:"owner" yaml:"owner"`
    GPUCount   int            `json:"gpu_count" yaml:"gpu_count"` // GPU count for resource sharing
    ResourceID string         `json:"resource_id" yaml:"resource_id"`
}

// SubmitGPUJobRequest is the message to request GPU resources for a job
type SubmitGPUJobRequest struct {
    JobID       string         `json:"job_id" yaml:"job_id"`
    GPURequired int            `json:"gpu_required" yaml:"gpu_required"`
    UserAddress sdk.AccAddress `json:"user_address" yaml:"user_address"`
}
