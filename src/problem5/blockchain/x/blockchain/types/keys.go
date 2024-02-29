package types

const (
	// ModuleName defines the module name
	ModuleName = "blockchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blockchain"

	// PostKey defines the prefix for a post
	PostKey = "Post/value/"

	// PostCountKey defines the ID of the latest post added to the store
	PostCountKey = "Post/count/"
)

var (
	ParamsKey = []byte("p_blockchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
