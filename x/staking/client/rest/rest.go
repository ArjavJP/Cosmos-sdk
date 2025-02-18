package rest

import (
	"github.com/gorilla/mux"

	"https://github.com/ArjavJP/Cosmos-sdk/client/rest"

	"https://github.com/ArjavJP/Cosmos-sdk/client"
)

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}
