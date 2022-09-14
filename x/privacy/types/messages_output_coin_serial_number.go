package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateOutputCoinSerialNumber = "create_output_coin_serial_number"
	TypeMsgUpdateOutputCoinSerialNumber = "update_output_coin_serial_number"
	TypeMsgDeleteOutputCoinSerialNumber = "delete_output_coin_serial_number"
)

var _ sdk.Msg = &MsgCreateOutputCoinSerialNumber{}

func NewMsgCreateOutputCoinSerialNumber(creator string) *MsgCreateOutputCoinSerialNumber {
	return &MsgCreateOutputCoinSerialNumber{
		Creator: creator,
	}
}

func (msg *MsgCreateOutputCoinSerialNumber) IsPrivacy() bool {
	return true
}

func (msg *MsgCreateOutputCoinSerialNumber) Route() string {
	return RouterKey
}

func (msg *MsgCreateOutputCoinSerialNumber) Type() string {
	return TypeMsgCreateOutputCoinSerialNumber
}

func (msg *MsgCreateOutputCoinSerialNumber) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOutputCoinSerialNumber) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOutputCoinSerialNumber) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateOutputCoinSerialNumber{}

func NewMsgUpdateOutputCoinSerialNumber(creator string) *MsgUpdateOutputCoinSerialNumber {
	return &MsgUpdateOutputCoinSerialNumber{
		Creator: creator,
	}
}

func (msg *MsgUpdateOutputCoinSerialNumber) IsPrivacy() bool {
	return true
}

func (msg *MsgUpdateOutputCoinSerialNumber) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOutputCoinSerialNumber) Type() string {
	return TypeMsgUpdateOutputCoinSerialNumber
}

func (msg *MsgUpdateOutputCoinSerialNumber) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOutputCoinSerialNumber) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOutputCoinSerialNumber) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteOutputCoinSerialNumber{}

func NewMsgDeleteOutputCoinSerialNumber(creator string) *MsgDeleteOutputCoinSerialNumber {
	return &MsgDeleteOutputCoinSerialNumber{
		Creator: creator,
	}
}

func (msg *MsgDeleteOutputCoinSerialNumber) IsPrivacy() bool {
	return true
}

func (msg *MsgDeleteOutputCoinSerialNumber) Route() string {
	return RouterKey
}

func (msg *MsgDeleteOutputCoinSerialNumber) Type() string {
	return TypeMsgDeleteOutputCoinSerialNumber
}

func (msg *MsgDeleteOutputCoinSerialNumber) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteOutputCoinSerialNumber) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteOutputCoinSerialNumber) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
