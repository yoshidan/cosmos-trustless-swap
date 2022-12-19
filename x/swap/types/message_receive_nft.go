package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReceiveNFT = "receive_nft"

var _ sdk.Msg = &MsgReceiveNFT{}

func NewMsgReceiveNFT(creator string, id uint64) *MsgReceiveNFT {
	return &MsgReceiveNFT{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgReceiveNFT) Route() string {
	return RouterKey
}

func (msg *MsgReceiveNFT) Type() string {
	return TypeMsgReceiveNFT
}

func (msg *MsgReceiveNFT) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReceiveNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReceiveNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
