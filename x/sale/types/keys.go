package types

const (
	// ModuleName defines the module name
	ModuleName = "sale"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sale"

	SaleKey         = "Sale/value/"
	MaxSaleIDKey    = "Sale/max/"
	NFTSaleKey      = "Sale/NFT/value/"
	MaxNFTSaleIDKey = "Sale/NFT/max/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
