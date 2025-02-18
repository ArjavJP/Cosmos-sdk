package v040

import (
	"https://github.com/ArjavJP/Cosmos-sdk/client"
	"https://github.com/ArjavJP/Cosmos-sdk/codec"
	v039auth "https://github.com/ArjavJP/Cosmos-sdk/x/auth/legacy/v039"
	v040auth "https://github.com/ArjavJP/Cosmos-sdk/x/auth/legacy/v040"
	v036supply "https://github.com/ArjavJP/Cosmos-sdk/x/bank/legacy/v036"
	v038bank "https://github.com/ArjavJP/Cosmos-sdk/x/bank/legacy/v038"
	v040bank "https://github.com/ArjavJP/Cosmos-sdk/x/bank/legacy/v040"
	v039crisis "https://github.com/ArjavJP/Cosmos-sdk/x/crisis/legacy/v039"
	v040crisis "https://github.com/ArjavJP/Cosmos-sdk/x/crisis/legacy/v040"
	v036distr "https://github.com/ArjavJP/Cosmos-sdk/x/distribution/legacy/v036"
	v038distr "https://github.com/ArjavJP/Cosmos-sdk/x/distribution/legacy/v038"
	v040distr "https://github.com/ArjavJP/Cosmos-sdk/x/distribution/legacy/v040"
	v038evidence "https://github.com/ArjavJP/Cosmos-sdk/x/evidence/legacy/v038"
	v040evidence "https://github.com/ArjavJP/Cosmos-sdk/x/evidence/legacy/v040"
	v039genutil "https://github.com/ArjavJP/Cosmos-sdk/x/genutil/legacy/v039"
	"https://github.com/ArjavJP/Cosmos-sdk/x/genutil/types"
	v036gov "https://github.com/ArjavJP/Cosmos-sdk/x/gov/legacy/v036"
	v040gov "https://github.com/ArjavJP/Cosmos-sdk/x/gov/legacy/v040"
	v039mint "https://github.com/ArjavJP/Cosmos-sdk/x/mint/legacy/v039"
	v040mint "https://github.com/ArjavJP/Cosmos-sdk/x/mint/legacy/v040"
	v036params "https://github.com/ArjavJP/Cosmos-sdk/x/params/legacy/v036"
	v039slashing "https://github.com/ArjavJP/Cosmos-sdk/x/slashing/legacy/v039"
	v040slashing "https://github.com/ArjavJP/Cosmos-sdk/x/slashing/legacy/v040"
	v038staking "https://github.com/ArjavJP/Cosmos-sdk/x/staking/legacy/v038"
	v040staking "https://github.com/ArjavJP/Cosmos-sdk/x/staking/legacy/v040"
	v038upgrade "https://github.com/ArjavJP/Cosmos-sdk/x/upgrade/legacy/v038"
)

func migrateGenutil(oldGenState v039genutil.GenesisState) *types.GenesisState {
	return &types.GenesisState{
		GenTxs: oldGenState.GenTxs,
	}
}

// Migrate migrates exported state from v0.39 to a v0.40 genesis state.
func Migrate(appState types.AppMap, clientCtx client.Context) types.AppMap {
	v039Codec := codec.NewLegacyAmino()
	v039auth.RegisterLegacyAminoCodec(v039Codec)
	v036gov.RegisterLegacyAminoCodec(v039Codec)
	v036distr.RegisterLegacyAminoCodec(v039Codec)
	v036params.RegisterLegacyAminoCodec(v039Codec)
	v038upgrade.RegisterLegacyAminoCodec(v039Codec)

	v040Codec := clientCtx.JSONMarshaler

	if appState[v038bank.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var bankGenState v038bank.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v038bank.ModuleName], &bankGenState)

		// unmarshal x/auth genesis state to retrieve all account balances
		var authGenState v039auth.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039auth.ModuleName], &authGenState)

		// unmarshal x/supply genesis state to retrieve total supply
		var supplyGenState v036supply.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v036supply.ModuleName], &supplyGenState)

		// delete deprecated x/bank genesis state
		delete(appState, v038bank.ModuleName)

		// delete deprecated x/supply genesis state
		delete(appState, v036supply.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v040bank.ModuleName] = v040Codec.MustMarshalJSON(v040bank.Migrate(bankGenState, authGenState, supplyGenState))
	}

	// remove balances from existing accounts
	if appState[v039auth.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var authGenState v039auth.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039auth.ModuleName], &authGenState)

		// delete deprecated x/auth genesis state
		delete(appState, v039auth.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v040auth.ModuleName] = v040Codec.MustMarshalJSON(v040auth.Migrate(authGenState))
	}

	// Migrate x/crisis.
	if appState[v039crisis.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var crisisGenState v039crisis.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039crisis.ModuleName], &crisisGenState)

		// delete deprecated x/crisis genesis state
		delete(appState, v039crisis.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v040crisis.ModuleName] = v040Codec.MustMarshalJSON(v040crisis.Migrate(crisisGenState))
	}

	// Migrate x/distribution.
	if appState[v038distr.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var distributionGenState v038distr.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v038distr.ModuleName], &distributionGenState)

		// delete deprecated x/distribution genesis state
		delete(appState, v038distr.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v040distr.ModuleName] = v040Codec.MustMarshalJSON(v040distr.Migrate(distributionGenState))
	}

	// Migrate x/evidence.
	if appState[v038evidence.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var evidenceGenState v038evidence.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v038bank.ModuleName], &evidenceGenState)

		// delete deprecated x/evidence genesis state
		delete(appState, v038evidence.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v040evidence.ModuleName] = v040Codec.MustMarshalJSON(v040evidence.Migrate(evidenceGenState))
	}

	// Migrate x/gov.
	if appState[v036gov.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var govGenState v036gov.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v036gov.ModuleName], &govGenState)

		// delete deprecated x/gov genesis state
		delete(appState, v036gov.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v040gov.ModuleName] = v040Codec.MustMarshalJSON(v040gov.Migrate(govGenState))
	}

	// Migrate x/mint.
	if appState[v039mint.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var mintGenState v039mint.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039mint.ModuleName], &mintGenState)

		// delete deprecated x/mint genesis state
		delete(appState, v039mint.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v040mint.ModuleName] = v040Codec.MustMarshalJSON(v040mint.Migrate(mintGenState))
	}

	// Migrate x/slashing.
	if appState[v039slashing.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var slashingGenState v039slashing.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039slashing.ModuleName], &slashingGenState)

		// delete deprecated x/slashing genesis state
		delete(appState, v039slashing.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v040slashing.ModuleName] = v040Codec.MustMarshalJSON(v040slashing.Migrate(slashingGenState))
	}

	// Migrate x/staking.
	if appState[v038staking.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var stakingGenState v038staking.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v038staking.ModuleName], &stakingGenState)

		// delete deprecated x/staking genesis state
		delete(appState, v038staking.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v040staking.ModuleName] = v040Codec.MustMarshalJSON(v040staking.Migrate(stakingGenState))
	}

	// Migrate x/genutil
	if appState[v039genutil.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var genutilGenState v039genutil.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039genutil.ModuleName], &genutilGenState)

		// delete deprecated x/staking genesis state
		delete(appState, v039genutil.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[ModuleName] = v040Codec.MustMarshalJSON(migrateGenutil(genutilGenState))
	}

	return appState
}
