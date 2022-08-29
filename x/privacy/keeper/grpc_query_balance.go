package keeper

import (
	"context"
	"fmt"
	"strconv"

	"privacy/x/privacy/repos/coin"
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

	var balance uint64

	for _, outputCoin := range outputCoins {
		o := &coin.Coin{}
		err := o.SetBytes(outputCoin.Value)
		if err != nil {
			return nil, err
		}
		o, err = o.Decrypt(&keySet)
		if err == nil {
			temp := balance + o.GetValue()
			if temp < balance {
				return nil, fmt.Errorf("balance is out of range")
			}
			balance = temp
		}
	}

	ctx.Logger().Info("[incognito] balance")
	ctx.Logger().Info(strconv.Itoa(int(balance)))

	return &types.QueryBalanceResponse{
		Value: balance,
	}, nil
}
