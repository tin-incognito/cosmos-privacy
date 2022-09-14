package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateCommitmentIndex = "create_commitment_index"
	TypeMsgUpdateCommitmentIndex = "update_commitment_index"
	TypeMsgDeleteCommitmentIndex = "delete_commitment_index"
)

var _ sdk.Msg = &MsgCreateCommitmentIndex{}

func NewMsgCreateCommitmentIndex(
	creator string,
	index string,

) *MsgCreateCommitmentIndex {
	return &MsgCreateCommitmentIndex{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateCommitmentIndex) IsPrivacy() bool {
	return true
}

func (msg *MsgCreateCommitmentIndex) Route() string {
	return RouterKey
}

func (msg *MsgCreateCommitmentIndex) Type() string {
	return TypeMsgCreateCommitmentIndex
}

func (msg *MsgCreateCommitmentIndex) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCommitmentIndex) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCommitmentIndex) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCommitmentIndex{}

func NewMsgUpdateCommitmentIndex(
	creator string,
	index string,

) *MsgUpdateCommitmentIndex {
	return &MsgUpdateCommitmentIndex{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateCommitmentIndex) IsPrivacy() bool {
	return true
}

func (msg *MsgUpdateCommitmentIndex) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCommitmentIndex) Type() string {
	return TypeMsgUpdateCommitmentIndex
}

func (msg *MsgUpdateCommitmentIndex) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCommitmentIndex) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCommitmentIndex) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCommitmentIndex{}

func NewMsgDeleteCommitmentIndex(
	creator string,
	index string,

) *MsgDeleteCommitmentIndex {
	return &MsgDeleteCommitmentIndex{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteCommitmentIndex) IsPrivacy() bool {
	return true
}

func (msg *MsgDeleteCommitmentIndex) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCommitmentIndex) Type() string {
	return TypeMsgDeleteCommitmentIndex
}

func (msg *MsgDeleteCommitmentIndex) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCommitmentIndex) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCommitmentIndex) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
