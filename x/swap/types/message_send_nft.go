package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendNFT = "send_nft"

var _ sdk.Msg = &MsgSendNFT{}

func NewMsgSendNFT(creator string, receiver string, classId string, nftId string, amountToReceive string) *MsgSendNFT {
	return &MsgSendNFT{
		Creator:         creator,
		Receiver:        receiver,
		ClassId:         classId,
		NftId:           nftId,
		AmountToReceive: amountToReceive,
	}
}

func (msg *MsgSendNFT) Route() string {
	return RouterKey
}

func (msg *MsgSendNFT) Type() string {
	return TypeMsgSendNFT
}

func (msg *MsgSendNFT) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
