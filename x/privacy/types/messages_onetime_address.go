package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateOnetimeAddress = "create_onetime_address"
	TypeMsgUpdateOnetimeAddress = "update_onetime_address"
	TypeMsgDeleteOnetimeAddress = "delete_onetime_address"
)

var _ sdk.Msg = &MsgCreateOnetimeAddress{}

func NewMsgCreateOnetimeAddress(
	creator string,
	index string,

) *MsgCreateOnetimeAddress {
	return &MsgCreateOnetimeAddress{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateOnetimeAddress) IsPrivacy() bool {
	return true
}

func (msg *MsgCreateOnetimeAddress) Route() string {
	return RouterKey
}

func (msg *MsgCreateOnetimeAddress) Type() string {
	return TypeMsgCreateOnetimeAddress
}

func (msg *MsgCreateOnetimeAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOnetimeAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOnetimeAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateOnetimeAddress{}

func NewMsgUpdateOnetimeAddress(
	creator string,
	index string,

) *MsgUpdateOnetimeAddress {
	return &MsgUpdateOnetimeAddress{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateOnetimeAddress) IsPrivacy() bool {
	return true
}

func (msg *MsgUpdateOnetimeAddress) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOnetimeAddress) Type() string {
	return TypeMsgUpdateOnetimeAddress
}

func (msg *MsgUpdateOnetimeAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOnetimeAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOnetimeAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteOnetimeAddress{}

func NewMsgDeleteOnetimeAddress(
	creator string,
	index string,

) *MsgDeleteOnetimeAddress {
	return &MsgDeleteOnetimeAddress{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteOnetimeAddress) IsPrivacy() bool {
	return true
}

func (msg *MsgDeleteOnetimeAddress) Route() string {
	return RouterKey
}

func (msg *MsgDeleteOnetimeAddress) Type() string {
	return TypeMsgDeleteOnetimeAddress
}

func (msg *MsgDeleteOnetimeAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteOnetimeAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteOnetimeAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
