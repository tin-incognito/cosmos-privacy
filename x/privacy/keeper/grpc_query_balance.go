package keeper

import (
	"context"
	"math/big"

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

	balance := big.NewInt(0)

	for _, outputCoin := range outputCoins {
		o := &coin.Coin{}
		err := o.SetBytes(outputCoin.Value)
		//err := proto.Unmarshal(outputCoin.Value, o)
		if err != nil {
			return nil, err
		}
		o, err = o.Decrypt(&keySet)
		if err == nil {
			balance = balance.Add(balance, big.NewInt(0).SetUint64(o.GetValue()))
		}
	}

	ctx.Logger().Info("[incognito] 100")
	ctx.Logger().Info(balance.String())

	return &types.QueryBalanceResponse{
		Value: balance.Bytes(),
	}, nil
}
