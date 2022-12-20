package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelNFT = "cancel_nft"

var _ sdk.Msg = &MsgCancelNFT{}

func NewMsgCancelNFT(creator string, id string) *MsgCancelNFT {
	return &MsgCancelNFT{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgCancelNFT) Route() string {
	return RouterKey
}

func (msg *MsgCancelNFT) Type() string {
	return TypeMsgCancelNFT
}

func (msg *MsgCancelNFT) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
