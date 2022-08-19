package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAirdrop = "airdrop"

var _ sdk.Msg = &MsgAirdrop{}

func NewMsgAirdrop(creator string) *MsgAirdrop {
	return &MsgAirdrop{
		Creator: creator,
	}
}

func (msg *MsgAirdrop) Route() string {
	return RouterKey
}

func (msg *MsgAirdrop) Type() string {
	return TypeMsgAirdrop
}

func (msg *MsgAirdrop) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAirdrop) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAirdrop) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
