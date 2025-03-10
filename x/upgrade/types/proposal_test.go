package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ArjavJP/Cosmos-sdk/codec"
	codectypes "github.com/ArjavJP/Cosmos-sdk/codec/types"
	gov "github.com/ArjavJP/Cosmos-sdk/x/gov/types"
	clienttypes "github.com/ArjavJP/Cosmos-sdk/x/ibc/core/02-client/types"
	ibctmtypes "github.com/ArjavJP/Cosmos-sdk/x/ibc/light-clients/07-tendermint/types"
	"github.com/ArjavJP/Cosmos-sdk/x/upgrade/types"
)

type ProposalWrapper struct {
	Prop gov.Content
}

func TestContentAccessors(t *testing.T) {
	cases := map[string]struct {
		p     gov.Content
		title string
		desc  string
		typ   string
		str   string
	}{
		"upgrade": {
			p: types.NewSoftwareUpgradeProposal("Title", "desc", types.Plan{
				Name: "due_time",
				Info: "https://foo.bar",
				Time: mustParseTime("2019-07-08T11:33:55Z"),
			}),
			title: "Title",
			desc:  "desc",
			typ:   "SoftwareUpgrade",
			str:   "Software Upgrade Proposal:\n  Title:       Title\n  Description: desc\n",
		},
		"cancel": {
			p:     types.NewCancelSoftwareUpgradeProposal("Cancel", "bad idea"),
			title: "Cancel",
			desc:  "bad idea",
			typ:   "CancelSoftwareUpgrade",
			str:   "Cancel Software Upgrade Proposal:\n  Title:       Cancel\n  Description: bad idea\n",
		},
	}

	cdc := codec.NewLegacyAmino()
	gov.RegisterLegacyAminoCodec(cdc)
	types.RegisterLegacyAminoCodec(cdc)

	for name, tc := range cases {
		tc := tc // copy to local variable for scopelint
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.title, tc.p.GetTitle())
			assert.Equal(t, tc.desc, tc.p.GetDescription())
			assert.Equal(t, tc.typ, tc.p.ProposalType())
			assert.Equal(t, "upgrade", tc.p.ProposalRoute())
			assert.Equal(t, tc.str, tc.p.String())

			// try to encode and decode type to ensure codec works
			wrap := ProposalWrapper{tc.p}
			bz, err := cdc.MarshalBinaryBare(&wrap)
			require.NoError(t, err)
			unwrap := ProposalWrapper{}
			err = cdc.UnmarshalBinaryBare(bz, &unwrap)
			require.NoError(t, err)

			// all methods should look the same
			assert.Equal(t, tc.title, unwrap.Prop.GetTitle())
			assert.Equal(t, tc.desc, unwrap.Prop.GetDescription())
			assert.Equal(t, tc.typ, unwrap.Prop.ProposalType())
			assert.Equal(t, "upgrade", unwrap.Prop.ProposalRoute())
			assert.Equal(t, tc.str, unwrap.Prop.String())

		})

	}
}

// tests a software update proposal can be marshaled and unmarshaled, and the
// client state can be unpacked
func TestMarshalSoftwareUpdateProposal(t *testing.T) {
	cs, err := clienttypes.PackClientState(&ibctmtypes.ClientState{})
	require.NoError(t, err)

	// create proposal
	plan := types.Plan{
		Name:                "upgrade ibc",
		Height:              1000,
		UpgradedClientState: cs,
	}
	content := types.NewSoftwareUpgradeProposal("title", "description", plan)
	sup, ok := content.(*types.SoftwareUpgradeProposal)
	require.True(t, ok)

	// create codec
	ir := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(ir)
	clienttypes.RegisterInterfaces(ir)
	gov.RegisterInterfaces(ir)
	ibctmtypes.RegisterInterfaces(ir)
	cdc := codec.NewProtoCodec(ir)

	// marshal message
	bz, err := cdc.MarshalJSON(sup)
	require.NoError(t, err)

	// unmarshal proposal
	newSup := &types.SoftwareUpgradeProposal{}
	err = cdc.UnmarshalJSON(bz, newSup)
	require.NoError(t, err)

	// unpack client state
	_, err = clienttypes.UnpackClientState(newSup.Plan.UpgradedClientState)
	require.NoError(t, err)
}
