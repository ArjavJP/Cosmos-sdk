package rest

import (
	"github.com/gorilla/mux"

	"https://github.com/ArjavJP/Cosmos-sdk/client"
	"https://github.com/ArjavJP/Cosmos-sdk/client/rest"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
}
