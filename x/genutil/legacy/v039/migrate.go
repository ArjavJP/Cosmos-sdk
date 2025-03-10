package v039

import (
	"github.com/ArjavJP/Cosmos-sdk/client"
	"github.com/ArjavJP/Cosmos-sdk/codec"
	v038auth "github.com/ArjavJP/Cosmos-sdk/x/auth/legacy/v038"
	v039auth "github.com/ArjavJP/Cosmos-sdk/x/auth/legacy/v039"
	v036distr "github.com/ArjavJP/Cosmos-sdk/x/distribution/legacy/v036"
	"github.com/ArjavJP/Cosmos-sdk/x/genutil/types"
	v036gov "github.com/ArjavJP/Cosmos-sdk/x/gov/legacy/v036"
	v036params "github.com/ArjavJP/Cosmos-sdk/x/params/legacy/v036"
	v038upgrade "github.com/ArjavJP/Cosmos-sdk/x/upgrade/legacy/v038"
)

// Migrate migrates exported state from v0.38 to a v0.39 genesis state.
//
// NOTE: No actual migration occurs since the types do not change, but JSON
// serialization of accounts do change.
func Migrate(appState types.AppMap, _ client.Context) types.AppMap {
	v038Codec := codec.NewLegacyAmino()
	v038auth.RegisterLegacyAminoCodec(v038Codec)
	v036gov.RegisterLegacyAminoCodec(v038Codec)
	v036distr.RegisterLegacyAminoCodec(v038Codec)
	v036params.RegisterLegacyAminoCodec(v038Codec)
	v038upgrade.RegisterLegacyAminoCodec(v038Codec)

	v039Codec := codec.NewLegacyAmino()
	v039auth.RegisterLegacyAminoCodec(v039Codec)
	v036gov.RegisterLegacyAminoCodec(v039Codec)
	v036distr.RegisterLegacyAminoCodec(v039Codec)
	v036params.RegisterLegacyAminoCodec(v039Codec)
	v038upgrade.RegisterLegacyAminoCodec(v039Codec)

	// migrate x/auth state (JSON serialization only)
	if appState[v038auth.ModuleName] != nil {
		var authGenState v038auth.GenesisState
		v038Codec.MustUnmarshalJSON(appState[v038auth.ModuleName], &authGenState)

		delete(appState, v038auth.ModuleName) // delete old key in case the name changed
		appState[v039auth.ModuleName] = v039Codec.MustMarshalJSON(v039auth.Migrate(authGenState))
	}

	return appState
}
