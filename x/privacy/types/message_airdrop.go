package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgAirdrop = "airdrop"

var _ sdk.Msg = &MsgAirdrop{}

func NewMsgAirdrop(otaReceiver string, amount, info []byte) *MsgAirdrop {
	return &MsgAirdrop{
		OtaReceiver: otaReceiver,
		Amount:      amount,
		Info:        info,
	}
}

func (msg *MsgAirdrop) Route() string {
	return RouterKey
}

func (msg *MsgAirdrop) Type() string {
	return TypeMsgAirdrop
}

func (msg *MsgAirdrop) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgAirdrop) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAirdrop) ValidateBasic() error {
	return nil
}
