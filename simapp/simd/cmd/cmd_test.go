package cmd_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	svrcmd "https://github.com/ArjavJP/Cosmos-sdk/server/cmd"
	"https://github.com/ArjavJP/Cosmos-sdk/simapp"
	"https://github.com/ArjavJP/Cosmos-sdk/simapp/simd/cmd"
	"https://github.com/ArjavJP/Cosmos-sdk/x/genutil/client/cli"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := cmd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",        // Test the init cmd
		"simapp-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
	})

	require.NoError(t, svrcmd.Execute(rootCmd, simapp.DefaultNodeHome))
}
