package cli

import (
	"github.com/spf13/cobra"

	"https://github.com/ArjavJP/Cosmos-sdk/client"
	"https://github.com/ArjavJP/Cosmos-sdk/x/ibc/light-clients/06-solomachine/types"
)

// NewTxCmd returns a root CLI command handler for all solo machine transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "Solo Machine transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewCreateClientCmd(),
		NewUpdateClientCmd(),
		NewSubmitMisbehaviourCmd(),
	)

	return txCmd
}
