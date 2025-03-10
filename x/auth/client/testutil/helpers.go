package testutil

import (
	"fmt"
	"strings"

	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/ArjavJP/Cosmos-sdk/client"
	"github.com/ArjavJP/Cosmos-sdk/client/flags"
	"github.com/ArjavJP/Cosmos-sdk/crypto/keyring"
	"github.com/ArjavJP/Cosmos-sdk/testutil"
	clitestutil "github.com/ArjavJP/Cosmos-sdk/testutil/cli"
	"github.com/ArjavJP/Cosmos-sdk/x/auth/client/cli"
)

func TxSignExec(clientCtx client.Context, from fmt.Stringer, filename string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
		fmt.Sprintf("--from=%s", from.String()),
		fmt.Sprintf("--%s=%s", flags.FlagHome, strings.Replace(clientCtx.HomeDir, "simd", "simcli", 1)),
		fmt.Sprintf("--%s=%s", flags.FlagChainID, clientCtx.ChainID),
		filename,
	}

	cmd := cli.GetSignCommand()
	tmcli.PrepareBaseCmd(cmd, "", "")

	return clitestutil.ExecTestCLICmd(clientCtx, cmd, append(args, extraArgs...))
}

func TxBroadcastExec(clientCtx client.Context, filename string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		filename,
	}

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetBroadcastCommand(), append(args, extraArgs...))
}

func TxEncodeExec(clientCtx client.Context, filename string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
		filename,
	}

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetEncodeCommand(), append(args, extraArgs...))
}

func TxValidateSignaturesExec(clientCtx client.Context, filename string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
		fmt.Sprintf("--%s=%s", flags.FlagChainID, clientCtx.ChainID),
		filename,
	}

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetValidateSignaturesCommand(), args)
}

func TxMultiSignExec(clientCtx client.Context, from string, filename string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
		fmt.Sprintf("--%s=%s", flags.FlagChainID, clientCtx.ChainID),
		filename,
		from,
	}

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetMultiSignCommand(), append(args, extraArgs...))
}

func TxSignBatchExec(clientCtx client.Context, from fmt.Stringer, filename string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
		fmt.Sprintf("--from=%s", from.String()),
		filename,
	}

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetSignBatchCommand(), append(args, extraArgs...))
}

func TxDecodeExec(clientCtx client.Context, encodedTx string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
		encodedTx,
	}

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetDecodeCommand(), append(args, extraArgs...))
}

func QueryAccountExec(clientCtx client.Context, address fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{address.String(), fmt.Sprintf("--%s=json", tmcli.OutputFlag)}

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetAccountCmd(), append(args, extraArgs...))
}

func TxMultiSignBatchExec(clientCtx client.Context, filename string, from string, sigFile1 string, sigFile2 string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
		filename,
		from,
		sigFile1,
		sigFile2,
	}

	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetMultiSignBatchCmd(), args)
}

// DONTCOVER
