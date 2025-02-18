package cli

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
    "my-chain/x/gpuresource/types"
    "github.com/cosmos/cosmos-sdk/types"
)

// Query command to check GPU resources
func GetGPUResourcesCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "gpu-resources",
        Short: "Query the available GPU resources",
        RunE: func(cmd *cobra.Command, args []string) error {
            // Querying for GPU resources
            res, _, err := client.QueryWithData(fmt.Sprintf("custom/%s/gpu-resources", types.ModuleName), nil)
            if err != nil {
                return err
            }
            var queryRes types.QueryGPUResourcesResponse
            clientCtx.Codec.UnmarshalJSON(res, &queryRes)

            // Print the result
            fmt.Println("Available GPU Resources:")
            for _, resource := range queryRes.Resources {
                fmt.Printf("Owner: %s, GPU Count: %d\n", resource.Owner, resource.GPUCount)
            }
            return nil
        },
    }

    return cmd
}
