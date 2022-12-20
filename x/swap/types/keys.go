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

	SwapKey         = "github.com/yoshidan/cosmos-trustless-swap/value/"
	MaxSwapIDKey    = "github.com/yoshidan/cosmos-trustless-swap/max/"
	NFTSwapKey      = "github.com/yoshidan/cosmos-trustless-swap/NFT/value/"
	MaxNFTSwapIDKey = "github.com/yoshidan/cosmos-trustless-swap/NFT/max/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
