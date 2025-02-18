package main

import (
	"os"

	"github.com/ArjavJP/Cosmos-sdk/server"
	svrcmd "github.com/ArjavJP/Cosmos-sdk/server/cmd"
	"github.com/ArjavJP/Cosmos-sdk/simapp"
	"github.com/ArjavJP/Cosmos-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
