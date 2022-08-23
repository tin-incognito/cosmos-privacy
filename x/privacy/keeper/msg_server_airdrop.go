package keeper

import (
	"context"
	"math/big"

	"privacy/x/privacy/common"
	"privacy/x/privacy/models"
	"privacy/x/privacy/repos/coin"
	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
)

func (k msgServer) Airdrop(goCtx context.Context, msg *types.MsgAirdrop) (*types.MsgAirdropResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	privateKey := models.GeneratePrivateKey()
	otaReceiver := coin.OTAReceiver{}
	err := otaReceiver.FromString(msg.OtaReceiver)
	if err != nil {
		ctx.Logger().Error("[incognito] err:")
		ctx.Logger().Error(err.Error())
		return nil, err
	}
	amount := big.NewInt(0).SetBytes(msg.Amount)

	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		ctx.Logger().Error("[incognito] err:")
		ctx.Logger().Error(err.Error())
		return nil, err
	}
	hash := common.HashH(msgBytes)

	tx, err := models.BuildMintTx(ctx, &privateKey, otaReceiver, *amount, msg.Info, hash)
	if err != nil {
		ctx.Logger().Error("[incognito] err:")
		ctx.Logger().Error(err.Error())
		return nil, err
	}

	err = k.setPrivacyData(ctx, tx.Value)
	if err != nil {
		ctx.Logger().Error("[incognito] err:")
		ctx.Logger().Error(err.Error())
		return nil, err
	}

	return &types.MsgAirdropResponse{}, nil
}
