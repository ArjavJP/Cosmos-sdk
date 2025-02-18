package testutil

import (
	"fmt"

	"github.com/ArjavJP/Cosmos-sdk/client"
	"github.com/ArjavJP/Cosmos-sdk/client/flags"
	"github.com/ArjavJP/Cosmos-sdk/testutil"
	clitestutil "github.com/ArjavJP/Cosmos-sdk/testutil/cli"
	sdk "github.com/ArjavJP/Cosmos-sdk/types"
	stakingcli "github.com/ArjavJP/Cosmos-sdk/x/staking/client/cli"
)

var commonArgs = []string{
	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
}

// MsgRedelegateExec creates a redelegate message.
func MsgRedelegateExec(clientCtx client.Context, from, src, dst, amount fmt.Stringer,
	extraArgs ...string) (testutil.BufferWriter, error) {

	args := []string{
		src.String(),
		dst.String(),
		amount.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from.String()),
	}

	args = append(args, commonArgs...)
	return clitestutil.ExecTestCLICmd(clientCtx, stakingcli.NewRedelegateCmd(), args)
}

// MsgUnbondExec creates a unbond message.
func MsgUnbondExec(clientCtx client.Context, from fmt.Stringer, valAddress,
	amount fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {

	args := []string{
		valAddress.String(),
		amount.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from.String()),
	}

	args = append(args, commonArgs...)
	return clitestutil.ExecTestCLICmd(clientCtx, stakingcli.NewUnbondCmd(), args)
}
