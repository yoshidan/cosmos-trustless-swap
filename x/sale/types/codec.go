package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSell{}, "sale/Sell", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "sale/Cancel", nil)
	cdc.RegisterConcrete(&MsgBuy{}, "sale/Buy", nil)
	cdc.RegisterConcrete(&MsgSellNFT{}, "sale/SellNFT", nil)
	cdc.RegisterConcrete(&MsgBuyNFT{}, "sale/BuyNFT", nil)
	cdc.RegisterConcrete(&MsgCancelNFT{}, "sale/CancelNFT", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSell{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSellNFT{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyNFT{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelNFT{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
