package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "nestchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_nestchain"
)

var (
	ParamsKey = collections.NewPrefix("p_nestchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
