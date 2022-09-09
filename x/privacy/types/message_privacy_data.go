package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPrivacyData = "privacy_data"

var _ sdk.Msg = &MsgPrivacyData{}

func NewMsgPrivacyData(
	creator string,
	lockTime, fee uint64,
	info, sigPubKey, sig, proof []byte,
	txType int32, metadata []byte,
) *MsgPrivacyData {
	return &MsgPrivacyData{
		Creator:   creator,
		LockTime:  lockTime,
		Fee:       fee,
		Info:      info,
		SigPubKey: sigPubKey,
		Sig:       sig,
		TxType:    txType,
		Metadata:  metadata,
	}
}

func (msg *MsgPrivacyData) Route() string {
	return RouterKey
}

func (msg *MsgPrivacyData) Type() string {
	return TypeMsgPrivacyData
}

func (msg *MsgPrivacyData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPrivacyData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPrivacyData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
