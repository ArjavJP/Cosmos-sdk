package rest

import (
	"net/http"

	"github.com/ArjavJP/Cosmos-sdk/client"
	"github.com/ArjavJP/Cosmos-sdk/client/tx"
	"github.com/ArjavJP/Cosmos-sdk/types/rest"
	govrest "github.com/ArjavJP/Cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/ArjavJP/Cosmos-sdk/x/gov/types"
	paramscutils "github.com/ArjavJP/Cosmos-sdk/x/params/client/utils"
	"github.com/ArjavJP/Cosmos-sdk/x/params/types/proposal"
)

// ProposalRESTHandler returns a ProposalRESTHandler that exposes the param
// change REST handler with a given sub-route.
func ProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "param_change",
		Handler:  postProposalHandlerFn(clientCtx),
	}
}

func postProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req paramscutils.ParamChangeProposalReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		content := proposal.NewParameterChangeProposal(req.Title, req.Description, req.Changes.ToParamChanges())

		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, req.Proposer)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
