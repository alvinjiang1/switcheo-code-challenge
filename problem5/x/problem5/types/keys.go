package types

const (
	// ModuleName defines the module name
	ModuleName = "problem5"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_problem5"
)

var (
	ParamsKey = []byte("p_problem5")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
