package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgPrivacyData = "privacy_data"

var _ sdk.Msg = &MsgPrivacyData{}

func NewMsgPrivacyData(
	lockTime, fee uint64,
	info, sigPubKey, sig, proof []byte,
	txType int32, metadata []byte,
) *MsgPrivacyData {
	return &MsgPrivacyData{
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
	/*creator, err := sdk.AccAddressFromBech32(msg.Creator)*/
	/*if err != nil {*/
	/*panic(err)*/
	/*}*/
	//return []sdk.AccAddress{creator}
	return []sdk.AccAddress{}
}

func (msg *MsgPrivacyData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPrivacyData) ValidateBasic() error {
	/*_, err := sdk.AccAddressFromBech32(msg.Creator)*/
	/*if err != nil {*/
	/*return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)*/
	/*}*/
	fmt.Println("ValidateBasic")
	return nil
}

func (msg *MsgPrivacyData) IsPrivacy() bool {
	return true
}
