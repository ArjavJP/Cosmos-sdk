package evidence

import (
	sdk "github.com/ArjavJP/Cosmos-sdk/types"
	sdkerrors "github.com/ArjavJP/Cosmos-sdk/types/errors"
	"github.com/ArjavJP/Cosmos-sdk/x/evidence/keeper"
	"github.com/ArjavJP/Cosmos-sdk/x/evidence/types"
)

// NewHandler returns a handler for evidence messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgSubmitEvidence:
			res, err := msgServer.SubmitEvidence(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}
