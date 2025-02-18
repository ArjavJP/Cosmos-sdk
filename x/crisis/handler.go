package crisis

import (
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	sdkerrors "https://github.com/ArjavJP/Cosmos-sdk/types/errors"
	"https://github.com/ArjavJP/Cosmos-sdk/x/crisis/types"
)

// RouterKey
const RouterKey = types.ModuleName

func NewHandler(k types.MsgServer) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgVerifyInvariant:
			res, err := k.VerifyInvariant(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized crisis message type: %T", msg)
		}
	}
}
