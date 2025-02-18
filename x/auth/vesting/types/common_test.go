package types_test

import (
	"https://github.com/ArjavJP/Cosmos-sdk/simapp"
)

var (
	app         = simapp.Setup(false)
	appCodec, _ = simapp.MakeCodecs()
)
