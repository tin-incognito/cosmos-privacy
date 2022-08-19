package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateSerialNumber{}, "privacy/CreateSerialNumber", nil)
	cdc.RegisterConcrete(&MsgUpdateSerialNumber{}, "privacy/UpdateSerialNumber", nil)
	cdc.RegisterConcrete(&MsgDeleteSerialNumber{}, "privacy/DeleteSerialNumber", nil)
	cdc.RegisterConcrete(&MsgCreateOutputCoin{}, "privacy/CreateOutputCoin", nil)
	cdc.RegisterConcrete(&MsgUpdateOutputCoin{}, "privacy/UpdateOutputCoin", nil)
	cdc.RegisterConcrete(&MsgDeleteOutputCoin{}, "privacy/DeleteOutputCoin", nil)
	cdc.RegisterConcrete(&MsgCreateCommitment{}, "privacy/CreateCommitment", nil)
	cdc.RegisterConcrete(&MsgUpdateCommitment{}, "privacy/UpdateCommitment", nil)
	cdc.RegisterConcrete(&MsgDeleteCommitment{}, "privacy/DeleteCommitment", nil)
	cdc.RegisterConcrete(&MsgCreateCommitmentIndex{}, "privacy/CreateCommitmentIndex", nil)
	cdc.RegisterConcrete(&MsgUpdateCommitmentIndex{}, "privacy/UpdateCommitmentIndex", nil)
	cdc.RegisterConcrete(&MsgDeleteCommitmentIndex{}, "privacy/DeleteCommitmentIndex", nil)
	cdc.RegisterConcrete(&MsgCreateToken{}, "privacy/CreateToken", nil)
	cdc.RegisterConcrete(&MsgUpdateToken{}, "privacy/UpdateToken", nil)
	cdc.RegisterConcrete(&MsgDeleteToken{}, "privacy/DeleteToken", nil)
	cdc.RegisterConcrete(&MsgCreateOnetimeAddress{}, "privacy/CreateOnetimeAddress", nil)
	cdc.RegisterConcrete(&MsgUpdateOnetimeAddress{}, "privacy/UpdateOnetimeAddress", nil)
	cdc.RegisterConcrete(&MsgDeleteOnetimeAddress{}, "privacy/DeleteOnetimeAddress", nil)
	cdc.RegisterConcrete(&MsgCreateTx{}, "privacy/CreateTx", nil)
	cdc.RegisterConcrete(&MsgCreateTxPrivacyData{}, "privacy/CreateTxPrivacyData", nil)
	cdc.RegisterConcrete(&MsgUpdateTxPrivacyData{}, "privacy/UpdateTxPrivacyData", nil)
	cdc.RegisterConcrete(&MsgDeleteTxPrivacyData{}, "privacy/DeleteTxPrivacyData", nil)
	cdc.RegisterConcrete(&MsgAirdrop{}, "privacy/Airdrop", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSerialNumber{},
		&MsgUpdateSerialNumber{},
		&MsgDeleteSerialNumber{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOutputCoin{},
		&MsgUpdateOutputCoin{},
		&MsgDeleteOutputCoin{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCommitment{},
		&MsgUpdateCommitment{},
		&MsgDeleteCommitment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCommitmentIndex{},
		&MsgUpdateCommitmentIndex{},
		&MsgDeleteCommitmentIndex{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateToken{},
		&MsgUpdateToken{},
		&MsgDeleteToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOnetimeAddress{},
		&MsgUpdateOnetimeAddress{},
		&MsgDeleteOnetimeAddress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTx{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTxPrivacyData{},
		&MsgUpdateTxPrivacyData{},
		&MsgDeleteTxPrivacyData{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAirdrop{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
