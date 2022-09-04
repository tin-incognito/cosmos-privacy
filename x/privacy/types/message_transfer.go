package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgTransfer = "transfer"

var _ sdk.Msg = &MsgTransfer{}

func NewMsgTransfer() *MsgTransfer {
	return &MsgTransfer{}
}

func (msg *MsgTransfer) Route() string {
	return RouterKey
}

func (msg *MsgTransfer) Type() string {
	return TypeMsgTransfer
}

func (msg *MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransfer) ValidateBasic() error {
	return nil
}
