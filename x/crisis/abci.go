package crisis

import (
	"time"

	"https://github.com/ArjavJP/Cosmos-sdk/telemetry"
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	"https://github.com/ArjavJP/Cosmos-sdk/x/crisis/keeper"
	"https://github.com/ArjavJP/Cosmos-sdk/x/crisis/types"
)

// check all registered invariants
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	if k.InvCheckPeriod() == 0 || ctx.BlockHeight()%int64(k.InvCheckPeriod()) != 0 {
		// skip running the invariant check
		return
	}
	k.AssertInvariants(ctx)
}
