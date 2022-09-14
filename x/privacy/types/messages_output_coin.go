package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateOutputCoin = "create_output_coin"
	TypeMsgUpdateOutputCoin = "update_output_coin"
	TypeMsgDeleteOutputCoin = "delete_output_coin"
)

var _ sdk.Msg = &MsgCreateOutputCoin{}

func NewMsgCreateOutputCoin(
	creator string,
	index string,

) *MsgCreateOutputCoin {
	return &MsgCreateOutputCoin{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateOutputCoin) IsPrivacy() bool {
	return true
}

func (msg *MsgCreateOutputCoin) Route() string {
	return RouterKey
}

func (msg *MsgCreateOutputCoin) Type() string {
	return TypeMsgCreateOutputCoin
}

func (msg *MsgCreateOutputCoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOutputCoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOutputCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateOutputCoin{}

func (msg *MsgUpdateOutputCoin) IsPrivacy() bool {
	return true
}
func NewMsgUpdateOutputCoin(
	creator string,
	index string,

) *MsgUpdateOutputCoin {
	return &MsgUpdateOutputCoin{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateOutputCoin) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOutputCoin) Type() string {
	return TypeMsgUpdateOutputCoin
}

func (msg *MsgUpdateOutputCoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOutputCoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOutputCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteOutputCoin{}

func NewMsgDeleteOutputCoin(
	creator string,
	index string,

) *MsgDeleteOutputCoin {
	return &MsgDeleteOutputCoin{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteOutputCoin) Route() string {
	return RouterKey
}

func (msg *MsgDeleteOutputCoin) Type() string {
	return TypeMsgDeleteOutputCoin
}

func (msg *MsgDeleteOutputCoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteOutputCoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteOutputCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgDeleteOutputCoin) IsPrivacy() bool {
	return true
}
