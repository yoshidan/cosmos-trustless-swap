package types

const (
	// ModuleName defines the module name
	ModuleName = "swap"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_swap"

	SwapKey         = "swap/value/"
	MaxSwapIDKey    = "swap/max/"
	NFTSwapKey      = "swap/NFT/value/"
	MaxNFTSwapIDKey = "swap/NFT/max/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
