package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTxPrivacyData = "create_tx_privacy_data"
	TypeMsgUpdateTxPrivacyData = "update_tx_privacy_data"
	TypeMsgDeleteTxPrivacyData = "delete_tx_privacy_data"
)

var _ sdk.Msg = &MsgCreateTxPrivacyData{}

func NewMsgCreateTxPrivacyData(
	creator string,
	index string,

) *MsgCreateTxPrivacyData {
	return &MsgCreateTxPrivacyData{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateTxPrivacyData) IsPrivacy() bool {
	return true
}

func (msg *MsgCreateTxPrivacyData) Route() string {
	return RouterKey
}

func (msg *MsgCreateTxPrivacyData) Type() string {
	return TypeMsgCreateTxPrivacyData
}

func (msg *MsgCreateTxPrivacyData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTxPrivacyData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTxPrivacyData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTxPrivacyData{}

func NewMsgUpdateTxPrivacyData(
	creator string,
	index string,

) *MsgUpdateTxPrivacyData {
	return &MsgUpdateTxPrivacyData{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateTxPrivacyData) IsPrivacy() bool {
	return true
}

func (msg *MsgUpdateTxPrivacyData) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTxPrivacyData) Type() string {
	return TypeMsgUpdateTxPrivacyData
}

func (msg *MsgUpdateTxPrivacyData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTxPrivacyData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTxPrivacyData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTxPrivacyData{}

func NewMsgDeleteTxPrivacyData(
	creator string,
	index string,

) *MsgDeleteTxPrivacyData {
	return &MsgDeleteTxPrivacyData{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteTxPrivacyData) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTxPrivacyData) Type() string {
	return TypeMsgDeleteTxPrivacyData
}

func (msg *MsgDeleteTxPrivacyData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTxPrivacyData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTxPrivacyData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgDeleteTxPrivacyData) IsPrivacy() bool {
	return true
}
