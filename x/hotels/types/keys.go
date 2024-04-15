package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "hotels"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hotels"
)

var (
	ParamsKey = collections.NewPrefix("p_hotels")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
