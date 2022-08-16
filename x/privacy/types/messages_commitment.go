package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateCommitment = "create_commitment"
	TypeMsgUpdateCommitment = "update_commitment"
	TypeMsgDeleteCommitment = "delete_commitment"
)

var _ sdk.Msg = &MsgCreateCommitment{}

func NewMsgCreateCommitment(
	creator string,
	index string,

) *MsgCreateCommitment {
	return &MsgCreateCommitment{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateCommitment) Route() string {
	return RouterKey
}

func (msg *MsgCreateCommitment) Type() string {
	return TypeMsgCreateCommitment
}

func (msg *MsgCreateCommitment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCommitment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCommitment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCommitment{}

func NewMsgUpdateCommitment(
	creator string,
	index string,

) *MsgUpdateCommitment {
	return &MsgUpdateCommitment{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateCommitment) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCommitment) Type() string {
	return TypeMsgUpdateCommitment
}

func (msg *MsgUpdateCommitment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCommitment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCommitment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCommitment{}

func NewMsgDeleteCommitment(
	creator string,
	index string,

) *MsgDeleteCommitment {
	return &MsgDeleteCommitment{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteCommitment) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCommitment) Type() string {
	return TypeMsgDeleteCommitment
}

func (msg *MsgDeleteCommitment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCommitment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCommitment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
