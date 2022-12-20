package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSend{}, "github.com/yoshidan/cosmos-trustless-swap/Send", nil)
	cdc.RegisterConcrete(&MsgReceive{}, "github.com/yoshidan/cosmos-trustless-swap/Receive", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "github.com/yoshidan/cosmos-trustless-swap/Cancel", nil)
	cdc.RegisterConcrete(&MsgSendNFT{}, "github.com/yoshidan/cosmos-trustless-swap/SendNFT", nil)
	cdc.RegisterConcrete(&MsgCancelNFT{}, "github.com/yoshidan/cosmos-trustless-swap/CancelNFT", nil)
	cdc.RegisterConcrete(&MsgReceiveNFT{}, "github.com/yoshidan/cosmos-trustless-swap/ReceiveNFT", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReceive{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendNFT{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelNFT{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReceiveNFT{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
