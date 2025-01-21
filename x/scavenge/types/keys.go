package types

const (
	// ModuleName defines the module name
	ModuleName = "scavenge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_scavenge"
)

var (
	ParamsKey = []byte("p_scavenge")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
