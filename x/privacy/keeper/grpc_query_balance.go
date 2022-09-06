package keeper

import (
	"context"

	"privacy/x/privacy/common"
	"privacy/x/privacy/models"
	"privacy/x/privacy/repos/key"
	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/incognitochain/go-incognito-sdk-v2/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Balance(goCtx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	outputCoins := k.GetAllOutputCoin(ctx)

	keyWallet, err := wallet.Base58CheckDeserialize(req.PrivateKey)
	if err != nil {
		return nil, err
	}
	keySet := key.KeySet{}
	err = keySet.InitFromPrivateKeyByte(keyWallet.KeySet.PrivateKey)
	if err != nil {
		return nil, err
	}

	serialNumbers := k.GetAllSerialNumber(ctx)
	mSerialNumbers := make(map[string]types.SerialNumber)
	for _, v := range serialNumbers {
		hash := common.HashH(append([]byte{common.BoolToByte(v.IsConfidentialAsset)}, v.Value...))
		mSerialNumbers[hash.String()] = v
	}

	balance, err := models.BalanceByKeySet(outputCoins, keySet, mSerialNumbers)
	if err != nil {
		return nil, err
	}

	return &types.QueryBalanceResponse{
		Value: balance,
	}, nil
}
