package types

import (
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
)

// SupplyKeeper defines the expected supply keeper (noalias)
type SupplyKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
