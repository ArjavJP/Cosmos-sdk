package upgrade

import (
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
	sdkerrors "https://github.com/ArjavJP/Cosmos-sdk/types/errors"
	govtypes "https://github.com/ArjavJP/Cosmos-sdk/x/gov/types"
	"https://github.com/ArjavJP/Cosmos-sdk/x/upgrade/keeper"
	"https://github.com/ArjavJP/Cosmos-sdk/x/upgrade/types"
)

// NewSoftwareUpgradeProposalHandler creates a governance handler to manage new proposal types.
// It enables SoftwareUpgradeProposal to propose an Upgrade, and CancelSoftwareUpgradeProposal
// to abort a previously voted upgrade.
func NewSoftwareUpgradeProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.SoftwareUpgradeProposal:
			return handleSoftwareUpgradeProposal(ctx, k, c)

		case *types.CancelSoftwareUpgradeProposal:
			return handleCancelSoftwareUpgradeProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized software upgrade proposal content type: %T", c)
		}
	}
}

func handleSoftwareUpgradeProposal(ctx sdk.Context, k keeper.Keeper, p *types.SoftwareUpgradeProposal) error {
	return k.ScheduleUpgrade(ctx, p.Plan)
}

func handleCancelSoftwareUpgradeProposal(ctx sdk.Context, k keeper.Keeper, _ *types.CancelSoftwareUpgradeProposal) error {
	k.ClearUpgradePlan(ctx)
	return nil
}
