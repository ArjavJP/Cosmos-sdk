package slashing

import (
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	sdkerrors "https://github.com/ArjavJP/Cosmos-sdk/types/errors"
	"https://github.com/ArjavJP/Cosmos-sdk/x/slashing/keeper"
	"https://github.com/ArjavJP/Cosmos-sdk/x/slashing/types"
)

// NewHandler creates an sdk.Handler for all the slashing type messages
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		msgServer := keeper.NewMsgServerImpl(k)

		switch msg := msg.(type) {
		case *types.MsgUnjail:
			res, err := msgServer.Unjail(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}
