package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateTx = "create_tx"

var _ sdk.Msg = &MsgCreateTx{}

func NewMsgCreateTx(creator string) *MsgCreateTx {
	return &MsgCreateTx{
		Creator: creator,
	}
}

func (msg *MsgCreateTx) Route() string {
	return RouterKey
}

func (msg *MsgCreateTx) Type() string {
	return TypeMsgCreateTx
}

func (msg *MsgCreateTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
