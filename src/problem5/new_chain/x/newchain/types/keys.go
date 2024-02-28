package types

const (
	// ModuleName defines the module name
	ModuleName = "newchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_newchain"

    
)

var (
	ParamsKey = []byte("p_newchain")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
