package rest

import (
	"github.com/gorilla/mux"

	"github.com/ArjavJP/Cosmos-sdk/client/rest"

	"github.com/ArjavJP/Cosmos-sdk/client"
)

// RegisterRoutes registers REST routes for the upgrade module under the path specified by routeName.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}
