package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReceive = "receive"

var _ sdk.Msg = &MsgReceive{}

func NewMsgReceive(creator string, id uint64) *MsgReceive {
	return &MsgReceive{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgReceive) Route() string {
	return RouterKey
}

func (msg *MsgReceive) Type() string {
	return TypeMsgReceive
}

func (msg *MsgReceive) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReceive) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReceive) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
