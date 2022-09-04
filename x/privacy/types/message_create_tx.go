package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgCreateTx = "create_tx"

var _ sdk.Msg = &MsgCreateTx{}

func NewMsgCreateTx(value []byte) *MsgCreateTx {
	return &MsgCreateTx{
		Value: value,
	}
}

func (msg *MsgCreateTx) Route() string {
	return RouterKey
}

func (msg *MsgCreateTx) Type() string {
	return TypeMsgCreateTx
}

func (msg *MsgCreateTx) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgCreateTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTx) ValidateBasic() error {
	return nil
}
