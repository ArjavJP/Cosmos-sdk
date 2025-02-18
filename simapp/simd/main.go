package main

import (
	"os"

	"https://github.com/ArjavJP/Cosmos-sdk/server"
	svrcmd "https://github.com/ArjavJP/Cosmos-sdk/server/cmd"
	"https://github.com/ArjavJP/Cosmos-sdk/simapp"
	"https://github.com/ArjavJP/Cosmos-sdk/simapp/simd/cmd"
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
