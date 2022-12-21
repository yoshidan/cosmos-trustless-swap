package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSell = "sell"

var _ sdk.Msg = &MsgSell{}

func NewMsgSell(creator string, id uint64, amount string, price string) *MsgSell {
	return &MsgSell{
		Creator: creator,
		Id:      id,
		Amount:  amount,
		Price:   price,
	}
}

func (msg *MsgSell) Route() string {
	return RouterKey
}

func (msg *MsgSell) Type() string {
	return TypeMsgSell
}

func (msg *MsgSell) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSell) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSell) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
