package codec

import (
	codectypes "https://github.com/ArjavJP/Cosmos-sdk/codec/types"
	"https://github.com/ArjavJP/Cosmos-sdk/crypto/keys/ed25519"
	"https://github.com/ArjavJP/Cosmos-sdk/crypto/keys/multisig"
	"https://github.com/ArjavJP/Cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "https://github.com/ArjavJP/Cosmos-sdk/crypto/types"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.crypto.PubKey", (*cryptotypes.PubKey)(nil))
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ed25519.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &secp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &multisig.LegacyAminoPubKey{})
}
