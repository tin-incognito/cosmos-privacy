package keeper

import (
	"context"
	"math/big"

	"privacy/x/privacy/models"
	"privacy/x/privacy/repos/key"
	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/incognitochain/go-incognito-sdk-v2/wallet"
)

func (k msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	keyWallet, err := wallet.Base58CheckDeserialize(msg.PrivateKey)
	if err != nil {
		ctx.Logger().Error("[incognito] err:")
		ctx.Logger().Error(err.Error())
		return nil, err
	}
	keySet := key.KeySet{}
	err = keySet.InitFromPrivateKeyByte(keyWallet.KeySet.PrivateKey)
	if err != nil {
		ctx.Logger().Error("[incognito] err:")
		ctx.Logger().Error(err.Error())
		return nil, err
	}

	outputCoinLength := big.NewInt(0)
	outputCoinSerialNumber, _ := k.GetOutputCoinSerialNumber(ctx)
	outputCoinLength.SetBytes(outputCoinSerialNumber.Value)

	outputCoins := k.GetAllOutputCoin(ctx)
	mOutputCoins := make(map[string]types.OutputCoin)
	for _, v := range outputCoins {
		mOutputCoins[v.Index] = v
	}

	tx, err := models.BuildTransferTx(outputCoins, keySet, msg, *outputCoinLength, mOutputCoins, outputCoins)
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

	return &types.MsgTransferResponse{
		Msg: "Success",
	}, nil
}
