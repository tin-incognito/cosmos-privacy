package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSerialNumber = "create_serial_number"
	TypeMsgUpdateSerialNumber = "update_serial_number"
	TypeMsgDeleteSerialNumber = "delete_serial_number"
)

var _ sdk.Msg = &MsgCreateSerialNumber{}

func NewMsgCreateSerialNumber(
	creator string,
	index string,

) *MsgCreateSerialNumber {
	return &MsgCreateSerialNumber{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateSerialNumber) Route() string {
	return RouterKey
}

func (msg *MsgCreateSerialNumber) Type() string {
	return TypeMsgCreateSerialNumber
}

func (msg *MsgCreateSerialNumber) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSerialNumber) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSerialNumber) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSerialNumber{}

func NewMsgUpdateSerialNumber(
	creator string,
	index string,

) *MsgUpdateSerialNumber {
	return &MsgUpdateSerialNumber{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateSerialNumber) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSerialNumber) Type() string {
	return TypeMsgUpdateSerialNumber
}

func (msg *MsgUpdateSerialNumber) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSerialNumber) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSerialNumber) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSerialNumber{}

func NewMsgDeleteSerialNumber(
	creator string,
	index string,

) *MsgDeleteSerialNumber {
	return &MsgDeleteSerialNumber{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteSerialNumber) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSerialNumber) Type() string {
	return TypeMsgDeleteSerialNumber
}

func (msg *MsgDeleteSerialNumber) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSerialNumber) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSerialNumber) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
